package udpserver

import (
	"context"
	"log"
	"net"
	"time"
)

func (u *UdpServer) send() {
	hostName := "localhost"
	portNum := "8011"

	service := hostName + ":" + portNum

	RemoteAddr, err := net.ResolveUDPAddr("udp", service)

	//LocalAddr := nil
	// see https://golang.org/pkg/net/#DialUDP

	conn, err := net.DialUDP("udp", nil, RemoteAddr)

	// note : you can use net.ResolveUDPAddr for LocalAddr as well
	//        for this tutorial simplicity sake, we will just use nil

	if err != nil {
		log.Fatal(err)
	}

	// log.Printf("Established connection to %s \n", service)
	// log.Printf("Remote UDP address : %s \n", conn.RemoteAddr().String())
	// log.Printf("Local UDP client address : %s \n", conn.LocalAddr().String())

	defer conn.Close()

	// write a message to server
	message := []byte("Hello UDP server!")

	_, err = conn.Write(message)

	if err != nil {
		log.Println(err)
	}

	// // receive message from server
	// buffer := make([]byte, 1024)
	// n, addr, err := conn.ReadFromUDP(buffer)

	// fmt.Println("UDP Server : ", addr)
	// fmt.Println("Received from UDP server : ", string(buffer[:n]))

}

func (u *UdpServer) client(ctx context.Context) {
	ticker := time.NewTicker(time.Duration(2) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("client close")
			return
		case <-ticker.C:
			u.send()
		}
	}
}
