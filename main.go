package main

import (
	"path/filepath"

	"github.com/Phantas0s/gocket/cmd"
	"github.com/adrg/xdg"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	cmd.Execute()
}

func initConfig() {
	viper.AddConfigPath(filepath.Join(xdg.ConfigHome, "gocket"))
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	viper.AutomaticEnv()
	viper.ReadInConfig()
}
