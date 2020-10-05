package storage

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/fandu139/boilerplate/src/utils/awslib/credential"
)

// S3Manager ...
type S3Manager struct {
	AWS credential.SessionInterface
}

// S3ManagerHandler ...
func S3ManagerHandler() *S3Manager {
	return &S3Manager{
		AWS: credential.SessionHandler(),
	}
}

// S3ManagerInterface ...
type S3ManagerInterface interface {
	New() *s3.S3
}

// New ...
func (handler *S3Manager) New() *s3.S3 {
	cfg := handler.AWS.Sessions()
	return s3.New(session.New(), cfg)
}
