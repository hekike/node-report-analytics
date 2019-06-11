package db

import (
	"context"

	"github.com/hekike/node-report-analytics/pkg/model"
	elastic "github.com/olivere/elastic/v7"
)

const (
	indexName = "reports_index"
	docType   = "report"
)

// CreateReportIndexIfDoesNotExist ...
func CreateReportIndexIfDoesNotExist(
	ctx context.Context,
	client *elastic.Client,
) error {
	err := CreateIndexIfDoesNotExist(ctx, client, indexName)
	return err
}

// InsertReports ...
func InsertReports(
	ctx context.Context,
	elasticClient *elastic.Client,
	reports []*model.DiagnosticReport,
) error {
	for _, report := range reports {
		_, err := elasticClient.Index().Index(indexName).Type(docType).BodyJson(report).Do(ctx)
		if err != nil {
			return err
		}
	}

	// Flush data (need for refreshing data in index) after this command possible to do get.
	elasticClient.Flush().Index(indexName).Do(ctx)

	return nil
}
