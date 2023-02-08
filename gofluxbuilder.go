package gofluxbuilder

import (
	"context"
	"errors"
	"fmt"
	"github.com/influxdata/influxdb-client-go/api"
)

type Builder interface {
	Validate() error
	Parse() string
}

type Query struct {
	From   Builder
	Range  Builder
	Filter Builder
}

type QueryBuilder struct {
	query *Query
}

func throwError(name string, data interface{}) error {
	return errors.New(fmt.Sprintf("%s: %v", name, data))
}

func NewGoFluxQueryBuilder() *QueryBuilder {
	return &QueryBuilder{query: &Query{}}
}

func (q *QueryBuilder) From(from Builder) *QueryBuilder {
	q.query.From = from
	return q
}

func (q *QueryBuilder) Range(r Builder) *QueryBuilder {
	q.query.Range = r
	return q
}

func (q *QueryBuilder) Filter(filter Builder) *QueryBuilder {
	q.query.Filter = filter
	return q
}

func (q *QueryBuilder) Query(ctx context.Context, client *api.QueryAPI) (res *api.
	QueryTableResult, err error) {
	query, err := q.Build()
	if err != nil {
		return nil, throwError(queryRequestError, err)
	}
	return makeQuery(ctx, client, query)
}

func (q *QueryBuilder) Build() (string, error) {
	err := q.query.From.Validate()
	if err != nil {
		return "", throwError(queryValidationError, err.Error())
	}
	err = q.query.Range.Validate()
	if err != nil {
		return "", throwError(queryValidationError, err.Error())
	}
	err = q.query.Filter.Validate()
	if err != nil {
		return "", throwError(queryValidationError, err.Error())
	}
	var query string
	query += q.query.From.Parse()
	query += pipeGenerator()
	query += q.query.Range.Parse()
	query += pipeGenerator()
	query += q.query.Filter.Parse()
	return query, nil
}
