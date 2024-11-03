package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	count = 1000000
)

func readFromDateTable(db *sql.DB) {
	//query
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
}

func insertIntoTable(db *sql.DB, table string) (time.Duration, error) {
    start := time.Now()
	query := fmt.Sprintf("INSERT INTO %s (ID, Age) VALUES (?, ?)", table)
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return time.Since(start), err
	}

    for i := 1; i <= count; i++ {
		if strings.Contains(table, "1") {
			_, err = stmt.Exec(i+1, i+1)
		} else {
			_, err = stmt.Exec(strconv.Itoa(i+1), i+1)
		}
		
        if err != nil {
            log.Fatal(err)
            return time.Since(start), err
        }
    }
	stmt.Close()
    return time.Since(start), nil
}

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
	defer db.Close()

	// readFromDateTable(db)

	time1, err1 := insertIntoTable(db, "test_table1")
	if err1 == nil {
		fmt.Printf("Inserted %v records in %v\n", count, time1)
	}

	time2, err2 := insertIntoTable(db, "test_table2")
	if err2 == nil {
		fmt.Printf("Inserted %v records in %v\n", count, time2)
	}
}
	