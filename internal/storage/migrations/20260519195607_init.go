package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upInit, downInit)
}

func upInit(ctx context.Context, tx *sql.Tx) error {
	query := `
		CREATE TABLE tasks (
          id TEXT PRIMARY KEY,
          title TEXT NOT NULL,
          description TEXT,
          status TEXT NOT NULL,
          priority INTEGER DEFAULT 0,
          created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
          deleted_at DATETIME DEFAULT NULL,
          updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
        );	

	CREATE TRIGGER update_tasks_timestamp 
		BEFORE UPDATE ON tasks
		BEGIN
			UPDATE tasks SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
		END;
	`
	_, err := tx.ExecContext(ctx, query)
	return err
}

func downInit(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "DROP TRIGGER IF EXISTS update_tasks_timestamp;")
	_, err = tx.ExecContext(ctx, "DROP TABLE IF EXISTS tasks;")
	return err
}
