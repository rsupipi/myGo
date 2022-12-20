package main

import (
	"fmt"
	"github.com/rsupipi/myGo/models"
)

func main() {
	u := models.User{
		ID:        1,
		FirstName: "pipi",
		LastName:  "ruchi",
	}
	fmt.Println(u)

}
