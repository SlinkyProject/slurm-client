// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

func TestClientOptions_ApplyOptions(t *testing.T) {
	type fields struct {
		EnableFor       []object.Object
		DisableFor      []object.Object
		CacheSyncPeriod time.Duration
	}
	type args struct {
		opts []ClientOption
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ClientOptions
	}{
		{
			name:   "No options",
			fields: fields{},
			args:   args{},
			want:   &ClientOptions{},
		},
		{
			name:   "From options",
			fields: fields{},
			args: args{
				opts: []ClientOption{
					&ClientOptions{
						EnableFor: []object.Object{
							&types.V0043Node{},
						},
						DisableFor: []object.Object{
							&types.V0042Node{},
						},
						CacheSyncPeriod: 2 * time.Second,
					},
				},
			},
			want: &ClientOptions{
				EnableFor: []object.Object{
					&types.V0043Node{},
				},
				DisableFor: []object.Object{
					&types.V0042Node{},
				},
				CacheSyncPeriod: 2 * time.Second,
			},
		},
		{
			name: "Overwrite existing options",
			fields: fields{
				EnableFor: []object.Object{
					&types.V0043Node{},
				},
				DisableFor: []object.Object{
					&types.V0042Node{},
				},
				CacheSyncPeriod: 2 * time.Second,
			},
			args: args{
				opts: []ClientOption{
					&ClientOptions{
						EnableFor: []object.Object{
							&types.V0043Node{},
						},
						DisableFor: []object.Object{
							&types.V0042Node{},
						},
						CacheSyncPeriod: 2 * time.Second},
				},
			},
			want: &ClientOptions{
				EnableFor: []object.Object{
					&types.V0043Node{},
					&types.V0043Node{},
				},
				DisableFor: []object.Object{
					&types.V0042Node{},
					&types.V0042Node{},
				},
				CacheSyncPeriod: 2 * time.Second,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &ClientOptions{
				EnableFor:       tt.fields.EnableFor,
				DisableFor:      tt.fields.DisableFor,
				CacheSyncPeriod: tt.fields.CacheSyncPeriod,
			}
			got := o.ApplyOptions(tt.args.opts)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestCreateOptions_ApplyOptions(t *testing.T) {
	type fields struct {
		Allocation bool
	}
	type args struct {
		opts []CreateOption
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateOptions
	}{
		{
			name:   "No options",
			fields: fields{},
			args:   args{},
			want:   &CreateOptions{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &CreateOptions{}
			got := o.ApplyOptions(tt.args.opts)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestDeleteOptions_ApplyOptions(t *testing.T) {
	type fields struct {
	}
	type args struct {
		opts []DeleteOption
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DeleteOptions
	}{
		{
			name:   "No options",
			fields: fields{},
			args:   args{},
			want:   &DeleteOptions{},
		},
		{
			name:   "From options",
			fields: fields{},
			args: args{
				opts: []DeleteOption{
					&DeleteOptions{},
				},
			},
			want: &DeleteOptions{},
		},
		{
			name:   "Overwrite existing options",
			fields: fields{},
			args: args{
				opts: []DeleteOption{
					&DeleteOptions{},
				},
			},
			want: &DeleteOptions{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &DeleteOptions{}
			got := o.ApplyOptions(tt.args.opts)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestGetOptions_ApplyOptions(t *testing.T) {
	type fields struct {
		SkipCache    bool
		RefreshCache bool
	}
	type args struct {
		opts []GetOption
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *GetOptions
	}{
		{
			name:   "No options",
			fields: fields{},
			args:   args{},
			want:   &GetOptions{},
		},
		{
			name:   "From options",
			fields: fields{},
			args: args{
				opts: []GetOption{
					&GetOptions{
						SkipCache:    true,
						RefreshCache: true,
					},
				},
			},
			want: &GetOptions{
				SkipCache:    true,
				RefreshCache: true,
			},
		},
		{
			name: "Overwrite existing options",
			fields: fields{
				SkipCache:    true,
				RefreshCache: true,
			},
			args: args{
				opts: []GetOption{
					&GetOptions{
						SkipCache:    false,
						RefreshCache: false,
					},
				},
			},
			want: &GetOptions{
				SkipCache:    false,
				RefreshCache: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &GetOptions{
				SkipCache:    tt.fields.SkipCache,
				RefreshCache: tt.fields.RefreshCache,
			}
			got := o.ApplyOptions(tt.args.opts)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestListOptions_ApplyOptions(t *testing.T) {
	type fields struct {
		SkipCache    bool
		RefreshCache bool
	}
	type args struct {
		opts []ListOption
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ListOptions
	}{
		{
			name:   "No options",
			fields: fields{},
			args:   args{},
			want:   &ListOptions{},
		},
		{
			name:   "From options",
			fields: fields{},
			args: args{
				opts: []ListOption{
					&ListOptions{
						SkipCache:    true,
						RefreshCache: true,
					},
				},
			},
			want: &ListOptions{
				SkipCache:    true,
				RefreshCache: true,
			},
		},
		{
			name: "Overwrite existing options",
			fields: fields{
				SkipCache:    true,
				RefreshCache: true,
			},
			args: args{
				opts: []ListOption{
					&ListOptions{
						SkipCache:    false,
						RefreshCache: false,
					},
				},
			},
			want: &ListOptions{
				SkipCache:    false,
				RefreshCache: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &ListOptions{
				SkipCache:    tt.fields.SkipCache,
				RefreshCache: tt.fields.RefreshCache,
			}
			got := o.ApplyOptions(tt.args.opts)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestUpdateOptions_ApplyOptions(t *testing.T) {
	type fields struct {
	}
	type args struct {
		opts []UpdateOption
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateOptions
	}{
		{
			name:   "No options",
			fields: fields{},
			args:   args{},
			want:   &UpdateOptions{},
		},
		{
			name:   "From options",
			fields: fields{},
			args: args{
				opts: []UpdateOption{
					&UpdateOptions{},
				},
			},
			want: &UpdateOptions{},
		},
		{
			name:   "Overwrite existing options",
			fields: fields{},
			args: args{
				opts: []UpdateOption{
					&UpdateOptions{},
				},
			},
			want: &UpdateOptions{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &UpdateOptions{}
			got := o.ApplyOptions(tt.args.opts)
			require.Equal(t, tt.want, got)
		})
	}
}
