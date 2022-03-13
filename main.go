package main

import (
	"project/config"
	"project/route"
)

func main() {
	config.GetClient()
	e := route.New()
	e.Logger.Fatal(e.Start(":8000"))
}
