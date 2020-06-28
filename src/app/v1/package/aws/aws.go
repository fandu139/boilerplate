package aws

import (
	"fmt"

	"github.com/sofyan48/boilerplate/src/utils/awslib"
	"github.com/sofyan48/boilerplate/src/utils/awslib/sqs"
	"github.com/sofyan48/boilerplate/src/utils/awslib/storage"
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
