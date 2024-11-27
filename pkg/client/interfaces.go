// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-FileCopyrightText: Copyright 2018 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"time"

	"k8s.io/client-go/tools/cache"

	"github.com/SlinkyProject/slurm-client/pkg/object"
)

// Reader knows how to read and list Slurm objects.
type Reader interface {
	// Get retrieves an obj for the given object key from the Slurm Cluster.
	// obj must be a struct pointer so that obj can be updated with the response
	// returned by the Server.
	Get(ctx context.Context, key object.ObjectKey, obj object.Object, opts ...GetOption) error

	// List retrieves list of objects for a given namespace and list options. On a
	// successful call, Items field in the list will be populated with the
	// result returned from the server.
	List(ctx context.Context, list object.ObjectList, opts ...ListOption) error
}

// Writer knows how to create, delete, and update Slurm objects.
type Writer interface {
	// Create saves the object obj in the Slurm cluster. obj must be a
	// struct pointer so that obj can be updated with the content returned by the Server.
	Create(ctx context.Context, obj object.Object, req any, opts ...CreateOption) error

	// Update updates the given obj in the Slurm cluster. obj must be a
	// struct pointer so that obj can be updated with the content returned by the Server.
	Update(ctx context.Context, obj object.Object, req any, opts ...UpdateOption) error

	// Delete deletes the given obj from Slurm cluster.
	Delete(ctx context.Context, obj object.Object, opts ...DeleteOption) error
}

// Client knows how to perform CRUD operations on Slurm objects.
type Client interface {
	Reader
	Writer
	Informers

	GetServer() string
	GetToken() string
}

// Informers knows how to create or fetch informers for different Objects.
// It's safe to call GetInformer from multiple threads.
type Informers interface {
	// GetInformer fetches or constructs an informer for the given objectType that corresponds to the
	// non-list version of the resource.
	GetInformer(objectType object.ObjectType) InformerCache

	// Start runs all the informers known to this cache until the context is closed.
	// It blocks.
	Start(ctx context.Context)

	// Stop runs all the informers known to this cache until the context is closed.
	// It blocks.
	Stop()
}

type InformerCache interface {
	Informer
	Reader
	Writer
}

// Informer - informer allows you interact with the underlying informer.
type Informer interface {
	// SetEventHandler sets an event handler to the informer. When cache changes,
	// event handlers will trigger.
	SetEventHandler(handler cache.ResourceEventHandler)

	// UnsetEventHandler unsets the event handler of the informer.
	UnsetEventHandler()

	// HasSynced return true if the informers underlying store has synced.
	HasSynced() (bool, error)

	// HasStarted return true if the informers underlying store has been started.
	HasStarted() bool

	// WaitForSyncList return nil if the informers were able to sync.
	WaitForSyncList(ctx context.Context, interval time.Duration) error

	// WaitForSyncGet return nil if the informer was able to sync.
	WaitForSyncGet(ctx context.Context, interval time.Duration) error

	// Run starts and runs the informer, returning after it stops.
	// The informer will be stopped when stopCh is closed.
	Run(stopCh <-chan struct{})
}
