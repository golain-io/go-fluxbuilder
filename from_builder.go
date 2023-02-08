package gofluxbuilder

import "errors"

type FromBuilder struct {
	Bucket   string
	BucketID string
	Host     string
	Org      string
	OrgID    string
	Token    string
}

func (f FromBuilder) Validate() error {
	if f.Bucket != "" && f.BucketID != "" {
		return errors.New("bucket and bucketID should be mutually exclusive")
	}
	if f.Bucket == "" && f.BucketID == "" {
		return errors.New("from needs bucket or bucketID")
	}
	return nil
}

func (f FromBuilder) Parse() string {
	return commonGenerator("from", f)
}