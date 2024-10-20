package database

import (
	"passfish/internal/models"
)

// Creates the credentials table if it does not exist.
func (db *Db) createCredentialsTable() error {
	sqlStmt := `
	create table if not exists credentials (
		title text not null primary key,
		username text not null,
		password blob not null,
		last_accessed datetime default current_timestamp
	);
	`

	_, err := db.Conn.Exec(sqlStmt)
	return err
}

func (db *Db) InsertCredentials(creds models.BaseCredentials) error {
	sqlStmt := `
	insert into credentials (title, username, password) values (?, ?, ?);
	`

	_, err := db.Conn.Exec(sqlStmt, creds.Title, creds.Username, creds.Password)
	return err
}

func (db *Db) NumberOfCredentials() int {
	sqlStmt := `
	select count(*) from credentials;
	`
	row := db.Conn.QueryRow(sqlStmt)

	var count int
	row.Scan(&count)
	return count
}

func (db *Db) GetTitles() ([]string, error) {
	sqlStmt := `
	select title from credentials;
	`
	rows, err := db.Conn.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize the slice with a capacity of 50 (arbitrary).
	titles := make([]string, 0, 50)
	for rows.Next() {
		var title string
		err = rows.Scan(&title)
		if err != nil {
			return nil, err
		}

		titles = append(titles, title)
	}

	return titles, nil
}

func (db *Db) GetCredentials(title string) (*models.Credentials, error) {
	sqlStmt := `
	select title, username, password, last_accessed from credentials where title = ?;
	`
	row := db.Conn.QueryRow(sqlStmt, title)

	var creds models.Credentials
	err := row.Scan(&creds.Base.Title, &creds.Base.Username, &creds.Base.Password, &creds.LastAccessed)
	if err != nil {
		return nil, err
	}

	return &creds, nil
}

func (db *Db) MarkAccessed(title string) error {
	sqlStmt := `
	update credentials set last_accessed = current_timestamp where title = ?;
	`

	_, err := db.Conn.Exec(sqlStmt, title)
	return err
}

func (db *Db) UpdateCredentials(creds models.BaseCredentials) error {
	sqlStmt := `
	update credentials set username = ?, password = ? where title = ?;
	`

	_, err := db.Conn.Exec(sqlStmt, creds.Username, creds.Password, creds.Title)
	return err
}

func (db *Db) DeleteCredentials(title string) error {
	sqlStmt := `
	delete from credentials where title = ?;
	`

	_, err := db.Conn.Exec(sqlStmt, title)
	return err
}
