package udpserver

import "net"

func (u *UdpServer) save(s string) {
	// u.logger.Info().Msgf(s)
	// u.mfile.Write()
}

func (u *UdpServer) savebuf(buf []byte, addr *net.UDPAddr) {
	// u.logger.Info().Msgf(s)
	src_pc := false

	if len(buf) < 2 {
		u.addr_nv = addr
		return
	}
	s2 := "osdp_log_on"
	s3 := "Cmd_for_ulo" //  "Cmd_for_ulog " + cmd + "\r\n";
	s1 := ""
	if len(buf) > 10 {
		s1 = string(buf[:11])
	}
	// u.logger.Info().Msgf(" s1=%s ", s1)
	if s1 == s2 {
		u.mux.Lock()
		u.cmd = ""
		u.cnt_ans = 0
		u.mux.Unlock()
		u.logger.Info().Msgf(" osdp_log_on ")
	}
	if s1 == s3 {
		src_pc = true
		len := len(buf) - 2
		str_cmd := string(buf[13:len])
		u.mux.Lock()
		u.cmd = str_cmd
		u.cnt_ans = 0
		u.mux.Unlock()
		u.logger.Info().Msgf("Cmd_for_ulog %s", str_cmd)
	}

	if src_pc == true {
		u.addr_pc = addr
	} else {
		u.addr_nv = addr
	}
	u.mfile.Write(buf)
}
