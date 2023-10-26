package main

import (
	"fmt"
	interfaced "packagestudy/interface/test01/interface_01"
)

type A interface {
	Say()
}

type integer int

func (i *integer) Say() {
	fmt.Print("i说他等于:", *i)
}

type B interface {
	hello()
}
type C interface {
	bye()
}

type monster struct {
}

func (m *monster) hello() {
	fmt.Print("我是B接口的HELLO")
}

func (m *monster) bye() {
	fmt.Print("我是C方法的bye")
}

func main() {
	//另外就是，go里面的接口都是隐式实现的
	//但是java就是现实的用的impliment
	//1.接口不能自己实现，必须和对象实例绑定使用
	//2.接口中的方法都没有实现体
	//3.实现接口的意思是实现接口的所有方法
	phone := interfaced.Phone{}
	phone.Start()
	fmt.Println()
	phone.Stop()

	//4.一个自定义类型只有实现了某个接口，才能将该自定义类型的实例赋值给接口类型
	//5.只要是自定义的数据类型，就可以实现接口，不只是结构体
	var i integer = 1
	var b A = &i
	b.Say()
	fmt.Println()

	//6.一个自定义类型可以实现多个接口
	//7.接口中不可以有任何变量
	//8.接口也可以继承别的的接口
	mon := monster{}
	mon.hello()
	fmt.Println()
	mon.bye()
}
