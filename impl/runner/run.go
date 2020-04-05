package runner

import (
	"context"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"
)

var markerLine = "-----------------------------------------------------------------------------"

func logSetup(nm string) (logger *log.Logger, lf *os.File) {
	bn := path.Base(nm)
	logfilename := bn + ".log"
	lf, _ = os.Create(logfilename)
	logger = log.New(lf, "", 0)
	logger.Println(markerLine)
	logger.Printf("Script : %s", nm)
	logger.Printf("Started: %s", time.Now().Format(time.ANSIC))
	logger.Println(markerLine)

	return logger, lf
}

func logResults(logger *log.Logger, res string) {
	logger.Println(res)
}

func Run(args []string) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	logger, lf := logSetup(args[0])
	defer lf.Close()

	var startTime = time.Now()

	cmd := exec.CommandContext(ctx, args[0], strings.Join(args[1:], " "))

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	io.Copy(lf, stdout)
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	logger.Println(markerLine)
	var exectime = time.Now().Sub(startTime)
	logger.Printf("Terminated %s Duration %s", time.Now().Format(time.ANSIC), exectime.String())
	logger.Println(markerLine)

	//log.Printf("%s", string(out))
}
