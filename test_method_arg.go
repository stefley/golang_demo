// package main

// import "fmt"

// type Person struct {
// 	name string
// }

// func test1() {

// 	p1 := Person{
// 		name: "tom",
// 	}
// 	p2 := &Person{
// 		name: "tom",
// 	}

// 	fmt.Printf("p1: %T\n", p1) // p1: main.Person
// 	fmt.Printf("p2: %T\n", p2) // p2: *main.Person
// }

// func showPerson1(per Person) {
// 	per.name = "jack"
// }
// func showPerson2(per *Person) {
// 	// per.name 自动解引用 等价(*per).name
// 	per.name = "jack"
// }

// func (per Person) showPerson3() {
// 	per.name = "tom...."
// }
// func (per *Person) showPerson4() {
// 	per.name = "tom...."
// }
// func main() {
// 	p1 := Person{
// 		name: "tom",
// 	}

// 	p2 := &Person{
// 		name: "tom",
// 	}
// 	showPerson1(p1)
// 	showPerson2(p2)
// 	fmt.Printf("p1: %v\n", p1)
// 	fmt.Printf("p2: %v\n", *p2)

// 	p1.showPerson3()
// 	fmt.Printf("p1: %v\n", p1) // p1: {tom}
// 	fmt.Println("-----------------")
// 	p2.showPerson4()
// 	fmt.Printf("p2: %v\n", p2) // p2: {tom....}
// }
