package main

import (
	"fmt"
	"sort"
)

func main() {
	//创建方式1
	var a = make(map[int]int)
	fmt.Println(a)

	//创建方式2
	b := make(map[string]int)
	fmt.Println(b)

	//创建方式3，代加赋值操作
	d := map[string]int{"cui": 1, "shuai": 2}
	fmt.Println(d)

	//也可以new,但是容易空指针异常，不建议使用，这里就有不写了
	//例如： a :=new(map[int]int)

	//数组容器
	//make(map[type]type,capital)
	c := make(map[int]int, 10)
	fmt.Println(c)

	//map也可以被定义为切片
	//简单的对比一下，map的创建和数组的差不了太多
	//所以可以简单的可以画等号

	//map集合相关操作
	//1.根据key寻找value
	val, ok := d["cui"]
	if ok {
		fmt.Println(val)
	}

	//2.删除操作
	delete(d, "shuai")
	fmt.Println(d)

	//3.遍历
	for j, i := range d {
		fmt.Println(j, i)
	}

	//4.获取值
	fmt.Println(d["cui"])

	//实现map的键值互换
	//简单来说就是创建一个新的切片，反过来赋值就可以了

	//5.切片排序
	barVal := map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}

	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	//排序没什么高级，将键值其中一个封装到数组里面
	//然后对数组进行排序
	for _, i := range keys {
		s := barVal[i]
		fmt.Println(s)
	}
}
