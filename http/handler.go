package http

import (
	"net/http"
)

func Handler(s *HttpServer) {
	s.logger.Info("Server listening")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		s.logger.Info("Status OK")
	})
	http.ListenAndServe(":3000", nil)
}
