package controllers

import "net/http"

func RegisterController() {
	uc := newUserController()

	/*		**** Hander interface ****
			I'm expecting a handler interface for the second parameter. Since I matched that method signature of the interface
			then, Go is going to recognize that, it does in fact has something that it can use as and HTTP handler.
	*/
	http.Handle("/users", *uc)  // providing a route
	http.Handle("/users/", *uc) // "/users" and "/users/" are different routes

}
