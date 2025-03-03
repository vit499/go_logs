package udpserver

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"
)

func (u *UdpServer) sendClient() {
	hostName := "localhost"
	portNum := "8011"

	service := hostName + ":" + portNum

	RemoteAddr, err := net.ResolveUDPAddr("udp", service)

	//LocalAddr := nil
	// see https://golang.org/pkg/net/#DialUDP

	conn, err := net.DialUDP("udp", nil, RemoteAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// write a message to server
	s := fmt.Sprintf("\r\n xxxxxxxxxxx aaaaaaaaaaaaa ccccccccccc 0000000000000 xxxxxxxxxxxxxxx bbbbbbbbbbbbbbb xxxxxxxxxx %d", u.cnt)
	u.cnt++
	message := []byte(s)

	_, err = conn.Write(message)

	if err != nil {
		log.Println(err)
	}

}

func (u *UdpServer) client(ctx context.Context) {
	ticker := time.NewTicker(time.Duration(1) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("client close")
			return
		case <-ticker.C:
			u.sendClient()
		}
	}
}
