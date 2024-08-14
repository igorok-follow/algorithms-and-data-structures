package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
)

func main_() {
	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)

	ch1 <- 1
	ch1 <- 2
	ch2 <- 3
	ch2 <- 4
	close(ch1)
	close(ch2)

	result := mergeChans[int](ch1, ch2)
	for v := range result {
		log.Println(v)
	}
}

func mergeChans[T any](chans ...chan T) chan T {
	result := make(chan T)
	wg := sync.WaitGroup{}

	for _, ch := range chans {
		ch := ch
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range ch {
				result <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

func main() {
	channels := make([]<-chan int, 100)

	for i := 0; i < 100; i++ {
		ch := make(chan int, 2)

		ch <- rand.Intn(10)
		ch <- rand.Intn(10)
		close(ch)
		channels[i] = ch
	}

	ctx, cancel := context.WithCancel(context.Background())

	chans := fanIn(channels, cancel)
	for {
		select {
		case <-ctx.Done():
			log.Println("end")
			return
		case x, ok := <-chans:
			if ok {
				log.Println(x)
			}
		}
	}
}

func fanIn(inputs []<-chan int, cancel context.CancelFunc) <-chan int {
	c := make(chan int)
	wg := sync.WaitGroup{}

	for _, v := range inputs {
		v := v
		wg.Add(1)
		go func() {
			defer wg.Done()
			c <- <-v
		}()
	}

	go func() {
		wg.Wait()
		cancel()
		close(c)
	}()

	return c
}
