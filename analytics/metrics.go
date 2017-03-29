package analytics

import (
	c "github.com/tylerconlee/SummitAPI/config"
)

type MetricTypes struct {
	Dimensions string
	Metrics string
}

func Query() {

}

func getConfig() (params MetricTypes) {
	config := c.InitConfig()
	params.Dimensions = config.Analytics.Dimensions
	params.Metrics = config.Analytics.Metrics
	return params
}