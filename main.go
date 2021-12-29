package main

import (
	"go-web-app/routes"
	"net/http"
)

func main() {
	routes.MountRoutes()
	http.ListenAndServe(":8001", nil)
}
