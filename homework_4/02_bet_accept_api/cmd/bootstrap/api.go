package bootstrap

import (
	"github.com/streadway/amqp"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bet_accept_api/cmd/config"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bet_accept_api/internal/api"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bet_accept_api/internal/api/controllers"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bet_accept_api/internal/api/controllers/validators"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bet_accept_api/internal/domain/services"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bet_accept_api/internal/infrastructure/rabbitmq"
)

func newBetReceivedValidator() *validators.BetReceivedValidator {
	return validators.BetReceivedUpdateValidator()
}

func newBetReceivedPublisher(publisher rabbitmq.QueuePublisher) *rabbitmq.BetReceivedPublisher {
	return rabbitmq.NewBetReceivedPublisher(
		config.Cfg.Rabbit.PublisherExchange,
		config.Cfg.Rabbit.PublisherBetReceivedQueue,
		config.Cfg.Rabbit.PublisherMandatory,
		config.Cfg.Rabbit.PublisherImmediate,
		publisher,
	)
}

func newBetReceivedService(publisher services.BetReceivedPublisher) *services.BetReceivedService {
	return services.NewBetReceivedService(publisher)
}

func newController(betReceivedValidator controllers.BetReceivedValidator, betReceivedService controllers.BetReceivedService) *controllers.Controller {
	return controllers.NewController(betReceivedValidator, betReceivedService)
}

// Api bootstraps the http server.
func Api(rabbitMqChannel *amqp.Channel) *api.WebServer {
	betReceivedValidator := newBetReceivedValidator()
	betReceivedPublisher := newBetReceivedPublisher(rabbitMqChannel)
	betReceivedService := newBetReceivedService(betReceivedPublisher)
	controller := newController(betReceivedValidator, betReceivedService)

	return api.NewServer(config.Cfg.Api.Port, config.Cfg.Api.ReadWriteTimeoutMs, controller)
}
