package configuracion

type Config struct {
	HttpAddress  string `mapstructure:"HTTP_ADDRESS"`
	AllowOrigins string `mapstructure:"ALLOW_ORIGINS"`
}
