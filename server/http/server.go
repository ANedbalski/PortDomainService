package http

import (
	chilogger "github.com/766b/chi-logger"
	"github.com/go-chi/chi"

	"go.uber.org/zap"
	"net/http"
)

type Server struct {
	server *http.Server
	log    *zap.SugaredLogger
}

func New(host string, log *zap.SugaredLogger) *Server {
	mux := chi.NewMux()
	mux.Use(chilogger.NewZapMiddleware("router", log.Desugar()))

	s := &Server{log: log}

	mux.HandleFunc("/ports", s.PortsHandler)
	s.server = &http.Server{Addr: host, Handler: mux}
	return s
}

func (s *Server) GetHTTPServer() *http.Server {
	return s.server
}

// authenticate user
// placeholder
func (s *Server) validateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
