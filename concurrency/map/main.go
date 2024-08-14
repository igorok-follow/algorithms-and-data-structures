package main

import (
	"log"
	"sync"
)

func main() {
	var (
		length = 1000
		data   = make(map[int]int, length)
		mutex  sync.RWMutex
	)

	// concurrency write error
	wg := &sync.WaitGroup{}
	wg.Add(length)
	for i := 0; i < length; i++ {
		i := i
		go func() {
			defer wg.Done()

			defer mutex.Unlock()
			mutex.Lock()
			data[i] = i
		}()
	}

	// concurrency read write error
	wg.Add(length)
	for i := 0; i < length; i++ {
		i := i
		go func() {
			defer wg.Done()

			defer mutex.RUnlock()
			mutex.RLock()
			log.Println(data[i])
		}()
	}

	wg.Wait()
}
