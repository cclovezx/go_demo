package main

import "fmt"

func main() {
	//方式一
	var a [5]int
	fmt.Print(a)
	//方式二
	b := [...]int{1, 2, 3, 4, 5}
	fmt.Print(b)
	c := [...]string{"a", "b", "c", "d"}
	fmt.Println(c)

	//如果数组值已经提前知道了，那么可以通过 数组常量 的方法来初始化数组，而不用依次使用 []= 方法（所有的组成元素都有相同的常量语法）。
	// var arrAge = [5]int{18, 20, 15, 22, 16}
	// var arrLazy = [...]int{5, 6, 7, 8, 22}
	// var arrLazy = []int{5, 6, 7, 8, 22}
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	// var arrKeyValue = []string{3: "Chris", 4: "Ron"}

	for i := 0; i < len(arrKeyValue); i++ {
		fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
	}

	//
}
