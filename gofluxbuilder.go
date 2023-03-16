package gofluxbuilder

import (
	"context"
	"fmt"

	"github.com/influxdata/influxdb-client-go/v2/api"
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
	Max    Builder
	Min    Builder
	Mean   Builder
	Sort   Builder
	Limit  Builder
	Count  Builder
	Group  Builder
}

// QueryBuilder is the persistent struct that allows us to hold the information
type QueryBuilder struct {
	query *Query
}

func throwError(name string, data interface{}) error {
	return fmt.Errorf("%s: %v", name, data)
}

// NewGoFluxQueryBuilder is the constructor to build flux queries
func NewGoFluxQueryBuilder() *QueryBuilder {
	return &QueryBuilder{query: &Query{
		From:   FromBuilder{},
		Range:  RangeBuilder{},
		Filter: nil,
		Max:    nil,
		Min:    nil,
		Mean:   nil,
		Sort:   nil,
		Limit:  nil,
		Count:  nil,
	}}
}

// From allows to define from parameters of flux query
func (q *QueryBuilder) From(from Builder) *QueryBuilder {
	q.query.From = from
	return q
}

// Range allows to define range parameters of flux query
func (q *QueryBuilder) Range(r Builder) *QueryBuilder {
	q.query.Range = r
	return q
}

// Filter allows to define filters of flux query
func (q *QueryBuilder) Filter(filter Builder) *QueryBuilder {
	q.query.Filter = filter
	return q
}

// Max allows to define max of flux query
func (q *QueryBuilder) Max(max ...Builder) *QueryBuilder {
	if len(max) == 0 {
		q.query.Max = MaxBuilder{}
		return q
	}
	q.query.Max = max[0]
	return q
}

// Min allows to define min of flux query
func (q *QueryBuilder) Min(min ...Builder) *QueryBuilder {
	if len(min) == 0 {
		q.query.Min = MinBuilder{}
		return q
	}
	q.query.Min = min[0]
	return q
}

// Mean allows to define mean of flux query
func (q *QueryBuilder) Mean(mean ...Builder) *QueryBuilder {
	if len(mean) == 0 {
		q.query.Mean = MeanBuilder{}
		return q
	}
	q.query.Mean = mean[0]
	return q
}

// Sort allows to define sort of flux query
func (q *QueryBuilder) Sort(sort Builder) *QueryBuilder {
	q.query.Sort = sort
	return q
}

// Limit allows to define limit of flux query
func (q *QueryBuilder) Limit(limit Builder) *QueryBuilder {
	q.query.Limit = limit
	return q
}

// Count allows to define count of flux query
func (q *QueryBuilder) Count(mean ...Builder) *QueryBuilder {
	if len(mean) == 0 {
		q.query.Mean = CountBuilder{}
		return q
	}
	q.query.Mean = mean[0]
	return q
}

// Group allows to define group of flux query
func (q *QueryBuilder) Group(group Builder) *QueryBuilder {
	q.query.Group = group
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
	if q.query.Filter != nil {
		err = q.query.Filter.Validate()
		if err != nil {
			return "", throwError(queryValidationError, err.Error())
		}
	}
	if q.query.Limit != nil {
		err = q.query.Filter.Validate()
		if err != nil {
			return "", throwError(queryValidationError, err.Error())

		}
	}
	var query string
	query += q.query.From.Parse()
	query += pipeGenerator()
	query += q.query.Range.Parse()
	if q.query.Filter != nil {
		query += pipeGenerator()
		query += q.query.Filter.Parse()
	}
	if q.query.Max != nil {
		query += pipeGenerator()
		query += q.query.Max.Parse()
	}
	if q.query.Min != nil {
		query += pipeGenerator()
		query += q.query.Min.Parse()
	}
	if q.query.Mean != nil {
		query += pipeGenerator()
		query += q.query.Mean.Parse()
	}
	if q.query.Sort != nil {
		query += pipeGenerator()
		query += q.query.Sort.Parse()
	}
	if q.query.Limit != nil {
		query += pipeGenerator()
		query += q.query.Limit.Parse()
	}
	if q.query.Count != nil {
		query += pipeGenerator()
		query += q.query.Count.Parse()
	}
	if q.query.Group != nil {
		query += pipeGenerator()
		query += q.query.Group.Parse()
	}
	return query, nil
}
