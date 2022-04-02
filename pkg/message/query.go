package message

import "time"

// QueryRepository _
type QueryRepository interface {
	FindMessages() ([]QueryMessage, error)
}

// QueryMessage _
type QueryMessage struct {
	Id        uint      `json:"id"`
	User      string    `json:"user"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
