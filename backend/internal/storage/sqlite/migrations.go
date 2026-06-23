package sqlite

import "database/sql"

func Migrate(db *sql.DB) error {

	query := `
	CREATE TABLE IF NOT EXISTS identity (
		id INTEGER PRIMARY KEY CHECK (id = 1),

		public_key BLOB NOT NULL,

		private_key BLOB NOT NULL,

		created_at DATETIME NOT NULL
	);
	`

	_, err := db.Exec(query)

	return err
}
