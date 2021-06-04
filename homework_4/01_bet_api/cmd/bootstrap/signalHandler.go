package bootstrap

import (
	"github.com/superbet-group/code-cadets-2021/homework4/01_bet_api/internal/tasks"
)

// SignalHandler bootstraps the signal handler.
func SignalHandler() *tasks.SignalHandler {
	return tasks.NewSignalHandler()
}
