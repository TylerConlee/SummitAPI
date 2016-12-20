package analytics

import "google.golang.org/api/analyticsreporting/v4"
import "fmt"

func MakeRequest() analyticsreporting.GetReportsRequest {

	metric := analyticsreporting.Metric{}
	metric.Expression = "ga:users"
	metrics := []*analyticsreporting.Metric{}
	metrics = append(metrics, &metric)

	dimension := analyticsreporting.Dimension{}
	dimension.Name = "ga:browser"
	dimensions := []*analyticsreporting.Dimension{}
	dimensions = append(dimensions, &dimension)
	fmt.Printf("Making request")

	request := analyticsreporting.ReportRequest{}
	request.Metrics = metrics
	request.Dimensions = dimensions
	request.ViewId = "93745560"

	requests := []*analyticsreporting.ReportRequest{}
	requests = append(requests, &request)

	report := analyticsreporting.GetReportsRequest{}
	report.ReportRequests = requests
	return report
}

func ResponseParser(response []*analyticsreporting.ReportRow) {
	for _, row := range response {
		fmt.Printf("Dimensions: %v", row.Dimensions)
		for _, metric := range row.Metrics {
			fmt.Printf("Metric: %v", metric.Values)
		}
	}
}
