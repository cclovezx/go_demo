package main

import (
	"encoding/json"
	"fmt"
)

// 定义结构体
type Student struct {
	name string
	age  int
}

// json字符串反序列成struct
func unmarshalStruct() {
	str := "{\"name\": \"hah\",\"age\": 34}"
	//反序列化
	var student Student
	err := json.Unmarshal([]byte(str), &student)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(&student)
}

func main() {
	unmarshalStruct()
}
