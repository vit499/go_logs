package udpserver

import (
	"context"
	"fmt"
	"go_logs/internal/file"
	"go_logs/pkg/config"
	"go_logs/pkg/logger"
	"log"
	"net"
	"sync"
)

type UdpServer struct {
	host    string
	port    int
	logger  *logger.Logger
	mfile   *file.Mfile
	cnt     int
	cmd     string
	mux     sync.Mutex
	cnt_ans int
	cnt_pc  int
	addr_nv *net.UDPAddr
	addr_pc *net.UDPAddr
	pc_en   bool
}

func New(ctx context.Context, logger *logger.Logger, mfile *file.Mfile) {
	cfg := config.Get()
	host := cfg.UdpHost
	port := cfg.UdpPort

	u := &UdpServer{
		host:    host,
		port:    port,
		logger:  logger,
		mfile:   mfile,
		cnt:     100,
		cmd:     "",
		cnt_ans: 0,
		cnt_pc:  0,
		pc_en:   false,
	}

	go u.udp_start(ctx)
	// go u.client(ctx)
	go u.ans(ctx)
}

// func (u *UdpServer) udp_server() {
// 	u.logger.Info().Msgf("udp server starting %s:%d", u.host, u.port)
// 	listener, _ := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP(u.host), Port: u.port}) // открываем слушающий UDP-сокет
// 	for {
// 		u.handleClient(listener) // обрабатываем запрос клиента
// 	}
// }

func (u *UdpServer) handleClient(conn *net.UDPConn) {
	buf := make([]byte, 128) // буфер для чтения клиентских данных

	len, addr, err := conn.ReadFromUDP(buf) // читаем из сокета
	if err != nil {
		fmt.Println(err)
		return
	}

	from_pc := u.savebuf(buf[:len], addr)

	if u.cmd != "" {

		bufcmd := []byte(u.cmd)
		// _, err := conn.WriteToUDP(bufcmd, u.addr_nv) // пишем в сокет
		// if err != nil {
		// 	u.logger.Info().Msgf("err send ans, err=%v ", err)
		// }
		u.send(conn, bufcmd, u.addr_nv)
		u.logger.Info().Msgf("send cmd to addr=%s, cmd:%s ", addr, string(bufcmd))
		u.cmd = ""
		u.cnt_ans = 0
	}
	if from_pc == 0 && u.pc_en { // получено от nv, отправляется в pc
		u.send(conn, buf[:len], u.addr_pc)
	}
}

func (u *UdpServer) udp_start(ctx context.Context) {
	service := fmt.Sprintf("%s:%d", u.host, u.port)

	udpAddr, err := net.ResolveUDPAddr("udp4", service)

	if err != nil {
		log.Fatal(err)
	}

	// setup listener for incoming UDP connection
	ln, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		log.Fatal(err)
	}

	u.logger.Info().Msgf("udp server start %s:%d", u.host, u.port)

	defer ln.Close()

	for {
		select {
		case <-ctx.Done():
			// log.Printf("go return")
			u.logger.Info().Msgf("udp server stop ")
			return
		default:
			u.handleClient(ln)
		}
		// wait for UDP client to connect
	}
}

func (u *UdpServer) send(conn *net.UDPConn, buf []byte, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP(buf, addr) // пишем в сокет
	if err != nil {
		u.logger.Info().Msgf("err send ans, err=%v ", err)
	}
}
