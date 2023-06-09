package main

import (
	"localhackathon/controller"
	"localhackathon/dao"
	"log"
	"net/http"
)

func enableCORS(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func corsHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			// CORS用のヘッダーを設定してOPTIONSメソッドへの応答を返す
			enableCORS(w, r)
			return
		}
		// 他のリクエストに対してもCORSヘッダーを設定する
		enableCORS(w, r)
		next(w, r)
	}
}

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
		// channel_idをもらうとメッセージ表示
		controller.GetMessageByChannelIDHandler(w, r)
	case http.MethodPost:
		// channelを作る
		controller.PostChannelHandler(w, r)
	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// メッセージのeditの値も持ってくる
}

func Handler3(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		controller.PostUserHandler(w, r)
	case http.MethodGet:
		controller.GetUser(w, r)
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
	case http.MethodDelete:
		controller.DeleteMessageHandler(w, r)
	case http.MethodPut:
		controller.UpdateMessageHandler(w, r)
	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// 編集できるように
	// メッセージの削除
}
func Handler5(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		controller.GetChannels(w)
	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

func main() {
	// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
	http.HandleFunc("/user", corsHandler(Handler1))
	// 動作確認OK
	http.HandleFunc("/channel", corsHandler(Handler2))
	// 動作確認OK
	http.HandleFunc("/users", corsHandler(Handler3))
	// OK
	http.HandleFunc("/message", corsHandler(Handler4))

	http.HandleFunc("/channels", corsHandler(Handler5))

	// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
	dao.CloseDBWithSysCall()
	// 参加したり抜けたりできるようにする

	// 9000番ポートでリクエストを待ち受ける
	log.Println("Listening...")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal(err)
	}
}
