package cmd

import (
	"github.com/spf13/cobra"
	//"gitlab.com/RajaSrinivasan/exec/impl/version"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the execution of a command",
	Long: `
	Stop the execution of a command. Argument is the id of the command
	`,
	Run:  Stop,
	Args: cobra.MinimumNArgs(1),
}

func init() {
	rootCmd.AddCommand(stopCmd)
}

func Stop(cmd *cobra.Command, args []string) {

}
