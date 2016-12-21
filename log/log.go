package log

import log "github.com/mgutz/logxi/v1"

var Logger log.Logger

func InitLog() {
	// create a logger with a unique identifier which
	// can be enabled from environment variables
	Logger = log.New("SummitAPI")
}
