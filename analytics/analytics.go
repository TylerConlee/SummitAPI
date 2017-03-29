package analytics

import (
	"fmt"
)

func Init() {
	// Run the ConnectAnalytics function that authenticates with Google using
	// oAuth and preps the Service for use
	ConnectAnalytics()

	// Get a list of the account, property and profile IDs from Google
	// analytics for the email address that's been authenticated
	// TODO: Add email config option for account changes
	GetID()

	// Output the grabbed profiles
	fmt.Printf("%+v", Profiles[0])
}