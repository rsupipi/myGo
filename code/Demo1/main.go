package main

import (
	"github.com/rsupipi/myGo/controllers"
	"net/http"
)

func main() {
	controllers.RegisterController()
	http.ListenAndServe(":3000", nil)
}
