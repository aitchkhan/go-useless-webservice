package main

import (
	"net/http"

	"github.com/aitchkhan/go-useless-webservice/controllers"
	"github.com/aitchkhan/go-useless-webservice/models"
)


var u = models.User{
 ID:        1,
 FirstName: "Haroon",
 LastName:  "Khan",
}
func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":8000", nil)
}

