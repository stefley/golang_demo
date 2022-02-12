/**
*switch
*支持多条件匹配
*不同case之间不使用break，默认只会执行一个case
*如果要执行多个case，需要使用fallthrough关键字，也可以用break终止
*分支可以使用表达式
 */
// package main

// import "fmt"

// func main() {
// 	// f2()
// 	// f3()
// 	f4()
// }
// func f1() {
// 	grade := "A"
// 	switch grade {
// 	case "A":
// 		fmt.Println("优秀")
// 	case "B":
// 		fmt.Println("良好")
// 	case "C":
// 		fmt.Println("及格")
// 	default:
// 		fmt.Println("不及格")
// 	}
// }
// func f2() {
// 	day := 3
// 	switch day {
// 	case 1, 2, 3, 4, 5:
// 		fmt.Println("工作日")
// 	case 6, 7:
// 		fmt.Println("休息日")
// 	default:
// 		fmt.Println("非法输入")
// 	}
// }

// func f3() {

// 	score := 88
// 	// 什么条件都没有默认true
// 	switch {
// 	case score >= 60:
// 		fmt.Println("通过")
// 	case score >= 80 && score < 90:
// 		fmt.Println("优秀")
// 	default:
// 		fmt.Println("其他")
// 	}
// }

// func f4() {
// 	num := 100
// 	switch num {
// 	case 100:
// 		fmt.Println("100")
// 		fallthrough
// 	case 200:
// 		fmt.Println("200")
// 	case 300:
// 		fmt.Println("300")
// 	}
// }
