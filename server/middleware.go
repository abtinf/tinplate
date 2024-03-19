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

func onlyWhenReady(s *server, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !s.ready.Load() {
			http.Error(w, "service not ready", http.StatusServiceUnavailable)
			s.log.Info("service called when not ready", "method", r.Method, "url", r.URL.Redacted())
			return
		}
		handler.ServeHTTP(w, r)
	})
}
