package controller

import (
	"localhackathon/dao"
	"log"
	"net/http"
)

func DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	messageID := r.URL.Query().Get("id")
	if messageID == "" {
		log.Printf("fail: Missing message_id parameter\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := dao.DeleteMessageByID(messageID)
	if err != nil {
		log.Printf("fail: dao.DeleteMessageByID, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
