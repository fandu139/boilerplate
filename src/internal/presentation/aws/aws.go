package aws

import (
	"fmt"

	"github.com/fandu139/boilerplate/src/internal/presentation"
	"github.com/fandu139/boilerplate/src/utils/awslib"
	"github.com/fandu139/boilerplate/src/utils/awslib/sqs"
	"github.com/fandu139/boilerplate/src/utils/awslib/storage"
)

// AwsPackage ...
type aws struct {
	SQS sqs.SQSManagerInterface
	S3  storage.S3ManagerInterface
}

// New ...
func New() presentation.AWS {
	return &aws{
		SQS: awslib.New().SQS(),
		S3:  awslib.New().S3(),
	}
}

// S3Version ...
func (a *aws) S3Version() {
	client := a.S3.New()
	fmt.Println(client.APIVersion)
}

// SQSVersion ..
func (a *aws) SQSVersion() {
	client := a.SQS.New()
	fmt.Println(client.APIVersion)
}
