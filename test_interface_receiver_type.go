// package main

// import "fmt"

// type Pet interface {
// 	eta(string) string
// }

// type Dog struct {
// 	name string
// }

// func (d Dog) eat(name string) {
// 	d.name = "poli"
// 	fmt.Printf("name: %v\n", name)
// 	fmt.Println("eat....")
// }
// func (d *Dog) run(name string) {
// 	d.name = "poil"
// 	fmt.Printf("name: %v\n", name)
// 	fmt.Println("run......")
// }
// func main() {
// 	d := &Dog{
// 		name: "tol",
// 	}
// 	d.eat("meat")
// 	d.run("fast...")
// 	fmt.Printf("d: %v\n", d)
// }
