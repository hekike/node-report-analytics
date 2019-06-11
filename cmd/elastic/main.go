package main

import (
	"context"
	"fmt"
	"os"

	"github.com/hekike/node-report-analytics/pkg/db"
	"github.com/hekike/node-report-analytics/pkg/report"
)

func main() {
	dir := os.Args[1]
	reports, err := report.ReadDir(dir)

	if err != nil {
		panic(err)
	}

	for _, report := range reports {
		fmt.Printf(
			"%s, %s, %s\n",
			report.Hash,
			report.Header.DumpEventTime,
			report.JavaScriptStack.Message,
		)
	}

	ctx := context.Background()

	// init Elastic client
	elasticClient, err := db.NewElasticClient(ctx)
	if err != nil {
		panic(err)
	}

	err = db.CreateReportIndexIfDoesNotExist(ctx, elasticClient)
	if err != nil {
		panic(err)
	}

	err = db.InsertReports(ctx, elasticClient, reports)
	if err != nil {
		panic(err)
	}
}
