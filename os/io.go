package main

import (
	"fmt"
	"strings"
)

func main() {
	//type Reader interface {
	//	Read(P []byte) (n int, err error)
	//}
	//type Writer interface {
	//	Write(P []byte) (n int, err error)
	//}
	r := strings.NewReader("hello world")
	buf := make([]byte, 20)
	r.Read(buf)
	fmt.Println("string(buf):", string(buf))
}
