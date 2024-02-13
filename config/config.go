package config

import (
	"sessionmanagement/api/model/dto"
	"fmt"

	"github.com/spf13/viper"
)

var DatabaseConfig dto.Database

func LoadEnv() {

	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath("../.config/")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	if err := viper.Unmarshal(&DatabaseConfig); err != nil {
		fmt.Println("Error While Decoding .env File")
	}

}
