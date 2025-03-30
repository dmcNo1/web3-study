package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

var proCh = make(chan int, 5)
var exitCh = make(chan bool, 1)

// 通过channel来实现生产者和消费者
// channel的阻塞现象：
// 1、单独在主线程中操作channel，写满了就会报错，没有数据去读取也会报错
// 2、只有在协程中操作过channel，写满了就会阻塞，没有睡觉去读取也会阻塞
func main() {
	// testCh()
	// testChannelNoBuffer()
	// testTransferChannel()
	testSelect()
}

func testCh() {
	go produceCh()
	go consumeCh()
	<-exitCh
}

func produceCh() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		fmt.Println("producer produce num:", num)
		proCh <- num
	}
	close(proCh)
	fmt.Println("producer produce over!")
}

func consumeCh() {
	for {
		if num, ok := <-proCh; ok {
			fmt.Println("consumer consume num:", num)
		} else {
			break
		}
	}
	exitCh <- true
	fmt.Println("consumer consume over!")
}

var noBufferCh = make(chan int, 0)

// 通过无缓冲的方式使用channel，这种方式的话，容量只能为1，这样理解：
// 1、有缓冲的channel具有异步能力，写入多个读一个或者读多个
// 2、无缓冲的channel具有同步能力，读一个写一个
func testChannelNoBuffer() {
	// 只有两端同是准备好才不会报错
	go func() {
		fmt.Println(<-noBufferCh)
		fmt.Println(<-noBufferCh)
	}()
	noBufferCh <- 1
	// noBufferCh <- 2
	// 同一个协程中使用会报错
	// fmt.Println(<-noBufferCh)
	for {
	}
}

// 双向管道转换为单向管道，需要注意的是，channel之间的赋值是地址传递，channel、inChan、outChan三个管道的底层指向相同的容器
func testTransferChannel() {
	channel := make(chan int, 2)
	var inChan chan<- int
	var outChan <-chan int
	inChan = channel
	outChan = channel
	inChan <- 1
	// <-inChan，达咩
	fmt.Println(<-outChan)
	// outChan<-2，达咩
}

// 测试select关键字
// select之后必须是一个IO操作，如果当前没有满足的情况，就会陷入阻塞
// select多用于多路监听和超时处理
func testSelect() {
	myCh := make(chan int, 5)
	exitCh := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			myCh <- i
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case num := <-myCh:
				fmt.Println(num)
			case <-time.After(time.Second * 2):
				exitCh <- true
				// 结束当前的协程
				runtime.Goexit()
			}
		}
	}()

	<-exitCh
	fmt.Println("program go end")
}
