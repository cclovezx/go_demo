package main

//多态数组体现多态

import "fmt"

type USB interface {
	Start()
	End()
}

type Phone struct{}
type Camera struct{}
type Computer struct{}

func (p *Phone) Start() {
	fmt.Println("手机连接上了")
}
func (p *Phone) End() {
	fmt.Println("手机拔出来了")
}
func (p Phone) Call() {
	fmt.Print("手机开始打电话")
}
func (c *Camera) Start() {
	fmt.Println("相机连接上了")
}
func (c *Camera) End() {
	fmt.Println("相机拔出来了")
}
func (c *Computer) Start() {
	fmt.Println("电脑连接上了")
}
func (c *Computer) End() {
	fmt.Println("电脑拔出来了")
}

//重点来了
func (c Computer) Work(usb USB) {
	usb.Start()
	//类型断言判断
	if phone, ok := usb.(*Phone); ok == true {
		phone.Call()
	}
	usb.End()
}

func main() {
	//定义一个USb接口的数组可以存放结构体变量
	//这里就体现了多态数组
	var usbArr [3]USB
	//不使用多态的话根本不可能装这么多类型
	usbArr[0] = &Phone{}
	usbArr[1] = &Camera{}
	usbArr[2] = &Computer{}
	fmt.Print(usbArr)

	var computer Computer
	for _, v := range usbArr {
		computer.Work(v)
	}
}
