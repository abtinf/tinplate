// Implement a production-ready GRPC/HTTP server
package server

import (
	"time"
)

func (s *server) liveMonitor() {
	t := time.Tick(s.config.monitorInterval)
	for range t {
		if s.live.CompareAndSwap(false, true) {
			s.log.Info("server health changed", "live", true)
		}
	}
}

func (s *server) readyMonitor() {
	t := time.Tick(s.config.monitorInterval)
	for range t {
		ready := s.httpServerAvailable.Load() && s.databaseAvailable.Load()
		if s.ready.CompareAndSwap(!ready, ready) {
			s.log.Info("server health changed", "ready", ready)
		}
	}
}

func (s *server) dbMonitor() {
	t := time.Tick(s.config.monitorInterval)
	for range t {
		dbAvailable := true
		if s.databaseAvailable.CompareAndSwap(!dbAvailable, dbAvailable) {
			s.log.Info("database health changed", "available", dbAvailable)
		}
	}
}
