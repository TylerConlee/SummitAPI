package analytics

import "google.golang.org/api/analyticsreporting/v4"

func MakeRequest() {
	metric := &analyticsreporting.Metric{}
	metric.Expression = "ga:users"
	request := analyticsreporting.ReportRequest{}
	request.Metrics = metric
}
