package main

import (
	"fmt"
	"packagestudy/struct/test3/model"
)

//工厂模式进行实例化

func main() {
	s := model.NewStudent("cui", 16, 90.21)
	s.SetName("cui")

	s.SetAge(78)
	s.SetScore(90.32)
	fmt.Println(s.GetName(), s.GetAge(), s.GetScore())

}
