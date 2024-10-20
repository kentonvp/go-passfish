package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Db is a database connection.
type Db struct {
	Conn *sql.DB
}

// New creates a new database connection obviously.
func New(dbPath string) (*Db, error) {
	dbConn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	db := Db{
		Conn: dbConn,
	}

	return &db, nil
}

// Closes the database connection.
func (db *Db) Close() {
	db.Conn.Close()
}

func (db *Db) CreateTables() error {
  if err := db.createCredentialsTable(); err != nil {
    return err
  }

  if err := db.createPassphraseTable(); err != nil {
    return err
  }

  return nil
}
