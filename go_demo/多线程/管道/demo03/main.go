/* Go 也允许指定 channel 的缓冲大小，很简单，就是 channel 可以存储多少元素.
ch:= make (chan bool, 4)，创建了可以存储 4 个元素的 bool 型 channel。
在这个 channel 中，前 4 个元素可以无阻塞的写入。
当写入第 5 个元素时，代码将会阻塞，直到其他 goroutine 从 channel 中读取一些元素，腾出空间。
*/
//ch := make(chan type, value)

package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// 如果想要停止线程，可以使用内置函数
// close
func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
