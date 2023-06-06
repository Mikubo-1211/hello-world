package controller

import (
	"localhackathon/dao"
	"log"
	"net/http"
)

type Users struct {
	UsersID      string `json:"user_id"`
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
	UserEmail    string `json:"user_email"`
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var users Users
	users.UserName = r.URL.Query().Get("user_name")
	users.UserPassword = r.URL.Query().Get("user_password")
	users.UserEmail = r.URL.Query().Get("user_email")

	if users.UserName == "" {
		log.Printf("fail: Channel name is required\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := dao.CreateUsers(users.UserName, users.UserPassword, users.UserEmail)
	if err != nil {
		log.Printf("fail: dao.InsertChannel, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
