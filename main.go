// SummitAPI creates a web server that handles requests made by SummitDaemon
// Designed to run in parallel in a containerized fashion
// SummitAPI is required to be running in order to run SummitDaemon

package main

import "github.com/tylerconlee/SummitAPI/analytics"
import "fmt"

func main() {
	// Set initial configurations, use overrides from environment variables
	config := NewConfig()

	// Use configuration to create a new connection
	config.DatabaseConnection.DatabaseHost = "localhost"

	// Use connection to be passed to listen for data push/pull requests
	analytics.ConnectAnalytics()
	println("Connected to Google Analytics")
	analytics.GetID()
	fmt.Printf("%v", analytics.Profiles)

	// Listen and serve HTTP requests from the PHP app
}
