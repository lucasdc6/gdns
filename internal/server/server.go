// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package server define the DNS server
package server

import (
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/lucasdc6/gdns/pkg/config"
	"github.com/lucasdc6/gdns/pkg/errors"
	"github.com/lucasdc6/gdns/pkg/parser"
)

// Server - Configuration for the DNS server
type Server struct {
	Host              string
	Port              int
	Mode              string
	WG                *sync.WaitGroup
	ConfigurationFile string
	Configuration     config.Configuration
	Verbose           string
}

func startUDPServer(server Server) {
	addr := net.UDPAddr{
		IP:   net.ParseIP(server.Host),
		Port: server.Port,
	}

	ser, err := net.ListenUDP("udp", &addr)

	if err != nil {
		log.Fatalf("Error starting the server: %v\n", err)
		os.Exit(errors.StartingServer)
	}

	log.Printf("UDP Server started at %s:%d\n", server.Host, server.Port)
	listenUDPPackages(ser)
}

func sendUDP(dstIP string, dstPort int, data []byte) []byte {
	p := make([]byte, 512)
	conn, err := net.ListenPacket("udp", ":0")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	dst, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", dstIP, dstPort))
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.WriteTo(data, dst)
	if err != nil {
		log.Fatal(err)
	}

	num, _, err := conn.ReadFrom(p)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Readed %d, with data %v", num, p)
	return p
}

func listenUDPPackages(conn *net.UDPConn) {

	p := make([]byte, 512)

	for {
		_, remoteaddr, err := conn.ReadFromUDP(p)

		if err != nil {
			log.Fatalf("Error retriving UDP package: %v", err)
			os.Exit(errors.RetrivingUDPPackage)
		}
		log.Printf("UDP Query recived from %v", remoteaddr)
		log.Printf("Data:\n%s\n", hex.Dump(p))

		parser.ParseDNSQuery(p)
		if true {
			log.Printf("Send query to authoritative server")
			res := sendUDP("8.8.8.8", 53, p)
			log.Printf("Data:\n%s\n", hex.Dump(res))
			parser.ParseDNSQuery(res)
		}
	}
}

func starTCPServer(server Server) {
	addr := net.TCPAddr{
		IP:   net.ParseIP(server.Host),
		Port: server.Port,
	}

	ser, err := net.ListenTCP("tcp", &addr)

	if err != nil {
		log.Fatalf("Error starting the server: %v\n", err)
		os.Exit(errors.StartingServer)
	}

	log.Printf("TCP Server started at %s:%d\n", server.Host, server.Port)
	listenTCPData(ser)
}

func listenTCPData(server *net.TCPListener) {
	p := make([]byte, 512)

	for {
		conn, err := server.Accept()
		remoteaddr := conn.RemoteAddr()

		if err != nil {
			log.Fatalf("Error when try to establish connection: %v", err)
			os.Exit(errors.EstablishingTCPConn)
		}
		log.Printf("TCP Query recived from %v", remoteaddr)

		_, err = conn.Read(p)

		if err != nil {
			log.Fatalf("Error retriving data: %v", err)
			os.Exit(errors.RetrivingTCPData)
		}

		log.Printf("Data: %+v\n", p)

		parser.ParseDNSQuery(p[2:])
	}
}

// Start - Start the DNS server
func Start(server Server) {
	defer server.WG.Done()

	server.Configuration = config.Load(server.ConfigurationFile)
	if server.Verbose == "All" {
		log.Printf("Started in verbose mode")
	}

	if server.Mode == "udp" {
		startUDPServer(server)

		return
	}

	starTCPServer(server)
}
