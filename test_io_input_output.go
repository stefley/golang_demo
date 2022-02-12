// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// 	"strings"
// )

// /*
// 实现io.Reader io.Writer 接口的类型
// os.File
// strings.Reader io.Reader
// bufio.Reader/Writer
// bytes.Buffer
// bytes.Reader io.Reader
// compress/gzip.Reader/Writer
// crypto/cipher.StreamReader/StreamWriter
// crypto/tls.Conn
// encoding/csv.Reader/Writer
// */

// func strReader() {
// 	s := strings.NewReader("hello world")
// 	buf := make([]byte, 10)
// 	s.Read(buf)
// 	fmt.Printf("string(buf): %v\n", string(buf))
// }
// func copy() {
// 	// copy
// 	r := strings.NewReader("some io.Reader stream to be read\n")
// 	fmt.Printf("r: %v\n", r) // r: &{some io.Reader stream to be read 0 -1}
// 	if _, err := io.Copy(os.Stdout, r); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func copyBuf() {
// 	r1 := strings.NewReader("first reader\n")
// 	r2 := strings.NewReader("scond reader\n")
// 	buf := make([]byte, 8)

// 	// buf is uesed here..
// 	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
// 		log.Fatal(err)
// 	}

// 	// ...reused here also. Non need to allocate an extra buffer
// 	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func limitReader() {
// 	r := strings.NewReader("some io.Reader stream to be read\n")
// 	lr := io.LimitReader(r, 4) // some
// 	if _, err := io.Copy(os.Stdout, lr); err != nil {
// 		log.Fatal(err)
// 	}
// }
// func multiReader() {
// 	r1 := strings.NewReader("first")
// 	r2 := strings.NewReader("second")
// 	r3 := strings.NewReader("third")
// 	r := io.MultiReader(r1, r2, r3)

// 	if _, err := io.Copy(os.Stdout, r); err != nil {
// 		log.Fatal(err)
// 	}
// }
// func multiWrite() {
// 	r := strings.NewReader("some io.Reader stream to read\n")
// 	var buf1, buf2 bytes.Buffer
// 	w := io.MultiWriter(&buf1, &buf2)
// 	if _, err := io.Copy(w, r); err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("buf1.String(): %v\n", buf1.String())
// 	fmt.Printf("buf2.String(): %v\n", buf2.String())
// }
// func pipe() {
// 	r, w := io.Pipe()
// 	go func() {
// 		fmt.Fprint(w, "some io.Reader stream to be read\n")
// 		w.Close()
// 	}()

// 	if _, err := io.Copy(os.Stdout, r); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func readAll() {
// 	r := strings.NewReader("Go is a  general-purpose language designed with systems programming in mind")

// 	b, err := io.ReadAll(r)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("b: %v\n", b)
// 	fmt.Printf("b: %s\n", b)
// }

// func sectionReader() {
// 	r := strings.NewReader("some io.Reader stream to be read\n")
// 	s := io.NewSectionReader(r, 5, 17)

// 	if _, err := io.Copy(os.Stdout, s); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func writeString() {
// 	if _, err := io.WriteString(os.Stdout, "hello world"); err != nil {
// 		log.Fatal(err)
// 	}
// }
// func main() {
// 	copy()
// 	copyBuf()
// 	limitReader()
// 	multiReader()
// 	multiWrite()
// 	pipe()
// 	readAll()
// 	sectionReader()
// 	writeString()
// }
