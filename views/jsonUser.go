package views

import (
	"abstractions/models"
	"encoding/json"
)

func JsonUser(u *models.User) []byte {
	serialized, _ := json.Marshal(u)

	return serialized
}
