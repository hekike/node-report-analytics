package main

import (
	"os"

	"github.com/hekike/node-report-analytics/pkg/model"
	"github.com/hekike/node-report-analytics/pkg/report"
	"github.com/hekike/node-report-analytics/pkg/terminal"
)

var timeFormat = "2006-01-02 15:04"
var messagePrintLengthMax = 25

func main() {
	dir := os.Args[1]
	reports, err := report.ReadDir(dir)

	if err != nil {
		panic(err)
	}

	aggregatedReports := model.AggregateDiagnosticReports(reports)
	terminal.RenderAggregatedDiagnosticReport(aggregatedReports)
}

