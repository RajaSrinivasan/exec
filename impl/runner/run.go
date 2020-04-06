package runner

import (
	"log"
	"os"
	"os/exec"
	"path"
	"time"
)

var markerLine = "-----------------------------------------------------------------------------"

func logSetup(nm string, fullcmd string) (logger *log.Logger, lf *os.File) {
	bn := path.Base(nm)
	logfilename := bn + ".log"
	lf, _ = os.Create(logfilename)
	logger = log.New(lf, "", 0)
	logger.Println(markerLine)
	logger.Printf("Command : %s", fullcmd)
	logger.Printf("Started: %s", time.Now().Format(time.ANSIC))
	logger.Println(markerLine)

	return logger, lf
}

func Run(args []string) {

	var startTime = time.Now()

	cmd := exec.Command(args[0], args[1:]...)

	logger, lf := logSetup(args[0], cmd.String())
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
