package runner

import (
	"log"
	"time"
)

func RunToDuration(dur time.Duration, repeat bool, args []string, msg chan string) {
	now := time.Now()
	at := now.Add(dur)
	log.Printf("Will wait till %s", at.Format(time.ANSIC))

	ticker := time.NewTicker(dur)
	defer ticker.Stop()

	lastExec := now
	for {
		select {
		case cmd := <-msg:
			switch cmd {
			case "quit":
				break
			case "status":
				nextDue := lastExec.Add(dur)
				log.Printf("Execution Count %d; Next execution at %s", buildNo, nextDue.Format(time.ANSIC))
				continue
			}

		case t := <-ticker.C:
			buildNo = buildNo + 1
			log.Printf("Execution %d at : %s", buildNo, t.Format(time.ANSIC))
			lastExec = t
			Run(args)
		}
		if !repeat {
			log.Printf("Repeat not requested. Quitting")
			break
		}
	}

}
