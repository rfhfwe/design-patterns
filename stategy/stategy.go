package main

import "fmt"

type EvictionAlgo interface {
	evict(c *Cache)
}

type Fifo struct {
}

type Lru struct {
}

type Lfu struct {
}

func (l *Fifo) evict(c *Cache) {
	fmt.Println("Fifo")
}

func (l *Lru) evict(c *Cache) {
	fmt.Println("Lru")
}

func (l *Lfu) evict(c *Cache) {
	fmt.Println("Lfu")
}
