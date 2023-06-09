package controller

import (
	"encoding/json"
	"fmt"
	"localhackathon/dao"
	"localhackathon/usecase"
	"log"
	_ "math/rand"
	"net/http"
	_ "time"
)

type IdInsert struct {
	Id string `json:"id"`
}

func HandlerPostUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var users UserResForHTTPGet
		if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
			fmt.Println(err)
			return
		}

		if err := usecase.ValidateUser(usecase.UserResForHTTPGet(users)); err != nil {
			log.Printf("fail: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := dao.InsertUser(dao.UserResForHTTPGet(users))
		if err != nil {
			log.Printf("fail: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// 自信なし
		var ID IdInsert
		ID.Id = id
		bytes, err := json.Marshal(ID)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
