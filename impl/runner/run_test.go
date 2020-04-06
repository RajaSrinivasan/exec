package runner

import (
	"testing"
)

func TestRun(t *testing.T) {
	args := []string{"go", "build", "../../..."}
	//args := []string{"../../build.sh"}
	Run(args)
}
