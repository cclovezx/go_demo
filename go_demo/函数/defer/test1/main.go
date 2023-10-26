package main

import "fmt"

//defer 关键字的作用简单来说被它修饰的肯定最后执行
//不用管太多先修饰的最后执行
//下面的代码简单展示了一下
func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func main() {
	b()
}
