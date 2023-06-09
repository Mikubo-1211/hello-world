package dao

import (
	"log"
)

func UpdateMessageContentAndEditFlag(messageID, content, edit string) error {

	// メッセージの更新クエリを実行
	_, err := db.Exec("UPDATE message SET content = ?, edit = ? WHERE id = ?", content, edit, messageID)
	if err != nil {
		log.Printf("fail: UpdateMessageContentAndEditFlag, %v\n", err)
		return err
	}

	return nil
}
