package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	//这样是可以的
	a = point

	var b Point
	//b=a 不可以赋值
	//b=a.(Point)
	b = a.(Point)
	fmt.Print(b)

	//类型断言
	var x interface{}
	var b2 float32 = 1.1
	x = b2

	y := x.(float32)
	fmt.Print(y)
}
