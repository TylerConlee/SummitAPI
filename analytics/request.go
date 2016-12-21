package analytics

import (
	"fmt"

	ga "google.golang.org/api/analytics/v3"
)

type Profile struct {
	accountID  string
	propertyID string
	profileID  string
}

var Profiles []Profile

func GetID() {
	var manage = ga.NewManagementService(Service)
	c, err := manage.Accounts.List().Do()
	if nil != err {
		fmt.Printf("%v", err)
	}
	for _, item := range c.Items {
		GetProperties(item.Id)
	}
}

func GetProperties(account string) {
	var manage = ga.NewManagementService(Service)
	c, err := manage.Webproperties.List(account).Do()
	if nil != err {
		fmt.Printf("%v", err)
	}
	for _, item := range c.Items {
		GetProfiles(item.AccountId, item.Id)
	}
}

func GetProfiles(account string, property string) {
	var manage = ga.NewManagementService(Service)
	c, err := manage.Profiles.List(account, property).Do()
	if nil != err {
		fmt.Printf("%v", err)
	}
	for _, item := range c.Items {
		var profile Profile
		profile.accountID = item.AccountId
		profile.propertyID = item.WebPropertyId
		profile.profileID = item.Id
		Profiles = append(Profiles, profile)
	}
}
