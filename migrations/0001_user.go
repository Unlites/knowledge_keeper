package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(up, down)
}

func up(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(20),
		password_hash VARCHAR(255)
	);`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func down(tx *sql.Tx) error {
	query := "DROP TABLE users;"
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
