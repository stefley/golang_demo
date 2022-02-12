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
// 	type Company struct {
// 		ID   int
// 		Name string
// 	}

// 	type Worker struct {
// 		gorm.Model
// 		Name      string
// 		CompanyID int
// 		Company   Company
// 	}

// 	db.AutoMigrate(&Company{}, &Worker{})

// }

// func test2() {
// 	type Company struct {
// 		ID   int
// 		Name string
// 	}

// 	// 重写外键
// 	type Worker struct {
// 		gorm.Model
// 		Name         string
// 		CompanyRefer int
// 		Company      Company `gorm:"foreignKey:CompanyRefer"` //使用CompanyRefer作为外键
// 	}

// 	db.AutoMigrate(&Company{}, &Worker{})

// }
// func main() {
// 	test1()
// 	test2()
// }
