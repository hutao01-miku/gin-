package main

import (
	"fmt"
	"io"
	"os"
)

func openCloseFile() {
	//只能读
	f, _ := os.Open("b.txt")
	fmt.Println(f.Name())
	err := f.Close()
	if err != nil {
		return
	}
	f2, err := os.OpenFile("b.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f2.Name())
		f2.Close()
	}
	f3, _ := os.CreateTemp("", "a.txt")
	fmt.Println(f3.Name())
}
func readOps() { //读取文件
	f, _ := os.Open("b.txt")
	f.Seek(3, 0)
	for {
		buf := make([]byte, 10)
		n, err := f.Read(buf)
		if err == io.EOF { //超出末尾数据时报错
			break
		}
		fmt.Println(n)
		fmt.Println(string(buf[:n]))
	}
}
func read() { //读取目录,并测试目录是否存在
	f, _ := os.Open("a")
	de, _ := f.ReadDir(-1)
	for _, v := range de {
		fmt.Println(v.IsDir())
		fmt.Println(v.Name())
	}
}
func main() {
	//openCloseFile()
	//readOps()
	//read()
}
