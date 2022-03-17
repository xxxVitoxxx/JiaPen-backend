package rest

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/xxxVitoxxx/JiaPen-backend/pkg/message"
	"github.com/xxxVitoxxx/JiaPen-backend/pkg/storage/fake"
)

func TestCreateMessage(t *testing.T) {
	r := gin.Default()
	fr := fake.NewMessageRepo()
	s := message.NewService(fr)
	handler := NewMessageHandler(s)
	handler.Router(r)

	t.Run("should response 201 when request current", func(t *testing.T) {
		input, _ := json.Marshal(CreateMessage{
			User:    "vito",
			Content: "今天天氣真好",
		})

		w := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodPost,
			"/message_api/message",
			bytes.NewBuffer(input),
		)

		assert.NoError(t, err)
		r.ServeHTTP(w, req)

		assert.Equal(t, 201, w.Code)
		fmt.Println("Body:", w.Body)
	})

	t.Run("when content more than 300 should return 400", func(t *testing.T) {
		fakeString := make([]byte, 301)
		input, _ := json.Marshal(CreateMessage{
			User:    "vito",
			Content: hex.EncodeToString(fakeString),
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(
			http.MethodPost,
			"/message_api/message",
			bytes.NewBuffer(input),
		)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)

	})
}
