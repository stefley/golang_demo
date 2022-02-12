// goto 语句通过标签进行代码间的无条件跳转。goto语句可以在快速跳出循环、避免重复退出上有一定帮助
// Go语言中使用goto语句能简化一些代码的实现过程。例如双层嵌套的for循环要退出时

// package main

// import "fmt"

// func test1() {
// 	a := 0
// 	if a == 1 {
// 		fmt.Println("1")
// 	} else {
// 		goto LABEL1
// 	}
// LABEL1:
// 	fmt.Println("next...")
// }

// // 跳出双层循环
// func test2() {
// 	for i := 0; i < 10; i++ {
// 		for j := 0; j < 10; j++ {
// 			if i >= 2 && j >= 2 {
// 				goto END
// 			}
// 			fmt.Printf("i: %v, j: %v\n", i, j)
// 		}
// 	}

// END:
// 	fmt.Println("END...")
// }
// func main() {
// 	test1()
// 	test2()
// }
