// SummitAPI creates a web server that handles requests made by SummitDaemon
// Designed to run in parallel in a containerized fashion
// SummitAPI is required to be running in order to run SummitDaemon

package main

import "google.golang.org/api/analyticsreporting/v4"

func main() {
	// Set initial configurations, use overrides from environment variables
	config := NewConfig()

	// Use configuration to create a new connection
	config.DatabaseConnection.DatabaseHost = "localhost"
	// Use connection to be passed to listen for data push/pull requests
	connectAnalytics()
	request := analyticsreporting.GetReportsRequest{
		 // ReportRequests: Requests, each request will have a separate
    // response.
    // There can be a maximum of 5 requests. All requests should have the
    // same
    // `dateRanges`, `viewId`, `segments`, `samplingLevel`, and
    // `cohortGroup`.
    ReportRequests []*ReportRequest `json:"reportRequests,omitempty"`

    // ForceSendFields is a list of field names (e.g. "ReportRequests") to
    // unconditionally include in API requests. By default, fields with
    // empty values are omitted from API requests. However, any non-pointer,
    // non-interface field appearing in ForceSendFields will be sent to the
    // server regardless of whether the field is empty or not. This may be
    // used to include empty fields in Patch requests.
    ForceSendFields []string `json:"-"`

    // NullFields is a list of field names (e.g. "ReportRequests") to
    // include in API requests with the JSON null value. By default, fields
    // with empty values are omitted from API requests. However, any field
    // with an empty value appearing in NullFields will be sent to the
    // server as null. It is an error if a field in this list has a
    // non-empty value. This may be used to include null fields in Patch
    // requests.
    NullFields []string `json:"-"`
	}
	results := AnalyticsReportingService.Reports.BatchGet()
	// Listen and serve HTTP requests from the PHP app
}
