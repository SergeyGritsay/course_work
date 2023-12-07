package app

import (
	"/enrollments-parser/pkg/config"

	logger "github.com/sirupsen/logrus"
)

func Start(configPath string) (err error) {
	logger.Println("Parser started")
	conf, err := config.ParseConfig(configPath)
	if err != nil {
		logger.Fatal(err)
		return err
	}

	logger.Println(conf)
	return err
}
