package database

import (
  "errors"
  "passfish/internal/passwords"
)

func (db *Db) createPassphraseTable() error {
  sqlStmt := `
  create table if not exists passphrase (
    passphrase blob not null
  );
  `
  _, err := db.Conn.Exec(sqlStmt)
  return err
}

func (db *Db) SetPassphrase(passphrase string) error {
  // Encrypt the passphrase and insert it into the database.
  if db.GetPassphrase() != "" {
    return errors.New("a passphrase already exists in the database")
  }

  sqlStmt := `
  insert into passphrase (passphrase) values (?);
  `

  e_passphrase := passwords.Encrypt(passphrase, passphrase)
  _, err := db.Conn.Exec(sqlStmt, e_passphrase)
  return err
}

func (db *Db) VerifyPassphrase(passphrase string) bool {
  // Check if the argument passphrase correctly decrypts the passphrase stored
  // in the database.
  if db.GetPassphrase() == "" {
    return false
  }

  // If the passphrase does not decrypt the stored passphrase, return false.
  if passwords.Decrypt(db.GetPassphrase(), passphrase) != passphrase {
    return false
  }

  return true
}

func (db *Db) GetPassphrase() string {
  sqlStmt := `
  select passphrase from passphrase;
  `
  row := db.Conn.QueryRow(sqlStmt)

  var e_passphrase string
  row.Scan(&e_passphrase)
  return e_passphrase
}

func (db *Db) DeletePassphrase() error {
  sqlStmt := `
  delete from passphrase;
  `
  _, err := db.Conn.Exec(sqlStmt)
  return err
}
