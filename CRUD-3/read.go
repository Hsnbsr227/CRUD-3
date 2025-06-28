package main

import (
	"fmt"
	"log"
)

func ReadUser() {
	query := "SELECT id,username,password FROM users"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Fail: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var username string
		var password string

		err := rows.Scan(&id, &username, &password)
		if err != nil {
			log.Fatal("Fail: ", err)
		}

		fmt.Printf("ID: %d | Username: %s | Password: %s\n", id, username, password)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal("Rows Fail!: ", err)
	}
}
