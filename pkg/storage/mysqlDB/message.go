package mysqlDB

import (
	"time"

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

type Message struct {
	Id        uint      `gorm:"auto_increment;primary_key" json:"id"`
	User      string    `gorm:"size:50;not null" json:"user"`
	Content   string    `gorm:"size:300;not null" json:"content"`
	CreatedAt time.Time `gorm:"type:datetime(0)" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime(0)" json:"updated_at"`
}

// CreateMessage _
func (m *MessageRepo) CreateMessage(input message.Message) error {
	msg := &Message{
		User:    input.User,
		Content: input.Content,
	}

	result := m.db.Create(msg)

	return result.Error
}
