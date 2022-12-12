package main

import (
	"net/http"

	"learn-go/simple-webserver/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
