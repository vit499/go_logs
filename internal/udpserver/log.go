package udpserver

import (
	"go_logs/pkg/utils"
	"net"
)

func (u *UdpServer) save(s string) {
	// u.logger.Info().Msgf(s)
	// u.mfile.Write()
}

func (u *UdpServer) savebuf(buf []byte, addr *net.UDPAddr) int {
	// u.logger.Info().Msgf(s)
	from_pc := false

	if len(buf) < 2 { // 's' - ping from nv
		u.addr_nv = addr
		return (0)
	}
	b_osdp_log_on := []byte("osdp_log_on")
	b_cmd_for_ulog := []byte("Cmd_for_ulog")
	b_pcping := []byte("pcping")

	// u.logger.Info().Msgf(" rec: %s", string(buf))

	if utils.StrNCmp(buf, b_pcping) == 0 { // ping from pc
		from_pc = true
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
	}

	if from_pc == true {
		u.addr_pc = addr
		u.pc_en = true
		u.cnt_pc = 0
	} else {
		u.addr_nv = addr
	}
	u.mfile.Write(buf)
	if from_pc {
		return (1)
	}
	return (0)
}
