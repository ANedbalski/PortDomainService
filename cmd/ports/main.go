package ports

import (
	"github.com/spf13/viper"
	cli "github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"ports/config"
	"strings"

	"log"
	"os"
)

func main() {
	app := &cli.App{
		Usage: "Run the ports service",
		Commands: []*cli.Command{
			serveCommand,
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "pds.yml",
				Usage:   "path to config file",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Application failed: %v", err)
	}

}

func initConfig(filePath, envPrefix string) (*config.Config, error) {
	cfg := (*config.Config)(nil)
	v := viper.New()
	v.SetConfigFile(filePath)
	if envPrefix != "" {
		v.SetEnvPrefix(envPrefix)
	}
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	err := v.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func initLogger(l zapcore.Level) (*zap.SugaredLogger, error) {
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(l),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
	lo, err := cfg.Build()
	return lo.Sugar(), err
}
