package terminal

import (
	"os"

	"github.com/hekike/node-report-analytics/pkg/model"
	"github.com/jedib0t/go-pretty/table"
)

var timeFormat = "2006-01-02 15:04"
var messagePrintLengthMax = 25


// RenderAggregatedDiagnosticReport renders aggregated diagnostic reports
func RenderAggregatedDiagnosticReport (
	aggregated map[string]*model.AggregatedDiagnosticReport,
) {
	// Create console table
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"Hash",
		"Count",
		"Message",
		"First Occurence",
		"Last Occurence",
	})
	for key, value := range aggregated {
		messagePrintLength := len(value.Message)
		if len(value.Message) > messagePrintLengthMax {
			messagePrintLength = messagePrintLengthMax
		}
		t.AppendRow([]interface{}{
			// use short hash
			key[:7],
			value.Count,
			// limit, message can be long
			value.Message[:messagePrintLength],
			value.FirstOccurence.Format(timeFormat),
			value.LastOccurence.Format(timeFormat),
		})
	}
	t.Render()
}
