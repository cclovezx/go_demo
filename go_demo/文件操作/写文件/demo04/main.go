package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 打开存在文件，读写并追加
func main() {
	filepath := "d:/abc.txt"
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("文件打开错误")
	}
	defer file.Close()
	//读取原来的文件内容
	reader := bufio.NewReader(file)
	for {
		str1, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("读到文件末尾了")
			break
		}
		fmt.Print(str1)
	}

	str := "gege\t\n"
	write := bufio.NewWriter(file)
	write.WriteString(str)
	write.Flush()
}
