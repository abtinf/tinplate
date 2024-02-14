package server

import (
	"net/http"
)

func upgradeHandler(base http.Handler, upgrade http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && r.Header.Get("Content-Type") == "application/grpc" {
			upgrade.ServeHTTP(w, r)
		} else {
			base.ServeHTTP(w, r)
		}
	})
}

func logger(s *server, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.log.Info("request", "method", r.Method, "url", r.URL.Redacted())
		handler.ServeHTTP(w, r)
	})
}
