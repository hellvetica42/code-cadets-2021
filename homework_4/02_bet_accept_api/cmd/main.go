package main

import (
	"log"

	"github.com/superbet-group/code-cadets-2021/homework_4/02_bet_accept_api/cmd/bootstrap"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bet_accept_api/cmd/config"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bet_accept_api/internal/tasks"
)

func main() {
	log.Println("Bootstrap initiated")

	config.Load()

	rabbitMqChannel := bootstrap.RabbitMq()
	signalHandler := bootstrap.SignalHandler()
	api := bootstrap.Api(rabbitMqChannel)

	log.Println("Bootstrap finished. Event API is starting")

	tasks.RunTasks(signalHandler, api)

	log.Println("Bet accept API finished gracefully")
}
