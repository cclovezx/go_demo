package main

import (
	"bufio"
	"fmt"
	"os"
)

//文件写入操作
//name文件名称  flag打开文件方式 perm用于linux
//func OpenFile(name string, flag int, perm FileMode) (*File, error) {
//	...
//}

/*
文件打开模式
O_RDONLY int = syscall.O_RDONLY // 只读模式
O_WRONLY int = syscall.O_WRONLY // 只写模式
O_RDWR int = syscall.O_RDWR // 读写模式
O_APPEND int = syscall.O_APPEND //追加模式
O_CREATE int = syscall.O_CREAT // 创建模式，如果文件不存在的话
O_TRUNC int = syscall.O_TRUNC // 覆盖模式
*/

func main() {
	//创建一个新文件
	//1.打开文件
	filepath := "d:/abc.txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	//准备写入
	str := "cuichi\r\n" //\r\n代表换行
	//写入时，使用带缓存的writer
	writer := bufio.NewWriter(file)
	writer.WriteString(str)

	//因为上述是带缓存的，因此调用上述方法写入文件的时候是先写在缓存的
	//所以要调用Flush函数，将缓存的数据写入到文件
	writer.Flush()
}
