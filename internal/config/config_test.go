package config_test

import (
	"fmt"
	"os"
	"passfish/internal/config"
	"path"
	"testing"
)

func TestNewConfig(t *testing.T) {
	// Create a temporary directory
	dir, err := os.MkdirTemp(".", "t_*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	// Create a temporary file
	file, err := os.Create(path.Join(dir, "config.yaml"))
	if err != nil {
		t.Fatal(err)
	}

	dbPath := path.Join(dir, "test_db.sql")
	// Write the YAML content to the file
	_, err = file.WriteString(fmt.Sprintf("\ndbPath: %s\n", dbPath))
	if err != nil {
		t.Fatal(err)
	}
	file.Close()

	// Create a new Config struct
	cfg, err := config.New(file.Name())
	if err != nil {
		t.Error("Expected nil, got an error")
	}

	// Check the values of the Config struct
	if cfg.DbPath != dbPath {
		t.Errorf("Expected \"DbPath\" to be %s, got %s", dbPath, cfg.DbPath)
	}
}

func TestNewConfigNoDbPath(t *testing.T) {
	// Create a temporary directory
	dir, err := os.MkdirTemp(".", "t_*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	// Create a temporary file
	file, err := os.Create(path.Join(dir, "config.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	file.Close()

	cfg, err := config.New(file.Name())
	if cfg != nil {
		t.Errorf("Expected nil, got %v", cfg)
	}
	if err == nil {
		t.Error("Expected an error, got nil")
	}
}

func TestNewConfigFileNotExist(t *testing.T) {
	_, err := config.New("nonexistent.yaml")
	if err == nil {
		t.Error("Expected an error, got nil")
	}
}
