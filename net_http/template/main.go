package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayHello)
	//把sayHello这个函数与'/'路由绑定,意思就是当访问'/'这个路由时,就会执行sayHello这个函数
	err := http.ListenAndServe(":9090", nil)
	//启动服务
	if err != nil {
		fmt.Println("http server start failed,err:", err)
		return
	}
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	//2解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	//解析hello.tmpl模板文件
	if err != nil {
		fmt.Println("parse template failed,err:", err)
		return
	}
	//3渲染模板
	err = t.Execute(w, "沙河小王子") //传入的是一个字符串
	if err != nil {
		fmt.Println("render template failed,err:", err)
		return
	}
}
