package main

import "fmt"

type ChatGPT struct {
}

type WenXin struct {
}

type XunFei struct {
}

func (c *ChatGPT) Chat() {
	fmt.Println("chatGPT chat...")
}

func (w *WenXin) ChatStream() {
	fmt.Println("wen xin chat...")
}

func (x *XunFei) ChatWithBot() {
	fmt.Println("xun fei chat...")
}
