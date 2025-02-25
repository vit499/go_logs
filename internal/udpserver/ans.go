package udpserver

import (
	"context"
	"log"
	"time"
)

func (u *UdpServer) ans(ctx context.Context) {
	ticker := time.NewTicker(time.Duration(10) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("ans close")
			return
		case <-ticker.C:
			u.mux.Lock()
			u.cnt_ans++
			if u.cnt_ans > 3 {
				// u.logger.Info().Msgf("ans timer")
				u.cmd = "test 00"
			}
			u.mux.Unlock()
		}
	}
}
