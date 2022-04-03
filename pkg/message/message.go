package message

// Message _
type Message struct {
	User    string
	Content string
}

// Service _
type Service struct {
	r Repository
}

// NewService will return interface for service
func NewService(r Repository) *Service {
	return &Service{r}
}

// CreateMessage _
func (s *Service) CreateMessage(message Message) error {
	return s.r.CreateMessage(message)
}

// UpdateMessage _
type UpdateMessage struct {
	Content string `json:"content"`
}

// UpdateMessage _
func (s *Service) UpdateMessage(id uint, content UpdateMessage) error {
	return s.r.UpdateMessage(id, content)
}

// DeleteMessage _
func (s *Service) DeleteMessage(id uint) error {
	return s.r.DeleteMessage(id)
}
