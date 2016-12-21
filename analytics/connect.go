package analytics

import (
	"os"

	log "github.com/tylerconlee/SummitAPI/log"
	ga "google.golang.org/api/analytics/v3"
)

var Service *ga.Service

func ConnectAnalytics() *ga.Service {
	log.InitLog("Analytics")

	var err error
	oauthHTTPClient := OauthConnect()
	Service, err = ga.New(oauthHTTPClient)
	log.Logger.Debug("oauth client connected to Google Analytics")
	if nil != err {
		os.Exit(1)
	}
	return Service
}
