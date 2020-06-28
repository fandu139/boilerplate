package sqs

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/sofyan48/boilerplate/src/utils/awslib/credential"
)

// SQSManager ...
type SQSManager struct {
	AWS credential.SessionInterface
}

// SQSManagerHandler ...
func SQSManagerHandler() *SQSManager {
	return &SQSManager{
		AWS: credential.SessionHandler(),
	}
}

// SQSManagerInterface ...
type SQSManagerInterface interface {
	New() *sqs.SQS
}

// New ...
func (handler *SQSManager) New() *sqs.SQS {
	cfg := handler.AWS.Sessions()
	return sqs.New(session.New(), cfg)
}
