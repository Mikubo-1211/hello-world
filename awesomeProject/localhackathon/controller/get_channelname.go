package controller

import (
	"encoding/json"
	"localhackathon/dao"
	"log"
	"net/http"
)

func GetChannels(w http.ResponseWriter) {
	channels, err := dao.GetChannels()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Channels := make([]Channel, len(channels))
	for i, channel := range channels {
		Channels[i] = Channel{
			ChannelId:   channel.ChannelID,
			ChannelName: channel.ChannelName,
			Description: channel.Description,
		}
	}
	bytes, err := json.Marshal(Channels)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
