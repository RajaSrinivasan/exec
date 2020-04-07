package runner

import (
	"log"
	"testing"
	"time"
)

func TestRunToDuration(t *testing.T) {

	args := []string{"ls"}
	dur, _ := time.ParseDuration("1m")
	comm := make(chan string)
	go RunToDuration(dur, true, args, comm)
	log.Printf("Sleeping for 5 minutes")
	time.Sleep(5 * dur)
	log.Printf("Woke up. Sending request to stop")
	comm <- "stop"
}
