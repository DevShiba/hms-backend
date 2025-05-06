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

		if err := db.Ping(); err != nil{
        log.Fatal("Failed to ping Postgres:", err)
		}

		return db
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