package awslib

import (
	"github.com/fandu139/boilerplate/src/utils/awslib/sqs"
	"github.com/fandu139/boilerplate/src/utils/awslib/storage"
)

// Contract ...
type Contract interface {
	S3() *storage.S3Manager
	SQS() *sqs.SQSManager
}
