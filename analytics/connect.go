package analytics

import (
	"os"

	"google.golang.org/api/analyticsreporting/v4"
)

var AnalyticsReportingService *analyticsreporting.Service

func connectAnalytics() {
	var err error
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./gcredentials.json")
	oauthHTTPClient := OauthConnect()
	AnalyticsReportingService, err = analyticsreporting.New(oauthHTTPClient)
	if nil != err {
		return
	}
}
