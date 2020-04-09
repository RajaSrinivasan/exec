package cmd

import (
	"fmt"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"gitlab.com/RajaSrinivasan/exec/impl/runner"
)

var cfgFile string
var verbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute specified command",
	Long: `
	Execute specified command according to the requested schedule
	`,
	Version: "v0.1.0",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli/cli.yaml)")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "be verbose")

}

func initConfig() {

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		home = "./"
	}

	logsDir := path.Join(home, "logs")
	viper.SetDefault("LogsDir", logsDir)

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		defconfigpath := path.Join(home, ".exec")
		viper.AddConfigPath(defconfigpath)
		viper.AddConfigPath("./etc")
		viper.SetConfigName("exec")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	runner.LogsDir = viper.GetString("LogsDir")
}
