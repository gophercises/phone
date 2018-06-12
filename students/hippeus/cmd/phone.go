package main

import (
	"fmt"
	"log"

	phone "github.com/gophercises/phone/students/hippeus"
	"github.com/gophercises/phone/students/hippeus/db"
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "admin"
	passwd = "admin123"
	dbname = "phone"
)

func main() {

	conn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, passwd, host, port, dbname)
	db, err := db.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	numbers := make(map[string]int)
	records, err := db.AllPhones()
	if err != nil {
		log.Fatal(err)
	}
	for _, n := range records {
		if normnum := phone.Normalize(n.Number); normnum != n.Number {
			err = db.UpdatePhone(n.UID, normnum)
			n.Number = normnum
			if err != nil {
				log.Fatal(err)
			}
		}
		if _, ok := numbers[n.Number]; !ok {
			numbers[n.Number] = n.UID
		} else {
			if err = db.DeletePhone(n.UID); err != nil {
				log.Println(err)
			}
		}
	}
}
