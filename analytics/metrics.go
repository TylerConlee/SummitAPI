package analytics

import (
	c "github.com/tylerconlee/SummitAPI/config"
	"github.com/tylerconlee/SummitAPI/log"
	ga "google.golang.org/api/analytics/v3"
)

type MetricTypes struct {
	Dimensions string
	Metrics string
}

var (
	startDate = "30daysAgo"
	endDate = "yesterday"
	dimset MetricTypes
)

func loadQuery() {
	config := c.InitConfig()
	dimset.Dimensions = config.Analytics.Dimensions
	dimset.Metrics = config.Analytics.Metrics
}

func Query(profile string) {
	g := ga.NewDataGaService(Service)
	req := g.Get("ga:" + profile, startDate, endDate, dimset.Metrics)
	data, err := req.Do()
	if nil != err {
		log.Logger.Fatal("Unable to process Analytics results")
	}

	log.Logger.Debug(data)

}

