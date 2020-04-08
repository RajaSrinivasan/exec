package cmd

import (
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"gitlab.com/RajaSrinivasan/exec/impl/runner"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the command",
	Long: `
	Run the command
	`,
	Run:  Run,
	Args: cobra.MinimumNArgs(1),
}

var clock string
var timer string
var repeat bool

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.PersistentFlags().StringVarP(&clock, "clock", "c", "", "execute when the clock matches hh:mm:ss")
	runCmd.PersistentFlags().StringVarP(&timer, "timer", "t", "", "when this timer expires.")
	runCmd.PersistentFlags().BoolVarP(&repeat, "repeat", "p", false, "repeat. repeat the wait after completion")

}

func Run(cmd *cobra.Command, args []string) {
	log.Printf("Invoking runner with args %v\n", args)
	ch := make(chan string)

	if len(clock) > 0 {
		clockspec, err := time.Parse("15:04:05", clock)
		if err != nil {
			log.Printf("%s", err)
			os.Exit(1)
		}
		if len(timer) > 0 {
			log.Printf("Timer spec ignored. Clock overrides the timer")
		}

		if repeat {
			go runner.RunToClock(clockspec, repeat, args, ch)
			for {
				time.Sleep(5 * time.Minute)
				log.Printf("Asking for status updates")
				ch <- "status"
			}
		} else {
			runner.RunToClock(clockspec, repeat, args, ch)
		}
	}

	if len(timer) > 0 {
		timerspec, err := time.ParseDuration(timer)
		if err != nil {
			log.Printf("%s", err)
			os.Exit(1)
		}

		if repeat {
			go runner.RunToDuration(timerspec, repeat, args, ch)
			for {
				log.Printf("Asking for status updates")
				ch <- "status"
				time.Sleep(5 * timerspec)
			}
		} else {
			runner.RunToDuration(timerspec, repeat, args, ch)
		}
	}
	runner.Run(args)
}
