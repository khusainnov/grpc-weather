package http

import "net/http"

func (s *Server) setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/alive", func(w http.ResponseWriter, r *http.Request) {
		s.cfg.L.Info("alive endpoint")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<h1>Server is working</h1>"))
		return
	})

	return mux
}
