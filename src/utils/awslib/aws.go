package awslib

import (
	"github.com/sofyan48/boilerplate/src/utils/awslib/sqs"
	"github.com/sofyan48/boilerplate/src/utils/awslib/storage"
)

// AWS ...
type AWS struct{}

// AWSHandler ...
func AWSHandler() *AWS {
	return &AWS{}
}

// AWSInterface ...
type AWSInterface interface {
	S3() *storage.S3Manager
	SQS() *sqs.SQSManager
}

// S3 ...
func (a *AWS) S3() *storage.S3Manager {
	return storage.S3ManagerHandler()
}

// SQS ...
func (a *AWS) SQS() *sqs.SQSManager {
	return sqs.SQSManagerHandler()
}
