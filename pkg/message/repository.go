package message

// Repository _
type Repository interface {
	CreateMessage(message Message) error
	UpdateMessage(id uint, content UpdateMessage) error
	DeleteMessage(id uint) error
}
