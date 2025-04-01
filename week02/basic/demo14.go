package main

import (
	"fmt"
	"time"
)

func main() {
	// testAfter()
	testTricker()
}

// 一次性定时器
func testAfter() {
	start := time.Now()
	fmt.Println("start time:", start)
	// 底层就是timer.C，return NewTimer(d).C
	timer := time.After(time.Second * 3)
	fmt.Println("code before time end")
	end := <-timer
	fmt.Println("end time:", end)
	fmt.Println("code after time end")
}

// 周期性定时器
// ticker存储的数据就是time
//
//	type Ticker struct {
//		C          <-chan Time // The channel on which the ticks are delivered.
//		initTicker bool
//	}
func testTricker() {
	ticker := time.NewTicker(time.Second)
	for {
		// 周期性获取，但是在系统写入数据之前会阻塞
		t := <-ticker.C
		fmt.Println(t)
	}
}
