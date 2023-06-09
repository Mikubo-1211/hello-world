package controller

import (
	"encoding/json"
	"localhackathon/dao"
	"net/http"
)

type User struct {
	UserID       string `json:"user_id"`
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
	UserEmail    string `json:"user_email"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("user_email")
	if email == "" {
		http.Error(w, "Missing email parameter", http.StatusBadRequest)
		return
	}
	user, err := dao.Getuser(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u := &User{
		UserID:       user.UserID,
		UserName:     user.UserName,
		UserPassword: user.UserPassword,
		UserEmail:    user.UserEmail,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)

}
