package chat

type ChatFacotry struct {
	chatMap map[string]ChatHandler
}

func (f *ChatFacotry) GetHandler(chatType string) ChatHandler {
	return f.chatMap[chatType]
}

func InitFactory() *ChatFacotry {
	return &ChatFacotry{
		chatMap: make(map[string]ChatHandler, 3),
	}
}

func (f *ChatFacotry) Register(chatType string, handler ChatHandler) error {
	f.chatMap[chatType] = handler
	return nil
}
