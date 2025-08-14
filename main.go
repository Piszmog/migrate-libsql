package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"io/fs"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	url := flag.String("url", "", "LibSQL database URL (required)")
	token := flag.String("token", "", "LibSQL authentication token (required)")
	migrationsDir := flag.String("migrations", "", "Path to migrations directory (required)")
	direction := flag.String("direction", "up", "Migration direction: up or down (default: up)")
	steps := flag.Int("steps", 1, "Number of steps for down migration (default: 1)")
	flag.Parse()

	if *url == "" {
		fmt.Printf("Error: -url flag is required\n")
		flag.Usage()
		os.Exit(1)
	}

	if *token == "" {
		fmt.Printf("Error: -token flag is required\n")
		flag.Usage()
		os.Exit(1)
	}

	if *migrationsDir == "" {
		fmt.Printf("Error: -migrations flag is required\n")
		flag.Usage()
		os.Exit(1)
	}

	db, err := newDB(*url, *token)
	if err != nil {
		fmt.Printf("Error: failed to create DB connection: %+v\n", err)
		return
	}
	defer func() {
		_ = db.Close()
	}()

	migrationsFS := os.DirFS(*migrationsDir)

	m, err := newMigration(db, migrationsFS)
	if err != nil {
		fmt.Printf("Error: failed to create migration: %+v\n", err)
		return
	}
	defer func() {
		_, _ = m.Close()
	}()

	switch *direction {
	case "up":
		err = m.Up()
		if err != nil {
			fmt.Printf("Error: failed to migrate up: %+v\n", err)
			return
		}
		fmt.Println("Migration up completed successfully")
	case "down":
		if *steps <= 0 {
			fmt.Printf("Error: steps must be greater than 0 for down migration\n")
			return
		}
		err = m.Steps(-*steps)
		if err != nil {
			fmt.Printf("Error: failed to migrate down %d steps: %+v\n", *steps, err)
			return
		}
		fmt.Printf("Migration down %d steps completed successfully\n", *steps)
	default:
		fmt.Printf("Error: invalid direction '%s'. Must be 'up' or 'down'\n", *direction)
	}
}

func newDB(url string, token string) (*sql.DB, error) {
	fullURL := url + "?authToken=" + token
	db, err := sql.Open("libsql", fullURL)
	if err != nil {
		return nil, err
	}
	if err = db.PingContext(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	return db, nil
}

func newMigration(db *sql.DB, migrations fs.FS) (*migrate.Migrate, error) {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to create database driver: %w", err)
	}

	iofsDriver, err := iofs.New(migrations, "migrations")
	if err != nil {
		return nil, fmt.Errorf("failed to create iofs: %w", err)
	}
	defer func() {
		_ = iofsDriver.Close()
	}()

	m, err := migrate.NewWithInstance("iofs", iofsDriver, "sqlite3", driver)
	if err != nil {
		return nil, fmt.Errorf("failed to create migration: %w", err)
	}
	return m, nil
}
