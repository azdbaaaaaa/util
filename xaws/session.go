package xaws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// NewSession returns a new session
func NewSession(c *Config) (cli *s3.S3) {
	creds := credentials.NewStaticCredentials(c.AccessKey, c.SecretKey, "")
	conf := &aws.Config{
		Credentials:      creds,
		Region:           aws.String(c.Region),
		S3ForcePathStyle: aws.Bool(true),
	}
	sess := session.Must(session.NewSession(conf))
	cli = s3.New(sess)
	return
}
