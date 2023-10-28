package main

import (
	"fmt"
	"os"
)

func main() {
	//获得所有环境变量***
	s := os.Environ()
	fmt.Printf("s:%v\n", s)
	//获得某个环境变量
	s2 := os.Getenv("GOPATH")
	fmt.Println(s2)
	fmt.Println(os.Getenv("JAVA_HOME"))
	//s, b := os.LookupEnv("GOPATH")
	//if b {
	//	fmt.Println(s)
	//}
	//设置环境变量
	os.Setenv("env1", "aaa")
	s2 = os.Getenv("env1")
	fmt.Println(s2)
	fmt.Println("-----------------")
	//查找指定环境变量的值
	s3, b := os.LookupEnv("env1")
	fmt.Println(b)
	fmt.Println(s3)
	//替换环境变量
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/user/gopher")
	fmt.Println(os.ExpandEnv("$NAME lives in ${BURROW}."))

	//清空环境变量
	os.Clearenv()
}
