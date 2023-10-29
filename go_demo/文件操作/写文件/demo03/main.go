package main

import (
	"bufio"
	"fmt"
	"os"
)

// 打开存在文件，在内容后面添加内容
func main() {
	filepath := "d:/abc.txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开错误")
	}
	defer file.Close()
	str := "wuhu\t\n"
	write := bufio.NewWriter(file)
	write.WriteString(str)
	write.Flush()
}
