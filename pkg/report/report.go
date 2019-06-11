package report

import (
	"encoding/json"

	"github.com/hekike/node-report-analytics/pkg/classify"
	"github.com/hekike/node-report-analytics/pkg/model"
)

// FromJSON parses a JSON diagnostic report
func FromJSON(jsonReport string) (*model.DiagnosticReport, error) {
	report := model.DiagnosticReport{}
	err := json.Unmarshal([]byte(jsonReport), &report)

	if err == nil {
		hash, err := classify.HashStack(report.JavaScriptStack)
		if err != nil {
			return nil, err
		}
		report.Hash = hash
	}
	return &report, err
}
