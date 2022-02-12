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
// 	Languages []Language `gorm:"many2many:user_languages;"`
// }
// type Language struct {
// 	gorm.Model
// 	Name string
// }

// func createTable() {
// 	db.AutoMigrate(&User{}, &Language{})
// }

// func insert() {
// 	l := Language{
// 		Name: "jp",
// 	}
// 	l2 := Language{
// 		Name: "zh",
// 	}
// 	user := User{
// 		Languages: []Language{l, l2},
// 	}
// 	db.Create(&l)
// 	db.Create(&user)
// }
// func main() {
// 	createTable()
// 	insert()
// }
