package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var (
		unique   = make(map[int]struct{})
		capacity = 1000

		wg    = &sync.WaitGroup{}
		mutex = &sync.Mutex{}

		uniqueChan = make(chan int, capacity)
	)

	input := make([]int, 0, capacity)
	for i := 0; i < capacity; i++ {
		input = append(input, rand.Intn(10))
	}

	wg.Add(capacity)
	for _, v := range input {
		v := v
		go func() {
			defer wg.Done()
			defer mutex.Unlock()
			mutex.Lock()

			if _, ok := unique[v]; !ok {
				unique[v] = struct{}{}

				uniqueChan <- v
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("unique values from chan")
		for i := 0; i < len(unique); i++ {
			fmt.Print(<-uniqueChan)
		}
		fmt.Println()
	}()

	wg.Wait()
	close(uniqueChan)

	fmt.Println("map: ", unique)
	fmt.Println("map len: ", len(unique))

}

// concurrency for by chan???
