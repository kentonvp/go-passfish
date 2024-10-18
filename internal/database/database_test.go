package database_test

import (
	"fmt"
	"os"
	"passfish/internal/database"
	"passfish/internal/models"
	"path"
	"testing"
)

func mockDbPath() string {
	dir, err := os.MkdirTemp(".", "t_*")
	if err != nil {
		panic(err)
	}
	return path.Join(dir, "test.db")
}

func cleanUp(dir string) {
	os.RemoveAll(dir)
}

const nTestUsers int = 10

func mockDb() (*database.Db, string) {
	dbPath := mockDbPath()
	db, err := database.NewDB(dbPath)
	if err != nil {
		panic(err)
	}

	if err := db.CreateCredentialsTable(); err != nil {
		panic(err)
	}

	for i := 0; i < nTestUsers; i++ {
		if err := db.InsertCredentials(models.Credentials{
			Title:    fmt.Sprintf("loginTitle%d", i),
			Username: fmt.Sprintf("user%d", i),
			Password: "password",
		}); err != nil {
			panic(err)
		}
	}

	return db, dbPath
}

func TestNewDb(t *testing.T) {
	dbPath := mockDbPath()
	defer cleanUp(path.Dir(dbPath))

	db, err := database.NewDB(dbPath)
	if err != nil {
		t.Error("Expected nil, got an error")
	}
	db.Close()
}

func TestCreateCredentialsTable(t *testing.T) {
	dbPath := mockDbPath()
	defer cleanUp(path.Dir(dbPath))

	db, _ := database.NewDB(dbPath)
	defer db.Close()

	if err := db.CreateCredentialsTable(); err != nil {
		t.Error("Expected nil, got an error")
	}
}

func TestInsertCredential(t *testing.T) {
	dbPath := mockDbPath()
	defer cleanUp(path.Dir(dbPath))

	db, _ := database.NewDB(dbPath)
	defer db.Close()

	err := db.CreateCredentialsTable()
	if err != nil {
		t.Error("Unexpected error: creating credentials table in TestInsertCredential", err)
	}

	creds := models.Credentials{
		Title:    "loginTitle",
		Username: "user1",
		Password: "password",
	}

	if err := db.InsertCredentials(creds); err != nil {
		t.Error("Expected nil, got an error")
	}
}

func TestGetTitles(t *testing.T) {
	db, dbPath := mockDb()
	defer cleanUp(path.Dir(dbPath))
	defer db.Close()

	titles, err := db.GetTitles()
	if err != nil {
		t.Error("Expected nil, got an error")
	}

	if len(titles) != nTestUsers {
		t.Error("Expected 1, got", len(titles))
	}
}

func TestGetCredentials(t *testing.T) {
	db, dbPath := mockDb()
	defer cleanUp(path.Dir(dbPath))
	defer db.Close()

	creds, err := db.GetCredentials("loginTitle0")
	if err != nil {
		t.Error("Expected nil, got an error")
	}

	if creds.Username != "user0" {
		t.Error("Expected \"user0\", got ", creds.Username)
	}

	creds, err = db.GetCredentials("loginTitle9")
	if err != nil {
		t.Error("Expected nil, got an error")
	}

	if creds.Username != "user9" {
		t.Error("Expected \"user9\", got ", creds.Username)
	}
}

func TestUpdateCredentials(t *testing.T) {
	db, dbPath := mockDb()
	defer cleanUp(path.Dir(dbPath))
	defer db.Close()

	// Confirm that the credentials start as lowercase.
	creds, err := db.GetCredentials("loginTitle0")
	if err != nil {
		t.Error("Expected nil, got an error")
	}
	if creds.Username != "user0" {
		t.Error("Expected \"user0\", got ", creds.Username)
	}
	if creds.Password != "password" {
		t.Error("Expected \"password\", got ", creds.Password)
	}

	// Update the \"loginTitle0\" credentials.
	newCreds := models.Credentials{
		Title:    "loginTitle0",
		Username: "USER0",
		Password: "PASSWORD",
	}
	if err := db.UpdateCredentials(newCreds); err != nil {
		t.Error("Expected nil, got an error")
	}

	// Retrieve and check the updated credentials.
	creds, err = db.GetCredentials("loginTitle0")
	if err != nil {
		t.Error("Expected nil, got an error")
	}

	if creds.Username != "USER0" {
		t.Error("Expected \"USER0\", got ", creds.Username)
	}
	if creds.Password != "PASSWORD" {
		t.Error("Expected \"PASSWORD\", got ", creds.Password)
	}
}

func TestDeleteCredentials(t *testing.T) {
	db, dbPath := mockDb()
	defer cleanUp(path.Dir(dbPath))
	defer db.Close()

	// Confirm that the credentials exist.
	creds, err := db.GetCredentials("loginTitle0")
	if err != nil {
		t.Error("Expected nil, got an error")
	}
	if creds.Username != "user0" {
		t.Error("Expected \"user0\", got ", creds.Username)
	}

	// Delete the credentials.
	if err := db.DeleteCredentials("loginTitle0"); err != nil {
		t.Error("Expected nil, got an error")
	}

	// Confirm that the credentials no longer exist.
	_, err = db.GetCredentials("loginTitle0")
	if err == nil {
		t.Error("Expected an error, got nil")
	}
}
