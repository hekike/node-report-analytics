package db

import (
	"context"
	"errors"

	elastic "github.com/olivere/elastic/v7"
)

// CreateIndexIfDoesNotExist ...
func CreateIndexIfDoesNotExist(
	ctx context.Context,
	client *elastic.Client,
	indexName string,
) error {
	exists, err := client.IndexExists(indexName).Do(ctx)
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	res, err := client.CreateIndex(indexName).Do(ctx)

	if err != nil {
		return err
	}

	if !res.Acknowledged {
		return errors.New("createIndex was not acknowledged")
	}

	return nil
}

// NewElasticClient ...
func NewElasticClient(ctx context.Context) (*elastic.Client, error) {
	client, err := elastic.NewSimpleClient()
	if err != nil {
		return nil, err
	}

	_, _, err = client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}
