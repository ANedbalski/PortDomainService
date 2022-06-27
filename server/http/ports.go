package http

import (
	"fmt"
	"net/http"
)

func (s *Server) PortsHandler(w http.ResponseWriter, r *http.Request) {
	err := s.validateUser(w, r)
	if err != nil {
		s.log.Errorw("error validate identity", "req", r, "err", err)
		_, err := fmt.Fprintln(w, err)
		if err != nil {
			s.log.Errorw("error sending response", "req", r, "err", err)
		}
		return
	}
}
