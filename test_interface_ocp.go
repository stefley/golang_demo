// // ocp (Open-Close Principle) 开闭原则 对扩展开放对修改关闭
// package main

// import "fmt"

// type Pet interface {
// 	eat()
// 	sleep()
// }

// type Dog struct {
// }
// type Cat struct {
// }

// func (dog Dog) eat() {
// 	fmt.Println("dog eat")
// }
// func (dog Dog) sleep() {
// 	fmt.Println("dog sleep")
// }
// func (cat Cat) eat() {
// 	fmt.Println("cat eat")
// }
// func (cat Cat) sleep() {
// 	fmt.Println("cat sleep")
// }

// type Person struct {
// }

// func (per Person) care(pet Pet) {
// 	pet.eat()
// 	pet.sleep()
// }
// func main() {
// 	dog := Dog{}
// 	cat := Cat{}
// 	per := Person{}
// 	per.care(dog)
// 	per.care(cat)
// }
