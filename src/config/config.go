package config

import (
	"github.com/aynp/store-passwords/src/database"
	"github.com/spf13/viper"
	"log"
)

func SetConfig() {
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Could not read the config file: %v", err)
	}

	database.Connect(viper.Get("password_db_url").(string))
	database.AutoMigrate()
}
