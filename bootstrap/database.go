package bootstrap

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func NewPostgresDatabase(env *Env) *sql.DB {
    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        env.DBHost, env.DBPort, env.DBUser, env.DBPass, env.DBName,
    )

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	maxRetries := 10
	retryInterval := time.Second * 3

	for i := 0; i < maxRetries; i++ {
		err := db.Ping()
		if err == nil {
			log.Println("Successfully connected to PostgreSQL database")
			return db
		}
		
		log.Printf("Failed to ping Postgres (attempt %d/%d): %v", i+1, maxRetries, err)
		
		if i < maxRetries-1 {
			log.Printf("Retrying in %v...", retryInterval)
			time.Sleep(retryInterval)
		}
	}
	
	log.Fatal("Failed to connect to database after multiple attempts. Giving up.")
	return nil 
}

func ClosePostgresDB(db *sql.DB){
	if db == nil {
		return
	}
	if err := db.Close(); err != nil {
		log.Println("Error closing Postgres connection:", err)
	}

	log.Println("Connection to Postgres closed.")
}