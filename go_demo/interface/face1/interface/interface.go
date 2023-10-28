package interfacetest1_test

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
	usb.End()
}
