package model

import (
	"time"
)

// AggregatedDiagnosticReport stores aggregated data about JavaScript stacks
type AggregatedDiagnosticReport struct {
	Message        string
	Count          int
	FirstOccurence time.Time
	LastOccurence  time.Time
}

// AggregateDiagnosticReports aggregate reports by hash
func AggregateDiagnosticReports(
	reports []*DiagnosticReport,
) map[string]*AggregatedDiagnosticReport {
	ad := make(map[string]*AggregatedDiagnosticReport)

	// Aggregate reports
	for _, report := range reports {
		hash := report.Hash

		// Already exists by hash
		if ad[hash] == nil {
			ad[hash] = &AggregatedDiagnosticReport{
				Message:        report.JavaScriptStack.Message,
				Count:          1,
				FirstOccurence: report.Header.DumpEventTime,
				LastOccurence:  report.Header.DumpEventTime,
			}
		} else {
			eventTime := report.Header.DumpEventTime
			ad[hash].Count++
			if ad[hash].FirstOccurence.After(eventTime) {
				ad[hash].FirstOccurence = eventTime
			}
			if ad[hash].LastOccurence.Before(eventTime) {
				ad[hash].LastOccurence = eventTime
			}
		}
	}

	return ad
}
