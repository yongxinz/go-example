package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	sync.RWMutex
	Map map[string]string
}

func NewSafeMap() *SafeMap {
	sm := new(SafeMap)
	sm.Map = make(map[string]string)
	return sm
}

func (sm *SafeMap) ReadMap(key string) string {
	sm.RLock()
	value := sm.Map[key]
	sm.RUnlock()
	return value
}

func (sm *SafeMap) WriteMap(key string, value string) {
	sm.Lock()
	sm.Map[key] = value
	sm.Unlock()
}

func main() {
	safeMap := NewSafeMap()

	var wg sync.WaitGroup

	// 启动多个goroutine进行写操作
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			safeMap.WriteMap(fmt.Sprintf("name%d", i), fmt.Sprintf("John%d", i))
		}(i)
	}

	wg.Wait()

	// 启动多个goroutine进行读操作
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(safeMap.ReadMap(fmt.Sprintf("name%d", i)))
		}(i)
	}

	wg.Wait()
}
