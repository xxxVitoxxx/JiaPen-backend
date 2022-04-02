package fake

import (
	"github.com/xxxVitoxxx/JiaPen-backend/pkg/message"
)

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

// UpdateMessage _
func (m *MessageRepo) UpdateMessage(id uint, content message.UpdateMessage) error {
	message := m.Messages[int(id)]
	message.Content = content.Content
	m.Messages[int(id)] = message

	return nil
}
