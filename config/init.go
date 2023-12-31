package config

import (
	"log"

	"github.com/spf13/viper"
)

const Directory = "./config"

func getConfigFiles() []string {
	return []string{"parser_config", "bot_config", "bot_messages"}
}

func Init() {
	viper.AddConfigPath(Directory)

	for _, filePath := range getConfigFiles() {
		viper.SetConfigName(filePath)
		err := viper.MergeInConfig()
		if err != nil {
			log.Fatal(err)
		}
	}
}
