package controller

import (
	"localhackathon/dao"
	"log"
	"net/http"
)

type Message struct {
	UserID    string `json:"user_id"`
	ChannelID string `json:"channel_id"`
	Content   string `json:"content"`
}

func PostMessageHandler(w http.ResponseWriter, r *http.Request) {
	var message Message
	message.Content = r.URL.Query().Get("content")
	message.ChannelID = r.URL.Query().Get("channel_id")
	message.UserID = r.URL.Query().Get("user_id")
	if message.Content == "" {
		log.Printf("fail: Content is required\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if message.ChannelID == "" {
		log.Printf("fail: Channel is not chosen\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if message.UserID == "" {
		log.Printf("fail: Who are you\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := dao.InsertMessage(message.Content, message.ChannelID, message.UserID)
	if err != nil {
		log.Printf("fail: dao.InsertChannel, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
