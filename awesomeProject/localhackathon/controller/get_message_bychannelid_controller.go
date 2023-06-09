package controller

import (
	"encoding/json"
	"localhackathon/dao"
	"log"
	"net/http"
)

type Message2 struct {
	UserID    string `json:"user_id"`
	UserName  string `json:"user_name"`
	Id        string `json:"id"`
	Edit      string `json:"edit"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

func GetMessageByChannelIDHandler(w http.ResponseWriter, r *http.Request) {
	channelID := r.URL.Query().Get("channel_id")
	if channelID == "" {
		http.Error(w, "Missing channel_id parameter", http.StatusBadRequest)
		return
	}

	// データベースから指定されたチャンネルIDに関連するメッセージを取得
	messages, err := dao.GetMessagesByChannelID(channelID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Messages := make([]Message2, len(messages))
	for i, message := range messages {
		Messages[i] = Message2{
			UserID:    message.UserID,
			UserName:  message.Username,
			Id:        message.ID,
			Edit:      message.Edit,
			Content:   message.Content,
			CreatedAt: message.CreatedAtStr,
		}
	}
	bytes, err := json.Marshal(Messages)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)

}
