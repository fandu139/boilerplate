package awslib

import (
	"github.com/fandu139/boilerplate/src/utils/awslib/sqs"
	"github.com/fandu139/boilerplate/src/utils/awslib/storage"
)

// AWS ...
type aws struct{}

// New ...
func New() Contract {
	return &aws{}
}

// S3 ...
func (a *aws) S3() *storage.S3Manager {
	return storage.S3ManagerHandler()
}

// SQS ...
func (a *aws) SQS() *sqs.SQSManager {
	return sqs.SQSManagerHandler()
}
