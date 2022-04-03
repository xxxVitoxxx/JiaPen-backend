package message_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xxxVitoxxx/JiaPen-backend/pkg/message"
	"github.com/xxxVitoxxx/JiaPen-backend/pkg/storage/fake"
)

func TestCreateMessage(t *testing.T) {
	fr := fake.NewMessageRepo()
	s := message.NewService(fr)
	t.Run("create message", func(t *testing.T) {
		input := message.Message{
			User:    "vito",
			Content: "今天天氣真好呢...",
		}

		err := s.CreateMessage(input)
		assert.NoError(t, err)
		assert.Equal(t, input.User, fr.Messages[1].User)
		assert.Equal(t, input.Content, fr.Messages[1].Content)
	})

	t.Run("patch message", func(t *testing.T) {
		message := message.UpdateMessage{
			Content: "hey",
		}

		err := s.UpdateMessage(1, message)
		assert.NoError(t, err)
		assert.Equal(t, message.Content, fr.Messages[1].Content)
	})

	t.Run("delete message", func(t *testing.T) {
		err := s.DeleteMessage(1)
		assert.NoError(t, err)
		assert.Len(t, fr.Messages, 0)
	})
}
