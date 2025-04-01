package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// work()
	once()
}

func work() {
	size := 10
	// wg会等待size个数的协程执行完之后才往下执行，类似于CountDownLatch
	wg := sync.WaitGroup{}
	wg.Add(size)
	for i := 0; i < size; i++ {
		go func() {
			// 执行完之后，wg的计数器-1
			defer wg.Done()
			fmt.Printf("work-%d starting\n", i)
			time.Sleep(time.Second)
			fmt.Printf("work-%d done\n", i)
		}()
	}
	// 阻塞，直到wg的计数器为0
	wg.Wait()
	fmt.Println("end!")
}

func once() {
	size := 10
	wg := sync.WaitGroup{}
	wg.Add(size)
	// sync.Once只有一个对外暴露的方法Do，只能执行一次，多用于实现单例
	once := sync.Once{}
	print := func() { fmt.Println("once invoke") }
	for i := 0; i < size; i++ {
		go func() {
			// 执行完之后，wg的计数器-1
			defer wg.Done()
			once.Do(print)
		}()
	}
	// 阻塞，直到wg的计数器为0
	wg.Wait()
	fmt.Println("end!")
}
