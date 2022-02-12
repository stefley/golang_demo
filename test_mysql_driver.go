// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"time"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {
// 	db, err := sql.Open("mysql", "user:123456@/go_db")
// 	if err != nil {
// 		panic(err)
// 	}
// 	// See "Important settings" section.
// 	db.SetConnMaxLifetime(time.Minute * 3)
// 	db.SetMaxOpenConns(10)
// 	db.SetMaxIdleConns(10)

// 	fmt.Printf("db: %v\n", db)
// }
