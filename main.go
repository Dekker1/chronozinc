package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jjdekker/chronozinc/parsing"
	"github.com/jjdekker/chronozinc/runtime"
	"github.com/jjdekker/chronozinc/settings"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "chonozinc [description]",
	Short: "Test and time your minizinc models!",
	Long: `ChonoZinc is a small tool that wraps around mzn-fzn. It main purpose is to
	test MiniZinc models using different solvers. It run your your model on all or
	a sub-set of declared solvers and report on the statistics gathered. The data
	gathered is saved in every stage.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			switch filepath.Ext(file) {
			case ".mzn":
				viper.Set("models", append(viper.GetStringSlice("models"), file))
			case ".dzn":
				viper.Set("data", append(viper.GetStringSlice("data"), file))
			default:
				viper.SetConfigFile(args[0])
				fmt.Println("Using config file:", args[0])
				err := viper.MergeInConfig()
				if err != nil {
					panic(err)
				}
			}
		}

		solvers := settings.SolversFromViper()
		instances := settings.InstancesFromViper()
		runtime.RunAll(solvers, instances)
		parsing.ParseAll(solvers, instances)
	},
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	settings.SetViperDefaults()

	viper.SetConfigName("settings")                 // name of config file (without extension)
	viper.AddConfigPath("$HOME/.config/chronozinc") // add home directory as first search path
	viper.AddConfigPath("/etc/chronozinc")          // adds global machine configuration
	viper.SetEnvPrefix("czn")                       // set environment prefix
	viper.AutomaticEnv()                            // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.test.yaml)")
}
