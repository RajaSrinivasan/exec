package runner

import (
	"log"

	"time"
)

func RunToClock(clockspec time.Time, repeat bool, args []string, msg chan string) {
	now := time.Now()
	nowh, nowmin, nowsec := now.Clock()

	toclock := time.Date(now.Year(), now.Month(), now.Day(), clockspec.Hour(), clockspec.Minute(), clockspec.Second(), 0, now.Location())
	tomorrowclock := toclock.Add(24 * time.Hour)
	spech, specmin, specsec := clockspec.Clock()

	var waittill time.Time

	switch {
	case nowh > spech:
		log.Printf("Hour is already past. Have to be next day %s", tomorrowclock)
		waittill = tomorrowclock
	case nowh < spech:
		log.Printf("Still time today")
		waittill = toclock
	default:
		switch {
		case nowmin > specmin:
			log.Printf("Minute has already passed . %s", tomorrowclock)
			waittill = tomorrowclock
		case nowmin < specmin:
			log.Printf("Still time for the minute")
			waittill = toclock
		default:
			switch {
			case nowsec > specsec:
				log.Printf("Second already passed. %s", tomorrowclock)
				waittill = tomorrowclock
			case nowsec < specsec:
				log.Printf("Still seconds left")
				waittill = toclock
			default:
				log.Printf("Exact time. Start right now")
				waittill = toclock
			}
		}
	}

	initticker := time.NewTicker(1 * time.Minute)
	defer initticker.Stop()
	log.Printf("Waiting till %s", waittill.Format(time.ANSIC))
	var lastExec time.Time

	select {
	case cmd := <-msg:
		switch cmd {
		case "quit":
			return
		case "status":
			log.Printf("No Executions yet. waiting till %s", waittill.Format(time.ANSIC))
		}
	case tfirst := <-initticker.C:
		if tfirst.After(waittill) {
			buildNo = buildNo + 1
			log.Printf("Execution %d at : %s", buildNo, tfirst.Format(time.ANSIC))
			lastExec = time.Now()
			Run(args)
			initticker.Stop()
			break
		} else {
			log.Printf("Not time yet. %s Going back to sleep", tfirst.Format(time.ANSIC))
		}
	}
	if !repeat {
		log.Printf("Repeats not requested. Quitting")
		return
	}

	ticker := time.NewTicker(24 * time.Hour)
	for {
		select {
		case cmd := <-msg:
			switch cmd {
			case "quit":
				break
			case "status":
				nextDue := lastExec.Add(24 * time.Hour)
				log.Printf("Execution Count %d; Next execution at %s", buildNo, nextDue.Format(time.ANSIC))
				continue
			}

		case t := <-ticker.C:
			buildNo = buildNo + 1
			log.Printf("Execution %d at : %s", buildNo, t.Format(time.ANSIC))
			lastExec = t
			Run(args)
		}
	}

}
