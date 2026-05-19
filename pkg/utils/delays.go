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

func GetDayTime() string {
	loc := time.FixedZone("EET", 2*60*60)
	t := time.Now().In(loc)
	s1 := fmt.Sprintf("%04d%02d%02d%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	return s1
}

func GetDayTime_() []byte {
	loc := time.FixedZone("EET", 2*60*60)
	t := time.Now().In(loc)
	s1 := fmt.Sprintf("\r\n%02d/%02d %02d:%02d:%02d ", t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	b := []byte(s1)
	return b
}

// \r\nhh:mm:ss_
func GetTime() []byte {
	loc := time.FixedZone("EET", 2*60*60)
	t := time.Now().In(loc)
	s1 := fmt.Sprintf("\r\n%02d:%02d:%02d ", t.Hour(), t.Minute(), t.Second())
	b := []byte(s1)
	return b
}
