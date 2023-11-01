package main

import (
	"fmt"
	"os"
)

func main() {
	//遍历os.Args切片就可以获得所有命令函输入的参数值
	for _, v := range os.Args {
		fmt.Println(v)
	}
}
