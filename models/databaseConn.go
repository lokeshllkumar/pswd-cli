package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	Conn *sql.DB
}

func InitDB() (*DB, error){
	db, err := sql.Open("sqlite3", "./database/pswd_cli_local.db")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DB{Conn: db}, nil
}

func (db *DB) CloseDB() error {
	return db.Conn.Close()
}

func (db *DB) CreateTable() error {
	creationCmd := `CREATE TABLE IF NOT EXISTS local_passwords (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		service TEXT DEFAULT "login" NOT NULL,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		creation_time DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Conn.Exec(creationCmd)
	return err
}