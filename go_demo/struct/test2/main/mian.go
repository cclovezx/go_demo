package main

import (
	"fmt"
	"packagestudy/struct/test2/model"
)

// 注意不能饮用中文路径
func main() {
	//方式1：
	s := model.NewStudent("cui", 16, 90.21)
	s.SetName("cuichi")
	fmt.Printf("s.GetName(): %v\n", s.GetName())
}
