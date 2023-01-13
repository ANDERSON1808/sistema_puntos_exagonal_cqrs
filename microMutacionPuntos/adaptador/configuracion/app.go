package configuracion

type Config struct {
	HttpAddress  string `mapstructure:"HTTP_ADDRESS"`
	AllowOrigins string `mapstructure:"ALLOW_ORIGINS"`

	BootstrapServer  string `mapstructure:"KAFKA_BOOTSTRAPSERVERS"`
	SecurityProtocol string `mapstructure:"SECURITY_PROTOCOL"`
	SaslUsername     string `mapstructure:"KAFKA_CLUSTER_API_KEY"`
	SaslPassword     string `mapstructure:"KAFKA_CLUSTER_API_SECRET"`

	TopicNotificationGeneral string `mapstructure:"TOPIC_NOTIFICATION_GENERAL"`
}
