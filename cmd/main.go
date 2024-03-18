package main

import (
	"zumm/internal/middleware"
	"zumm/internal/model"
	"zumm/internal/route"
)

// main is the entry point for the application, it initialises middleware, opens a
// database connection and starts a http server on port 8080
func main() {
	middleware.StartLogger()
	model.ConnectDatabase()
	e := route.SetupRouter()
	e.Logger.Fatal(e.Start(":8080"))
}
