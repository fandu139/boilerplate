package aws

import (
	"fmt"

	"github.com/orn-id/orn-mn-boilerplate-go/src/utils/awslib"
	"github.com/orn-id/orn-mn-boilerplate-go/src/utils/awslib/sqs"
	"github.com/orn-id/orn-mn-boilerplate-go/src/utils/awslib/storage"
)

// AwsPackage ...
type AwsPackage struct {
	SQS sqs.SQSManagerInterface
	S3  storage.S3ManagerInterface
}

// AwsPackageHandler ...
func AwsPackageHandler() *AwsPackage {
	return &AwsPackage{
		SQS: awslib.AWSHandler().SQS(),
		S3:  awslib.AWSHandler().S3(),
	}
}

// AwsPackageInterface ..
type AwsPackageInterface interface {
	S3Version()
	SQSVersion()
}

// S3Version ...
func (aws *AwsPackage) S3Version() {
	client := aws.S3.New()
	fmt.Println(client.APIVersion)
}

// SQSVersion ..
func (aws *AwsPackage) SQSVersion() {
	client := aws.SQS.New()
	fmt.Println(client.APIVersion)
}
