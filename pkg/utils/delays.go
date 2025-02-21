package utils

import "time"

func D_1ms(ms int64) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
func D_100ms(ms int64) {
	time.Sleep(time.Duration(ms) * 100 * time.Millisecond)
}
func D_1s(sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
}
