// package main

// import "fmt"

// func main() {
// 	type Dog struct {
// 		name  string
// 		age   int
// 		color string
// 	}
// 	type Person struct {
// 		dog  Dog
// 		name string
// 		age  int
// 	}

// 	dog := Dog{
// 		name:  "poli",
// 		age:   2,
// 		color: "black",
// 	}
// 	per := Person{
// 		dog:  dog,
// 		name: "dav",
// 		age:  20,
// 	}
// 	fmt.Printf("per: %v\n", per)
// 	fmt.Printf("per.dog.name: %v\n", per.dog.name)
// }
