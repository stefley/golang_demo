// break关键字可以结束for、switch、select的代码块
/**
*单独在select中使用break和不适用break没有区别
单独在表达式switch语句，并没有fallthrough，使用break和不适用break没区别
单独在表达式switch语句，并且有fallthrough，使用break能终止fallthrough后面的case执行
带标签break可以跳出多层select/select作用域。让break更加灵活，写法更加简单灵活，不需要使用控制变量一曾一层跳出循环，没有带break只能跳出当前语句
*/
// package main

// import "fmt"

// func main() {
// 	f1()
// 	f2()
// 	f3()
// }

// func f1() {
// 	for i := 0; i <= 10; i++ {
// 		fmt.Println(i)
// 		break // 跳出循环
// 	}
// }

// func f2() {
// 	i := 2
// 	switch i {
// 	case 1:
// 		fmt.Println("1")
// 		break
// 		// fallthrough
// 	case 2:
// 		fmt.Println("2")
// 		// break
// 		fallthrough
// 	case 3:
// 		fmt.Println("3")

// 	}
// }

// // break到标签
// func f3() {
// MYLABEL:

// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("i: %v\n", i)
// 		if i >= 5 {
// 			break MYLABEL
// 		}
// 	}
// 	fmt.Println("END....")
// }
