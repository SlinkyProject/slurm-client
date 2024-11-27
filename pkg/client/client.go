// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"k8s.io/utils/set"

	v0040 "github.com/SlinkyProject/slurm-client/pkg/client/api/v0040"
	v0041 "github.com/SlinkyProject/slurm-client/pkg/client/api/v0041"
	"github.com/SlinkyProject/slurm-client/pkg/event"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

// Config holds the common attributes that can be passed to a Slurm client on
// initialization.
type Config struct {
	// The server with port (e.g. `http://localhost:8080/`).
	// +required
	Server string

	// The Slurm JWT token for authentication.
	// +required
	AuthToken string

	// HTTPClient is the HTTP client to use for requests.
	HTTPClient *http.Client
}

func validate(config *Config) error {
	switch {
	case config == nil:
		return fmt.Errorf("config cannot be nil")
	case config.Server == "":
		return fmt.Errorf("server cannot be empty")
	case config.AuthToken == "":
		return fmt.Errorf("authToken cannot be empty")
	}
	return nil
}

type client struct {
	mu        sync.RWMutex
	informers map[object.ObjectType]InformerCache
	uncached  set.Set[object.ObjectType]
	started   bool
	stopped   bool
	ctx       context.Context
	cancel    context.CancelFunc

	v0040Client v0040.ClientInterface
	v0041Client v0041.ClientInterface

	server          string
	authToken       string
	cacheSyncPeriod time.Duration
}

// NewClient initializes a client.
func NewClient(config *Config, opts ...ClientOption) (Client, error) {
	if err := validate(config); err != nil {
		return nil, err
	}

	// Apply options
	options := &ClientOptions{
		CacheSyncPeriod: defaultSyncPeriod,
	}
	options.ApplyOptions(opts)

	v0040Client, err := v0040.NewSlurmClient(config.Server, config.AuthToken, config.HTTPClient)
	if err != nil {
		return nil, fmt.Errorf("unable to create client: %v", err)
	}

	v0041Client, err := v0041.NewSlurmClient(config.Server, config.AuthToken, config.HTTPClient)
	if err != nil {
		return nil, fmt.Errorf("unable to create client: %v", err)
	}

	// create return client object
	client := &client{
		v0040Client:     v0040Client,
		v0041Client:     v0041Client,
		informers:       make(map[object.ObjectType]InformerCache),
		uncached:        make(set.Set[object.ObjectType]),
		authToken:       config.AuthToken,
		server:          config.Server,
		cacheSyncPeriod: options.CacheSyncPeriod,
	}

	for _, obj := range options.EnableFor {
		client.GetInformer(obj.GetType())
	}
	for _, obj := range options.DisableFor {
		client.uncached.Insert(obj.GetType())
	}

	return client, nil
}

// Create implements Client.
func (c *client) Create(
	ctx context.Context,
	obj object.Object,
	req any,
	opts ...CreateOption,
) error {
	// Apply options
	options := &CreateOptions{}
	options.ApplyOptions(opts)

	switch obj.(type) {
	default:
		return errors.New(http.StatusText(http.StatusNotImplemented))
	}
}

// Delete implements Client.
func (c *client) Delete(
	ctx context.Context,
	obj object.Object,
	opts ...DeleteOption,
) error {
	// Apply options
	options := &DeleteOptions{}
	options.ApplyOptions(opts)

	key := string(obj.GetKey())
	switch obj.(type) {
	case *types.V0040Node:
		return c.v0040Client.DeleteNode(ctx, key)

	case *types.V0041Node:
		return c.v0041Client.DeleteNode(ctx, key)

	default:
		return errors.New(http.StatusText(http.StatusNotImplemented))
	}
}

// Update implements Client.
func (c *client) Update(
	ctx context.Context,
	obj object.Object,
	req any,
	opts ...UpdateOption,
) error {
	// Apply options
	options := &UpdateOptions{}
	options.ApplyOptions(opts)

	key := string(obj.GetKey())
	switch o := obj.(type) {
	case *types.V0040Node:
		err := c.v0040Client.UpdateNode(ctx, key, req)
		if err != nil {
			return err
		}
		return c.Get(ctx, obj.GetKey(), o)

	case *types.V0041Node:
		err := c.v0041Client.UpdateNode(ctx, key, req)
		if err != nil {
			return err
		}
		return c.Get(ctx, obj.GetKey(), o)

	default:
		return errors.New(http.StatusText(http.StatusNotImplemented))
	}
}

