// /*
// 	// 将struct编码成json，可以接收任意类型
// 	func Marshal(v interface {}) ([]byte, error)
// 	// 将json转码成struct结构体
// 	func Unmarshal(data []byte, v interface {}) error

// 	// 从输入流读取并解析json
// 	type Decoder struct {	}
// 	// 写json到输出流
// 	type Encoder struct {	}
// */
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"
// )

// type Person struct {
// 	Name  string
// 	Age   int
// 	Email string
// }

// func test1() {
// 	p := Person{
// 		Name:  "tom",
// 		Age:   20,
// 		Email: "tom@gmail.com",
// 	}
// 	b, _ := json.Marshal(p)
// 	fmt.Printf("b: %s\n", b)
// }

// func test2() {
// 	b := []byte(`{"Name":"tom","Age":20,"Email":"tom@gmail.com"}`)
// 	var p Person
// 	json.Unmarshal(b, &p)
// 	fmt.Printf("p: %+v\n", p)
// }

// // 解析嵌套类型
// func test3() {
// 	b := []byte(`{"Name":"tom", "Age":20,"Email":"tom@gmail.com","Parents": ["tom", "kite"]}`)
// 	var f interface{}
// 	json.Unmarshal(b, &f)
// 	fmt.Printf("f: %v\n", f)

// }

// func test4() {
// 	type NPerson struct {
// 		Name   string
// 		Age    int
// 		Email  string
// 		Parent []string
// 	}
// 	p := NPerson{
// 		Name:   "jack",
// 		Age:    11,
// 		Email:  "jack@gg.com",
// 		Parent: []string{"big", "small"},
// 	}
// 	b, _ := json.Marshal(p)
// 	fmt.Printf("b: %v\n", string(b))
// }

// // io流 Reader Writer 可以扩展到http websocket等场景
// func test5() {
// 	// dec := json.NewDecoder(os.Stdin)
// 	// a.json: {"Name": "tom", "Age": 22, "Email": "tom@aa.com", "Parent": ["tom", "kite"]}
// 	f, err := os.Open("a.json")
// 	defer f.Close()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	dec := json.NewDecoder(f)
// 	fmt.Printf("dec: %v\n", dec)
// 	enc := json.NewEncoder(os.Stdout)
// 	fmt.Printf("enc: %v\n", enc)
// 	for {
// 		var v map[string]interface{}
// 		if err := dec.Decode(&v); err != nil {
// 			log.Println(err)
// 			return
// 		}
// 		fmt.Printf("v: %v\n", v)
// 		if err := enc.Encode(&v); err != nil {
// 			log.Println(err)
// 		}
// 	}
// 	/*
// 		输入 {“"Name": "tom", Age: 20, "Email": "tom@gmail.com"}
// 		输出
// 			v: map[Age:20 Email: tom@gmail.com Name: tom]
// 			{"Age": 20, "Email": "tom@gmail.com", "Name": "tom"}
// 	*/
// }

// func test6() {
// 	p := Person{
// 		Name:  "tom",
// 		Age:   20,
// 		Email: "tom@cc.com",
// 	}

// 	f, err := os.OpenFile("a.json", os.O_WRONLY, 0777)
// 	defer f.Close()
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	e := json.NewEncoder(f)
// 	e.Encode(p)
// }
// func main() {
// 	test1()
// 	test2()
// 	test3()
// 	test4()
// 	test5()
// 	test6()
// }
