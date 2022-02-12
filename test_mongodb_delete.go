// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// var client *mongo.Client
// var ctx = context.TODO()

// func initDB() {
// 	co := options.Client().ApplyURI("mongodb://localhost:27017")
// 	c, err := mongo.Connect(ctx, co)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err2 := c.Ping(ctx, nil)
// 	if err2 != nil {
// 		log.Fatal(err2)
// 	} else {
// 		fmt.Println("数据库连接成功!")
// 	}
// 	client = c
// }
// func delete() {
// 	initDB()
// 	defer client.Disconnect(ctx)
// 	c := client.Database("golang_db").Collection("student")
// 	dr, err := c.DeleteMany(ctx, bson.D{{"age", 20}})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("dr.DeletedCount: %v\n", dr.DeletedCount)
// }
// func main() {
// 	delete()
// }
