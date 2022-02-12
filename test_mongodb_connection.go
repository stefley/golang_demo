// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// var client *mongo.Client

// func initDB() {
// 	co := options.Client().ApplyURI("mongodb://localhost:2701")
// 	c, err := mongo.Connect(context.TODO(), co)
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		fmt.Printf("c: %v\n", c)
// 	}
// 	err2 := c.Ping(context.TODO(), nil)
// 	if err2 != nil {
// 		log.Fatal(err2)
// 	} else {
// 		fmt.Println("数据库连接成功!")
// 	}
// }
// func main() {
// 	initDB()
// }
