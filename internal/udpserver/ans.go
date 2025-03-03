package udpserver

import (
	"context"
	"log"
	"time"
)

func (u *UdpServer) ans(ctx context.Context) {
	ticker1 := time.NewTicker(time.Duration(10) * time.Second)
	defer ticker1.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("ans close")
			return
		case <-ticker1.C:
			u.mux.Lock()
			u.cnt_ans++
			if u.cnt_ans > 3 {
				// u.logger.Info().Msgf("ans timer")
				u.cmd = "test 00"
			}
			u.mux.Unlock()
			u.cnt_pc++
			if u.cnt_pc > 15 {
				u.cnt_pc = 0
				u.pc_en = false
			}
		}
	}
}
