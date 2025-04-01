package main

import (
	"fmt"
	"sync"
)

func main() {
	// test()
	testChangeMap()
}

func test() {
	m := sync.Map{}
	str := "abcabcd"
	for _, char := range str {
		cnt, loadFlag := m.Load(char)
		if !loadFlag {
			m.Store(char, 1)
		} else {
			m.Store(char, cnt.(int)+1)
		}
	}

	m.Range(func(key any, value any) bool {
		fmt.Printf("%c -> %d\n", key, value)
		return true
	})
}

var sm = sync.Map{}

func testChangeMap() {
	size := 2
	wg := sync.WaitGroup{}
	wg.Add(size)
	for i := 0; i < size; i++ {
		i := i
		go func() {
			defer wg.Done()
			changeMap(i)
		}()
	}
	wg.Wait()

	sm.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}

func changeMap(key int) {
	sm.Store(key, 1)
}
