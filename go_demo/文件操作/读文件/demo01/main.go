package main

import (
	"fmt"
	"os"
)

// 认识输入输出流
// 输入：文件到程序，也叫读文件
// 输出流：程序到文件，写入文件
func main() {
	//打开文件
	file, err := os.Open("D:/go_learn_wenjian/learning.txt")
	if err != nil {
		fmt.Println("open test is err:", err)
	}
	//输出文件
	fmt.Println(file)
	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close test err:", err)
	}
}
