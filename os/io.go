package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func testcopy() {
	r := strings.NewReader("hello world")
	io.Copy(os.Stdout, r)

	//if err!= nil {
	//	log.Fatal(err)
	//}
}
func test6() {
	s := strings.NewReader("ABC DEF GHI JKL")
	br := bufio.NewReader(s)
	w, _ := br.ReadBytes(' ')
	fmt.Println(string(w))
	w, _ = br.ReadBytes(' ')
	fmt.Println(string(w))
	w, _ = br.ReadBytes(' ')
	fmt.Println(string(w))
	w, _ = br.ReadBytes(' ')
	fmt.Println(string(w))
}
func test9() {
	f, _ := os.OpenFile("a.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString("hello world")
	w.Flush()

}
func test10() {
	s := strings.NewReader("ABC DEF GHI JKL")
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanWords)
	for bs.Scan() {
		fmt.Println(bs.Text())
	}
}

func main() {
	//type Reader interface {
	//	Read(P []byte) (n int, err error)
	//}
	//type Writer interface {
	//	Write(P []byte) (n int, err error)
	//}
	//r := strings.NewReader("hello world")
	//buf := make([]byte, 20)
	//r.Read(buf)
	//fmt.Println("string(buf):", string(buf))
	//s := strings.NewReader("some io.Reader stream to be read\n")
	//if _, err := io.Copy(os.Stdout, s); err != nil {
	//	log.Fatal("err")
	//}
	//r := strings.NewReader("hello world")
	//f, _ := os.Open("a.txt")
	//defer f.Close()
	//r2 := bufio.NewReader(f)
	//s, _ := r2.ReadString('\n')
	//fmt.Println(s)
	//test6()
	//test9()
	test10()
}
