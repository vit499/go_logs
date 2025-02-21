package config

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"sync"
)

type Config struct {
	UdpHost string
	// HttpPort string
	UdpPort int
}

var (
	config Config
	once   sync.Once
)

func Get() *Config {

	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Printf("Error loading .env file, default config")
		}

		config.UdpHost = os.Getenv("UDP_HOST")

		var s string
		s = os.Getenv("UDP_PORT")
		config.UdpPort, err = strconv.Atoi(s)
		if err != nil {
			config.UdpPort = 8011
		}

		b, err := json.MarshalIndent(config, "", "")
		if err != nil {
			log.Printf("json cfg err: %s", err.Error())
		}
		log.Printf("cfg: %s", string(b))
	})
	return &config
}
