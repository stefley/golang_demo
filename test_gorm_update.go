// package main

// import (
// 	"log"
// 	"time"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type User struct {
// 	gorm.Model
// 	Name     string
// 	Age      int
// 	Birthday time.Time
// 	Active   bool
// }

// func createTable() {
// 	db.AutoMigrate(&User{})
// }

// func insert() {
// 	user := User{
// 		Name:     "tom",
// 		Age:      20,
// 		Birthday: time.Now(),
// 		Active:   true,
// 	}
// 	db.Create(&user)
// }

// func update() {
// 	var user User
// 	db.Find(&user)

// 	user.Name = "big tom"
// 	user.Age = 30
// 	db.Save(&user)
// }

// func update2() {
// 	db.Model(&User{}).Where("active =?", true).Update("name", "jack")
// 	// UPDATE users SET name="jack",updated_at='2022-2-10 16:00:00' WHERE active=true;"
// }

// func update3() {
// 	var user User
// 	db.First(&user)
// 	db.Model(&user).Updates(User{Name: "toto", Age: 12, Active: false})
// 	// UPDATE users SET name='toto',age=12,updated_at='2022-2-10 16:00:00' WHERE id = 1;
// }

// func update4() {
// 	var user User
// 	db.First(&user)
// 	db.Model(&user).Updates(map[string]interface{}{"name": "tiny", "age": 3, "active": false})
// }

// func main() {
// 	createTable()
// 	insert()
// 	update()
// 	update2()
// 	update3()
// 	update4()
// }

// var db *gorm.DB

// func init() {
// 	dsn := "root:123456@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
// 	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Panicf("database connect failed: %v\n", err)
// 	}
// 	db = d
// }
