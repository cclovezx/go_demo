package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileapath := "d:/a.txt"
	filebpath := "d:/b.txt"
	//1.首先读取fileA文件的内容到内存
	strA, err := ioutil.ReadFile(fileapath)
	if err != nil {
		fmt.Println("读文件出错了")
		return
	}

	//2.将读取到的内容添加到b
	err2 := ioutil.WriteFile(filebpath, strA, 0666)
	if err2 != nil {
		fmt.Println("出错了")
	}

}

// 函数判断文件是否存在
// 如果返回错误为nil说明存在
// 如果返回值类型用IsNotExist判断为true说明不存在
// 如果返回别的的东西那就不知道存在不存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}
