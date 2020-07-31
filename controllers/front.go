package controllers

import "net/http"

// RegisterControllers registers the handlers
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
