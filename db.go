package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"
	"github.com/go-sql-driver/mysql"
)

// Connect to db
func ConnectDB() (*sql.DB, error) {
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        user, password, host, port, dbname,
    )
	
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    // Validate the connection is alive
    if err = db.Ping(); err != nil {
        return nil, err
    }

    log.Println("Successfully connected to MySQL database!")
    return db, nil
}
