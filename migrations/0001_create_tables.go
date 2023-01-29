package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(up0001, down0001)
}

func up0001(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(20) NOT NULL UNIQUE,
		password_hash VARCHAR(255) NOT NULL,
		refresh_token VARCHAR(255),
		token_expires_at INTEGER
	);`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func down0001(tx *sql.Tx) error {
	query := "DROP TABLE users;"
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
