package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xxxVitoxxx/JiaPen-backend/pkg/message"
)

type MessageHandler struct {
	s *message.Service
}

func NewMessageHandler(s *message.Service) *MessageHandler {
	return &MessageHandler{s}
}

func (h *MessageHandler) Router(r *gin.Engine) {
	api := r.Group("/message_api")
	api.POST("/message", h.CreateMessage)
}

type CreateMessage struct {
	User    string `json:"user" binding:"required"`
	Content string `json:"content" binding:"min=1,max=300"`
}

func (h *MessageHandler) CreateMessage(c *gin.Context) {
	req := new(CreateMessage)
	err := c.BindJSON(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})

		return
	}

	err = h.s.CreateMessage(message.Message{
		User:    req.User,
		Content: req.Content,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})

		return
	}

	c.Status(http.StatusCreated)
}
