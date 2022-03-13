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
