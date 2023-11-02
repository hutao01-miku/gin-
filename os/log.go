package main

import (
	"fmt"
	"log"
	"os"
)

func test1() {
	log.Print("my log")
	log.Printf("my log%d", 100)
	name := "tom"
	age := 20
	log.Println(name, " ", age)
}
func test2() {
	defer fmt.Println("panic 结束之后在执行。。。")
	log.Print("my log")
	//log.Print("my log%d", 100)
	//log.Fatal("my log")
	log.Panic("my panic")
	fmt.Println("end...")
}
func main() {
	i := log.Flags()
	fmt.Println(i)
	log.Print("my log")

}

func init() {
	//log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("[my log]") //前缀设置
	f, err := os.OpenFile("a.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		log.Fatal("日志文件错误")
	}
	log.SetOutput(f)
}
