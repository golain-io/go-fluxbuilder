# gofluxbuilder

A package to build and generate flux queries with ease!


## Usage

The current features include:

- From
- Range
- Filter 

The package is as easy to use as:

```go
result, error := gofluxbuilder.NewGoFluxQueryBuilder()
                 .From(gofluxbuilder.FromBuilder{Bucket: "birdy"})
                 .Range(gofluxbuilder.RangeBuilder{Start: "-4y"})
                 .Filter(gofluxbuilder.NewFilterBuilder().
			Equal("_measurement", "migration"))
                 .Query(context.Background(), &cl)
```

Currently the following are supported in `Filters`:

- `and`
- `or`
- `==`
- `>`
- `<`
- `>=`
- `<=`

