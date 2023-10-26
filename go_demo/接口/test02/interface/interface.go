package interfaced2

import "fmt"

type Student struct {
}

type A interface {
	Speak()
}
type B interface {
	Hello()
}
type C interface {
	A
	B
	Wa()
}

func (s *Student) Speak() {
	fmt.Println("我是小学生")
}

func (s *Student) Hello() {
	fmt.Println("hello!!")
}

func (s *Student) Wa() {
	fmt.Print("you you")
}
