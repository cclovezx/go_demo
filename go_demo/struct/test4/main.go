package main

import "fmt"

type student struct {
	name string
}

//实例化结构体的所有方式
func main() {
	//方式1：
	s1 := student{}
	s1.name = "cui"
	fmt.Println(s1)
	//方式二：
	s2 := student{"ha2"}
	fmt.Println(s2)
	//方式三：
	s3 := student{
		name: "ha3",
	}
	fmt.Println(s3)

	//方式四：
	var s4 student
	s4.name = "ha4"
	fmt.Println(s4)

	//方式五
	var s5 = student{
		name: "ha5",
	}
	fmt.Println(s5)

	//方式六：
	var s6 = student{"ha6"}
	fmt.Println(s6)
}
