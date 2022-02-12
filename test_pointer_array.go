// package main

// import "fmt"

// func main() {
// 	a := [3]int{1, 2, 3}
// 	var pa [3]*int
// 	fmt.Printf("pa: %v\n", pa) // [<nil> <nil> <nil>]

// 	for i := 0; i < len(a); i++ {
// 		pa[i] = &a[i]
// 	}

// 	fmt.Printf("pa: %v\n", pa) // [0xc000014138 0xc000014140 0xc000014148]

// 	for i := 0; i < len(pa); i++ {
// 		fmt.Printf("pa[%v]: %v\n", i, pa[i])
// 	}
// }
