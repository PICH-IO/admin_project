package database

import (
	"log"
	"os"
	"sync"

	"github.com/jmoiron/sqlx"
)

var (
	once sync.Once
	db   *sqlx.DB
)

func initializeDB() {

	DATABASE_URL := os.Getenv("DATABASE_URL")
	var errDB error
	db, errDB = sqlx.Connect("postgres", DATABASE_URL)
	if errDB != nil {
		log.Fatalln("Error connecting to the database:", errDB)
	}

	// Verify the connection to the database is still alive
	if err := db.Ping(); err != nil {
		defer db.Close()
		log.Fatalf("Failed to ping the database: %v", err)
	}
	// Set connection pool settings
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(0)
}

func GetDB() *sqlx.DB {
	// fmt.Println("db_connect_success")
	once.Do(initializeDB)
	return db
}
