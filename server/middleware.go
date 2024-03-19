package server

import (
	"net/http"
	"net/http/httputil"
	"net/url"
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

func reverseProxy(proxy *httputil.ReverseProxy) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})
}

func mustReverseProxy(s *server, rawURL string) http.Handler {
	url, err := url.Parse(rawURL)
	if err != nil {
		s.log.Error("failed to parse reverse proxy url", "rawURL", rawURL, "error", err)
		panic(rawURL)
	}
	proxy := &httputil.ReverseProxy{
		Rewrite: func(r *httputil.ProxyRequest) {
			r.SetURL(url)
		},
	}
	return reverseProxy(proxy)
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
