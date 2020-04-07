package runner

import (
	"log"

	"time"
)

func RunToClock(clockspec time.Time, repeat bool, args []string) {
	now := time.Now()
	nowh, nowmin, nowsec := now.Clock()

	toclock := time.Date(now.Year(), now.Month(), now.Day(), clockspec.Hour(), clockspec.Minute(), clockspec.Second(), 0, now.Location())
	tomorrowclock := toclock.Add(24 * time.Hour)
	spech, specmin, specsec := clockspec.Clock()

	switch {
	case nowh > spech:
		log.Printf("Hour is already past. Have to be next day %s", tomorrowclock)
	case nowh < spech:
		log.Printf("Still time today")
	default:
		switch {
		case nowmin > specmin:
			log.Printf("Minute has already passed . %s", tomorrowclock)
		case nowmin < specmin:
			log.Printf("Still time for the minute")
		default:
			switch {
			case nowsec > specsec:
				log.Printf("Second already passed. %s", tomorrowclock)
			case nowsec < specsec:
				log.Printf("Still seconds left")
			default:
				log.Printf("Exact time. Start right now")
			}
		}
	}
}
