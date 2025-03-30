package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// testLock()
	testProducerConsumer()
	for {
	}
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
