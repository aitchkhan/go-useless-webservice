package controllers

import (
	"net/http"
	"regexp"
)


type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from User controller"))
}

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^users/(\d+)/?`),
	}
}