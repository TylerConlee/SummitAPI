// SummitAPI creates a web server that handles requests made by SummitDaemon
// Designed to run in parallel in a containerized fashion
// SummitAPI is required to be running in order to run SummitDaemon

package main

import (
	analytics "github.com/tylerconlee/SummitAPI/analytics"
	c "github.com/tylerconlee/SummitAPI/config"
	log "github.com/tylerconlee/SummitAPI/log"
)

func main() {
	// Set initial configurations, use overrides from environment variables
	config := c.NewConfig()

	// Use configuration to create a new connection
	config.DatabaseConnection.DatabaseHost = "localhost"

	// Start logger
	log.InitLog("SummitAPI")

	log.Logger.Info("Application log initialized")

	// Run the ConnectAnalytics function that autheticates with Google using oAuth and preps the analytics.Service for use
	analytics.ConnectAnalytics()

	// Get a list of the account, property and profile IDs from Google analytics for the email address that's been autheticated
	analytics.GetID()

	// Output the grabbed profiles
	log.Logger.Debug(analytics.Profiles)
}
