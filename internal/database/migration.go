package database

func Migrate() error {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id TEXT PRIMARY KEY,
        username TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL
    );
    CREATE TABLE IF NOT EXISTS files (
        userId TEXT PRIMARY KEY,
        public BOOLEAN DEFAULT 0,
        type STRING NOT NULL,
        name STRING NOT NULL,
        entryId STRING NOT NULL UNIQUE
    );`
	_, err := DB.Exec(query)
	return err
}
