package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	addresses := map[int]string{
		0: "1.1.1.1",
		1: "8.8.8.8",
		2: "8.8.4.4",
		3: "192.168.0.1",
	}

	handlers := map[string]int{
		"get_route":  5,
		"post_route": 5,
	}
	whiteList := map[string]struct{}{
		"192.168.0.1": {},
	}

	ctx, cancel := context.WithCancel(context.Background())
	r := NewRateLimiter(handlers, whiteList, time.Second, time.Second)
	r.Start(ctx)

	for i := 0; i < 100; i++ {
		addr := addresses[rand.Intn(len(addresses))]
		log.Println(addr, r.Limit(addr, "get_route"))
		time.Sleep(time.Millisecond * 20)
	}

	cancel()
}

type Frame struct {
	counter  int
	lifetime int64
}

type RateLimiter struct {
	frame     map[string]*Frame
	handlers  map[string]int
	whiteList map[string]struct{}

	gcInterval time.Duration
	lifetime   time.Duration

	mu sync.RWMutex
}

func NewRateLimiter(handlers map[string]int, whiteList map[string]struct{}, lifetime, gcInterval time.Duration) *RateLimiter {
	return &RateLimiter{
		frame:     make(map[string]*Frame),
		handlers:  handlers,
		whiteList: whiteList,

		gcInterval: gcInterval,
		lifetime:   lifetime,

		mu: sync.RWMutex{},
	}
}

func (r *RateLimiter) Start(ctx context.Context) {
	gc := NewGarbageCollector(r.gcInterval, r)
	gc.Start(ctx)
}

func (r *RateLimiter) GetOutdatedKeys() []string {
	defer r.mu.Unlock()
	r.mu.Lock()

	keys := make([]string, 0)
	for k, v := range r.frame {
		if v.lifetime <= time.Now().Unix() {
			keys = append(keys, k)
		}
	}

	return keys
}

func (r *RateLimiter) Flush(keys []string) {
	defer r.mu.Unlock()
	r.mu.Lock()

	for _, v := range keys {
		delete(r.frame, v)
	}
}

func (r *RateLimiter) Limit(ip, handler string) bool {
	defer r.mu.Unlock()
	r.mu.Lock()

	var ok bool
	if _, ok = r.whiteList[ip]; ok {
		return true
	}

	var ttlLimit int
	if ttlLimit, ok = r.handlers[handler]; !ok {
		return false
	}

	if _, ok = r.frame[ip+handler]; !ok {
		r.frame[ip+handler] = &Frame{
			lifetime: time.Now().Add(r.lifetime).Unix(),
		}
		return true
	}

	if r.frame[ip+handler].counter > ttlLimit {
		return false
	}

	r.frame[ip+handler].counter += 1

	return true
}

type Gc struct {
	interval    time.Duration
	rateLimiter *RateLimiter
}

func NewGarbageCollector(interval time.Duration, limiter *RateLimiter) *Gc {
	return &Gc{
		interval:    interval,
		rateLimiter: limiter,
	}
}

func (g *Gc) Start(ctx context.Context) {
	go func() {
		for {
			select {
			case <-time.After(g.interval):
				keys := g.rateLimiter.GetOutdatedKeys()
				g.rateLimiter.Flush(keys)
				log.Println("flushed")
			case <-ctx.Done():
				return
			}
		}
	}()
}
