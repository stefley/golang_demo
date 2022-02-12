// // gorm  orm(object relational mapping 对象关系映射)
// package main

// import (
// 	"fmt"
// 	"log"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type Product struct {
// 	gorm.Model
// 	Code  string
// 	Price uint
// }

// func create(db *gorm.DB) {
// 	// 创建表
// 	db.AutoMigrate(&Product{})
// }

// func inset(db *gorm.DB) {
// 	// 插入数据
// 	prod := Product{
// 		Code:  "1001",
// 		Price: 100,
// 	}
// 	db.Create(&prod)
// }

// func query(db *gorm.DB) {
// 	var p Product
// 	// 查询
// 	db.First(&p, 1) // 以主键为1查询
// 	fmt.Printf("p: %+v\n", p)
// 	db.First(&p, "code =?", "1001") // 条件查询，查询code为1001
// 	fmt.Printf("p: %v\n", p)
// 	fmt.Printf("p.Price: %v\n", p.Price)
// }

// func update(db *gorm.DB) {
// 	// 更新
// 	var p Product
// 	db.First(&p, 1)
// 	db.Model(&p).Update("Price", 200)
// 	fmt.Printf("p: %v\n", p)
// 	// update多个字段
// 	// db.Model(&p).Updates(Product{Price: 1001, Code: "1002"}) // 仅更新非0字段
// 	db.Model(&p).Updates(map[string]interface{}{"Price": 1003, "Code": "1003"})
// }

// func del(db *gorm.DB) {
// 	// 删除
// 	var p Product
// 	db.First(&p, 1)
// 	db.Delete(&p, 1)
// }
// func main() {
// 	dsn := "root:123456@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Panic("数据库连接失败：%v\n", err)
// 	}
// 	create(db)
// 	inset(db)
// 	query(db)
// 	update(db)
// 	del(db)
// }
