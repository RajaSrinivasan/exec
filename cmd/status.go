package cmd

import (
	"github.com/spf13/cobra"
	//"gitlab.com/RajaSrinivasan/exec/impl/version"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Report the current status of all the jobs",
	Long: `
	Report the current status of all the jobs
	`,
	Run: Status,
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

func Status(cmd *cobra.Command, args []string) {
	//version.Report()
}
