package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().Unix()))

	length := 1000

	ctx, cancel := context.WithCancel(context.Background())

	pool := NewPool()
	pool.Start(ctx)

	for i := 0; i < length; i++ {
		pool.receiver <- rand.Intn(13)
	}
	close(pool.receiver)
	cancel()

	pool.wg.Wait()
}

type Pool struct {
	broadcast chan int
	receiver  chan int
	wg        *sync.WaitGroup
}

func NewPool() *Pool {
	return &Pool{
		broadcast: make(chan int),
		receiver:  make(chan int),
		wg:        new(sync.WaitGroup),
	}
}

func (p *Pool) Start(ctx context.Context) {
	for i := 0; i < 5; i++ {
		w := NewWorker(p.broadcast, p.wg)
		w.Start(ctx)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(p.broadcast)
				return
			case t := <-p.receiver:
				p.broadcast <- t
			}
		}
	}()
}

type Worker struct {
	broadcast chan int
	wg        *sync.WaitGroup
}

func NewWorker(broadcast chan int, wg *sync.WaitGroup) *Worker {
	return &Worker{broadcast: broadcast, wg: wg}
}

func (w *Worker) Start(ctx context.Context) {
	w.wg.Add(1)

	go func() {
		defer w.wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case t := <-w.broadcast:
				go func() {
					log.Println(t, "is", factorial(t))
					time.Sleep(time.Second)
				}()
			}
		}
	}()
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * factorial(n-1)
}
