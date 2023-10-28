package main

import (
	interfacetest1_test "packagestudy/interface/face1/interface"
)

func main() {
	//简单想象一下就是把手机插进电脑
	s := interfacetest1_test.Phone{}

	camera := interfacetest1_test.Camera{}

	computer := interfacetest1_test.Computer{}
	computer.Work(&s)
	computer.Work(&camera)
}
