// Creates a struct with the basic config with default configuration options
// set. Designed to be overridden potentially via configuration file or env var.

package main

// Config contains the possible configurable settings for the Summit application
type Config struct {

	// Provides API key for Google Analytics
	AnalyticsAPIKey string

	// Database connection info
	DatabaseConnection struct {
		// Store the DatabaseHost information
		DatabaseHost string
		// Store the DatabasePort information
		DatabasePort string
		// Store the DatabaseName information
		DatabaseName string
	}
}

// NewConfig creates an instance of the Config struct with the default values. // Values can be overridden after initialized.
func NewConfig() Config {
	c := Config{}

	c.AnalyticsAPIKey = ""
	c.DatabaseConnection.DatabaseHost = "localhost"
	c.DatabaseConnection.DatabasePort = "3306"
	c.DatabaseConnection.DatabaseName = "summit-ppc"

	return c
}
