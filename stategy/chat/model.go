package chat

import (
	"context"
	"fmt"
)

type ChatGPT struct {
}

type WenXin struct {
}

type XingHuo struct {
}

func (c *ChatGPT) Chat(content string, ctx context.Context) error {
	fmt.Println("chatGPT")
	return nil
}

func (c *WenXin) Chat(content string, ctx context.Context) error {
	fmt.Println("wenxin")
	return nil
}

func (c *XingHuo) Chat(content string, ctx context.Context) error {
	fmt.Println("xinghuo")
	return nil
}
