package model

type Message struct {
	ID           string // メッセージのID
	UserID       string // ユーザーのID
	Username     string
	ChannelID    string // チャンネルのID
	Content      string // メッセージの内容
	Edit         string // 編集フラグ
	CreatedAtStr string //作成日時
}

type Channels struct {
	ChannelID   string
	ChannelName string
	Description string
}
