package main

import (
	"errors"
	"fmt"

	"github.com/aitchkhan/go-useless-webservice/models"
)


var u = models.User{
 ID:        1,
 FirstName: "Haroon",
 LastName:  "Khan",
}
func main() {
	startWebServer(3000)
}


func startWebServer(port int) (int, error) {
	fmt.Println("Starting web server")
	fmt.Println("Started web server")
	if (port!=3000) {
		return 0, errors.New("Something went wrong")
	}
	return port, nil
}
