package controller

import (
	"encoding/json"
	"localhackathon/dao"
	"log"
	"net/http"
)

type UserRequest struct {
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
	UserEmail    string `json:"user_email"`
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var userReq UserRequest
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		log.Printf("fail: error decoding userrequest body, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if userReq.UserName == "" {
		log.Printf("fail: user name is required\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = dao.CreateUsers(userReq.UserName, userReq.UserPassword, userReq.UserEmail)
	if err != nil {
		log.Printf("fail: dao.InsertChannel, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
