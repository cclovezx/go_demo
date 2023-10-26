package main

import (
	interfaced2 "packagestudy/interface/test02/interface"
)

func main() {
	stu := interfaced2.Student{}

	//在这里不用太过于关注这些方法是那个接口的，能跑就行
	stu.Hello()
	stu.Speak()
	stu.Wa()
}
