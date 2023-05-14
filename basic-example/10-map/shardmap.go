package main

import (
	"fmt"
	"sync"
)

const N = 16

type SafeMap struct {
	maps  [N]map[string]string
	locks [N]sync.RWMutex
}

func NewSafeMap() *SafeMap {
	sm := new(SafeMap)
	for i := 0; i < N; i++ {
		sm.maps[i] = make(map[string]string)
	}
	return sm
}

func (sm *SafeMap) ReadMap(key string) string {
	index := hash(key) % N
	sm.locks[index].RLock()
	value := sm.maps[index][key]
	sm.locks[index].RUnlock()
	return value
}

func (sm *SafeMap) WriteMap(key string, value string) {
	index := hash(key) % N
	sm.locks[index].Lock()
	sm.maps[index][key] = value
	sm.locks[index].Unlock()
}

func hash(s string) int {
	h := 0
	for i := 0; i < len(s); i++ {
		h = 31*h + int(s[i])
	}
	return h
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
