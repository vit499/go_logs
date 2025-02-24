package udpserver

func (u *UdpServer) save(s string) {
	// u.logger.Info().Msgf(s)
	// u.mfile.Write()
}

func (u *UdpServer) savebuf(buf []byte) {
	// u.logger.Info().Msgf(s)
	u.mfile.Write(buf)
}
