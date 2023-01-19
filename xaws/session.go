package xaws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// NewSession returns a new session
func NewSession(c *Config) (sess *session.Session) {
	creds := credentials.NewStaticCredentials(c.AccessKey, c.SecretKey, "")
	conf := &aws.Config{
		Credentials:      creds,
		Region:           aws.String(c.Region),
		S3ForcePathStyle: aws.Bool(true),
	}
	return session.Must(session.NewSession(conf))
}
