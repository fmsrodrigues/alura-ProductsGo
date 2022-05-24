package main

import (
	"net/http"

	"localhost/alura-GoWebProducts/src/routes"

	_ "github.com/lib/pq"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8080", nil)
}
