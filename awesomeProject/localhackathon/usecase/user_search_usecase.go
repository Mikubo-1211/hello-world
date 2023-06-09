package usecase

import (
	"database/sql"
)

type UserResForHTTPGet struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func ExtractUsersFromRows(rows *sql.Rows) ([]UserResForHTTPGet, error) {
	users := make([]UserResForHTTPGet, 0)

	for rows.Next() {
		var u UserResForHTTPGet
		if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
