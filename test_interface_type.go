// package main

// import (
// 	"fmt"
// )

// type Player interface {
// 	playMusic()
// }
// type PlayVideo interface {
// 	playVideo()
// }
// type Mobile struct {
// }

// func (m Mobile) playMusic() {
// 	fmt.Println("play music....")
// }

// func (m Mobile) playVideo() {
// 	fmt.Println("play video.....")
// }

// // 多个类型实现同一接口
// type Pet interface {
// 	eat()
// }
// type Dog struct {
// }
// type Cat struct{}

// func (d Dog) eat() {
// 	fmt.Println("dog eat....")
// }
// func (c Cat) eat() {
// 	fmt.Println("cat eat....")
// }
// func main() {
// 	m := Mobile{}
// 	m.playMusic()
// 	m.playVideo()

// 	d := Dog{}
// 	d.eat()
// 	c := Cat{}
// 	c.eat()

// 	var pet Pet
// 	pet = Dog{}
// 	pet.eat()
// 	pet = Cat{}
// 	pet.eat()
// }
