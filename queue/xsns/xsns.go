package xsns

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/azdbaaaaaa/util/xaws"
)

type Config struct {
	xaws.Config
	TopicArn string `json:"topic_arn" mapstructure:"topic_arn"`
}

// NewSnsClient new aws sns and retry connection when has error.
func NewSnsClient(c *Config) (cli *sns.SNS) {
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
		EC2MetadataDisableTimeoutOverride: nil,
		UseDualStack:                      nil,
		SleepDelay:                        nil,
		DisableRestProtocolURICleaning:    nil,
		EnableEndpointDiscovery:           nil,
		DisableEndpointHostPrefix:         nil,
	}
	sess := session.Must(session.NewSession(conf))
	cli = sns.New(sess)
	return
}
