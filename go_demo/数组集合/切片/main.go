package main

import "fmt"

//这里面最重要的是多维切片

func main() {
	//创建方式1
	var a = []int{1, 2, 3, 4, 5}
	var arr = a[1:4]
	fmt.Print(arr)

	//使用make创建切片
	//格式： make([]type,len,capital)
	b := make([]int, 3)
	for i := range b {
		fmt.Print(i)
	}

	//new一个
	var p *[]int = new([]int) // *p == nil; with len and cap 0
	d := new([]int)
	fmt.Print(p, d)

	fmt.Println()
	fmt.Print(sum(a))

	//多维切片

	//注意一点每个维度都需要进行切片
	//举个栗子
	var m = make([][]int, 5)
	for i := 0; i < len(m); i++ {
		m[i] = make([]int, 5)
	}
	fmt.Print(m)

}

//将切片传递给函数
func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}
