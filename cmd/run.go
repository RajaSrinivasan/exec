package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/RajaSrinivasan/exec/impl/runner"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the command",
	Long: `
	Run the command
	`,
	Run: Run,
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func Run(cmd *cobra.Command, args []string) {
	fmt.Printf("Invoking runner with args %v", args)
	runner.Run(args)
}
