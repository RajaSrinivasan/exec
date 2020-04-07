package runner

import (
	"log"
	"time"
)

func RunToDuration(dur time.Duration, repeat bool, args []string) {
	now := time.Now()
	at := now.Add(dur)
	log.Printf("Will wait till %s", at.Format(time.ANSIC))
}
