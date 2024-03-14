package main

import (
	"zumm/routes"
)

func main() {
	e := routes.SetupRouter()
	e.Logger.Fatal(e.Start(":8080"))
}
