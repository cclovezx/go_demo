package main

import "fmt"

type B interface {
	BB()
}

type C interface {
	CC()
}

//A继承了BC
type A interface {
	B
	C
	AA()
}
type Student struct {
	Name string
}

func (s Student) AA() {
	fmt.Println("AA")
}
func (s Student) BB() {
	fmt.Println("BB")
}
func (s Student) CC() {
	fmt.Println("CC")
}

func main() {
	//跑一下
	var stu Student
	//stu同时实现了ABC
	var a A = stu
	//不能删除其中一个实现方法
	a.AA()
	a.BB()
	a.CC()

}
