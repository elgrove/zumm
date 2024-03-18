package main

import (
	"zumm/internal/route"
)

func main() {
	e := route.SetupRouter()
	e.Logger.Fatal(e.Start(":8080"))
}
