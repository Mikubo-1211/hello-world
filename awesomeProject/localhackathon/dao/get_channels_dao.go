package dao

import "localhackathon/model"

func GetChannels() ([]model.Channels, error) {
	query := "SELECT * FROM channels"

	// Execute the quer
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize a slice to store the messages
	channels := make([]model.Channels, 0)

	// Iterate over the query results
	for rows.Next() {
		var channel model.Channels
		err := rows.Scan(&channel.ChannelID, &channel.ChannelName, &channel.Description)
		if err != nil {
			return nil, err
		}
		channels = append(channels, channel)
	}

	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return channels, nil
}
