package runner

import (
	"testing"
	"time"
)

func TestRunToClock(t *testing.T) {

	args := []string{"ls"}
	clockspec, _ := time.Parse("15:04:05", "01:00:20")
	RunToClock(clockspec, false, args)
	clockspec2, _ := time.Parse("15:04:05", "01:30:20")
	RunToClock(clockspec2, false, args)
	clockspec3, _ := time.Parse("15:04:05", "01:15:00")
	RunToClock(clockspec3, false, args)

}
