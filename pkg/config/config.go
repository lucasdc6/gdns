// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package config define the internal configuration
// of the DNS server
package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/goccy/go-yaml"
	"github.com/lucasdc6/gdns/pkg/errors"
	"github.com/lucasdc6/gdns/pkg/types"
)

// Global - Define the struct of the global section
type Global struct {
}

// Zone - Define the struct of the Zones in the configuration
type Zone struct {
	Name    string   `yaml:"name" json:"name"`
	Records []Record `yaml:"records" json:"records"`
}

// Host - Define the struct of the Hosts in the configuration
type Record struct {
	Name  string      `yaml:"name" json:"name"`
	Type  types.QType `yaml:"type" json:"type"`
	Value string      `yaml:"value" json:"value"`
	TTL   int         `yaml:"ttl" json:"ttl"`
}

// Configuration - Define the general struct of the configuration file
type Configuration struct {
	Zones []Zone `yaml:"zones" json:"zones"`
}

func ReadConfigFile(path string) []byte {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		log.Printf("Error opening configuration file: %s", err)
		os.Exit(255)
	}
	return file
}

// Parse - Generate the internal configuration
func Parse(configStr []byte, format string) (config Configuration) {
	log.Tracef("Parsing %s file", format)

	if format == "" {
		log.Printf("Format not found, default to yaml")

		format = ".yaml"
	}

	switch format {
	case ".yaml", ".yml":
		err := yaml.Unmarshal(configStr, &config)

		if err != nil {
			log.Fatalf("Error reading yaml configuration: %v\n", err)
			os.Exit(errors.ReadingYAMLConfiguration)
		}
	case ".json":
		err := json.Unmarshal(configStr, &config)

		if err != nil {
			log.Fatalf("Error reading json configuration: %v\n", err)
			os.Exit(errors.ReadingJSONConfiguration)
		}
	}
	log.Printf("Loaded config: %+v", config)

	return config
}

func Load(path string) (config Configuration) {
	if path == "" {
		return Configuration{}
	}

	log.Infof("Server configuration file '%s'", path)

	file := ReadConfigFile(path)
	ext := filepath.Ext(path)

	log.Debugf("Server configuration format '%s'", ext)

	return Parse(file, ext)
}
