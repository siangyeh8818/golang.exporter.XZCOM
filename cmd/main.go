package main

import (
	"log"
	"os"

	export "github.com/siangyeh8818/golang.exporter.templeate/internal"
)

func main() {
	log.Println("Exporter is start ro running")
	nats_ip_env := os.Getenv("NATS_IP")
	if nats_ip_env == "" {
		nats_ip_env = "nats"
	}
	log.Printf("NATS_IP : %s \n", nats_ip_env)
	nats_port_env := os.Getenv("NATS_PORT")
	if nats_port_env == "" {
		nats_port_env = "8222"
	}
	log.Printf("NATS_PORT : %s \n", nats_port_env)
	export.Run_Exporter_Server()
}
