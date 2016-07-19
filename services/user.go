package services

import (
	"adwall/models"
)

func GetUserById(id int64) (*models.TestLh, error) {
	user := models.TestLh{}
	has, err := Ormer().Where("id = ?", id).Get(&user)
	if has {
		return &user, err
	}
	return nil, err
}