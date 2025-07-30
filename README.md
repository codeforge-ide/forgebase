# Forgebase

**Forgebase** is an open-source, AI-powered Firebase alternative written in Go. It aims to provide scalable, modular, and developer-friendly backend services for modern apps.

## Features (planned)
- Realtime database API
- Authentication & user management
- AI-powered data analysis
- Modular plugin system
- REST API (with future gRPC/WebSocket support)

## Getting Started

```sh
# Run locally
make run

# Run tests
make test

# Build Docker image
make docker
```

## Project Structure
- `cmd/forgebase` — Entrypoint
- `internal/api` — API server & handlers
- `internal/db` — Database abstraction
- `internal/ai` — AI module
- `internal/config` — Config loader
- `pkg` — Utilities
- `tests` — Tests

## Contributing
PRs and issues welcome!
