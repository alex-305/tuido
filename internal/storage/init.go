package storage

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pressly/goose/v3"

	_ "github.com/alex-305/tuido/internal/storage/migrations"
	_ "modernc.org/sqlite"
)

type Store struct {
	db *sql.DB
}

func NewStore() (*Store, error) {
	dbPath, err := GetDBPath()
	if err != nil {
		return nil, fmt.Errorf("failed to resolve db path: %w", err)
	}

	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create data directory: %w", err)
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	if err := goose.SetDialect("sqlite3"); err != nil {
		return nil, fmt.Errorf("failed to set goose dialect: %w", err)
	}

	if err := goose.Up(db, "internal/storage/migrations"); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}
	return &Store{
		db: db,
	}, nil
}
