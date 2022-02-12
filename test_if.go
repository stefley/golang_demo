// package main

// import "fmt"

// func main() {
// 	// a := 1
// 	// b := 2

// 	// if a > b {
// 	// 	fmt.Println("a > b")
// 	// } else {
// 	// 	fmt.Println("a < b")
// 	// }

// 	// if age := 20; age > 18 {
// 	// 	fmt.Println("成年人")
// 	// }
// 	// a := 100
// 	// if a {
// 	// 	fmt.Printf("aj", a)
// 	// }

// 	// var name string
// 	// var age int
// 	// var email string

// 	// fmt.Println("请输入name,age,email,用空格分割")
// 	// fmt.Scan(&name, &age, &email)
// 	// fmt.Printf("name: %v\n", name)
// 	// fmt.Printf("age: %v\n", age)
// 	// fmt.Printf("email: %v\n", email)

// 	// f()
// 	// f2()
// 	// f3()
// 	// f4()
// 	f5()
// }

// func f() {
// 	var num int
// 	fmt.Println("请输入一个数字：")
// 	fmt.Scan(&num)

// 	if num%2 == 0 {
// 		fmt.Println("偶数")
// 	} else {
// 		fmt.Println("奇数")
// 	}
// }

// func f2() {
// 	// score := 80
// 	// if score >= 60 && score <= 70 {
// 	// 	fmt.Println("C")
// 	// } else if score >= 70 && score <= 90 {
// 	// 	fmt.Println("B")
// 	// } else {
// 	// 	fmt.Println("A")
// 	// }
// 	if score := 92; score >= 60 && score <= 70 {
// 		fmt.Println("C")
// 	} else if score >= 70 && score <= 90 {
// 		fmt.Println("B")
// 	} else {
// 		fmt.Println("A")
// 	}
// }

// // 输入星期几的第一个字母来判断星期几，如第一个字母一样则判断第二个字母
// func f3() {
// 	// Monday Tuesday Wednesday Thursday Friday Saturday Sunday
// 	var c string
// 	fmt.Println("请输入一个字母:")
// 	fmt.Scan(&c)

// 	if c == "S" {
// 		fmt.Println("请输入第二个字母:")
// 		fmt.Scan(&c)
// 		if c == "a" {
// 			fmt.Println("Saturday")
// 		} else if c == "u" {
// 			fmt.Println("Sunday")
// 		} else {
// 			fmt.Println("输入错误")
// 		}
// 	} else if c == "T" {
// 		fmt.Println("请输入第二个字符:")
// 		fmt.Scan(&c)
// 		if c == "u" {
// 			fmt.Println("Tuesday")
// 		} else if c == "h" {
// 			fmt.Println("Thursday")
// 		} else {
// 			fmt.Println("输入错误")
// 		}
// 	} else if c == "M" {
// 		fmt.Println("Monday")
// 	} else if c == "W" {
// 		fmt.Println("Wednesday")
// 	} else if c == "F" {
// 		fmt.Println("Friday")
// 	} else {
// 		fmt.Println("输入错误")
// 	}
// }

// func f4() {
// 	a, b, c := 1, 2, 3
// 	if a > b {
// 		if a > c {
// 			fmt.Println("a")
// 		}
// 	} else {
// 		if b > c {
// 			fmt.Println("b")
// 		} else {
// 			fmt.Println("c")
// 		}
// 	}
// }

// func f5() {
// 	gender := "男"

// 	age := 20

// 	if gender == "男" {
// 		if age >= 18 {
// 			fmt.Println("成年男性")
// 		} else {
// 			fmt.Println("未成年男性")
// 		}
// 	} else {
// 		if age >= 18 {
// 			fmt.Println("成年女性")
// 		} else {
// 			fmt.Println("未成年女性")
// 		}
// 	}
// }
