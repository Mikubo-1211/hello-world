package dao

import (
	"localhackathon/model"
)

func GetMessagesByChannelID(channelID string) ([]model.Message, error) {
	// Assuming you have a database connection named "db" established

	// Prepare the SQL query
	query := "SELECT m.user_id, u.user_name, m.Id,m.edit,m.content, m.created_at FROM message AS m JOIN users AS u ON m.user_id = u.user_id WHERE m.channel_id = ? ORDER BY m.created_at ASC"

	// Execute the query
	rows, err := db.Query(query, channelID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize a slice to store the messages
	messages := make([]model.Message, 0)

	// Iterate over the query results
	for rows.Next() {
		var message model.Message
		err := rows.Scan(&message.UserID, &message.Username, &message.ID, &message.Edit, &message.Content, &message.CreatedAtStr)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
