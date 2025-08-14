# Agent Guidelines for migrate-libsql

## Build/Test/Lint Commands
- **Build**: `go build .` or `go install`
- **Test**: `go test -race ./...` (run all tests with race detection)
- **Lint**: `golangci-lint run` (uses .golangci.yml config)
- **Format**: `goimports -w .` and `gofmt -w .`
- **Download deps**: `go mod download`

## Code Style Guidelines
- **Language**: Go 1.24+ (see go.mod)
- **Imports**: Standard library first, then external packages, group with blank lines
- **Error handling**: Always wrap errors with fmt.Errorf and %w verb for context
- **Naming**: Use camelCase for local vars, PascalCase for exported funcs/types
- **Functions**: Keep functions focused, use early returns for error conditions
- **Variables**: Use descriptive names, avoid single-letter vars except for loops
- **Comments**: No comments needed unless explaining complex business logic

## Project Structure
- Single-file CLI tool in main.go
- Uses golang-migrate for database migrations
- LibSQL/Turso database support with authentication tokens
- File system-based migration sources

## Dependencies
- github.com/golang-migrate/migrate/v4 for migration logic
- github.com/tursodatabase/libsql-client-go for LibSQL driver