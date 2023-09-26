package database

import (
	"database/sql"
)

func ExecuteCommand(db *sql.DB, command string, args ...any) error {
	_, err := db.Exec(command, args...)
	return err
}

func CreateTables() error {
	tables := []string{
		`
        CREATE TABLE IF NOT EXISTS todo (
            id INTEGER PRIMARY KEY,
            title VARCHAR(128),
            description VARCHAR(1024),
			completed BOOLEAN
        )
    `,
	}
	for _, table := range tables {
		err := ExecuteCommand(DB, table)
		if err != nil {
			return err
		}
	}
	return nil
}
