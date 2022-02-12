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

// func test1() {
// 	type CreditCard struct {
// 		gorm.Model
// 		Number string
// 		UserID uint
// 	}
// 	type User struct {
// 		gorm.Model
// 		CreditCard CreditCard
// 	}

// 	db.AutoMigrate(&User{}, &CreditCard{})

// }

// type Toy struct {
// 	ID        int
// 	Name      string
// 	OwnerID   int
// 	OwnerType string
// }
// type Cat struct {
// 	ID   int
// 	Name string
// 	Toy  Toy `gorm:"polymorphic:Owner;"`
// }
// type Dog struct {
// 	ID   int
// 	Name string
// 	Toy  Toy `gorm:"polymorphic:Owner;"`
// }

// // 多态关联
// func test2() {
// 	db.AutoMigrate(&Toy{}, &Cat{}, &Dog{})
// }

// func test3() {
// 	// db.Create(&Dog{Name: "didi", Toy: Toy{Name: "toy1"}})
// 	db.Create(&Cat{Name: "pio", Toy: Toy{Name: "toy1"}})
// }
// func main() {
// 	test1()
// 	test2()
// 	test3()
// }
