package main

var (
	lfu  *Lfu
	lru  *Lru
	fifo *Fifo

	cache *Cache
)

func Init() {
	lfu = &Lfu{}
	lru = &Lru{}
	fifo = &Fifo{}
	cache = initCache(lfu)
}

func main() {
	Init()
	cache.add("a", "1")
	cache.add("b", "2")
	cache.setEvictionAlgo(lru)
	cache.add("c", "3")
	cache.add("d", "4")
	cache.setEvictionAlgo(fifo)
	cache.add("e", "5")

}
