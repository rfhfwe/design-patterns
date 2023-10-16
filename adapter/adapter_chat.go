package main

type Model interface {
	Chat()
}

type ChatService struct {
}

func (c *ChatService) ChatSrv(m Model) {
	m.Chat()
}

type ChatAdapter struct {
	wenxin *WenXin
}

func (c *ChatAdapter) Chat() {
	// 自己写转换逻辑
	c.wenxin.ChatStream()
}

type XunFeiChatAdapter struct {
	xunFei *XunFei
}

// Chat 适配器的Chat方法，本质上是在此做了sdk差异的转换
func (x *XunFeiChatAdapter) Chat() {
	// 差异转换
	x.xunFei.ChatWithBot()
}
