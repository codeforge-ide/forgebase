# AGENTS.md

## Build, Lint, and Test Commands
- **Run app:** `make run` or `go run ./cmd/forgebase`
- **Run all tests:** `make test` or `go test ./...`
- **Run a single test:** `go test -run TestName`
- **Build binary:** `make build` or `go build -o forgebase ./cmd/forgebase`
- **Build Docker image:** `make docker`

## Code Style Guidelines
- **Formatting:** Use `gofmt` (standard Go formatting)
- **Imports:** Group as stdlib, third-party, then local; separate blocks
- **Types:** Use clear, descriptive names; interfaces end with `-er` or describe capability
- **Naming:** CamelCase for exported, lowerCamel for unexported, ALL_CAPS for constants
- **Error Handling:** Always check errors, return early, wrap with context if needed
- **Comments:** Use full sentences; doc comments for exported items
- **Tests:** Place in `*_test.go`, use `t.Fatalf` for errors, prefer table-driven tests

No Cursor or Copilot rules are present.

notes:

1. run git add ., generate one sentence, concise commit message (no info about authors or new lines, just one sentence commit message) and push, afte  each code changes
2. remember, we are building a platform similar to firebase/supabase and need to carefully structure the architecture modularly.