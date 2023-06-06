package dao

import (
	"github.com/oklog/ulid"
	"time"
)

type Message struct {
	MessageID string
	UserID    string
	ChannelID string
	Content   string
	Edit      string
	CreatedAt string
}

func InsertMessage(content, channelID, userID string) (string, error) {
	// メッセージの作成時刻を現在の日本時間で取得
	now := time.Now().In(time.FixedZone("Asia/Tokyo", 9*60*60))

	// メッセージIDを生成
	messageID := generateMessageID()

	// メッセージを作成
	message := Message{
		MessageID: messageID,
		UserID:    userID,
		ChannelID: channelID,
		Content:   content,
		Edit:      "0",
		CreatedAt: now.Format("2006-01-02T15:04:05Z07:00"),
	}

	// メッセージをデータベースにインサート
	_, err := db.Exec("INSERT INTO message (id, user_id, channel_id, content, edit, created_at) VALUES (?, ?, ?, ?, ?, ?)", message.MessageID, message.UserID, message.ChannelID, message.Content, message.Edit, message.CreatedAt)
	if err != nil {
		return "", err
	}

	return messageID, nil
}

// メッセージIDを生成する関数
func generateMessageID() string {
	// ULID を生成
	id := ulid.MustNew(ulid.Now(), nil).String()
	return id
}
