package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GetClientId() string {
	rand.Seed(time.Now().UnixNano())
	r1 := rand.Intn(65000)
	r2 := rand.Intn(65000)
	//fmt.Printf("\r\n r2=%X%X", r1, r2)
	s := fmt.Sprintf("%s%X%X", "22223333", r1, r2)
	return s
}
func GetTopicSub(login, unit string) string {
	s := fmt.Sprintf("%s/%s/devsend/#", login, unit)
	//m.logger.Info().Msg("sub topic: %s ", s)
	return s
}
