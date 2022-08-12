package controllers

import "net/http"

var uc = newUserController()

func RegisterControllers() {
	http.Handle("/users/", *uc)
	http.Handle("/users", *uc)
}