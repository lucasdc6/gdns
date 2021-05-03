package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/lucasdc6/gdns/pkg/config"
	"github.com/pborman/getopt/v2"
)

func main() {
	fileFlag := getopt.StringLong("file", 'f', "", "Define the path to the configuration file")
	getopt.Parse()

	log.Printf("Check config")
	config.Load(*fileFlag)
}
