package sqlite

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func Open(
	path string,
) (*sql.DB, error) {

	dir := filepath.Dir(path)

	if err := os.MkdirAll(
		dir,
		0755,
	); err != nil {
		return nil, err
	}

	return sql.Open(
		"sqlite",
		path,
	)
}
