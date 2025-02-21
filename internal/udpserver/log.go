package udpserver

func (u *UdpServer) save(s string) {
	u.logger.Info().Msgf(s)
}
