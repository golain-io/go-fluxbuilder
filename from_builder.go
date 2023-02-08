package gofluxbuilder

type FromBuilder struct {
	Bucket   string
	BucketID string
	Host     string
	Org      string
	OrgID    string
	Token    string
}

// Validate is the validation impl of Builder for FromBuilder
func (f FromBuilder) Validate() error {
	if f.Bucket != "" && f.BucketID != "" {
		return throwError(fromBuilderError, "bucket and bucketID should be mutually"+
			" exclusive")
	}
	if f.Bucket == "" && f.BucketID == "" {
		return throwError(fromBuilderError, "from needs bucket or bucketID")
	}
	return nil
}

// Parse is the parsing impl of Builder for FromBuilder
func (f FromBuilder) Parse() string {
	return commonGenerator("from", f)
}
