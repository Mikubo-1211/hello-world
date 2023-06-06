package dao

import (
	"github.com/oklog/ulid"
	"log"
	"math/rand"
	"time"
)

func generateULID() (string, error) {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)

	id, err := ulid.New(ulid.Timestamp(time.Now()), entropy)
	if err != nil {
		log.Printf("failed to generate ULID: %v", err)
		return "", err
	}

	return id.String(), nil
}

func CreateChannel(channelName string, description string) (string, error) {
	// Generate a ULID for the channel ID
	channelID, _ := generateULID()

	tx, err := db.Begin()
	if err != nil {
		return "", err
	}
	defer tx.Rollback()

	// Assuming you have a database connection named "db" established

	// Prepare the SQL query

	// Execute the query
	_, err = db.Exec("INSERT INTO channels(channel_id,channel_name,description) VALUES (?, ?, ?)", channelID, channelName, description)
	if err != nil {
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		return "", err
	}

	return channelID, nil
}
