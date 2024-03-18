package main

import (
	"zumm/internal/model"
	"zumm/internal/route"
)

// main is the entry point for the application. It starts the router on port 8080.
func main() {
	model.ConnectDatabase()
	e := route.SetupRouter()
	e.Logger.Fatal(e.Start(":8080"))
}
