package fake

import "github.com/xxxVitoxxx/JiaPen-backend/pkg/message"

// QueryMessageRepo _
type QueryMessageRepo struct {
	QueryMessages []message.QueryMessage
}

// NewQueryMessageRepo _
func NewQueryMessageRepo() *QueryMessageRepo {
	return &QueryMessageRepo{}
}

// FindMessages _
func (q *QueryMessageRepo) FindMessages() ([]message.QueryMessage, error) {
	return q.QueryMessages, nil
}
