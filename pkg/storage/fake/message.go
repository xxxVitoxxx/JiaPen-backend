package fake

import (
	"github.com/xxxVitoxxx/JiaPen-backend/pkg/message"
)

// MessageRepo _
type MessageRepo struct {
	MessageId int
	Messages  map[int]*message.Message
}

// NewMessageRepo _
func NewMessageRepo() *MessageRepo {
	return &MessageRepo{
		MessageId: 0,
		Messages:  make(map[int]*message.Message),
	}
}

// CreateMessage _
func (m *MessageRepo) CreateMessage(input message.Message) error {
	m.MessageId++
	m.Messages[m.MessageId] = &input
	return nil
}

// UpdateMessage _
func (m *MessageRepo) UpdateMessage(id uint, content message.UpdateMessage) error {
	// map 的 struct 用指標就可以直接改值
	m.Messages[1].Content = content.Content

	return nil
}

// DeleteMessage _
func (m *MessageRepo) DeleteMessage(id uint) error {
	delete(m.Messages, int(id))

	return nil
}
