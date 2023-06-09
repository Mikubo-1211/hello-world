package dao

import (
	"log"
)

func DeleteMessageByID(messageID string) error {
	// データベース接続などの前処理

	// メッセージの削除クエリを実行
	_, err := db.Exec("DELETE FROM message WHERE id = ?", messageID)
	if err != nil {
		log.Printf("fail: DeleteMessageByID, %v\n", err)
		return err
	}

	// 後処理やエラーチェックなど

	return nil
}
