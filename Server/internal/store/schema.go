package store

import "database/sql"

func SetupDB(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			email TEXT UNIQUE NOT NULL,
			firstName TEXT NOT NULL,
			lastName TEXT NOT NULL,
			joinDate DATE NOT NULL,
			pwdHash TEXT NOT NULL,
			streetAddress TEXT,
			city TEXT,
			country TEXT,
			postCode TEXT
		);		
	`)

	return err
}
