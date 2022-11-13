package configs

import (
	"github.com/spf13/viper"
)

func InitConfig(path, file string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(file)
	return viper.ReadInConfig()
}
