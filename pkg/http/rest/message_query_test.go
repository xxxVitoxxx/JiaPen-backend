package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/xxxVitoxxx/JiaPen-backend/pkg/message"
	"github.com/xxxVitoxxx/JiaPen-backend/pkg/storage/fake"
)

func TestGetMessages(t *testing.T) {
	r := gin.Default()
	fr := fake.NewQueryMessageRepo()
	handler := NewQueryMessageHandler(fr)
	handler.Router(r)

	t.Run("get messages", func(t *testing.T) {
		createAt, _ := time.Parse(time.RFC3339, "2025-05-05T12:01:01")
		fr.QueryMessages = []message.QueryMessage{
			{
				Id:        1,
				User:      "vito",
				Content:   "今天天氣真好",
				CreatedAt: createAt,
				UpdatedAt: createAt,
			},
			{
				Id:        2,
				User:      "eve",
				Content:   "今天過得好嗎",
				CreatedAt: createAt,
				UpdatedAt: createAt,
			},
		}

		w := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodGet,
			"/message_api/messages",
			nil,
		)
		r.ServeHTTP(w, req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(
			t,
			`[
				{
					"id":1,
					"user":"vito",
					"content":"今天天氣真好",
					"created_at":"0001-01-01T00:00:00Z",
					"updated_at":"0001-01-01T00:00:00Z"
				},
				{
					"id":2,
					"user":"eve",
					"content":"今天過得好嗎",
					"created_at":"0001-01-01T00:00:00Z",
					"updated_at":"0001-01-01T00:00:00Z"
				}
			]`,
			w.Body.String(),
		)
	})
}
