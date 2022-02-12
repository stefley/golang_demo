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

// func find() {
// 	initDB()
// 	ctx := context.TODO()
// 	// 是否连接超时
// 	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	// defer cancel()
// 	// 关闭数据库连接
// 	defer client.Disconnect(ctx)
// 	c := client.Database("golang_db").Collection("student")
// 	cur, err := c.Find(ctx, bson.D{{"name", "tom"}})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer cur.Close(ctx)
// 	for cur.Next(ctx) {
// 		var result bson.D
// 		cur.Decode(&result)
// 		fmt.Printf("result: %v\n", result)
// 		fmt.Printf("result map: %v\n", result.Map())
// 	}

// }
// func initDB() {
// 	co := options.Client().ApplyURI("mongodb://localhost:27017")
// 	c, err := mongo.Connect(context.TODO(), co)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err2 := c.Ping(context.TODO(), nil)
// 	if err2 != nil {
// 		log.Fatal(err2)
// 	} else {
// 		fmt.Println("数据库连接成功！")
// 	}
// 	client = c
// }
// func main() {
// 	find()
// }
