// package main

// import "fmt"

// func test1() {
// 	fmt.Println("我既没有参数也没有返回")
// }

// func sum(a int, b int) (ret int) {
// 	ret = a + b
// 	return ret
// }
// func sum2(a int, b int) int {
// 	// ret := a + b
// 	// return ret
// 	return a + b
// }

// // 多个返回值
// func test2() (name string, age int) {
// 	name = "tom"
// 	age = 20
// 	// return name, age
// 	return // 等价于 return name,age
// }

// // return覆盖返回
// func test3() (name string, age int) {
// 	n := "name"
// 	a := 10
// 	return n, a
// }
// func test4() (string, int) {
// 	n := "t"
// 	a := 10
// 	return n, a
// }
// func main() {
// test1()
// n, a := test2()
// fmt.Printf("n: %v, a: %v\n", n, a)
// name, age := test3()
// fmt.Printf("name: %v, age: %v\n", name, age)
// _, age := test4()
// fmt.Printf("age: %v\n", age)
// }
