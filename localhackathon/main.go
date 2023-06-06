package main

import (
	"localhackathon/controller"
	"localhackathon/dao"
	"log"
	"net/http"
)

func Handler1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		controller.HandlerPostUser(w, r)
	case http.MethodGet:
		controller.HandlerGetUser(w, r)

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func Handler2(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		//channel_idをもらうとメッセージ表示
		controller.GetMessageByChannelIDHandler(w, r)
	case http.MethodPost:
		//channelを作る
		controller.PostChannelHandler(w, r)

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//メッセージのeditの値も持ってくる

}
func Handler3(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		controller.PostUserHandler(w, r)

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
func Handler4(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		controller.PostMessageHandler(w, r)

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//編集できるように
	//メッセージの削除
}

func main() {
	// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
	http.HandleFunc("/user", Handler1)
	//動作確認OK
	http.HandleFunc("/channel", Handler2)
	//動作確認OK
	http.HandleFunc("/users", Handler3)
	//OK
	http.HandleFunc("/message", Handler4)

	// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
	dao.CloseDBWithSysCall()
	//参加したり抜けたりできるようにする

	// 8000番ポートでリクエストを待ち受ける
	log.Println("Listening...")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal(err)
	}
}
