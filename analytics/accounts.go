package analytics

import (
	"fmt"

	"github.com/tylerconlee/SummitAPI/log"

	ga "google.golang.org/api/analytics/v3"
)

// Each account is stored as a Profile struct, with an individual profile ID,
// property and account. Accounts can have multiple properties, and each
// property can have multiple profiles.
// https://github.com/google/google-api-go-client/blob/master/analytics/v3/analytics-gen.go

type Profile struct {
	AccountID  string
	PropertyID string
	ProfileID  string
}

// Profiles is a slice of all available individual Profiles stored.
var Profiles []*Profile

func GetID() {
	// Start up a new module in the logger for Analytics
	log.InitLog("Analytics")

	// Create an instance of the management service
	// https://developers.google.com/analytics/devguides/config/mgmt/v3/
	var manage = ga.NewManagementService(Service)
	c, err := manage.Accounts.List().Do()
	if nil != err {
		fmt.Printf("%v", err)
	}
	log.Logger.Debug("Accounts pulled, response: ", c.ServerResponse, "results: ", c.TotalResults)

	for _, item := range c.Items {
		getProperties(item.Id)
	}
	log.Logger.Info("Analytics profiles retrieved")
}

func getProperties(account string) {
	// Create an instance of the management service
	// https://developers.google.com/analytics/devguides/config/mgmt/v3/
	var manage = ga.NewManagementService(Service)
	c, err := manage.Webproperties.List(account).Do()
	if nil != err {
		fmt.Printf("%v", err)
	}
	log.Logger.Debug("Property pulled for account: ", account, " response: ", c.ServerResponse, "results: ", c.TotalResults)
	for _, item := range c.Items {
		getProfiles(item.AccountId, item.Id)
	}
}

func getProfiles(account string, property string) {
	// Create an instance of the management service
	// https://developers.google.com/analytics/devguides/config/mgmt/v3/
	var manage = ga.NewManagementService(Service)
	c, err := manage.Profiles.List(account, property).Do()
	if nil != err {
		fmt.Printf("%v", err)
	}
	log.Logger.Debug("Profile pulled for account: ", account, " property: ", property, " response: ", c.ServerResponse, "results: ", c.TotalResults)
	for _, item := range c.Items {
		var profile Profile
		profile.AccountID = item.AccountId
		profile.PropertyID = item.InternalWebPropertyId
		profile.ProfileID = item.Id
		Profiles = append(Profiles, &profile)

	}
}
