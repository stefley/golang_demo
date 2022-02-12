// 结构体当作参数传递时
// 直接传递结构体，这是一个副本（拷贝），在函数内部不会改变外部结构体内容
// 传递结构体指针，函数内部可以改变结构体内容
// package main

// import "fmt"

// type Person struct {
// 	id  int
// 	age int
// }

// func showPerson(person Person) {
// 	person.id = 101
// }

// func showPerson2(person *Person) {
// 	person.id = 101
// }
// func main() {
// 	person := Person{
// 		id:  10,
// 		age: 18,
// 	}
// 	fmt.Printf("person: %v\n", person) // person: {10 18}
// 	fmt.Println("------------")
// 	showPerson(person)
// 	fmt.Printf("person: %v\n", person) // person: {10 18}
// 	fmt.Println("----------------------")
// 	showPerson2(&person)
// 	fmt.Printf("person: %v\n", person)
// }
