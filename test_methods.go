// package main

// import "fmt"

// type Person struct {
// 	name string
// 	age  int
// }

// // 属性和方法分开来写
// // (per Person)接收者 receiver
// func (per Person) eat() {
// 	fmt.Printf("%v,eat...\n", per.name)
// }
// func (per Person) sleep() {
// 	fmt.Printf("%v,sleep.....\n", per.name)
// }

// type Customer struct {
// 	name string
// }

// func (customer Customer) login(name string, pwd string) bool {
// 	fmt.Printf("customer.name: %v\n", customer.name)
// 	if name == "tom" && pwd == "123" {
// 		return true
// 	} else {
// 		return false
// 	}
// }
// func main() {
// 	per := Person{
// 		name: "Jack",
// 		age:  20,
// 	}
// 	per.eat()
// 	per.sleep()
// 	customer := Customer{
// 		name: "tom",
// 	}
// 	r := customer.login(customer.name, "123")
// 	fmt.Printf("r: %v\n", r)
// }
