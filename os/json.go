package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string   `json:"name"`
	Age    int      `json:"age"`
	Email  string   `json:"email"`
	Parent []string `json:"parent"`
}

func test12() {
	p := Person{
		Name:   "tom",
		Age:    20,
		Email:  "tom@qq.com",
		Parent: []string{"tom", "jerry"},
	}
	b, _ := json.Marshal(p)
	fmt.Println(string(b))
}
func test13() {
	jsonStr := []byte(`{"name":"Bob","age":30}`)
	var p Person
	err := json.Unmarshal(jsonStr, &p)
	if err != nil {
		return
	}
	fmt.Println(p)
}

// 解析嵌套类型（可以用json.RawMessage来解析未知类型）

func main() {
	test12()
	test13()
}
