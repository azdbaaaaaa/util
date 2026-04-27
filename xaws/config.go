package xaws

type Config struct {
	AccessKey string `json:"access_key" mapstructure:"access_key"`
	SecretKey string `json:"secret_key" mapstructure:"secret_key"`
	Region    string `json:"region" mapstructure:"region"`
}
