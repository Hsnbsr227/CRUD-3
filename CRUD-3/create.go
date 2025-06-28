package main

import (
	"fmt"
	"log"
)

func CreateUser(name string, password string) error {
	// Transaction başlat
	tx, err := BeginTransaction()
	if err != nil {
		log.Printf("Failed to start transaction: %v", err)
		return err
	}

	// Hata durumunda rollback yapmak için defer kullan
	defer func() {
		if err != nil {
			RollbackTransaction(tx)
		}
	}()

	query := "INSERT INTO users (username,password) VALUES ($1,$2)"
	_, err = tx.Exec(query, name, password)
	if err != nil {
		log.Printf("Failed to add user: %v", err)
		return err
	}

	// Transaction'ı commit et
	err = CommitTransaction(tx)
	if err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return err
	}

	fmt.Printf("User Added: %s\n", name)
	return nil
}
