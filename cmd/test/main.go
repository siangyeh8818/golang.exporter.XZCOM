package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"

	server "github.com/siangyeh8818/golang.exporter.XZCOM/internal/server"
	selenium "github.com/siangyeh8818/golang.exporter.XZCOM/internal/selenium"
)

func main() {
	go func{
		selenium.RunSelium()
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
