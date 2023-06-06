package usecase

import (
	"fmt"
)

func ValidateUser(users UserResForHTTPGet) error {
	if users.Name == "" || len(users.Name) > 50 {
		return fmt.Errorf("Name entered is invalid: %s", users.Name)
	}
	if users.Age < 20 || users.Age > 80 {
		return fmt.Errorf("Age entered is invalid: %d", users.Age)
	}
	return nil
}
