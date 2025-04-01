package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Unix()) // 单位为秒
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	// 格式化
	fmtTime := now.Format("2006/01/02 15:03:04")
	fmt.Println(fmtTime)

	// sleep
	time.Sleep(time.Second * 2)
	fmt.Println("sleep over")

	// 随机数，范围[0,10)
	rand.Seed(now.UnixNano())
	fmt.Println(rand.Intn(10))
}
