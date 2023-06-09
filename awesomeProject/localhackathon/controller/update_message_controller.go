package controller

import (
	"encoding/json"
	"localhackathon/dao"
	"log"
	"net/http"
)

func UpdateMessageHandler(w http.ResponseWriter, r *http.Request) {
	var message struct {
		ID      string `json:"id"`
		Content string `json:"content"`
	}

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		log.Printf("fail: json.NewDecoder, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if message.ID == "" {
		log.Printf("fail: Message ID is required\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if message.Content == "" {
		log.Printf("fail: Content is required\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = dao.UpdateMessageContentAndEditFlag(message.ID, message.Content, "1")
	if err != nil {
		log.Printf("fail: dao.UpdateMessageContentAndEditFlag, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
