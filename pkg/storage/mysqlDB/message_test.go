package mysqlDB

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xxxVitoxxx/JiaPen-backend/pkg/message"
	"github.com/xxxVitoxxx/JiaPen-backend/pkg/storage/conn"
)

func TestCreateMessage(t *testing.T) {
	db := conn.CheckConnect()
	repo := NewMessageRepo(db)
	db.AutoMigrate(&Message{})
	defer db.Migrator().DropTable(&Message{})

	t.Run("CreateMessage should add data and return nil when success", func(t *testing.T) {
		input := message.Message{
			User:    "vito",
			Content: "今天天氣真好呀",
		}
		err := repo.CreateMessage(input)
		assert.NoError(t, err)

		dbRepo := Message{}
		db.Find(&dbRepo)
		assert.Equal(t, input.User, dbRepo.User)
		assert.Equal(t, input.Content, dbRepo.Content)
	})

	t.Run("patch message", func(t *testing.T) {
		message := message.UpdateMessage{
			Content: "Hey",
		}

		err := repo.UpdateMessage(1, message)
		assert.NoError(t, err)

		dbRepo := Message{}
		db.Find(&dbRepo)
		assert.Equal(t, message.Content, dbRepo.Content)
	})

}
