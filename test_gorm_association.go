// package main

// import (
// 	"fmt"
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

// type Email struct {
// 	gorm.Model
// 	Email  string
// 	UserID uint
// }

// type User struct {
// 	gorm.Model
// 	Name            string
// 	BillingAddress  Address
// 	ShippingAddress Address
// 	Emails          []Email
// 	Languages       []Language `gorm:"many2many:user_languages;"`
// }

// type Address struct {
// 	gorm.Model
// 	Address1 string
// 	UserID   uint
// }

// type Language struct {
// 	gorm.Model
// 	Name string
// }

// func createTable() {
// 	db.AutoMigrate(&User{}, &Email{}, &Address{}, &Language{})
// }

// func test1() {
// 	user := User{
// 		Name:            "kak",
// 		BillingAddress:  Address{Address1: "b-address"},
// 		ShippingAddress: Address{Address1: "s-address"},
// 		Emails: []Email{
// 			{Email: "11.com"},
// 			{Email: "22.com"},
// 		},
// 		Languages: []Language{
// 			{Name: "ZH"},
// 			{Name: "EN"},
// 		},
// 	}
// 	// db.Create(&user)
// 	// 非创建而是更新时，更新关联数据
// 	// db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)

// 	//只创建name  insert info "users" (name) values ("kak", 1,2)
// 	// db.Select("Name").Create(&user)

// 	// skip create BillingAddress when creating a user
// 	// db.Omit("BillingAddress").Create(&user)

// 	// skip all associations when creating a user
// 	// db.Omit(clauses.associations).Create(&user)

// 	// 关联模式
// 	db.Model(&user).Association("Lanaguages")
// 	/*
// 		user是原模型，他的主键不能为空
// 		关系的字段名是Lanagus
// 		如果匹配了上面2个要求会开始关联模式，否则返错误
// 	*/

// }

// func test2() {
// 	var user User
// 	var languages []Language
// 	db.First(&user)
// 	db.Model(&user).Association("Languages").Find(&languages)
// 	fmt.Printf("languages: %v\n", languages)
// }

// func test3() {
// 	var user User
// 	db.First(&user)
// 	db.Model(&user).Association("Languages").Count()
// }
// func main() {
// 	createTable()
// 	test1()
// 	test2()
// }
