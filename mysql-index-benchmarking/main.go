package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }
    dsn := os.Getenv("MYSQL_DSN")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Connection failed with %v", err)
	} 
	if err := db.Ping(); err != nil {
        log.Fatalf("Error connecting to database: %v\n", err)
    }
    fmt.Println("Successfully connected to MySQL!")

	// Example of running a query
    rows, err := db.Query("SELECT * FROM mydate")
    if err != nil {
        log.Fatalf("Query error: %v\n", err)
    }
    defer rows.Close()

    // Fetch results
    for rows.Next() {
        var dt time.Time
        var ts time.Time
        if err := rows.Scan(&dt, &ts); err != nil {
            log.Fatalf("Row scan error: %v\n", err)
        }
        fmt.Printf("\nPRINTING dt: %v, ts: %v\n", dt, ts)
    }
	
	defer db.Close()
}
	