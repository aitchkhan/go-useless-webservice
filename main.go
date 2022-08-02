package main

import (
	"fmt"

	"github.com/aitchkhan/go-useless-webservice/models"
)

func main() {
	u := models.User{
		ID:        1,
		firstName: "Haroon",
		lastName:  "Khan",
	}
	fmt.Println(u)
}
