package dao

func CreateUsers(userName string, userPassword string, userEmail string) (string, error) {
	// Generate a ULID for the channel ID
	userID, _ := generateULID()

	tx, err := db.Begin()
	if err != nil {
		return "", err
	}
	defer tx.Rollback()

	// Assuming you have a database connection named "db" established

	// Prepare the SQL query

	// Execute the query
	_, err = db.Exec("INSERT INTO users(user_id,user_name,user_password,user_email) VALUES (?, ?, ?,?)", userID, userName, userPassword, userEmail)
	if err != nil {
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		return "", err
	}

	return userID, nil
}
