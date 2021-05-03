package main

import (
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/lucasdc6/gdns/internal/server"
	"github.com/lucasdc6/gdns/internal/usage"
	"github.com/pborman/getopt/v2"
)

func main() {
	hostFlag := getopt.StringLong("host", 'h', "127.0.0.1", "Define the udp service host")
	tcpHostFlag := getopt.StringLong("tcp-host", 0, "127.0.0.1", "Define the tcp service host")
	portFlag := getopt.IntLong("port", 'p', 3000, "Define the udp service port")
	tcpPortFlag := getopt.IntLong("tcp-port", 0, 3000, "Define the tcp service port")
	fileFlag := getopt.StringLong("file", 'f', "", "Define the path to the configuration file")
	manFlag := getopt.EnumLong("man", 'm', []string{"file-syntax"}, "", "Show usage for the following modules\n- file-syntax")
	modeFlag := getopt.EnumLong("mode", 0, []string{"tcp", "udp", "both"}, "udp", "Run the server in udp, tcp or both")
	verboseLevelFlag := getopt.EnumLong("verbose", 'v', []string{"All", "Info"}, "Info", "Set the verbose mode")
	helpFlag := getopt.BoolLong("help", '?', "Show this help")

	getopt.Parse()
	log.SetLevel(log.DebugLevel)

	// Check args
	if *helpFlag {
		getopt.Usage()
		return
	}

	if *manFlag != "" {
		usage.Usage(*manFlag)
		return
	}

	var wg sync.WaitGroup

	if *modeFlag == "udp" || *modeFlag == "both" {
		serverConfig := server.Server{
			Host:              *hostFlag,
			Port:              *portFlag,
			ConfigurationFile: *fileFlag,
			Mode:              "udp",
			WG:                &wg,
			Verbose:           *verboseLevelFlag,
		}

		serverConfig.WG.Add(1)

		go server.Start(serverConfig)
	}

	if *modeFlag == "tcp" || *modeFlag == "both" {
		serverConfig := server.Server{
			Host:              *tcpHostFlag,
			Port:              *tcpPortFlag,
			ConfigurationFile: *fileFlag,
			Mode:              "tcp",
			WG:                &wg,
			Verbose:           *verboseLevelFlag,
		}
		serverConfig.WG.Add(1)

		go server.Start(serverConfig)
	}

	wg.Wait()
	log.Printf("Shutting down")
}
