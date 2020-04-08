package runner

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

var markerLine = "-----------------------------------------------------------------------------"

func logFileName(name string, buildno int) string {
	nm := path.Join(LogsDir, path.Base(name))
	var lfn string
	if buildno > 0 {
		lfn = fmt.Sprintf("%s_%d.log", nm, buildno)
	} else {
		lfn = nm + ".log"
	}
	return lfn
}

func Create(nm string, fullcmd string, buildno int) (logger *log.Logger, lf *os.File) {
	logfilename := logFileName(nm, buildno)
	lf, err := os.Create(logfilename)
	if err != nil {
		log.Fatal(err)
	}
	logger = log.New(lf, "", 0)
	logger.Println(markerLine)
	logger.Printf("Command : %s", fullcmd)
	logger.Printf("Started: %s", time.Now().Format(time.ANSIC))
	logger.Println(markerLine)
	return logger, lf
}
