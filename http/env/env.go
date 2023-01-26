package env

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/spf13/viper"
)

type Configuration struct {
	AppPort      string         `mapstructure:"APP_PORT"`
	AppName      string         `mapstructure:"APP_NAME"`
	GinMode      string         `mapstructure:"GIN_MODE"`
	DBDebug      bool           `mapstructure:"DB_DEBUG"`
	DBDebugLevel int            `mapstructure:"DB_DEBUG_LEVEL"`
	Database     DatabaseConfig `mapstructure:",squash"`
}

var config *Configuration
var once sync.Once

var requiredEnvs = []string{
	"APP_HOST",
	"APP_PORT",
	"APP_NAME",
	"DB_HOST",
	"DB_PORT",
	"DB_USER",
	"DB_PASSWORD",
	"DB_NAME",
	"DB_SSL_MODE",
	"DB_DEBUG",
	"DB_DEBUG_LEVEL",
	"DB_URL",
	"GIN_MODE",
}

func Init() *Configuration {
	once.Do(func() {
		_, b, _, _ := runtime.Caller(0)
		basePath := filepath.Join(filepath.Dir(b), "..")

		viper.AddConfigPath(basePath)
		viper.AddConfigPath(fmt.Sprintf("%s/config", basePath))

		viper.SetConfigName(".dev")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("error while loading env file: ", err)
		}

		checkRequiredValues()
		config = new(Configuration)
		if err := viper.Unmarshal(config); err != nil {
			log.Fatal("error in unmarshalling the config: ", err)
		}
	})
	return config
}
func Config() *Configuration {
	return config
}

func checkRequiredValues() {
	for _, k := range requiredEnvs {
		if !viper.IsSet(k) {
			panic(k + ": Required key is not present")
		}
	}
}
