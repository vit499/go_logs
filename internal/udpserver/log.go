package udpserver

import (
	"go_logs/pkg/utils"
	"net"
)

func (u *UdpServer) save(s string) {
	// u.logger.Info().Msgf(s)
	// u.mfile.Write()
}

func (u *UdpServer) savebuf(buf []byte, addr *net.UDPAddr) ([]byte, int) {
	// u.logger.Info().Msgf(s)
	from_pc := false
	pc_ping := false
	dst := make([]byte, len(buf))

	copy(dst, buf)
	if len(buf) < 2 { // 's' - ping from nv
		u.addr_nv = addr
		return dst, 2 // чтобы не отправлять в компьютер ping от nv
	}
	b_osdp_log_on := []byte("osdp_log_on")
	b_cmd_for_ulog := []byte("Cmd_for_ulog")
	b_pcping := []byte("pcping")

	// u.logger.Info().Msgf(" rec: %s", string(buf))

	if utils.StrNCmp(buf, b_pcping) == 0 { // ping from pc
		from_pc = true
		pc_ping = true
		// u.logger.Info().Msgf(" pcping ")
	} else if utils.StrNCmp(buf, b_osdp_log_on) == 0 { // это начало лога, приходит от nv
		u.mux.Lock()
		u.cmd = ""
		u.cnt_ans = 0
		u.mux.Unlock()
		u.logger.Info().Msgf(" osdp_log_on ")
	} else if utils.StrNCmp(buf, b_cmd_for_ulog) == 0 { // команда от pc
		from_pc = true
		len := len(buf) - 2
		str_cmd := string(buf[13:len])
		u.mux.Lock()
		u.cmd = str_cmd
		u.cnt_ans = 0
		u.mux.Unlock()
		u.logger.Info().Msgf("Cmd_for_ulog %s", str_cmd)
	} else if utils.StrNCmpTime(buf) == 0 { // нужно изменить время
		buf_time := utils.GetTime() //   \r\nhh:mm:ss
		len_time := len(buf_time)
		for i := 0; i < len_time; i++ {
			dst[i] = buf_time[i]
		}
	} else if utils.StrNCmpOsdp(buf) == 0 { // нужно добавить время
		buf_time := utils.GetTime() //   \r\nhh:mm:ss
		len_time := len(buf_time)
		for i := 0; i < len_time; i++ {
			dst[i] = buf_time[i]
		}
		for i := 0; i < len(buf); i++ {
			dst[i+len_time] = buf[i]
		}
	} else if utils.StrNCmpStart(buf) == 0 { // нужно добавить дату
		buf_time := utils.GetDayTime_() //   \r\nhh:mm:ss
		len_time := len(buf_time)
		for i := 0; i < len_time; i++ {
			dst[i] = buf_time[i]
		}
		for i := 0; i < len(buf); i++ {
			dst[i+len_time] = buf[i]
		}
	}

	if from_pc == true {
		u.addr_pc = addr
		u.pc_en = true
		u.cnt_pc = 0
	} else {
		u.addr_nv = addr
	}
	if !pc_ping {
		u.mfile.Write(dst)
	}
	if from_pc {
		return dst, 1
	}
	return dst, 0
}
