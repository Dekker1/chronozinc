package settings

import "github.com/spf13/viper"

// SetViperDefaults set the default settings for the viper handler
func SetViperDefaults() {

	viper.SetDefault("processes", 1)
	viper.SetDefault("mznfzn", "mzn-fzn")
	viper.SetDefault("raw_dir", "raw")
	viper.SetDefault("flags", "--verbose --statistics")
	viper.SetDefault("output", "benchmark.csv")
	viper.SetDefault("comma", ",")
	viper.SetDefault("crlf", false)

}
