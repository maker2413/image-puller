package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/image-puller/")
	viper.AddConfigPath("$HOME/.config/image-puller/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if !viper.IsSet("images") {
		panic(fmt.Errorf("fatal no images provided in config file!"))
	}

	images := viper.GetStringSlice("images")

	fmt.Println(images)
}
