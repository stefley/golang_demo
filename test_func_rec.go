// package main

// import "fmt"

// // 没有递归出口会导致内存溢出
// func f1() {
// 	f1()
// }

// // 阶乘
// func f2() {
// 	var s int = 1
// 	// 5! 5x4x3x2x1
// 	for i := 1; i <= 5; i++ {
// 		s *= i
// 	}
// 	fmt.Printf("s: %v\n", s)
// }

// // 递归实现阶乘
// func f3(n int) int {
// 	if n == 1 {
// 		return 1
// 	} else {
// 		return n * f3(n-1)
// 	}
// }

// // 斐波那契 f(n) = f(n-1) +f(n-2) 并且f(2) = f(1) = 1
// func fib(n int) int {
// 	if n == 2 || n == 1 {
// 		return 1
// 	}
// 	return fib(n-1) + fib(n-2)
// }
// func main() {
// 	f2() // 120
// 	r := f3(5)
// 	fmt.Printf("r: %v\n", r) // 120
// }
