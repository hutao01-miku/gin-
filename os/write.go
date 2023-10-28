package main

import "os"

func write() { //写字节数组
	f, _ := os.OpenFile("b.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	f.Write([]byte("hello world"))
	f.Close()
}
func writeString() { //写字符串
	f, _ := os.OpenFile("b.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	f.WriteString("hello worldd effeeful world")
}
func writeAt() { //随机写
	f, _ := os.OpenFile("b.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	f.WriteAt([]byte("hello world"), 3)
	f.Close()
}

func main() {
	//write()
	//writeString()
	writeAt()
}
