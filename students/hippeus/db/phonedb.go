package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

// Database is a wrapper for a SQL-like database
type Database struct {
	db *sql.DB
}

// PhoneRecord represents phone numbers in a database
type PhoneRecord struct {
	UID    int
	Number string
}

// Open connects to SQL-like database using driverName.
// As establishing connection can fail, function makes multiple attempts before reporting an error.
func Open(driverName, dataSrc string) (*Database, error) {
	db, err := sql.Open(driverName, dataSrc)
	if err != nil {
		return nil, fmt.Errorf("could not connect to db: %v", err)
	}
	maxAtt := 10
	for i := 0; i < maxAtt; i++ {
		if err = db.Ping(); err != nil {
			if i == maxAtt {
				return nil, fmt.Errorf("could not connect to db: %v", err)
			}
			log.Printf("#%d Failed to connect, next attemp in 200ms\n", i)
			time.Sleep(200 * time.Millisecond)
		} else {
			break
		}
	}
	return &Database{db}, nil
}

// Close close connection with an external Database.
func (db *Database) Close() error {
	return db.db.Close()
}

// AllPhones retrieves all phone records from a database.
func (db *Database) AllPhones() ([]PhoneRecord, error) {

	rows, err := db.db.Query(`select * from phone_numbers`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var phones []PhoneRecord
	var ph PhoneRecord
	for rows.Next() {
		if err = rows.Scan(&ph.UID, &ph.Number); err != nil {
			log.Printf("Database read failure: %s. Skipping row.\n", err)
			continue
		}
		phones = append(phones, ph)
	}
	return phones, rows.Err()
}

// FindPhone look up database in order to find matching record.
// In case of more than match, only the first occurence is returned.
func (db *Database) FindPhone(number string) (PhoneRecord, error) {
	var ph PhoneRecord
	row := db.db.QueryRow("select id,phone from phone_numbers where phone=$1;", number)
	if err := row.Scan(&ph.UID, &ph.Number); err != nil {
		switch err {
		case sql.ErrNoRows:
			return PhoneRecord{}, sql.ErrNoRows
		default:
			return PhoneRecord{}, err
		}
	}
	return ph, nil
}

// UpdatePhone updates record in a database for a given id.
func (db *Database) UpdatePhone(id int, newNumber string) error {
	if _, err := db.db.Exec("update phone_numbers set id=$1, phone=$2 where id=$1", id, newNumber); err != nil {
		log.Printf("Update of a record %d failed: %v\n", id, err)
		return err
	}
	return nil
}

// DeletePhone permanently removes record from a database.
func (db *Database) DeletePhone(id int) error {
	if _, err := db.db.Exec("delete from phone_numbers where id=$1;", id); err != nil {
		log.Printf("Delition of an record %d failed: %v\n", id, err)
		return err
	}
	return nil
}
