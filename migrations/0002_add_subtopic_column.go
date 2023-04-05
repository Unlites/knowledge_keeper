package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(up0002, down0002)
}

func up0002(tx *sql.Tx) error {
	query := `ALTER TABLE records
	ADD COLUMN subtopic VARCHAR(100) DEFAULT '';`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func down0002(tx *sql.Tx) error {
	query := `ALTER TABLE records
	DROP COLUMN subtopic;`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
