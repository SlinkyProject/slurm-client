// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-FileCopyrightText: Copyright 2018 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"time"

	"github.com/SlinkyProject/slurm-client/pkg/object"
)

const (
	DefaultWatchInterval time.Duration = 30 * time.Second
)

// {{{ "Functional" Option Interfaces

// ClientOption is some configuration that modifies options for the client.
type ClientOption interface {
	// ApplyToCreate applies this configuration to the given client options.
	ApplyToClient(*ClientOptions)
}

// CreateOption is some configuration that modifies options for a create request.
type CreateOption interface {
	// ApplyToCreate applies this configuration to the given create options.
	ApplyToCreate(*CreateOptions)
}

// DeleteOption is some configuration that modifies options for a delete request.
type DeleteOption interface {
	// ApplyToDelete applies this configuration to the given delete options.
	ApplyToDelete(*DeleteOptions)
}

// GetOption is some configuration that modifies options for a get request.
type GetOption interface {
	// ApplyToGet applies this configuration to the given get options.
	ApplyToGet(*GetOptions)
}

// ListOption is some configuration that modifies options for a list request.
type ListOption interface {
	// ApplyToList applies this configuration to the given list options.
	ApplyToList(*ListOptions)
}

// UpdateOption is some configuration that modifies options for a update request.
type UpdateOption interface {
	// ApplyToUpdate applies this configuration to the given update options.
	ApplyToUpdate(*UpdateOptions)
}

// }}}

// {{{ Client Options

// ClientOptions contains options for creating a client to make requests to Slurm.
type ClientOptions struct {
	// EnableFor indicates the list of resources for the informer to ensure.
	EnableFor []object.Object

	// DisableFor indicates the list of resources for the informer to ignore.
	DisableFor []object.Object

	// CacheSyncPeriod is the time to wait before updating the cache
	CacheSyncPeriod time.Duration
}

// ApplyOptions applies the given create options on these options,
// and then returns itself (for convenient chaining).
func (o *ClientOptions) ApplyOptions(opts []ClientOption) *ClientOptions {
	for _, opt := range opts {
		opt.ApplyToClient(o)
	}
	return o
}

// ApplyToClient implements ClientOption.
func (o *ClientOptions) ApplyToClient(co *ClientOptions) {
	co.DisableFor = append(co.DisableFor, o.DisableFor...)
	co.EnableFor = append(co.EnableFor, o.EnableFor...)
	if o.CacheSyncPeriod <= 0 {
		o.CacheSyncPeriod = defaultSyncPeriod
	}
	co.CacheSyncPeriod = o.CacheSyncPeriod
}

var _ ClientOption = &ClientOptions{}

// }}}

// {{{ Create Options

// CreateOptions contains options for create requests. It's generally a subset
// of metav1.CreateOptions.
type CreateOptions struct {
}

// ApplyOptions applies the given create options on these options,
// and then returns itself (for convenient chaining).
func (o *CreateOptions) ApplyOptions(opts []CreateOption) *CreateOptions {
	for _, opt := range opts {
		opt.ApplyToCreate(o)
	}
	return o
}

// ApplyToCreate implements CreateOption.
func (o *CreateOptions) ApplyToCreate(co *CreateOptions) {
}

var _ CreateOption = &CreateOptions{}

// }}}

// {{{ Delete Options

// DeleteOptions contains options for delete requests. It's generally a subset
// of metav1.DeleteOptions.
type DeleteOptions struct {
}

// ApplyOptions applies the given delete options on these options,
// and then returns itself (for convenient chaining).
func (o *DeleteOptions) ApplyOptions(opts []DeleteOption) *DeleteOptions {
	for _, opt := range opts {
		opt.ApplyToDelete(o)
	}
	return o
}

var _ DeleteOption = &DeleteOptions{}

// ApplyToDelete implements DeleteOption.
func (o *DeleteOptions) ApplyToDelete(do *DeleteOptions) {
}

// }}}

// {{{ Get Options

// GetOptions contains options for get operation.
// Now it only has a Raw field, with support for specific resourceVersion.
type GetOptions struct {
	// SkipCache indicates not to use cache when fetching data.
	SkipCache bool

	// RefreshCache indicates to refresh the cache before reading from it.
	RefreshCache bool
}

var _ GetOption = &GetOptions{}

// ApplyToGet implements GetOption for GetOptions.
func (o *GetOptions) ApplyToGet(lo *GetOptions) {
	lo.SkipCache = o.SkipCache
	lo.RefreshCache = o.RefreshCache
}

// ApplyOptions applies the given get options on these options,
// and then returns itself (for convenient chaining).
func (o *GetOptions) ApplyOptions(opts []GetOption) *GetOptions {
	for _, opt := range opts {
		opt.ApplyToGet(o)
	}
	return o
}

// }}}

// {{{ List Options

// ListOptions contains options for limiting or filtering results.
// It's generally a subset of metav1.ListOptions, with support for
// pre-parsed selectors (since generally, selectors will be executed
// against the cache).
type ListOptions struct {
	// SkipCache indicates not to use cache when fetching data.
	SkipCache bool

	// RefreshCache indicates to refresh the cache before reading from it.
	RefreshCache bool
}

var _ ListOption = &ListOptions{}

// ApplyToList implements ListOption for ListOptions.
func (o *ListOptions) ApplyToList(lo *ListOptions) {
	lo.SkipCache = o.SkipCache
	lo.RefreshCache = o.RefreshCache
}

// ApplyOptions applies the given list options on these options,
// and then returns itself (for convenient chaining).
func (o *ListOptions) ApplyOptions(opts []ListOption) *ListOptions {
	for _, opt := range opts {
		opt.ApplyToList(o)
	}
	return o
}

// }}}

// {{{ Update Options

// UpdateOptions contains options for create requests.
type UpdateOptions struct {
}

// ApplyOptions applies the given update options on these options,
// and then returns itself (for convenient chaining).
func (o *UpdateOptions) ApplyOptions(opts []UpdateOption) *UpdateOptions {
	for _, opt := range opts {
		opt.ApplyToUpdate(o)
	}
	return o
}

var _ UpdateOption = &UpdateOptions{}

// ApplyToUpdate implements UpdateOption.
func (o *UpdateOptions) ApplyToUpdate(uo *UpdateOptions) {
}

// }}}
