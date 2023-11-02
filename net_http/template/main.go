package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http server start failed,err:", err)
		return
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//2解析模板

	//3渲染模板
}
