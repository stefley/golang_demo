// package main

// import (
// 	"log"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var db *gorm.DB

// func init() {
// 	dsn := "root:123456@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
// 	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	db = d
// }

// type User struct {
// 	gorm.Model
// 	Name string
// }

// func test1() {
// 	db.Session(&gorm.Session{DryRun: true}).AutoMigrate(&User{})
// }
// func main() {

// }
