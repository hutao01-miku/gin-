package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

func check(s string) (string, error) {
	if s == "" {
		err := errors.New("字符串不能为空")
		return "", err
	} else {
		return s, nil
	}
}
func test11() {
	//t := time.Now()
	//fmt.Printf("t:%T\n", t)
	//fmt.Println(t)
	//i := t.Year()
	//m := t.Month()
	//n := t.Day()
}
func test() {

}
func main() {
	p := Person{
		Name:  "tom",
		Age:   20,
		Email: "tom@qq.com",
	}
	b, _ := json.Marshal(p)
	fmt.Println(string(b))
	//s, err := check("234")
	//if err != nil {
	//	fmt.Println(err.Error())
	//} else {
	//	fmt.Println(s)
	//}

	//data := sort.Float64Slice{3.3434, 1.34324, 4.34, 1, 5, 9, 2, 6, 5, 3, 5}
	//sort.Sort(data)
	//fmt.Println(data)
	//test11()

}
