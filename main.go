// SummitAPI creates a web server that handles requests made by SummitDaemon
// Designed to run in parallel in a containerized fashion
// SummitAPI is required to be running in order to run SummitDaemon

package main

import "github.com/tylerconlee/SummitAPI/analytics"
import "os"
import "fmt"

func main() {
	// Set initial configurations, use overrides from environment variables
	config := NewConfig()

	// Use configuration to create a new connection
	config.DatabaseConnection.DatabaseHost = "localhost"
	// Use connection to be passed to listen for data push/pull requests
	analytics.ConnectAnalytics()
	println("Connected to Google Analytics")
	request := analytics.MakeRequest()
	results := analytics.AnalyticsReportingService.Reports.BatchGet(&request)
	response, err := results.Do()
	if nil != err {
		fmt.Print(err)
		os.Exit(1)
	}
	output := response.Reports[0].Data.Rows[0]
	if nil != err {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Printf("\n %v", output.Dimensions)
	fmt.Printf("\n %v", output.Metrics)
	// Listen and serve HTTP requests from the PHP app
}
