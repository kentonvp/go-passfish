package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DbPath       string `yaml:"dbPath"`
	DbPassphrase string `yaml:"dbPassphrase"`
}

type MissingFieldError struct {
	Field string
}

func (e *MissingFieldError) Error() string {
	return fmt.Sprintf("missing required field: \"%s\"", e.Field)
}

func verifyConfig(cfg *Config) error {
	if cfg.DbPath == "" {
		return &MissingFieldError{Field: "dbPath"}
	}

	if cfg.DbPassphrase == "" {
		return &MissingFieldError{Field: "dbPassphrase"}
	}

	return nil
}

func NewConfig(cfgPath string) (*Config, error) {
	var cfg Config
	reader, err := os.Open(cfgPath)
	if err != nil {
		return nil, err
	}

	if err := yaml.NewDecoder(reader).Decode(&cfg); err != nil {
		return nil, err
	}

	if err = verifyConfig(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
