package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.AddConfigPath("config")

	viper.SetConfigName("env")

	viper.SetConfigType("json")

	// Searches for config file in given paths and read it
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Confirm which config file is used
	fmt.Printf("Using s: %s\n", viper.ConfigFileUsed())
}
