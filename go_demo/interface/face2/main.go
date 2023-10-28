package main

import "fmt"

type Ainterface interface {
	Say()
}

type BInetrface interface {
	Hello()
}

type Student struct {
	Name string
}

func (stu Student) Say() {
	fmt.Printf("我叫%s", stu.Name)
}
func (stu Student) Hello() {
	fmt.Println("hello ")
}

type Interger int

func (i Interger) Say() {
	fmt.Println("我不是结构体")
}

func main() {
	//1.接口本身不能自己创建实例，但是可以指向一个了实现了该接口的自定义类型变量
	//例如这样
	var stu Student        //结构体变量，实现了Say()方法，就说明实现了接口
	var A Ainterface = stu //所以可以这样来写
	A.Say()
	//2. 接口中所有的方法都没有方法体，只有方法

	//3.实现了某个接口的意思是，某个方法实现接口了所有的方法

	//4.只要是自定义类型就可以实现接口，不一定要是结构体
	//例如：
	var a Interger
	var B Ainterface = a
	B.Say()

	//5.一个自定义对象可以实现多个接口
	//例如结构体stu 实现了A以及B
	var C BInetrface = stu
	C.Hello()

	//6.接口中不能有任何变量

	//7.一个接口可以继承别的接口，如果要实现这个接口就要将他继承的所有接口实现
	//例子可以看看face3

	//8.interface 是一个指针（引用类型），如果没有初始化就使用是错误的

	//9.空接口没有任何方法，所有方法都引用了它

}
