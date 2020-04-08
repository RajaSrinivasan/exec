package runner

import (
	"os/exec"
	"time"
)

var buildNo int

func Run(args []string) {

	var startTime = time.Now()

	cmd := exec.Command(args[0], args[1:]...)

	logger, lf := Create(args[0], cmd.String(), buildNo)
	defer lf.Close()

	var finalstat error
	out, finalstat := cmd.CombinedOutput()

	logger.Println(string(out))
	logger.Println(markerLine)
	var exectime = time.Now().Sub(startTime)
	if finalstat != nil {
		logger.Printf("Final Status: %s", finalstat)
	}

	logger.Printf("Terminated %s Duration %s", time.Now().Format(time.ANSIC), exectime.String())
	logger.Println(markerLine)

}
