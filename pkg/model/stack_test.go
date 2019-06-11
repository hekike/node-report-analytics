package model

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestAggregateDiagnosticReports(t *testing.T) {
	now := time.Now()
	then := now.AddDate(0, -1, 0)
	report1 := DiagnosticReport{
		Hash: "aaa",
		Header: DiagnosticReportHeader{DumpEventTime: now},
		JavaScriptStack: JavaScriptStack{Message: "My Error 1"},
	}
	report2 := DiagnosticReport{
		Hash: "aaa",
		Header: DiagnosticReportHeader{DumpEventTime: then},
		JavaScriptStack: JavaScriptStack{Message: "My Error 1"},
	}
	report3:= DiagnosticReport{
		Hash: "bbb",
		Header: DiagnosticReportHeader{DumpEventTime: then},
		JavaScriptStack: JavaScriptStack{Message: "My Error 2"},
	}

	// Expected
	expected := make(map[string]*AggregatedDiagnosticReport)
	expected["aaa"] = &AggregatedDiagnosticReport{
		Message:        "My Error 1",
		Count:          2,
		FirstOccurence: then,
		LastOccurence:  now,
	}
	expected["bbb"] = &AggregatedDiagnosticReport{
		Message:        "My Error 2",
		Count:          1,
		FirstOccurence: then,
		LastOccurence:  then,
	}

	result := AggregateDiagnosticReports([]*DiagnosticReport{
		&report1,
		&report2,
		&report3,
	})

	assert.Equal(t, expected, result)
}
