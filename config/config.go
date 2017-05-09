// Creates a struct with the basic config with default configuration options
// set. Designed to be overridden potentially via configuration file or env var.

package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

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

	// Set the log level. Default is set to Info
	LogLevel string
}

// NewConfig creates an instance of the Config struct with the default values.
// Values can be overridden after initialized.
func NewConfig() Config {
	c := Config{}

	c.AnalyticsAPIKey = "testkey"
	c.DatabaseConnection.DatabaseHost = "localhost"
	c.DatabaseConnection.DatabasePort = "3306"
	c.DatabaseConnection.DatabaseName = "summit-ppc"
	c.LogLevel = "info"

	return c
}

func InitConfig() (Config, error) {
	// Set the default path for the configuration file
	// TODO: include the ability to set this value with a command line arg?
	var configfile = "./config.toml"

	// Verify that the file exists and load into memory
	_, err := os.Stat(configfile)

	// If file does not exist, return a basic config with default values
	if err != nil {
		//log.Println("Config file is missing: ", configfile)
		// TODO: Create a config.toml if there isn't one present.
		return NewConfig(), nil
	}

	// Start with a default configuration
	var config = NewConfig()

	// Decode the TOML configuration file and use it to overwrite the default
	// configuation values
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	return config, nil
}
