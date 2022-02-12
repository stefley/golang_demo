// package main

// import "fmt"

// func test1() {
// 	name := "tom"
// 	age := "20"

// 	f1 := func() string {
// 		return name + age
// 	}
// 	msg := f1()
// 	fmt.Printf("msg: %v\n", msg)
// }

// func main() {
// 	max := func(a int, b int) int {
// 		if a > b {
// 			return a
// 		} else {
// 			return b
// 		}
// 	}
// 	r := max(1, 5)
// 	fmt.Printf("r: %v\n", r)

// 	匿名函数自执行
// 	r := func(a int, b int) int {
// 		if a > b {
// 			return a
// 		} else {
// 			return b
// 		}
// 	}(2, 4)
// 	fmt.Printf("r: %v\n", r)
// 	test1()
// }
