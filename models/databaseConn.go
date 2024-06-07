package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	Conn *sql.DB
}

func InitDB() (*DB, error) {
	db, err := sql.Open("sqlite3", "./database/pswd_cli_local.db")
	if err != nil {
		fmt.Printf("Error connecting to database\n")
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Printf("Error pinging database connection object\n")
		return nil, err
	}

	dbObj := DB{Conn: db}

	dbObj.CreateTable()

	return &dbObj, nil
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
