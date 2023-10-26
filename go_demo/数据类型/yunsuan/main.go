package main

import "fmt"

func main() {
	a := 10
	b := 8
	//加法
	fmt.Print(a + b)
	//减法
	fmt.Print(a - b)
	//乘法
	fmt.Print(a * b)
	//除法
	fmt.Print(a / b)

	//高级
	//与
	var c, d = true, false
	fmt.Print(c && d)
	//或
	fmt.Print(c || d)
	//非
	fmt.Print(!c, !d)

	//更高级
	//按位或
	fmt.Print(a & b)
	//按位与
	fmt.Print(a | b)
}
