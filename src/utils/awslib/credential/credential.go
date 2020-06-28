package credential

import (
	"os"

	AWS "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

type Session struct{}

// SessionHandler ...
func SessionHandler() *Session {
	return &Session{}
}

// SessionInterface ...
type SessionInterface interface {
	Sessions() *AWS.Config
}

// Sessions ...
func (aws *Session) Sessions() *AWS.Config {
	creds := credentials.NewStaticCredentials(
		os.Getenv("AWS_ACCESS_KEY"),
		os.Getenv("AWS_ACCESS_SECRET"), "")
	creds.Get()
	cfgAws := AWS.NewConfig().WithRegion(os.Getenv("AWS_ACCESS_AREA")).WithCredentials(creds)

	return cfgAws
}
