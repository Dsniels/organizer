package main

import (
	"log"
	"context"

	"github.com/dsniels/organizer/internal/app"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	log.Println("Starting....")
	app := app.NewApp()


	app.Tasks.ListTasks(context.Background())
}
