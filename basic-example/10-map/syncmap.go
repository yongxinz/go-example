package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	var wg sync.WaitGroup

	// 启动多个goroutine进行写操作
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Store(fmt.Sprintf("name%d", i), fmt.Sprintf("John%d", i))
		}(i)
	}

	wg.Wait()

	// 启动多个goroutine进行读操作
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			v, _ := m.Load(fmt.Sprintf("name%d", i))
			fmt.Println(v.(string))
		}(i)
	}

	wg.Wait()
}
