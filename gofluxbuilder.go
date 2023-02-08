package gofluxbuilder

import (
	"context"
	"errors"
	"fmt"
	"github.com/influxdata/influxdb-client-go/api"
)

// Builder - The common interface of all various Builder structs
type Builder interface {
	Validate() error
	Parse() string
}

type Query struct {
	From   Builder
	Range  Builder
	Filter Builder
}

// QueryBuilder is the persistent struct that allows us to hold the information
type QueryBuilder struct {
	query *Query
}

func throwError(name string, data interface{}) error {
	return errors.New(fmt.Sprintf("%s: %v", name, data))
}

// NewGoFluxQueryBuilder is the constructor to build flux queries
func NewGoFluxQueryBuilder() *QueryBuilder {
	return &QueryBuilder{query: &Query{}}
}

// From allows to define from parameters of flux query
func (q *QueryBuilder) From(from Builder) *QueryBuilder {
	q.query.From = from
	return q
}

// Range allows to define range paramteres of flux query
func (q *QueryBuilder) Range(r Builder) *QueryBuilder {
	q.query.Range = r
	return q
}

// Filter allows to define filters of flux query
func (q *QueryBuilder) Filter(filter Builder) *QueryBuilder {
	q.query.Filter = filter
	return q
}

// Query makes the request and executes the flux query on influxDB
func (q *QueryBuilder) Query(ctx context.Context, client *api.QueryAPI) (res *api.
	QueryTableResult, err error) {
	query, err := q.Build()
	if err != nil {
		return nil, throwError(queryRequestError, err)
	}
	return makeQuery(ctx, client, query)
}

// Build generates the flux query template as a string with validations
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
