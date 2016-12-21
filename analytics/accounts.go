package analytics

import (
	"fmt"

	"github.com/tylerconlee/SummitAPI/log"

	ga "google.golang.org/api/analytics/v3"
)

type Profile struct {
	accountID  string
	propertyID string
	profileID  string
}

var Profiles []Profile

func GetID() {
	log.InitLog("Analytics")
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
	var manage = ga.NewManagementService(Service)
	c, err := manage.Profiles.List(account, property).Do()
	if nil != err {
		fmt.Printf("%v", err)
	}
	log.Logger.Debug("Profile pulled for account: ", account, " property: ", property, " response: ", c.ServerResponse, "results: ", c.TotalResults)
	for _, item := range c.Items {
		var profile Profile
		profile.accountID = item.AccountId
		profile.propertyID = item.InternalWebPropertyId
		profile.profileID = item.Id
		Profiles = append(Profiles, profile)

	}
}
