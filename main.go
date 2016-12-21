// SummitAPI creates a web server that handles requests made by SummitDaemon
// Designed to run in parallel in a containerized fashion
// SummitAPI is required to be running in order to run SummitDaemon

package main

import (
	"os"

	analytics "github.com/tylerconlee/SummitAPI/analytics"
	c "github.com/tylerconlee/SummitAPI/config"
	log "github.com/tylerconlee/SummitAPI/log"
)

func main() {
	os.Setenv("LOGXI", "*")

	// Set initial configurations, use overrides from environment variables
	config := c.NewConfig()

	// Use configuration to create a new connection
	config.DatabaseConnection.DatabaseHost = "localhost"

	// Start logger
	log.InitLog("SummitAPI")

	log.Logger.Info("Application log initialized")

	analytics.ConnectAnalytics()
	analytics.GetID()
	log.Logger.Debug(analytics.Profiles)

