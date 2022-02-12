// package main

// import (
// 	"bufio"
// 	"bytes"
// 	"fmt"
// 	"io"
// 	"os"
// 	"strings"
// )

// const (
// 	defaultBufSize = 4096
// )

// func testReader() {

// 	r := strings.NewReader("hello world")
// 	r2 := bufio.NewReader(r)
// 	s, _ := r2.ReadString('\n')
// 	fmt.Printf("s: %v\n", s)

// 	f, _ := os.Open("a.txt")
// 	defer f.Close()
// 	r3 := bufio.NewReader(f)
// 	s2, _ := r3.ReadString('\n')
// 	fmt.Printf("s2: %v\n", s2)
// }
// func testRead() {
// 	s := strings.NewReader("abdfjklsjfldsjlf")
// 	br := bufio.NewReader(s)
// 	p := make([]byte, 4)

// 	for {
// 		n, err := br.Read(p)
// 		if err == io.EOF {
// 			break
// 		} else {
// 			fmt.Printf("string(p):%v\n", string(p[0:n]))
// 		}
// 	}
// }
// func testReadByte() {
// 	s := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
// 	br := bufio.NewReader(s)

// 	c, _ := br.ReadByte()
// 	fmt.Printf("c: %c\n", c)

// 	c, _ = br.ReadByte()
// 	fmt.Printf("c: %c\n", c)

// 	br.UnreadByte()
// 	c, _ = br.ReadByte()
// 	fmt.Printf("c: %c\n", c)
// }

// func testReadRune() {
// 	s := strings.NewReader("你好，世界")
// 	br := bufio.NewReader(s)

// 	c, size, _ := br.ReadRune()
// 	fmt.Printf("%c %v\n", c, size)

// 	c, size, _ = br.ReadRune()
// 	fmt.Printf("%c %v\n", c, size)

// 	br.UnreadRune()
// 	c, size, _ = br.ReadRune()
// 	fmt.Printf("%c %v\n", c, size)
// }
// func testReadLine() {
// 	s := strings.NewReader("dsljkflsjflaps;dfksjfojsadf\nsldfjldssjfdk\nfdlsjfl\n")
// 	br := bufio.NewReader(s)

// 	w, isPrefix, _ := br.ReadLine()
// 	fmt.Printf("%q %v\n", w, isPrefix)

// 	w, isPrefix, _ = br.ReadLine()
// 	fmt.Printf("%q %v\n", w, isPrefix)

// 	w, isPrefix, _ = br.ReadLine()
// 	fmt.Printf("%q %v\n", w, isPrefix)

// 	w, isPrefix, _ = br.ReadLine()
// 	fmt.Printf("%q %v\n", w, isPrefix)

// }

// func testReadSlice() {
// 	s := strings.NewReader("AB,CD,EF")
// 	br := bufio.NewReader(s)

// 	w, _ := br.ReadSlice(',')
// 	fmt.Printf("w: %q\n", w)

// 	w, _ = br.ReadSlice(',')
// 	fmt.Printf("w: %q\n", w)

// 	w, _ = br.ReadSlice(',')
// 	fmt.Printf("w: %q\n", w)
// }

// func testReadBytes() {
// 	s := strings.NewReader("ABC DEF MK")
// 	br := bufio.NewReader(s)

// 	w, _ := br.ReadBytes(' ')
// 	fmt.Printf("%q\n", w)
// 	w, _ = br.ReadBytes(' ')
// 	fmt.Printf("%q\n", w)
// 	w, _ = br.ReadBytes(' ')
// 	fmt.Printf("%q\n", w)
// }

// func testReadString() {
// 	s := strings.NewReader("asdgg ldsjflk dlkjfs")
// 	br := bufio.NewReader(s)

// 	w, _ := br.ReadString(' ')
// 	fmt.Printf("%q\n", w)
// 	w, _ = br.ReadString(' ')
// 	fmt.Printf("%q\n", w)
// 	w, _ = br.ReadString(' ')
// 	fmt.Printf("%q\n", w)
// }

// func testWriteTo() {
// 	s := strings.NewReader("sldjflkslf dslkfj dslk")
// 	br := bufio.NewReader(s)
// 	b := bytes.NewBuffer(make([]byte, 0))

// 	br.WriteTo(b)
// 	fmt.Printf("b: %s\n", b)

// 	s2 := strings.NewReader("new string")
// 	br2 := bufio.NewReader(s2)
// 	f, _ := os.OpenFile("a.txt", os.O_RDWR, 0777)
// 	defer f.Close()
// 	br2.WriteTo(f)

// }

// func testWrite() {
// 	f, _ := os.OpenFile("a.txt", os.O_RDWR, 0777)
// 	defer f.Close()
// 	w := bufio.NewWriter(f)
// 	w.WriteString("hello golang")
// 	w.Flush() // 刷新缓存区
// }

// func testWriteReset() {
// 	b := bytes.NewBuffer(make([]byte, 0))
// 	w := bufio.NewWriter(b)
// 	w.WriteString("abc")
// 	c := bytes.NewBuffer(make([]byte, 0))
// 	w.Reset(c)
// 	w.WriteString("123")
// 	w.Flush()
// 	fmt.Printf("b: %v\n", b)
// 	fmt.Printf("c: %v\n", c)
// }
// func testWriteFrom() {
// 	b := bytes.NewBuffer(make([]byte, 0))
// 	bw := bufio.NewWriter(b)
// 	fmt.Printf("bw.Available(): %v\n", bw.Available())
// 	fmt.Printf("bw.Buffered(): %v\n", bw.Buffered())

// 	bw.WriteString("abcABC")
// 	fmt.Printf("bw.Available(): %v\n", bw.Available())
// 	fmt.Printf("bw.Buffered(): %v\n", bw.Buffered())
// 	fmt.Printf("b: %q\n", b)

// 	bw.Flush()
// 	fmt.Printf("bw.Available(): %v\n", bw.Available())
// 	fmt.Printf("bw.Buffered(): %v\n", bw.Buffered())
// 	fmt.Printf("b: %q\n", b)
// }

// func testReadWrite() {
// 	b := bytes.NewBuffer(make([]byte, 0))
// 	w := bufio.NewWriter(b)
// 	s := strings.NewReader("123")
// 	r := bufio.NewReader(s)
// 	rw := bufio.NewReadWriter(r, w)
// 	p, _ := rw.ReadString('\n')
// 	fmt.Printf("p: %s\n", p)
// 	rw.WriteString("asdf")
// 	rw.Flush()
// 	fmt.Printf("b: %v\n", b)
// }

// func testScanner() {
// 	s := strings.NewReader("ABC DEF IK")
// 	bs := bufio.NewScanner(s)
// 	bs.Split(bufio.ScanWords)
// 	for bs.Scan() {
// 		fmt.Println(bs.Text())
// 	}
// }
// func testScanRune() {
// 	s := strings.NewReader("HEllo 世界！")
// 	bs := bufio.NewScanner(s)
// 	bs.Split(bufio.ScanRunes)
// 	for bs.Scan() {
// 		fmt.Printf("bs.Text(): %s\n", bs.Text())
// 	}
// }
// func main() {
// testReader()
// testRead()
// testReadByte()
// testReadRune()
// testReadLine()
// testReadSlice()
// testReadBytes()
// testReadString()
// testWriteTo()
// testWrite()
// testWriteReset()
// testWriteFrom()
// testReadWrite()
// testScanner()
// testScanRune()
// }
