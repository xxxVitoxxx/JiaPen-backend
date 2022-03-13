package message

// Repository _
type Repository interface {
	CreateMessage(message Message) error
}
