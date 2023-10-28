package main

import (
	"fmt"
	"os"
)

// 创建文件
func createFile() {
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name():%v\n", f.Name())
	}
}

// 创建目录
func createDir() {
	err := os.Mkdir("b", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 删除目录或文件
func removeDir() {
	err := os.Remove("b")
	//err2 := os.RemoveAll("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
func chWd() { //切换目录
	fmt.Println(os.Getwd())
	err := os.Chdir("D:/")
	if err != nil {
		fmt.Printf("err:%v\n\n", err)
	}
	fmt.Println(os.Getwd())
	s := os.TempDir()
	fmt.Println(s)
}
func rename() {
	err := os.Rename("a.txt", "b.txt")
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
}
func readFile() { //读取文件
	b, err := os.ReadFile("b.txt")
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	fmt.Printf("b:%v\n", string(b[:]))
}
func writeFile() { //写入文件
	err := os.WriteFile("b.txt", []byte("hello world"), os.ModePerm)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
}
func main() {
	//createFile()
	//createDir()
	//removeDir()
	//chWd()
	//rename()
	//readFile()
	writeFile()
}
