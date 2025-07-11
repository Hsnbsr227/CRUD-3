package main

import (
	"fmt"
	"log"
)

func UpdateUser(id int, newName string, newPassword string) error {
	
	tx, err := BeginTransaction()
	if err != nil {
		log.Printf("Failed to start transaction: %v", err)
		return err
	}

	
	defer func() {
		if err != nil {
			RollbackTransaction(tx)
		}
	}()

	query := "UPDATE users SET username=$1, password =$2 WHERE id=$3"
	_, err = tx.Exec(query, newName, newPassword, id)
	if err != nil {
		log.Printf("Failed to update user: %v", err)
		return err
	}

	
	err = CommitTransaction(tx)
	if err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return err
	}

	fmt.Printf("User (ID: %d) successfully updated.\n", id)
	return nil
}
