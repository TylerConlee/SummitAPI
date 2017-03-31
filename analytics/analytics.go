package analytics

import (
	log "github.com/tylerconlee/SummitAPI/log"

)

func Init() {
	// Start the log module for Analytics
	log.InitLog("Analytics")


	// Run the ConnectAnalytics function that authenticates with Google using
	// oAuth and preps the Service for use
	ConnectAnalytics()
	log.Logger.Debug("Analytics connected")

	// Get a list of the account, property and profile IDs from Google
	// analytics for the email address that's been authenticated
	// TODO: Add email config option for account changes
	GetID()
	log.Logger.Debug("IDs retrieved")

	// Output the grabbed profiles
	log.Logger.Debug(Profiles[1].AccountID)
	loadQuery()


}

func Harvest(){
	for _, profile := range Profiles {
		Query(profile.ProfileID)
	}
}