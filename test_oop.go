// package main

// import "fmt"

// type Person struct {
// 	Name string // 公共属性首字母大写，其他包可以访问
// 	age  int
// }

// func (per Person) eat() {
// 	fmt.Println("eat...")
// }
// func (per Person) sleep() {
// 	fmt.Println("sleep....")
// }
// func (per Person) work() {
// 	fmt.Println("work....")
// }
// func main() {
// 	per := Person{
// 		Name: "tom",
// 		age:  18,
// 	}
// 	per.eat()
// 	per.sleep()
// 	per.work()
// }
