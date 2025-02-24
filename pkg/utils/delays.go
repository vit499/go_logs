package utils

import (
	"fmt"
	"time"
)

func D_1ms(ms int64) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
func D_100ms(ms int64) {
	time.Sleep(time.Duration(ms) * 100 * time.Millisecond)
}
func D_1s(sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
}

func GetTime() string {
	// dv_time=20140101120029
	t := time.Now()
	s1 := fmt.Sprintf("%04d%02d%02d%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	return s1
}
