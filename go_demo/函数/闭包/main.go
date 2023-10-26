package main

import (
	"fmt"
	"os"
)

//闭包就是创建内部函数类似于in文件夹的函数

func main() {
	f()
}
func f() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d ", i) } //此例子中只是为了演示匿名函数可分配不同的内存地址，在现实开发中，不应该把该部分信息放置到循环中。
		g(i)
		fmt.Fprintf(os.Stdout, " - g is of type %v and has value %v\n", []any{g, g}...)
	}
}
