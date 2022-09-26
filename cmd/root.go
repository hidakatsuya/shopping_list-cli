package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "shopping_list-cli",
	Short: "A command line interface for Shopping List",
	Long:  "A command line interface for Shopping List.",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initEnv)
}

func initEnv() {
	viper.SetEnvPrefix("shopping_list")
	viper.BindEnv("url")
	viper.BindEnv("api_key")
}
