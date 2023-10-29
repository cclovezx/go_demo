package main

import (
	"fmt"
	"io/ioutil"
)

// 直接读取一整个文件，不需要打开文件，
// 一般用来读取小文件
// 相关函数：ioutil.ReadFile
func main() {
	byte, err := ioutil.ReadFile("D:/go_learn_wenjian/learning.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", byte)
	fmt.Printf("%s", string(byte))
}
