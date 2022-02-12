// package main

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var db *gorm.DB

// type User struct {
// 	gorm.Model
// 	Name     string
// 	Age      uint
// 	Birthday time.Time
// }

// func select1() {
// 	var user User
// 	db.First(&user)
// 	fmt.Printf("user: %v\n", user)
// 	db.Take(&user)
// 	fmt.Printf("user: %v\n", user)
// 	db.Last(&user)
// 	fmt.Printf("user: %v\n", user)
// }

// func select2() {
// 	var user User
// 	db.First(&user, "3")
// 	// select * from users where id = 3
// 	db.First(&user, 3)
// 	// select * from users where id = 3
// 	var users []User
// 	db.Find(&users, []int{1, 2, 3})
// 	// select * from users where id in (1,2,3)
// 	for _, user := range users {
// 		fmt.Printf("user.ID: %v\n", user.ID)
// 	}
// }

// func init() {
// 	dsn := "root:123456@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
// 	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Panicf("database connect failed: %v\n", err)
// 	}
// 	db = d
// }

// func select3() {
// 	var users []User
// 	result := db.Find(&users)
// 	fmt.Printf("users: %v\n", users)
// 	fmt.Printf("result.RowsAffected: %v\n", result.RowsAffected)
// }
// func main() {
// 	select1()
// 	select2()
// 	select3()
// }
