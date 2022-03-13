package fake

import "github.com/xxxVitoxxx/JiaPen-backend/pkg/message"

// MessageRepo _
type MessageRepo struct {
	MessageId int
	Messages  map[int]message.Message
}

// NewMessageRepo _
func NewMessageRepo() *MessageRepo {
	return &MessageRepo{
		MessageId: 0,
		Messages:  make(map[int]message.Message),
	}
}

// CreateMessage _
func (m *MessageRepo) CreateMessage(input message.Message) error {
	m.MessageId++
	m.Messages[m.MessageId] = input
	return nil
}
