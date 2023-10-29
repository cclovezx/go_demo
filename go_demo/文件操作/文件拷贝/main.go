package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 相关函数：func Copy(dst Writer,src Reader)(written int64,err error)
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcfile, err1 := os.Open(srcFileName)
	if err1 != nil {
		fmt.Println(err1)
	}
	//通过srcfile,获取到reader
	reader := bufio.NewReader(srcfile)
	defer srcfile.Close()
	//打开如果不存在创建一个
	dstfile, err2 := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	defer dstfile.Close()

	if err2 != nil {
		fmt.Println(err)
		return
	}

	//通过dstfile 获取到writer
	writer := bufio.NewWriter(dstfile)
	return io.Copy(writer, reader)

}

// Copy在io包
func main() {

	//调用我们写的函数然后进行拷贝
	fileA := "d:/a.txt"
	fileB := "d:/b.txt"
	_, err := CopyFile(fileA, fileB)
	if err != nil {
		fmt.Println("copy 错误")
	} else {
		fmt.Println("copy 完成")
	}

}
