package interfaced

import "fmt"

type USB interface {
	Start()
	Stop()
}

type Phone struct {
}

func (p *Phone) Start() {
	fmt.Print("手机启动了！")
}

func (p *Phone) Stop() {
	fmt.Print("手机关闭了！")
}
