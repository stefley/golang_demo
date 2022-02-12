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

// func initDB() {
// 	ctx := context.TODO()
// 	co := options.Client().ApplyURI("mongodb://localhost:27017")
// 	c, err := mongo.Connect(ctx, co)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err2 := c.Ping(ctx, nil)
// 	if err2 != nil {
// 		log.Fatal(err2)
// 	} else {
// 		fmt.Println("数据库连接成功！")
// 	}
// 	client = c
// }

// func update() {
// 	initDB()
// 	ctx := context.TODO()
// 	defer client.Disconnect(ctx)
// 	c := client.Database("golang_db").Collection("student")
// 	update := bson.D{{"$set", bson.D{{"name", "kite"}, {"age", 22}}}}
// 	ur, err := c.UpdateMany(ctx, bson.D{{"name", "kita"}}, update)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("ur.MatchedCount: %v\n", ur.MatchedCount)
// }
// func main() {
// 	update()
// }
