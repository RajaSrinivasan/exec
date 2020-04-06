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

var clock string
var timer string
var repeat bool

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.PersistentFlags().StringVarP(&clock, "clock", "c", "", "execute when the clock matches [[hh:mm:]]ss")
	runCmd.PersistentFlags().StringVarP(&timer, "timer", "t", "", "when this timer expires [[hh:]mm:]ss")
	runCmd.PersistentFlags().BoolVarP(&repeat, "repeat", "p", false, "repeat. repeat the wait after completion")

}

func Run(cmd *cobra.Command, args []string) {
	fmt.Printf("Invoking runner with args %v\n", args)
	runner.Run(args)
}
