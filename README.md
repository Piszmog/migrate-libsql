# migrate-libsql

A simple command-line tool for running database migrations on LibSQL/Turso databases using [golang-migrate](https://github.com/golang-migrate/migrate).

## Features

- Run migrations up (apply all pending) or down (rollback specific steps)
- Support for LibSQL URLs with authentication tokens
- File system-based migration sources
- Built with Go for cross-platform compatibility

## Installation

Download the latest binary from the [releases page](https://github.com/your-username/migrate-libsql/releases) or build from source:

```bash
go install github.com/your-username/migrate-libsql@latest
```

## Usage

```bash
migrate-libsql -url <libsql-url> -token <auth-token> -migrations <migrations-dir> [options]
```

### Flags

| Flag | Default | Required | Description |
|------|---------|----------|-------------|
| `-url` | - | ✅ | LibSQL database URL (e.g., `libsql://database-name.turso.io`) |
| `-token` | - | ✅ | LibSQL authentication token |
| `-migrations` | - | ✅ | Path to directory containing migration files |
| `-direction` | `up` | ❌ | Migration direction: `up` (apply) or `down` (rollback) |
| `-steps` | `1` | ❌ | Number of steps for down migration |

### Examples

```bash
# Apply all pending migrations
migrate-libsql -url "libsql://my-db.turso.io" -token "your-token" -migrations "./migrations"

# Rollback the last migration
migrate-libsql -url "libsql://my-db.turso.io" -token "your-token" -migrations "./migrations" -direction down

# Rollback the last 3 migrations  
migrate-libsql -url "libsql://my-db.turso.io" -token "your-token" -migrations "./migrations" -direction down -steps 3
```

## Migration Files

Migration files should follow the standard golang-migrate naming convention:

```
migrations/
├── 001_initial_schema.up.sql
├── 001_initial_schema.down.sql
├── 002_add_users_table.up.sql
└── 002_add_users_table.down.sql
```

## Building

```bash
go build -o migrate-libsql .
```

## License

MIT
