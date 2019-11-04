package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 通过异步来阻塞主线程
func Syn() {
	group := sync.WaitGroup{}
	group.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("GO2-", i)
			group.Done()
		}()
	}
	group.Wait()
}

// 通道缓存, 功能类似于 java 的 CountDownLatch
func Channel() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//c := make(chan bool, 1)
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		var value = i
		go func() {
			fmt.Println("GO1-", value)
			time.Sleep(2 * time.Second)
			c <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-c
	}
}

func main() {

	Channel()
	Syn()

	// TODO select 的使用（用于阻塞主线程）

}
