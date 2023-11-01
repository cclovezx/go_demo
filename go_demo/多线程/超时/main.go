package main

import "time"

/* 有时候会出现 goroutine 阻塞的情况，那么我们如何避免整个程序进入阻塞的情况呢？我们可以利用 select 来设置超时，通过如下的方式实现： */
func main() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				println(v)
			case <-time.After(5 * time.Second):
				println("timeout")
				o <- true
				// 此处的break只是跳出了select循环，并未终止for循环，要用return才能终止这个子进程
				break
			}
		}
	}()
	<-o
}

/* runtime 包中有几个处理 goroutine 的函数：

Goexit

退出当前执行的 goroutine，但是 defer 函数还会继续调用

Gosched

让出当前 goroutine 的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行。

NumCPU

返回 CPU 核数量

NumGoroutine

返回正在执行和排队的任务总数

GOMAXPROCS

用来设置可以并行计算的 CPU 核数的最大值，并返回之前的值
*/
