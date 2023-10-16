package main

func main() {
	//cli := Client{}
	//mac := &Mac{}
	//
	//cli.InsertLightningConnectorIntoComputer(mac)
	//
	//win := &Windows{}
	//winAdapter := &WindowsAdapter{
	//	windowsMachine: win,
	//}
	//cli.InsertLightningConnectorIntoComputer(winAdapter)

	// chat adapter test...
	Init() // 初始化
	// 调用chat业务服务
	chatService(gpt)
	chatService(winAdapter)
	chatService(xunAdapter)
}

var (
	gpt        *ChatGPT
	chatSrv    *ChatService
	wen        *WenXin
	winAdapter *ChatAdapter
	xunfei     *XunFei
	xunAdapter *XunFeiChatAdapter
)

// Init Bean初始化 Spring容器注解 go的wire等等
func Init() {
	chatSrv = &ChatService{}
	gpt = &ChatGPT{}
	wen = &WenXin{}
	xunfei = &XunFei{}
	winAdapter = &ChatAdapter{wen}
	xunAdapter = &XunFeiChatAdapter{xunfei}
}

// 模拟业务逻辑
func chatService(m Model) {
	// 业务逻辑...
	m.Chat()
	// 业务逻辑...
}
