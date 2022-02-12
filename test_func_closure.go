// 闭包
// package main

// import "fmt"

// func add() func(y int) int {
// 	var x int
// 	return func(y int) int {
// 		fmt.Printf("x: %v\n", x)
// 		x += y
// 		return x
// 	}
// }

// func calc(base int) (func(a int) int, func(a int) int) {
// 	add := func(a int) int {
// 		base += a
// 		return base
// 	}
// 	sub := func(a int) int {
// 		base -= a
// 		return base
// 	}
// 	return add, sub
// }
// func main() {
// 	f := add()
// 	r := f(10)
// 	fmt.Printf("r: %v\n", r)
// 	r = f(10)
// 	fmt.Printf("r: %v\n", r)
// 	fmt.Println("---------------------")

// 	f1 := add()
// 	r = f1(100)
// 	fmt.Printf("r: %v\n", r)
// 	r = f(10)
// 	fmt.Printf("r: %v\n", r)

// 	f1, f2 := calc(20)
// 	r1 := f1(5)
// 	fmt.Printf("r1: %v\n", r1)
// 	r2 := f2(5)
// 	fmt.Printf("r2: %v\n", r2)
// }
