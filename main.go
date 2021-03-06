// SummitAPI is a simple app that, using the email address set in the oauth
// settings, will poll available Google Analytics accounts and return data for
// these accounts, storing it into a database

package main

import (
	"github.com/tylerconlee/SummitAPI/analytics"
	c "github.com/tylerconlee/SummitAPI/config"
	"github.com/tylerconlee/SummitAPI/log"
)

func main() {
	// Set initial configurations, use overrides from environment variables
	config := c.InitConfig()

	// Use configuration to create a new connection
	// TODO: Actually do something here to establish a database connection
	config.DatabaseConnection.DatabaseHost = "localhost"

	// Start logger
	log.InitLog("SummitAPI")

	log.Logger.Info("Application log initialized")

	// Get a list of the account, property and profile IDs from Google
	// analytics for the email address that's been authenticated

	analytics.Init()

	StartHarvest()
}

func StartHarvest() {
	analytics.Harvest()
}
