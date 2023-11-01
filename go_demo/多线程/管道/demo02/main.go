package main

import "fmt"

/* Go 提供了一个很好的通信机制 channel。channel 可以与 Unix shell 中的双向管道做类比：
可以通过它发送或者接收值。这些值只能是特定的类型：
 channel 类型。定义一个 channel 时，也需要定义发送到 channel 的值的类型。
 注意，必须使用 make 创建 channel：
*/
/*
ci := make(chan int)
cs := make(chan string)
cf := make(chan interface{})
*/

//channel 通过操作符 <- 来接收和发送数据
/*
ch <- v     发送 v 到 channel ch.
v := <-ch   从 ch 中接收数据，并赋值给v
*/
func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
