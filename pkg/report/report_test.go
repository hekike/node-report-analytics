package report

import (
	"testing"
	"time"
	"github.com/hekike/node-report-analytics/pkg/model"
	"github.com/hekike/node-report-analytics/pkg/classify"
	"github.com/stretchr/testify/assert"
)

func TestFromJSON(t *testing.T) {
	json := `
	{
		"header": {
			"event": "My Error",
			"dumpEventTime": "2019-06-09T12:45:42Z",
			"processId": 47640
		},
		"javaScriptStack": {
			"message": "My Error",
			"stack": ["foo", "bar"]
		}
	}`
	report, err := FromJSON(json)
	if err != nil {
		panic(err)
	}

	dumpEventTime, err := time.Parse(time.RFC3339, "2019-06-09T12:45:42Z")
	if err != nil {
		panic(err)
	}

	stack := model.JavaScriptStack{
		Message: "My Error",
		Stack: []string{"foo", "bar"},
	}
	hash, err := classify.HashStack(stack)
	if err != nil {
		panic(err)
	}
	expected := &model.DiagnosticReport{
		Hash: hash,
		Header: model.DiagnosticReportHeader{
			Event: "My Error",
			DumpEventTime: dumpEventTime,
			ProcessID: 47640,
			// we don't test for the rest of the fields here
			// ...
		},
		JavaScriptStack: stack,
	}
	assert.Equal(t, expected, report)
}
