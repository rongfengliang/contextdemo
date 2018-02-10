package main

import (
	"context"
	"log"
	"sync"
	"time"
)

type user struct {
	name  string
	appid string
}

const (
	// MAXSIZE set max run goroutine
	MAXSIZE = 500000
)

var (
	wg sync.WaitGroup
)

func main() {

	wg.Add(MAXSIZE)
	ctx := context.Background()
	ctx2, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	for i := 0; i < MAXSIZE; i++ {
		go func() {
			v := login(ctx2, user{name: "dalong", appid: "dd"})
			log.Println(v)
		}()
	}
	wg.Wait()
}

func login(ctx context.Context, v interface{}) (result interface{}) {
	//TODO
	defer wg.Done()
	//log.Println(v)
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		//log.Println(ctx.Err())
		//wg.Done()
		result = ctx.Err()
	default:
		result = v
		//wg.Done()
	}
	return
}
