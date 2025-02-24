package udpserver

import (
	"context"
	"fmt"
	"go_logs/internal/file"
	"go_logs/pkg/config"
	"go_logs/pkg/logger"
	"log"
	"net"
)

type UdpServer struct {
	host   string
	port   int
	logger *logger.Logger
	mfile  *file.Mfile
}

func New(ctx context.Context, logger *logger.Logger, mfile *file.Mfile) {
	cfg := config.Get()
	host := cfg.UdpHost
	port := cfg.UdpPort

	u := &UdpServer{
		host:   host,
		port:   port,
		logger: logger,
		mfile:  mfile,
	}

	go u.udp_start(ctx)
	go u.client(ctx)
}

func (u *UdpServer) udp_server() {
	u.logger.Info().Msgf("udp server starting %s:%d", u.host, u.port)
	listener, _ := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP(u.host), Port: u.port}) // открываем слушающий UDP-сокет
	for {
		u.handleClient(listener) // обрабатываем запрос клиента
	}
}

func (u *UdpServer) handleClient(conn *net.UDPConn) {
	buf := make([]byte, 128) // буфер для чтения клиентских данных

	len, addr, err := conn.ReadFromUDP(buf) // читаем из сокета
	if err != nil {
		fmt.Println(err)
		return
	}

	if len > 10000 {
		u.logger.Info().Msgf("rec %s", addr) // чтобы не ругался на addr
	}
	// log.Printf("r:%d addr=%s, r:%s ", len, addr, string(buf[:len]))
	// u.logger.Info().Msgf("rec %s mes:%s", addr, string(buf[:len]))
	// conn.WriteToUDP(append([]byte("Hello, you said: "), buf[:readLen]...), addr) // пишем в сокет
	// s := string(buf[:len])
	u.savebuf(buf[:len])
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
