package analytics

import (
	"os"

	ga "google.golang.org/api/analytics/v3"
)

var Service *ga.Service

func ConnectAnalytics() *ga.Service {
	var err error
	//os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./gcredentials.json")
	oauthHTTPClient := OauthConnect()
	Service, err = ga.New(oauthHTTPClient)
	if nil != err {
		os.Exit(1)
	}
	return Service
}
