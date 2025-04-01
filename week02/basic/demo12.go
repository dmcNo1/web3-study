package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// testLock()
	// testProducerConsumer()
	// for {
	// }
	// testRWLock()
	testCond()
}

var lock = sync.Mutex{}

func testLock() {
	go func() { printer("hello") }()
	go func() { printer("world") }()
	for {
	}
}

func printer(str string) {
	// 加锁
	lock.Lock()
	for _, char := range str {
		fmt.Printf("%c", char)
		time.Sleep(time.Millisecond * 500)
	}
	// 释放锁
	lock.Unlock()
}

var sce = make([]int, 10)

func testProducerConsumer() {
	go consume()
	go produce()
}

func produce() {
	lock.Lock()
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		sce[i] = num
		fmt.Println("producer produce num:", num)
		time.Sleep(time.Millisecond * 500)
	}
	lock.Unlock()
}

func consume() {
	lock.Lock()
	for i := 0; i < 10; i++ {
		num := sce[i]
		fmt.Println("consumer consume num:", num)
	}
	lock.Unlock()
}

type counter struct {
	count  int
	rwLock sync.RWMutex // 读写锁
}

func (c *counter) getCount() int {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	return c.count
}

func (c *counter) add() {
	c.rwLock.Lock()
	defer c.rwLock.Unlock()
	c.count++
}

// Mutex和RWMutex都不是递归锁，不可重入
func testRWLock() {
	wg := sync.WaitGroup{}
	size := 100
	wg.Add(size)
	counter := counter{count: 0, rwLock: sync.RWMutex{}}
	for i := 0; i < size; i++ {
		go func() {
			defer wg.Done()
			counter.getCount()
			counter.add()
		}()
	}

	wg.Wait()
	fmt.Println("count =", counter.count)
}

// cond，可以理解为Java中的Condition
func testCond() {
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}
	condition := sync.NewCond(&lock)

	size := 10
	wg.Add(size + 1)
	for i := 0; i < size; i++ {
		go func() {
			defer wg.Done()
			condition.L.Lock()
			fmt.Printf("%d ready\n", i)
			// Wait实际上是会先解锁condition.L，再阻塞当前goroutine
			// 这样其它goroutine调用上面的cond.L.Lock()才能加锁成功，才能进一步执行到Wait方法，
			// 等待被Broadcast或者signal唤醒。
			// Wait被Broadcast或者Signal唤醒的时候，会再次对condition.L加锁，加锁后Wait才会return
			// 简而言之，Wait会释放锁资源，被唤醒之后，当前协程又会再次获取到condition.L的锁资源
			condition.Wait()
			fmt.Printf("%d done\n", i)
			condition.L.Unlock()
		}()
	}

	time.Sleep(time.Second * 2)
	go func() {
		defer wg.Done()
		condition.Broadcast()
	}()
	wg.Wait()
}
