package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//读取文件的内容并且显示在终端
//函数：os.Open ;file Close ;bufio.NewReader();reader.ReadString

func main() {
	file, err := os.Open("D:/go_learn_wenjian/learning.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	//创建一个带缓冲的reader
	//文档中写的缓冲区默认值是4096
	reader := bufio.NewReader(file)

	//循环读取文件
	for {
		//\n读到一个换行符就结束
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("读到文件末尾了")
			break
		} else {
			fmt.Println(str)
		}
	}
	fmt.Println("文件读取结束")
}
