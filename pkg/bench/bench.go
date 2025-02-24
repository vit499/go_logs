package bench

import (
	"log"
	"time"
)

func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func Duration(msg string, start time.Time) {
	t := time.Since(start)
	// if t > 1*time.Second {
	log.Printf("%v: %v\n", msg, t)
	// }
	// log.Printf("%v: %v\n", msg, time.Since(start))
}
