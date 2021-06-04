package bootstrap

import (
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bet_accept_api/internal/tasks"
)

// SignalHandler bootstraps the signal handler.
func SignalHandler() *tasks.SignalHandler {
	return tasks.NewSignalHandler()
}
