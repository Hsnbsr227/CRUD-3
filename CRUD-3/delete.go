package main

import "log"

func DeleteUser(id int) error {
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

	query := "DELETE FROM users WHERE id=$1"
	_, err = tx.Exec(query, id)
	if err != nil {
		log.Printf("Failed to delete user: %v", err)
		return err
	}

	// Transaction'ı commit et
	err = CommitTransaction(tx)
	if err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return err
	}

	log.Printf("User (ID: %d) successfully deleted.\n", id)
	return nil
}
