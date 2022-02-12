// package main

// import "fmt"

// // 参数必须指定类型
// func sum(a int, b int) int {
// 	return a + b
// }

// func f1(a int) {
// 	// 参数传递的是一个copy, 引用类型传递的是指针地址，修改会改变
// 	a = 100
// }
// func f2(s []int) {
// 	s[0] = 1000
// }

// // 变长参数
// func f3(args ...int) {
// 	fmt.Printf("args: %v ,%T\n", args, args)
// 	for _, v := range args {
// 		fmt.Printf("v: %v\n", v)
// 	}
// }

// func f4(name string, ok bool, args ...int) {
// 	fmt.Printf("name: %v\n", name)
// 	fmt.Printf("ok: %v\n", ok)
// 	fmt.Printf("args: %v\n", args)
// }
// func main() {
// 	x := 200
// 	f1(x)
// 	fmt.Printf("x: %v\n", x) // 200
// 	s := []int{1, 2, 3}
// 	f2(s)
// 	fmt.Printf("s: %v\n", s) // [1000 2 3]

// 	f3(1, 2, 3, 4, 5)
// 	f4("tom", true, 0, 2, 3, 4)

// }
