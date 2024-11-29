package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

var (
	dbInstance *sql.DB
	dbOnce     sync.Once
)

func GetDBInstance() *sql.DB {
	dbOnce.Do(func() {
		fmt.Println("Initializing database connection...")

		connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_SSL"))

		var err error
		dbInstance, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal("Error opening database connection:", err)
		}

		dbInstance.SetMaxOpenConns(25)
		dbInstance.SetMaxIdleConns(25)
		dbInstance.SetConnMaxLifetime(5 * time.Minute)

		if err = dbInstance.Ping(); err != nil {
			log.Fatal("Error connecting to the database:", err)
		}
	})
	return dbInstance
}
