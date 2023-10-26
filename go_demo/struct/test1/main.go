package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score float64
}

//如果我们要给结构体绑定方法的话，看我操作

func (s *Student) Speak() {
	fmt.Print("我叫：", s.Name)
}

func main() {
	//声明方法1：
	var s Student
	s.Name = "cuichi"
	s.Age = 43
	s.Score = 80.90
	fmt.Println(s)

	//声明方法2:
	d := Student{}
	d.Name = "haha"
	fmt.Println(d)

}
