package gofluxbuilder

import (
	"context"

	"github.com/influxdata/influxdb-client-go/v2/api"
)

func makeQuery(ctx context.Context, client *api.QueryAPI, query string) (res *api.QueryTableResult, err error) {
	return (*client).Query(ctx, query)
}
