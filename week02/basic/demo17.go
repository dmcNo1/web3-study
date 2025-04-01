package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// atomic包的使用，底层原理都是基于CPU的CAS操作来完成的
func main() {
	testSwap()
	testCompareAndSwap()
	testAdd()
	testStore()
	testValue()
}

func testSwap() {
	var dst int32 = 100
	var newVal int32 = 200
	old := atomic.SwapInt32(&dst, newVal)
	fmt.Printf("oldVal = %d, newVal = %d\n", old, dst)
}

func testCompareAndSwap() {
	var dst int32 = 100
	var newVal int32 = 200
	// 返回dst指向的值
	oldVal := atomic.LoadInt32(&dst)
	// CAS，只有dst等于oldVal的时候才会交换成功，返回值为交换结果
	swapped := atomic.CompareAndSwapInt32(&dst, oldVal, newVal)
	fmt.Printf("swapped result = %v, old value = %d, new value = %d\n", swapped, oldVal, newVal)
}

func testAdd() {
	var sum int32 = 0
	max := 100
	wg := sync.WaitGroup{}
	wg.Add(max)
	for i := 1; i <= max; i++ {
		go func(num int32) {
			defer wg.Done()
			atomic.AddInt32(&sum, num)
		}(int32(i))
	}
	wg.Wait()
	fmt.Printf("sum = %d\n", sum)
}

func testStore() {
	var dst int32 = 100
	var newVal int32 = 200
	// 把dst地址的值修改成newVal
	atomic.StoreInt32(&dst, newVal)
	fmt.Printf("value after store = %d\n", dst)
}

func testValue() {
	// config变量用来存放该服务的配置信息
	config := atomic.Value{}
	// 初始化时从别的地方加载配置文件，并存到config变量里
	config.Store(loadConfig())
	go func() {
		for {
			// 每10s更新一次配置
			time.Sleep(time.Second * 10)
			config.Store(loadConfig())
		}
	}()

	for i := 0; i < 10; i++ {
		go func() {
			for r := range requests() {
				// 对应于取值操作 c := config
				// 由于Load()返回的是一个interface{}类型，所以我们要先强制转换一下
				c := config.Load().(map[string]string)
				// 这里是根据配置信息处理请求的逻辑...
				_, _ = r, c
			}
		}()
	}
}

func loadConfig() map[string]string {
	// 从数据库或者文件系统中读取配置信息，然后以map的形式存放在内存里
	return make(map[string]string)
}

func requests() chan int {
	// 将从外界中接收到的请求放入到channel里
	return make(chan int)
}
