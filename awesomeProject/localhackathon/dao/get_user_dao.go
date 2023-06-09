package dao

import "localhackathon/model"

func Getuser(email string) (model.User, error) {
	query := "SELECT * FROM users WHERE user_email = ?"

	row := db.QueryRow(query, email)
	user := model.User{}
	err := row.Scan(&user.UserID, &user.UserName, &user.UserPassword, &user.UserEmail) // Adjust the fields as per your user model
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
