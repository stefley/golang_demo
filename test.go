// package main

// import "fmt"

// func main() {
// 	// var (
// 	// 	name string
// 	// 	age  int
// 	// )
// 	// name = "zyq"
// 	// age = 3
// 	// var b bool = true
// 	// var msg = "hello go"
// 	// var name, age, b, msg = "zyq", 3, true, "hello go"
// 	name := "zyq"
// 	age := 3
// 	b := true
// 	msg := "hello go"
// 	_, age2 := getNameAndAge()
// 	// const PI float64 = 3.14
// 	// const PI2  = 3.14
// 	// const (
// 	// 	PI3 = 3.14
// 	// 	PI4 = 3.14
// 	// )
// 	// const PI5,PI6 = 3.14,3.14
// 	// const (
// 	// 	a1 = iota // 0
// 	// 	a2 = iota // 1
// 	// 	a3 = iota // 2
// 	// )
// 	// const (
// 	// 	a1 = iota // 0
// 	// 	_ // 忽略
// 	// 	a2 = iota // 2
// 	// 	a3 = iota //3
// 	// )
// 	const (
// 		a1 = iota // 0
// 		a2 = 100  // 100 中间插队
// 		a3 = iota // 2
// 	)
// 	fmt.Printf("name: %v\n", name)
// 	fmt.Printf("age: %v\n", age)
// 	fmt.Printf("b: %v\n", b)
// 	fmt.Printf("msg: %v\n", msg)
// 	fmt.Printf("age2: %v\n", age2)
// 	fmt.Printf("a1: %v\n", a1)
// 	fmt.Printf("a2: %v\n", a2)
// 	fmt.Printf("a3: %v\n", a3)
// }

// func getNameAndAge() (string, int) {
// 	return "zyq", 3
// }
