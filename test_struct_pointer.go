// package main

// import (
// 	"fmt"
// )

// func test1() {
// 	var name string
// 	name = "tom"
// 	var p_name *string
// 	p_name = &name
// 	fmt.Printf("p_name: %v\n", p_name)
// 	fmt.Printf("p_name: %v\n", *p_name)
// }
// func test2() {
// 	type Person struct {
// 		id   int
// 		name string
// 		age  int
// 	}
// 	var p_person *Person
// 	tom := Person{
// 		id:   101,
// 		name: "tom",
// 		age:  18,
// 	}
// 	p_person = &tom
// 	fmt.Printf("p_person: %p\n", p_person)
// 	fmt.Printf("p_person: %v\n", *p_person)
// }

// 使用new关键字创建结构体指针
// func test3() {
// 	type Person struct {
// 		id, age int
// 		name    string
// 	}
// 	var tom = new(Person)
// 	// (*tom).id = 101
// 	tom.id = 101
// 	tom.age = 19
// 	tom.name = "tom"
// 	fmt.Printf("tom: %p\n", tom)
// 	fmt.Printf("tom.name: %v\n", tom.name)
// 	fmt.Printf("tom: %v\n", *tom)
// }
// func main() {
// 	test3()
// }
