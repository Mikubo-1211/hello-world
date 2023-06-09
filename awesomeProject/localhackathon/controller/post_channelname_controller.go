package controller

import (
	"localhackathon/dao"
	"log"
	"net/http"
)

type Channel struct {
	ChannelId   string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	Description string `json:"description"`
}

func PostChannelHandler(w http.ResponseWriter, r *http.Request) {
	var channel Channel
	channel.ChannelName = r.URL.Query().Get("channel_name")
	channel.Description = r.URL.Query().Get("description")

	if channel.ChannelName == "" {
		log.Printf("fail: Channel name is required\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := dao.CreateChannel(channel.ChannelName, channel.Description)
	if err != nil {
		log.Printf("fail: dao.InsertChannel, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
