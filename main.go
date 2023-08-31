package main

import (
	"log"

	"github.com/nggrjh/travel-planner/internal/application"
)

func main() {
	app, err := application.New()
	if err != nil {
		log.Fatal(err)
	}
	defer app.Close()

	app.Start()

	app.WaitForShutdown()
}
