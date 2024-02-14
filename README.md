# Tinplate - A Go Service Template

An implementation of a production-ready service written in Go.

Not ready yet.

## Features
* `manage.sh` script to handle common tasks
* Configuration with flowdown from defaults, environment variables, and command line arguments
* Structured logging with redaction of configuration secrets
* Probe monitor endpoints for startup, liveness, and readiness
* Graceful shutdown
* Static embedding of assets
  * Self-hosting, generated swagger and swagger-ui: `http://localhost:8080/static/http/swagger-ui/`
* Prometheus support
* Database support (in-memory SQLite)
  * Generate database code using sqlc
* HTTP endpoints generated using protobufs
  * Get Example: `curl http://localhost:8080/api/ExampleGet?name=tinplate`
  * Post Example: `curl -X 'POST' 'http://localhost:8080/api/ExamplePost' -d '{"name": "tinplate"}'`
  * Download example: `curl http://localhost:8080/api/download`
* Middleware
  * Logging 
  * Protocol upgrade (HTTP and GRPC served on same port)
* Unencrypted HTTP/2


## TODO
* Endpoint Example
  * basic auth
  * JWT
  * Reverse proxy
  * GRPC/HTTP mapped code returns
  * Custom HTTP status code return
  * File upload
  * Business logic calls
* HTTP
  * Favicon
* Interop
  * Client libraries
    * React
    * Go
    * Postman
* Middleware
  * CORS
* Database Support
  * Postgres
  * DB migrations
  * DB leadership
  * Replace sqlite with embedded postgres
* TLS support
  * Certificate loading
  * Use h2c if no cert available
  * Let's Encrypt
* Security & Auditability
  * Rate limiting
  * Audit log
  * RBAC
  * OAUTH/JWT validation
* Resiliency
  * Service dependency monitoring
    * DB
* Observability
  * Instrument calls and add custom metrics to prometheus
  * Tracing middleware
  * Print source location on error
* Configuration
  * Improve argument parsing
    * Port should be a type that enforces limit
    * Examples for parsing user input
  * Feature flags
  * Feature entitlements
* Testing
  * Unit tests
  * Integration tests
* Documentation
  * Instructions for modifying the database
  * Instructions for adding config parameters
  * Add comments and generate godoc
* Developer QoL
  * Visual Studio Code config (f5 to run, environment, etc)
  * Hot reload for server and UI
  * Admin interface/API
  * Basic UI
  * Move main to cmd
  * Tailscale
  * More swagger annotations
  * Github workflows
  * Special github files
* Build and depoy
  * Embed git commit hash in build
  * Helm charts/kubernetes/local deploy support
  * Add build to manage.sh
  * Dockerfile
* Bugs
  * Fix static file path (includes http right now)
  * Don't list directory contents of static

## Getting Started
* Install sqlc
  * `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`
* Install protobuf compiler
  * `brew install protobuf@3`
  * Add the `protoc` binary to your path: `export PATH="/opt/homebrew/opt/protobuf@3/bin:$PATH"`
  * `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
  * `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
  * Add the go plugins in to your path so `protoc` can find them: `export PATH="$PATH:$(go env GOPATH)/bin"`
  * `go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest`
  * `go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest`
* Modifying endpoints
  * Edit `proto/api.proto`
  * Run `./manage.sh -p`
  * Add the endpoints to `server/endpoints.go`
