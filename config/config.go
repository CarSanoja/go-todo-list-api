package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port    string
	CSVFile string
}

var config Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}

	config = Config{
		Port:    viper.GetString("port"),
		CSVFile: viper.GetString("csv_file"),
	}
}

func GetConfig() Config {
	return config
}
