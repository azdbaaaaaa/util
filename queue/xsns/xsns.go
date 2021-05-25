package xsns

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type Config struct {
	AccessKey string `json:"access_key" mapstructure:"access_key"`
	SecretKey string `json:"secret_key" mapstructure:"secret_key"`
	Region    string `json:"region" mapstructure:"region"`
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
