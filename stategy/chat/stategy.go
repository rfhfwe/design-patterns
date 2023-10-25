package chat

import "context"

type ChatHandler interface {
	Chat(content string, ctx context.Context) error
}
