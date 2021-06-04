package bootstrap

import (
	"github.com/superbet-group/code-cadets-2021/homework4/01_bet_api/cmd/config"
	"github.com/superbet-group/code-cadets-2021/homework4/01_bet_api/internal/api"
	"github.com/superbet-group/code-cadets-2021/homework4/01_bet_api/internal/api/controllers"
	"github.com/superbet-group/code-cadets-2021/homework4/01_bet_api/internal/api/controllers/validators"
	"github.com/superbet-group/code-cadets-2021/homework4/01_bet_api/internal/domain/mappers"
	"github.com/superbet-group/code-cadets-2021/homework4/01_bet_api/internal/domain/services"
	"github.com/superbet-group/code-cadets-2021/homework4/01_bet_api/internal/infrastructure/sqlite"
)

func newBetRequestValidator() *validators.BetRequestValidator{
	return validators.NewBetRequestValidator()
}

func newController(betRequestValidator controllers.BetRequestValidator, betRepository controllers.BetRepository, mapper controllers.DtoBetMapper) *controllers.Controller {
	return controllers.NewController(betRequestValidator, betRepository, mapper)
}

func newBetMapper() *mappers.BetMapper {
	return mappers.NewBetMapper()
}

func newBetRepository(dbExecutor sqlite.DatabaseExecutor, betMapper sqlite.BetMapper) *sqlite.BetRepository {
	return sqlite.NewBetRepository(dbExecutor, betMapper)
}

func newBetService() *services.BetService {
	return services.NewBetService()
}

// Api bootstraps the http server.
func Api(db sqlite.DatabaseExecutor) *api.WebServer {
	mapper := newBetMapper()
	betRepository := newBetRepository(db, mapper)

	betUpdateValidator := newBetRequestValidator()
	//betService := newBetService(
	controller := newController(betUpdateValidator, betRepository, mapper)

	return api.NewServer(config.Cfg.Api.Port, config.Cfg.Api.ReadWriteTimeoutMs, controller)
}
