package main

import (
	"net"
	"time"

	"github.com/patrickmn/go-cache"

	crawler "github.com/siangyeh8818/golang.exporter.XZCOM/internal/crawler"
	server "github.com/siangyeh8818/golang.exporter.XZCOM/internal/server"
)

func main() {

	c := cache.New(5*time.Minute, 10*time.Minute)

	c.Set("foo", "bar", cache.DefaultExpiration)

	go func() {
		crawler.RunSelium()
	}()

	server.Run_Exporter_Server()
}

func PickUnusedPort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	port := l.Addr().(*net.TCPAddr).Port
	if err := l.Close(); err != nil {
		return 0, err
	}
	return port, nil
}
