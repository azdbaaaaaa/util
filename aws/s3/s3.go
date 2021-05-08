package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Config s3 config.
type Config struct {
	AccessKey string
	SecretKey string
	Region    string
	Force     bool
}

// NewS3Client new aws s3 and retry connection when has error.
func NewS3Client(c *Config) (cli *s3.S3) {
	creds := credentials.NewStaticCredentials(c.AccessKey, c.SecretKey, "")
	conf := &aws.Config{
		CredentialsChainVerboseErrors:     nil,
		Credentials:                       creds,
		Endpoint:                          nil,
		EndpointResolver:                  nil,
		EnforceShouldRetryCheck:           nil,
		Region:                            aws.String(c.Region),
		DisableSSL:                        nil,
		HTTPClient:                        nil,
		LogLevel:                          nil,
		Logger:                            nil,
		MaxRetries:                        nil,
		Retryer:                           nil,
		DisableParamValidation:            nil,
		DisableComputeChecksums:           nil,
		S3ForcePathStyle:                  aws.Bool(c.Force),
		S3Disable100Continue:              nil,
		S3UseAccelerate:                   nil,
		S3DisableContentMD5Validation:     nil,
		EC2MetadataDisableTimeoutOverride: nil,
		UseDualStack:                      nil,
		SleepDelay:                        nil,
		DisableRestProtocolURICleaning:    nil,
		EnableEndpointDiscovery:           nil,
		DisableEndpointHostPrefix:         nil,
	}
	sess := session.Must(session.NewSession(conf))
	cli = s3.New(sess)
	return
}
