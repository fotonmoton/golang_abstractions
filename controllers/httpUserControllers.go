package controllers

import (
	"abstractions/models"
	"abstractions/views"
	"net/http"
)

func ChangeUserAgeController(w http.ResponseWriter, r *http.Request, u *models.User, age int) {
	u.ChangeAge(age)

	ShowUserController(w, r, u)
}

func ShowUserController(w http.ResponseWriter, r *http.Request, u *models.User) {
	w.Write(views.JsonUser(u))
}
