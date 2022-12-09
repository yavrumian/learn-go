package main

import (
	"fmt"
	"test-mod/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Vahe",
		LastName:  "Yavrumian",
	}
	fmt.Println(u)
}
