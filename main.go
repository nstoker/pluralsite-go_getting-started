package main

import (
	"net/http"

	"github.com/pluralsite-go_getting-started/controllers"
)

func main() {
	controllers.RegisterControllers()

	http.ListenAndServe(":3000", nil)
}
