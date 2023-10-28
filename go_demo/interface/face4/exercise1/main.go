package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}

// 声明一个Hero结构体切片实现
type HeroSlice []Hero

// 实现interface接口
func (H HeroSlice) Len() int {
	return len(H)
}

// less决定你用什么标准进行排序
// 按年龄从小到大排序
func (H HeroSlice) Less(i, j int) bool {
	return H[i].Age < H[j].Age
}

// 交换
func (H HeroSlice) Swap(i, j int) {
	temp := H[i]
	H[i] = H[j]
	H[j] = temp
}

func main() {
	//简单的数组排序
	var arr = []int{1, 5, 3, 4, 2}
	sort.Ints(arr)
	fmt.Println(arr)

	//切片排序
	//测试看看可不可以对结构体切片进行排序
	var heroSlice HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		//heroSlice[i]=hero
		heroSlice = append(heroSlice, hero)
	}

	//排序前顺序
	for _, i := range heroSlice {
		fmt.Print(i, " ")
	}
	fmt.Println()
	//排序后
	sort.Sort(heroSlice)
	for _, i := range heroSlice {
		fmt.Print(i, " ")
	}

}
