// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	type args struct {
		config *Config
		opts   []ClientOption
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "empty",
			wantErr: true,
		},
		{
			name: "with token",
			args: args{
				config: &Config{
					Server: "http://bar",
				},
			},
			wantErr: true,
		},
		{
			name: "with server",
			args: args{
				config: &Config{
					AuthToken: "foo",
				},
			},
			wantErr: true,
		},
		{
			name: "valid",
			args: args{
				config: &Config{
					AuthToken: "foo",
					Server:    "http://bar",
				},
			},
		},
		{
			name: "valid",
			args: args{
				config: &Config{
					AuthToken: "foo",
					Server:    "http://bar",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewClient(tt.args.config, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
