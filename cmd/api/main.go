package main

import (
	"log"
	"task-api/internal/app"
)

func main() {
	log.Println("NEW VERSION RUNNING")
	app := app.NewApp()
	log.Fatal(app.Run())
}
