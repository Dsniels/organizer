package app

import (
	"log"

	"github.com/dsniels/organizer/internal/client"
	"github.com/dsniels/organizer/internal/service"
)

type App struct {
	Tasks service.ITask
}

func NewApp() *App {
	log.Println("Getting ms client....")
	msClient := client.GetMSGraphClient()
	log.Println("Getting task services....")
	tasks := service.NewMSTaskService(msClient)
	return &App{Tasks: tasks}

}
