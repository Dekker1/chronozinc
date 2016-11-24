package main

import "github.com/spf13/viper"

func setDefaults() {

	viper.SetDefault("Processes", 1)
	viper.SetDefault("RawDir", "raw")
	viper.SetDefault("Solvers", map[string]interface{}{})

}
