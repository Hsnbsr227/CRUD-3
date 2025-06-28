package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnDb() {
	
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: .env file not found, using system environment variables")
	}

	
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	ConnStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	var dbErr error
	db, dbErr = sql.Open("postgres", ConnStr)
	if dbErr != nil {
		log.Fatal("Connection Error: ", dbErr)
	}

	dbErr = db.Ping()
	if dbErr != nil {
		log.Fatal("Connection Fail: ", dbErr)
	}

	fmt.Print("Successfully Connected")
}

// Transaction başlatma fonksiyonu
func BeginTransaction() (*sql.Tx, error) {
	return db.Begin()
}

// Transaction'ı commit etme fonksiyonu
func CommitTransaction(tx *sql.Tx) error {
	return tx.Commit()
}

// Transaction'ı rollback etme fonksiyonu
func RollbackTransaction(tx *sql.Tx) error {
	return tx.Rollback()
}
