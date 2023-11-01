/*
	如果存在多个 channel 的时候，我们该如何操作呢，

Go 里面提供了一个关键字 select，通过 select 可以监听 channel 上的数据流动。
select 默认是阻塞的，只有当监听的 channel 中有发送或接收可以进行时才会运行，
当多个 channel 都准备好的时候，select 是随机的选择一个执行的。
*/
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
			//只有监听到quit上面有数据流动的时候，才会执行
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

/* 在 select 里面还有 default 语法，
elect 其实就是类似 switch 的功能，
default 就是当监听的 channel 都没有准备好的时候，
默认执行的（select 不再阻塞等待 channel）
*/
/*

select {
	case i := <-c:
    // use i
	default:
    // 当 c 阻塞的时候执行这里
}
*/
