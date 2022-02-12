// // 通过结构体嵌套模拟继承
// package main

// import "fmt"

// type Animal struct {
// 	name string
// 	age  int
// }

// func (ani Animal) eat() {
// 	fmt.Printf("%v, eat....\n", ani.name)
// }
// func (ani Animal) sleep() {
// 	fmt.Printf("%v, sleep...\n", ani.name)
// }

// type Dog struct {
// 	ani   Animal
// 	color string
// }
// type Cat struct {
// 	Animal
// 	color string
// }

// func main() {
// 	dog := Dog{
// 		ani: Animal{
// 			name: "dd",
// 			age:  1,
// 		},
// 		color: "black",
// 	}
// 	cat := Cat{
// 		Animal{
// 			name: "cc",
// 			age:  3,
// 		},
// 		"蓝色",
// 	}
// 	dog.ani.eat()
// 	dog.ani.sleep()
// 	cat.eat()
// 	cat.sleep()
// }
