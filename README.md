# Gonfoot - A Footgun-Free Go Service Template

An implementation of a production-ready service written in Go.

Will be ready soon.

## Features

- Observability
  - Prometheus support
  - Structured logging with redaction of configuration secrets
- Ops
  - Probe monitor endpoints for startup, liveness, and readiness
  - Graceful shutdown
- Database support
  - Postgres
  - Embedded postgres for development
  - Migration leadership
- Code Generation
  - HTTP endpoints generated using protobufs
    - Unencrypted HTTP/2
  - Database code generated using sqlc
- Middleware
  - Logging
  - Protocol upgrade (HTTP and GRPC served on same port)
- Resiliency
  - Service dependency monitoring
    - DB
- Build and deploy
  - Static asset hosting
    - Generated swagger and self-hosted swagger-ui: `http://localhost:8080/static/swagger-ui/`
    - Does not expose directory contents of static assets
    - Root favicon
- Developer QoL
  - Visual Studio Code config (f5 to run, environment, etc)
  - `manage.sh` script to handle common tasks
  - Configuration with flowdown from defaults, environment variables, and command line arguments

## Roadmap

### Soon

- Middleware
  - CORS
- Developer QoL
  - Move main to cmd

### Later

- Documentation
  - Instructions for modifying the database
  - Instructions for adding config parameters
  - Add comments and generate godoc
- TLS support
  - Only use h2c if no cert available
  - Let's Encrypt
- Build and depoy
  - Embed git commit hash in build
  - Add build to manage.sh
  - Dockerfile
- Middleware
  - Rate limiting
- Database Support
  - Monitor database for migration changes
- Observability
  - Print source location on error
- Define "production-ready"

### Future

- Special GitHub files
  - Example workflow
  - CODE_OF_CONDUCT.md
  - CONTRIBUTING.md
  - .github/dependabot.yml
  - .github/CHANGELOG
  - .github/SUPPORT.md
  - .github/SECURITY.md
  - .github/CODEOWNERS
  - .github/ISSUE_TEMPLATE
  - .github/PULL_REQUEST_TEMPLATE
- Security & Auditability
  - Audit log
  - RBAC
  - OAUTH/JWT validation
- Developer QoL
  - Hot reload for server and UI
  - Admin interface/API
  - Basic UI
  - Tailscale
  - Postman
  - CLI client
  - Unit tests
  - More swagger annotations
- Quality
  - Integration tests
- Code Generation
  - React client lib
  - Go client lib
- Observability
  - Tracing
  - Instrument calls and add custom metrics to prometheus
- Configuration
  - Improve argument parsing
    - Port should be a type that enforces limit
    - Examples for parsing user input
  - Feature flags
  - Feature entitlements
- TLS support
  - Certificate loading
- Build & Deploy
  - Helm charts/kubernetes/local deploy support
  - Package as standalone app
  - Minify static assets
- Endpoint Example
  - basic auth
  - JWT
  - Reverse proxy
  - GRPC/HTTP mapped code returns
  - Custom HTTP status code return
  - File upload
  - Business logic calls
  - Cookie/session support with encryption
- Web UI
  - Basic UI
  - i18n

## Getting Started

- Install sqlc
  - `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`
- Install protobuf compiler
  - `brew install protobuf@3`
  - Add the `protoc` binary to your path: `export PATH="/opt/homebrew/opt/protobuf@3/bin:$PATH"`
  - `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
  - `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
  - Add the go plugins in to your path so `protoc` can find them: `export PATH="$PATH:$(go env GOPATH)/bin"`
  - `go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest`
  - `go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest`
- Modifying endpoints
  - Edit `proto/api.proto`
  - Run `./manage.sh -p`
  - Add the endpoints to `server/endpoints.go`
