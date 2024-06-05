package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "api-gorm-muti-database",
	Short: "api-gorm-muti-database root",
	Long:  "api-gorm-muti-database root.",
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
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "etc/etc.yaml", "config file (default is project root dir etc/etc.yaml")
}

func initConfig() {
	if len(os.Args) <= 1 || os.Args[1] != serverCmd.Name() {
		return
	}

	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		_, _ = fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		cfgFile = viper.ConfigFileUsed()
	} else {
		cobra.CheckErr(err)
	}
}
