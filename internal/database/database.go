package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Db is a database connection.
type Db struct {
	Conn *sql.DB
}

// NewDB creates a new database connection.
func NewDB(dbPath string) (*Db, error) {
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
