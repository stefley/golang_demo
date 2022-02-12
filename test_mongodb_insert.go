// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type Student struct {
// 	Name string
// 	Age  int
// }

// func insert() {
// 	initDB()
// 	s1 := Student{
// 		"tom",
// 		20,
// 	}
// 	collection := clinet.Database("golang_db").Collection("student")
// 	ior, err := collection.InsertOne(context.TODO(), s1)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Printf("ior.InsertedID: %v\n", ior.InsertedID)
// 	}
// }

// func insetMany() {
// 	initDB()

// 	c := clinet.Database("golang_db").Collection("student")
// 	s1 := Student{
// 		Name: "kita",
// 		Age:  10,
// 	}
// 	s2 := Student{
// 		"rose",
// 		12,
// 	}

// 	imr, err := c.InsertMany(context.TODO(), []interface{}{s1, s2})
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// 	fmt.Printf("imr.InsertedIDs: %v\n", imr.InsertedIDs)
// }

// var clinet *mongo.Client

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
// 		fmt.Println("数据库连接成功")
// 	}
// 	clinet = c
// }
// func main() {
// 	insert()
// 	insetMany()
// }
