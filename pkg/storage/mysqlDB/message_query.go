package mysqlDB

import (
	"github.com/xxxVitoxxx/JiaPen-backend/pkg/message"
	"gorm.io/gorm"
)

// QueryMessageRepo _
type QueryMessageRepo struct {
	db *gorm.DB
}

// NewQueryMessageRepo will return instance of QueryMessageRepo
func NewQueryMessageRepo(db *gorm.DB) message.QueryRepository {
	return &QueryMessageRepo{db}
}

// FindMessages will find all of the information of the message
func (q *QueryMessageRepo) FindMessages() ([]message.QueryMessage, error) {
	messages := []message.QueryMessage{}
	err := q.db.Model(Message{}).Scan(&messages).Error

	return messages, err
}
