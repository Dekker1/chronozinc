package main

import (
	"fmt"
	"os"

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
	Run: func(cmd *cobra.Command, args []string) {},
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName("config")                   // name of config file (without extension)
	viper.AddConfigPath("$HOME/.config/chronozinc") // add home directory as first search path
	viper.AddConfigPath(".")                        // add current directory as an alternative
	viper.SetEnvPrefix("czn")                       // set environment prefix
	viper.AutomaticEnv()                            // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("No config file found; using ENV and defaults")
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
