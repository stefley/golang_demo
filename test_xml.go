// package main

// import (
// 	"encoding/xml"
// 	"fmt"
// 	"io/ioutil"
// )

// type Person struct {
// 	XMLName xml.Name `xml: "person"`
// 	Name    string   `xml: "name"`
// 	Age     int      `xml: "age"`
// 	Email   string   `xml: "Email"`
// }

// func test1() {
// 	p := Person{
// 		Name:  "tom",
// 		Age:   20,
// 		Email: "tom@tt.com",
// 	}
// 	// b, _ := xml.Marshal(p)
// 	// 有缩进格式
// 	b, _ := xml.MarshalIndent(p, " ", "	")
// 	fmt.Printf("b: %s\n", b)
// }

// func test2() {
// 	type Email struct {
// 		Where string `xml:"where,attr"`
// 		Addr  string
// 	}
// 	type Address struct {
// 		City, State string
// 	}
// 	type Result struct {
// 		XMLName xml.Name `xml:"Person"`
// 		Name    string   `xml:"FullName"`
// 		Phone   string
// 		Email   []Email
// 		Groups  []string `xml:"Group>Value"`
// 		Address
// 	}
// 	v := Result{Name: "none", Phone: "none"}

// 	data := `
// 		<Person>
// 			<FullName>Grace R. Emlin</FullName>
// 			<Company>Example Inc.</Company>
// 			<Email where="home">
// 				<Addr>gre@example.com</Addr>
// 			</Email>
// 			<Email where='work'>
// 				<Addr>gre@work.com</Addr>
// 			</Email>
// 			<Group>
// 				<Value>Friends</Value>
// 				<Value>Squash</Value>
// 			</Group>
// 			<City>Hanga Roa</City>
// 			<State>Easter Island</State>
// 		</Person>
// 	`
// 	err := xml.Unmarshal([]byte(data), &v)
// 	if err != nil {
// 		fmt.Printf("error: %v", err)
// 		return
// 	}
// 	fmt.Printf("v: %s\n", v)
// 	fmt.Printf("XMLName: %#v\n", v.XMLName)
// 	fmt.Printf("Name: %q\n", v.Name)
// 	fmt.Printf("Phone: %q\n", v.Phone)
// 	fmt.Printf("Email: %v\n", v.Email)
// 	fmt.Printf("Groups: %v\n", v.Groups)
// 	fmt.Printf("Address: %v\n", v.Address)
// }

// func test3() {
// 	b, _ := ioutil.ReadFile("a.xml")

// 	var per Person
// 	xml.Unmarshal(b, &per)
// 	fmt.Printf("per: %v\n", per)
// }
// func main() {
// 	test2()
// 	test3()
// }
