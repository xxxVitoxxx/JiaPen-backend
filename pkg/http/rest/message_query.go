package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xxxVitoxxx/JiaPen-backend/pkg/message"
)

// QueryMessageHandler _
type QueryMessageHandler struct {
	r message.QueryRepository
}

// NewQueryMessageHandler will return instance of QueryMessageHandler
func NewQueryMessageHandler(r message.QueryRepository) *QueryMessageHandler {
	return &QueryMessageHandler{r}
}

// Router _
func (h *QueryMessageHandler) Router(r *gin.Engine) {
	api := r.Group("/message_api")
	api.GET("/messages", h.GetMessages)
}

// GetMessages _
func (h *QueryMessageHandler) GetMessages(c *gin.Context) {
	result, err := h.r.FindMessages()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, result)
}
