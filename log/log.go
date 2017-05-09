package log

import (
	"os"

	log "github.com/op/go-logging"
	c "github.com/tylerconlee/SummitAPI/config"
)

var Logger *log.Logger

var format = log.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{module} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func InitLog(module string) {
	// Load configuration to be able to grab log level
	config := c.InitConfig()

	// Create new log.Logger with the module name
	Logger = log.MustGetLogger(module)

	// Tell the log where to output
	// TODO: turn this into a configurable setting
	backend := log.NewLogBackend(os.Stderr, "", 0)

	// Assign the format to the log
	formatter := log.NewBackendFormatter(backend, format)

	// Filter the log based off of the configured log level
	l := log.AddModuleLevel(formatter)

	// Load the log level from the configuration
	level, _ := log.LogLevel(config.LogLevel)

	// Set the log level to the configured log level
	l.SetLevel(level, "")

	// Tell the log instance to use the leveled logs
	log.SetBackend(l)
}
