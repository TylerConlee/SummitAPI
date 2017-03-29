// SummitAPI is a simple app that, using the email address set in the oauth
// settings, will poll available Google Analytics accounts and return data for
// these accounts, storing it into a database

package main

import (
	analytics "github.com/tylerconlee/SummitAPI/analytics"
	c "github.com/tylerconlee/SummitAPI/config"
	log "github.com/tylerconlee/SummitAPI/log"
	"fmt"
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

	// Run the ConnectAnalytics function that autheticates with Google using
	// oAuth and preps the analytics.Service for use
	analytics.ConnectAnalytics()

	// Get a list of the account, property and profile IDs from Google
	// analytics for the email address that's been autheticated
	// TODO: Add email config option for account changes
	analytics.GetID()

	// Output the grabbed profiles
	fmt.Printf("%+v\n", analytics.Profiles[0].PropertyID)
	log.Logger.Debug(analytics.Profiles[1])
}
