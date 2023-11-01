package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Count struct {
	Engcount   int
	Kongcount  int
	Numcount   int
	Othercount int
}

func main() {

	filename := "D:/a.txt"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	var count Count

	rider := bufio.NewReader(file)
	for {
		str, err1 := rider.ReadString('\n')
		if err1 == io.EOF {
			fmt.Println("读取到文件末尾了")
			break
		}

		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.Engcount++
			case v == ' ' || v == '\t':
				count.Kongcount++
			case v >= '0' && v < '9':
				count.Numcount++
			default:
				count.Othercount++
			}
		}
	}
	fmt.Println(count.Engcount, count.Kongcount, count.Numcount, count.Othercount)

}
