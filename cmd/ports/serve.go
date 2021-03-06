package ports

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"ports/domain/services"
	"ports/domain/services/importer"
	"ports/repository/memory"
	http_srv "ports/server/http"
	"ports/task"
	"sync"
	"syscall"
)

var (
	serveCommand = &cli.Command{
		Name:    "serve",
		Aliases: []string{"s"},
		Usage:   "start the server",
		Action:  serveAction,
	}
)

func serveAction(c *cli.Context) error {
	cfg, err := initConfig(c.String("config"), "pds")
	if err != nil {
		return fmt.Errorf("error initialising config: %w", err)
	}

	log, err := initLogger(zap.PanicLevel)
	if err != nil {
		return fmt.Errorf("error initialising logger: %w", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//instantiate services

	portService, err := services.NewPort(services.WithInMemoryPortRepository(memory.NewPort()))
	if err != nil {
		return fmt.Errorf("cannot instantiate Port Service: %w", err)
	}

	portImportService, err := services.NewPortImport(
		services.WithInMemoryPortImportRepository(memory.NewPort()),
		services.WithStreamingPortImporter(importer.NewJSONStream()),
		services.WithSugaredLogger(log))
	if err != nil {
		return fmt.Errorf("cannot instantiate Port Service: %w", err)
	}

	var wg sync.WaitGroup

	//setup event dispatcher
	// currently launching in the serve process, because test implementation
	// implements only chan version of queue
	// supposed to live in the separate process
	dispatcher := task.NewDispatcher(portImportService, log)
	dispatcher.Sub()
	dispatcher.Run(ctx)

	//Setup http rest api server
	restSrv := http_srv.New(cfg.Server.HTTP.Service, log, portService)
	startHTTPServer(&wg, restSrv.GetHTTPServer(), log)

	go func() {
		graceful := true
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

		for {
			sig := <-sig

			if !graceful {
				log.Info("caught another exit signal, now hard dying", "signal", sig)
				os.Exit(1)
			}
			graceful = false

			go func() {
				log.Infow("starting graceful shutdown", "signal", sig)

				cancel()

				if restSrv.GetHTTPServer() != nil {
					if err := restSrv.GetHTTPServer().Shutdown(context.Background()); err != nil {
						log.Errorw("error shutting down callback listener", "err", err)
					} else {
						log.Infow("callback server shutdown complete")
					}
				}
			}()
		}
	}()

	wg.Wait()
	return nil
}

func startHTTPServer(wg *sync.WaitGroup, server *http.Server, log *zap.SugaredLogger) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		log.Infow("HTTP server starting", "addr", server.Addr)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
}
