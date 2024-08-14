package main

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"time"
)

var errorTimeout = errors.New("request timeout")

type response struct {
	value int
	err   error
}

func main() {
	rand.New(rand.NewSource(time.Now().Unix()))

	var (
		respChan    = make(chan int)
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*2)
	)
	go call(respChan)

	select {
	case <-ctx.Done():
		log.Println(errorTimeout.Error())
	case v := <-respChan:
		log.Println("response:", v)
	}

	cancel()
}

func main_() {
	rand.New(rand.NewSource(time.Now().Unix()))

	var (
		respChan    = make(chan response)
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*2)
	)
	go call_(ctx, respChan)

	v := <-respChan
	log.Println(v.value, v.err)

	cancel()
}

func call(resp chan int) {
	time.Sleep(time.Second * time.Duration(rand.Intn(4)))

	resp <- rand.Intn(10)
}

func call_(ctx context.Context, resp chan response) {
	select {
	case <-ctx.Done():
		resp <- response{
			err: errorTimeout,
		}
	case <-time.After(time.Second * time.Duration(rand.Intn(4))):
		resp <- response{
			value: rand.Intn(10),
		}
	}
}
