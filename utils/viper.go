package utils

import (
	"strings"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// err := viper.ReadInConfig()
	// if err != nil {
	// 	log.Panic(err)
	// }
}
