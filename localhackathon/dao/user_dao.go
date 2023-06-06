package dao

import (
	"database/sql"
	"github.com/oklog/ulid"
	"log"
	"math/rand"
	"time"
)

type UserResForHTTPGet struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetUserJoho(name string) (*sql.Rows, error) {
	rows, err := db.Query("SELECT id, name, age FROM user WHERE name = ?", name)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func InsertUser(users UserResForHTTPGet) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		return "", err
	}
	defer tx.Rollback()

	t := time.Now()

	// ランダムなエンタロピーを生成
	entropy := rand.New(rand.NewSource(t.UnixNano()))

	// ULIDを生成
	id := ulid.MustNew(ulid.Timestamp(t), entropy).String()

	_, err = db.Exec("INSERT INTO user(id,name,age) VALUES (?, ?, ?)", id, users.Name, users.Age)
	if err != nil {
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		return "", err
	}

	log.Printf("id: %v\n", id)
	return id, nil
}
