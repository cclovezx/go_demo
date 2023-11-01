package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name     string `json:"name"` //这个意识是固定序列化后的值
	Age      int
	Birthday string
}

func testStruct() {
	student := Student{
		Name:     "hah",
		Age:      43,
		Birthday: "12.07",
	}
	//序列化
	data, err := json.Marshal(student)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

}

func main() {
	testStruct()
}
