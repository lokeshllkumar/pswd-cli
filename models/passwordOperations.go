package models

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type PasswordEntry struct {
	ID             int       `json:"id"`
	Service        string    `json:"service"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	TimeOfCreation time.Time `json:"createdAt"`
}

func (db *DB) AddPassword(service string, username string, encPassword string) error {
	insertQuery := `INSERT INTO local_passwords (service, username, password) VALUES (?, ?, ?);`
	_, err := db.Conn.Exec(insertQuery, service, username, encPassword)
	return err
}

func (db *DB) GetPassword(service string, username string) (string, error) {
	getQuery := `SELECT id, service, username, password, creation_time FROM local_passwords WHERE service = ? AND username = ?;`
	rows, err := db.Conn.Query(getQuery, service, username)
	if err != nil {
		fmt.Printf("Error running SELECT query in database\n")
		return "", err
	}
	defer rows.Close()

	var res []PasswordEntry

	for rows.Next() {
		var record PasswordEntry
		var creationTime string

		if err := rows.Scan(&record.ID, &record.Service, &record.Username, &record.Password, &creationTime); err != nil {
			fmt.Printf("Error reading data retrieved from database\n")
			return "", err
		}

		/*
			record.TimeOfCreation, err = time.Parse("2006-01-02 15:04:08", creationTime)
			if err != nil {
				fmt.Println("Error, could not parse data storing time of entry")
				return "", err
			}

		*/

		res = append(res, record)
	}

	jsonRes, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("Error parsing JSON data\n")
		return "", err
	}

	return string(jsonRes), nil
}

func (db *DB) GetPasswords(service string) (string, error) {
	getQuery := `SELECT id, service, username, password, creation_time FROM local_passwords WHERE service = ?;`
	rows, err := db.Conn.Query(getQuery, service)
	if err != nil {
		fmt.Printf("Error running SELECT query in database\n")
		return "", err
	}
	defer rows.Close()

	var res []PasswordEntry

	for rows.Next() {
		var record PasswordEntry
		var creationTime string

		if err := rows.Scan(&record.ID, &record.Service, &record.Username, &record.Password, &creationTime); err != nil {
			fmt.Printf("Error reading data retrieved from database\n")
			return "", err
		}

		/*
			record.TimeOfCreation, err = time.Parse("2006-01-02 15:04:08", creationTime)
			if err != nil {
				fmt.Println("Error, could not parse data storing time of entry")
				return "", err
			}

		*/

		res = append(res, record)
	}

	jsonRes, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("Error parsing JSON data\n")
		return "", err
	}

	return string(jsonRes), nil
}

func (db *DB) ListPasswords() (string, error) {
	getQuery := `SELECT id, service, username, password, creation_time FROM local_passwords;`
	rows, err := db.Conn.Query(getQuery)
	if err != nil {
		fmt.Printf("Error running SELECT query in database\n")
		return "", err
	}
	defer rows.Close()

	var res []PasswordEntry
	for rows.Next() {
		var record PasswordEntry
		var creationTime string

		if err := rows.Scan(&record.ID, &record.Service, &record.Username, &record.Password, &creationTime); err != nil {
			fmt.Printf("Error reading data retrieved from database\n")
			return "", err
		}

		/*
			record.TimeOfCreation, err = time.Parse("2006-01-02 15:04:08", creationTime)
			if err != nil {
				fmt.Println("Error, could not parse data storing time of entry")
				return "", err
			}
		*/

		res = append(res, record)
	}

	jsonRes, err := json.Marshal(res)
	if err != nil {
		return "", err
	}

	return string(jsonRes), nil
}

func (db *DB) DeletePassword(service string, username string) error { // for a single entry belonging to a service
	deleteQuery := `DELETE FROM local_passwords WHERE service = ? AND username = ?;`
	_, err := db.Conn.Exec(deleteQuery, service, username)
	return err
}

func (db *DB) DeletePasswords(service string) error { // for all entries belonging to a service
	deleteQuery := `DELETE FROM local_passwords WHERE service = ?;`
	_, err := db.Conn.Exec(deleteQuery, service)
	return err
}

func (db *DB) UpdatePassword(service string, username string, newPassword string) error {
	updateQuery := `UPDATE local_passwords SET password = ? WHERE service = ? AND username = ?;`
	_, err := db.Conn.Exec(updateQuery, newPassword, service, username)
	return err
}
