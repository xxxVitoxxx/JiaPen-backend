package mysqlDB

import (
	"github.com/xxxVitoxxx/JiaPen-backend/pkg/message"
	"gorm.io/gorm"
)

// MessageRepo _
type MessageRepo struct {
	db *gorm.DB
}

// NewMessageRepo _
func NewMessageRepo(db *gorm.DB) message.Repository {
	return &MessageRepo{db}
}

// CreateMessage _
func (m *MessageRepo) CreateMessage(input message.Message) error {
	// TODO: create message
	return nil
}
