// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/api/equality"

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
			if got := o.ApplyOptions(tt.args.opts); !equality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("ClientOptions.ApplyOptions() = %v, want %v", got, tt.want)
			}
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
			if got := o.ApplyOptions(tt.args.opts); !equality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("CreateOptions.ApplyOptions() = %v, want %v", got, tt.want)
			}
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
			if got := o.ApplyOptions(tt.args.opts); !equality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("DeleteOptions.ApplyOptions() = %v, want %v", got, tt.want)
			}
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
			if got := o.ApplyOptions(tt.args.opts); !equality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("GetOptions.ApplyOptions() = %v, want %v", got, tt.want)
			}
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
			if got := o.ApplyOptions(tt.args.opts); !equality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("ListOptions.ApplyOptions() = %v, want %v", got, tt.want)
			}
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
			if got := o.ApplyOptions(tt.args.opts); !equality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("UpdateOptions.ApplyOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
