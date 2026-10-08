// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-FileCopyrightText: Copyright 2018 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"

	"github.com/SlinkyProject/slurm-client/pkg/cache"
	"github.com/SlinkyProject/slurm-client/pkg/client/token"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

// ObjectKey identifies a Slurm Object.
type ObjectKey = object.ObjectKey

// ObjectKeyFromObject returns the ObjectKey given a object.Object.
func ObjectKeyFromObject(obj object.Object) ObjectKey {
	return obj.GetKey()
}

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
	// By default, Create reads the complete object back after creation. If that read
	// fails, Create returns the read error and obj retains the identity Slurm accepted.
	// Do not resubmit the create request in this case; retry Get using that identity.
	// Set CreateOptions.SkipReadAfterCreate to return the identity without reading back.
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

	SetServer(server string)
	GetServer() string

	// SetTokenProvider replaces the TokenProvider used by subsequent requests.
	// It is safe to call concurrently with client requests and informers.
	SetTokenProvider(tokenProvider token.Provider)

	// GetToken returns the latest token successfully resolved by the provider.
	GetToken() string

	// Versioned returns a handle to call a specific raw versioned client.
	Versioned() VersionedClients
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
	Indexer
}

// IndexFunc extracts zero or more secondary keys from an object. It must be
// deterministic, must not modify the object, and must not call back into the cache.
// Return no keys for objects that do not belong in the index.
type IndexFunc func(object.Object) []string

// Indexer supports local secondary indexes over an informer's objects.
type Indexer interface {
	// AddIndex registers an index and indexes any existing objects atomically.
	// Empty names, nil extractors, and duplicate names are rejected.
	AddIndex(name string, extract IndexFunc) error

	// ByIndex appends deep copies of matching cached objects to list. It has the
	// same synchronization semantics as a cached List and makes no live request.
	// The list must match the informer's object type. An unknown index is an error;
	// a known index with no matches returns an empty result, not ErrNotFound.
	// A cache miss is not proof that an object does not exist on the server.
	ByIndex(ctx context.Context, name, value string, list object.ObjectList) error
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

	// Run starts and runs the informer, returning after it stops.
	// The informer will be stopped when stopCh is closed.
	Run(stopCh <-chan struct{})
}
