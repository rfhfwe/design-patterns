package main

import (
	"context"
	"go.mod/stategy/chat"
)

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
	// 测试 chat
	Chat()
}

func Chat() {
	chatgpt := &chat.ChatGPT{}
	wenxin := &chat.WenXin{}
	xingHuo := &chat.XingHuo{}
	facotry := chat.InitFactory()
	facotry.Register("gpt", chatgpt)
	facotry.Register("wenxin", wenxin)
	facotry.Register("xinghuo", xingHuo)

	chat := facotry.GetHandler("wenxin")
	chat.Chat("123", context.Background())
}
