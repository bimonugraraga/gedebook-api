package env

type DatabaseConfig struct {
	Host       string `mapstructure:"DB_HOST"`
	Port       string `mapstructure:"DB_PORT"`
	User       string `mapstructure:"DB_USER"`
	Password   string `mapstructure:"DB_PASSWORD"`
	Name       string `mapstructure:"DB_NAME"`
	SSLMode    string `mapstructure:"DB_SSL_MODE"`
	URL        string `mapstructure:"DB_URL"`
	Debug      bool   `mapstructure:"DB_DEBUG"`
	DebugLevel int    `mapstructure:"DB_DEBUG_LEVEL"`
}
