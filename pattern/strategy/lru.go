package main

import "fmt"

type lru struct{}

func (l *lru) evict(c *cache) {
	fmt.Println("Eviction by lru strategy")
}