// Get implements Client.
func (c *client) Get(
	ctx context.Context,
	key object.ObjectKey,
	obj object.Object,
	opts ...GetOption,
) error {
	// Apply options
	options := &GetOptions{}
	options.ApplyOptions(opts)

	if !options.SkipCache {
		objectType := obj.GetType()
		objectType = object.ObjectType(strings.TrimSuffix(string(objectType), "List"))
		informerCache := c.GetInformer(objectType)
		if informerCache.HasStarted() && !c.uncached.Has(objectType) {
			return informerCache.Get(ctx, key, obj, opts...)
		}
	}

	switch o := obj.(type) {
	case *types.V0040ControllerPing:
		out, err := c.v0040Client.GetControllerPing(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0040JobInfo:
		out, err := c.v0040Client.GetJobInfo(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0040Node:
		out, err := c.v0040Client.GetNode(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0040PartitionInfo:
		out, err := c.v0040Client.GetPartitionInfo(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out

	case *types.V0041ControllerPing:
		out, err := c.v0041Client.GetControllerPing(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0041JobInfo:
		out, err := c.v0041Client.GetJobInfo(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0041Node:
		out, err := c.v0041Client.GetNode(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0041PartitionInfo:
		out, err := c.v0041Client.GetPartitionInfo(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out

	default:
		return errors.New(http.StatusText(http.StatusNotImplemented))
	}

	return nil
}

// List implements Client.
func (c *client) List(
	ctx context.Context,
	list object.ObjectList,
	opts ...ListOption,
) error {
	// Apply options
	options := &ListOptions{}
	options.ApplyOptions(opts)

	if !options.SkipCache {
		objectType := list.GetType()
		objectType = object.ObjectType(strings.TrimSuffix(string(objectType), "List"))
		informerCache := c.GetInformer(objectType)
		if informerCache.HasStarted() && !c.uncached.Has(objectType) {
			return informerCache.List(ctx, list, opts...)
		}
	}

	// Determine ObjectList type
	switch objList := list.(type) {
	case *types.V0040ControllerPingList:
		out, err := c.v0040Client.ListControllerPing(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0040JobInfoList:
		out, err := c.v0040Client.ListJobInfo(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0040NodeList:
		out, err := c.v0040Client.ListNodes(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0040PartitionInfoList:
		out, err := c.v0040Client.ListPartitionInfo(ctx)
		if err != nil {
			return err
		}
		*objList = *out

	case *types.V0041ControllerPingList:
		out, err := c.v0041Client.ListControllerPing(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0041JobInfoList:
		out, err := c.v0041Client.ListJobInfo(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0041NodeList:
		out, err := c.v0041Client.ListNodes(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0041PartitionInfoList:
		out, err := c.v0041Client.ListPartitionInfo(ctx)
		if err != nil {
			return err
		}
		*objList = *out

	default:
		return errors.New(http.StatusText(http.StatusNotImplemented))
	}

	return nil
}

// GetServer returns the client server.
func (c *client) GetServer() string {
	return c.server
}

// GetToken returns the client token.
func (c *client) GetToken() string {
	return c.authToken
}

// GetInformer implements Client.
func (c *client) GetInformer(objectType object.ObjectType) InformerCache {
	objectType = object.ObjectType(strings.TrimSuffix(string(objectType), "List"))

	c.mu.RLock()
	if informerCache, ok := c.informers[objectType]; ok {
		defer c.mu.RUnlock()
		return informerCache
	}
	c.mu.RUnlock()

	c.mu.Lock()
	defer c.mu.Unlock()
	// Reverify existance after aquiring lock
	if informerCache, ok := c.informers[objectType]; ok {
		return informerCache
	}
	// Ensure informer cache exists
	c.informers[objectType] = &informerCache{
		reader:     c,
		writer:     c,
		objectType: objectType,
		cache:      make(map[object.ObjectKey]*cacheEntry),
		syncPeriod: c.cacheSyncPeriod,
		eventCh:    make(chan event.Event),
		syncCh:     make(chan bool),
		syncObjCh:  make(chan object.ObjectKey),
	}
	return c.informers[objectType]
}

// Start implements Client.
func (c *client) Start(ctx context.Context) {
	c.mu.Lock()
	if c.started || c.stopped || ctx.Err() != nil {
		defer c.mu.Unlock()
		return
	}
	c.ctx, c.cancel = context.WithCancel(ctx)
	ticker := time.NewTicker(defaultSyncPeriod)
	defer ticker.Stop()
	stopCh := make(chan struct{})
	c.started = true
	c.stopped = false
	c.mu.Unlock()

	for {
		c.mu.Lock()
		select {
		case <-ctx.Done():
			defer c.mu.Unlock()
			c.started = false
			c.stopped = true
			close(stopCh)
			return
		case <-c.ctx.Done():
			// c.ctx is a copy of ctx and technically not the same context for
			// determining when Done() is emitted.
			defer c.mu.Unlock()
			c.started = false
			c.stopped = true
			close(stopCh)
			return
		default:
			// do not block
		}
		for objectType, ic := range c.informers {
			if c.uncached.Has(objectType) {
				continue
			}
			go ic.Run(stopCh)
		}
		c.mu.Unlock()

		select {
		case <-ticker.C:
			// wait for tick
		case <-ctx.Done():
			c.mu.Lock()
			defer c.mu.Unlock()
			c.started = false
			c.stopped = true
			close(stopCh)
			return
		case <-c.ctx.Done():
			// c.ctx is a copy of ctx and technically not the same context for
			// determining when Done() is emitted.
			c.mu.Lock()
			defer c.mu.Unlock()
			c.started = false
			c.stopped = true
			close(stopCh)
			return
		}
	}
}

// Stop implements Client.
func (c *client) Stop() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.stopped = true
	if !c.started {
		return
	}

	c.cancel()
	c.started = false
}
