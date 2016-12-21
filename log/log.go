package log

import (
	"os"

	log "github.com/op/go-logging"
)

var Logger *log.Logger

var format = log.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{module} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func InitLog(module string) {
	Logger = log.MustGetLogger(module)
	// create a logger with a unique identifier which
	// can be enabled from environment variables
	backend := log.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := log.NewBackendFormatter(backend, format)
	log.SetBackend(backendFormatter)
}
