package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	RDSHostName    string `mapstructure:"RDS_HOSTNAME"`
	RDSPort        string `mapstructure:"RDS_PORT"`
	RDSUser        string `mapstructure:"RDS_USER"`
	RDSPassword    string `mapstructure:"RDS_PASSWORD"`
	RDSDBName      string `mapstructure:"RDS_DB_NAME"`
	Secret         string `mapstructure:"SECRET"`
	AllowedOrigins string `mapstructure:"ALLOW_ORIGIN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
