package main

import (
	"bufio"
	"fmt"
	"os"
)

// 打开存在文件，将内容覆盖掉
func main() {
	filepath := "d:/abc.txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("文件打开错误")
	}
	defer file.Close()
	str := "haha\t\n"
	write := bufio.NewWriter(file)
	write.WriteString(str)
	write.Flush()
}
