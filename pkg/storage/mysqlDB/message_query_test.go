package mysqlDB

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xxxVitoxxx/JiaPen-backend/pkg/message"
	"github.com/xxxVitoxxx/JiaPen-backend/pkg/storage/conn"
)

func TestFindMessages(t *testing.T) {
	db := conn.CheckConnect()
	repo := NewMessageRepo(db)
	queryRepo := NewQueryMessageRepo(db)
	db.AutoMigrate(&Message{})
	defer db.Migrator().DropTable(&Message{})

	t.Run("Set test data", func(t *testing.T) {
		messages := message.Message{
			User:    "vito",
			Content: "今天天氣真好",
		}
		_ = repo.CreateMessage(messages)

		messages = message.Message{
			User:    "eve",
			Content: "你今天要去哪",
		}
		_ = repo.CreateMessage(messages)
	})

	t.Run("Find messages", func(t *testing.T) {
		messages, err := queryRepo.FindMessages()
		assert.NoError(t, err)
		assert.Len(t, messages, 2)
		assert.Equal(t, "vito", messages[0].User)
		assert.Equal(t, "你今天要去哪", messages[1].Content)
	})
}
