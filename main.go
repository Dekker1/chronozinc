package main

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chonozinc [description]",
	Short: "Test and time your minizinc models!",
	Long: `ChonoZinc is a small tool that wraps around mzn-fzn. It main purpose is to
	test MiniZinc models using different solvers. It run your your model on all or
	a sub-set of declared solvers and report on the statistics gathered. The data
	gathered is saved in every stage.`,
	Run: func(cmd *cobra.Command, args []string) {},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
