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

	//msg := make(chan string)
	/*if repeat {
		go func() {
			for {
				msg <- "status"
				time.Sleep(30 * time.Second)
			}
		}()
	}*/

	lastExec := now
	numexecs := 0
	for {
		select {
		case cmd := <-msg:
			//fmt.Println("Command ", cmd)
			switch cmd {
			case "quit":
				break
			case "status":
				nextDue := lastExec.Add(dur)
				log.Printf("Execution Count %d Next execution at %s", numexecs, nextDue.Format(time.ANSIC))
				continue
			}

		case t := <-ticker.C:
			log.Println("Execution at : ", t)
			lastExec = t
			Run(args)
			numexecs = numexecs + 1
		}
		if !repeat {
			log.Printf("Repeat not requested. Quitting")
			break
		}
	}

}
