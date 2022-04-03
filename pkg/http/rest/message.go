package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xxxVitoxxx/JiaPen-backend/pkg/message"
)

// MessageHandler _
type MessageHandler struct {
	s *message.Service
}

// NewMessageHandler _
func NewMessageHandler(s *message.Service) *MessageHandler {
	return &MessageHandler{s}
}

// Router _
func (h *MessageHandler) Router(r *gin.Engine) {
	api := r.Group("/message_api")
	api.POST("/message", h.CreateMessage)
	api.PATCH("/message/:message_id", h.UpdateMessage)
}

// CreateMessage _
type CreateMessage struct {
	User    string `json:"user" binding:"required"`
	Content string `json:"content" binding:"min=1,max=300"`
}

// CreateMessage _
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

// UpdateMessage _
type UpdateMessage struct {
	Content string `json:"content" binding:"min=1,max=300"`
}

// UpddateMessage _
func (h *MessageHandler) UpdateMessage(c *gin.Context) {
	req := UpdateMessage{}
	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	id := c.Param("message_id")
	messageId, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	message := message.UpdateMessage{
		Content: req.Content,
	}
	err = h.s.UpdateMessage(uint(messageId), message)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.Status(http.StatusNoContent)
}
