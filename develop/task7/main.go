package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(3*time.Second),
		sig(4*time.Second),
		sig(1*time.Second),
	)

	fmt.Printf("fone after %v\n", time.Since(start))
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	trigger := sync.WaitGroup{}
	result := make(chan interface{})
	trigger.Add(len(channels))
	for _, channel := range channels {
		go func(channel <-chan interface{}) {
			for ok := range channel {
				result <- ok
			}
			defer trigger.Done()
			defer fmt.Println("i'm done")
		}(channel)
	}

	go func() {
		trigger.Wait()
		defer close(result)
	}()
	// trigger.Wait()
	// defer close(result)
	fmt.Println("done")
	return result
}
