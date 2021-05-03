// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package config define the internal configuration
// of the DNS server
package config

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		configStr []byte
		format    string
	}

	tests := []struct {
		name       string
		args       args
		wantConfig Configuration
	}{
		{
			name: "Empty configuration (YAML)",
			args: args{
				configStr: []byte{},
				format:    "yaml",
			},
			wantConfig: Configuration{},
		},
		{
			name: "Empty configuration (JSON)",
			args: args{
				configStr: []byte{},
				format:    "json",
			},
			wantConfig: Configuration{},
		},
		{
			name: "Empty configuration (XML)",
			args: args{
				configStr: []byte{},
				format:    "xml",
			},
			wantConfig: Configuration{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotConfig := Parse(tt.args.configStr, tt.args.format); !reflect.DeepEqual(gotConfig, tt.wantConfig) {
				t.Errorf("Parse() = %v, want %v", gotConfig, tt.wantConfig)
			}
		})
	}
}
