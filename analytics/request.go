package analytics

import "google.golang.org/api/analyticsreporting/v4"
import "fmt"

func MakeRequest() analyticsreporting.GetReportsRequest {

	metric := analyticsreporting.Metric{}
	metric.Expression = "ga:users"
	metrics := []*analyticsreporting.Metric{}
	metrics = append(metrics, &metric)
	fmt.Printf("Making request")

	request := analyticsreporting.ReportRequest{}
	request.Metrics = metrics
	request.ViewId = "93745560"

	requests := []*analyticsreporting.ReportRequest{}
	requests = append(requests, &request)

	report := analyticsreporting.GetReportsRequest{}
	report.ReportRequests = requests
	return report
}
