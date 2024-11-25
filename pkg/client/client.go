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
	slurmtypes "github.com/SlinkyProject/slurm-client/pkg/types"
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
	clients   []SlurmClientInterface
	clientUse int

	mu        sync.RWMutex
	informers map[object.ObjectType]InformerCache
	uncached  set.Set[object.ObjectType]
	started   bool
	ctx       context.Context
	cancel    context.CancelFunc

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

	slurmClientV0041, err := v0041.NewSlurmClient(config.Server, config.AuthToken, config.HTTPClient)
	if err != nil {
		return nil, fmt.Errorf("unable to create client: %v", err)
	}

	slurmClientV0040, err := v0040.NewSlurmClient(config.Server, config.AuthToken, config.HTTPClient)
	if err != nil {
		return nil, fmt.Errorf("unable to create client: %v", err)
	}

	// create return client object
	client := &client{
		clients: []SlurmClientInterface{
			slurmClientV0041,
			slurmClientV0040,
		},
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
	opts ...CreateOption,
) error {
	panic("unimplemented")
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

	switch o := obj.(type) {
	case *slurmtypes.Node:
		err := c.deleteNode(ctx, string(o.GetKey()))
		if err != nil {
			return err
		}
	default:
		return errors.New(http.StatusText(http.StatusNotImplemented))
	}

	return nil
}

// DeleteAllOf implements Client.
func (c *client) DeleteAllOf(
	ctx context.Context,
	obj object.Object,
	opts ...DeleteAllOfOption,
) error {
	panic("unimplemented")
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
	case *slurmtypes.Node:
		node, err := c.getNode(ctx, string(key))
		if err != nil {
			return err
		}
		(*o) = (*node)
	case *slurmtypes.Ping:
		ping, err := c.getPing(ctx, string(key))
		if err != nil {
			return err
		}
		(*o) = (*ping)
	case *slurmtypes.JobInfo:
		node, err := c.getJobInfo(ctx, string(key))
		if err != nil {
			return err
		}
		(*o) = (*node)
	case *slurmtypes.PartitionInfo:
		node, err := c.getPartitionInfo(ctx, string(key))
		if err != nil {
			return err
		}
		(*o) = (*node)
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
	case *slurmtypes.NodeList:
		nodeList, err := c.listNodes(ctx)
		if err != nil {
			return err
		}
		(*objList) = (*nodeList)
	case *slurmtypes.PingList:
		pingList, err := c.listPing(ctx)
		if err != nil {
			return err
		}
		(*objList) = (*pingList)
	case *slurmtypes.JobInfoList:
		jobInfoList, err := c.listJobInfos(ctx)
		if err != nil {
			return err
		}
		(*objList) = (*jobInfoList)
	case *slurmtypes.PartitionInfoList:
		partitionInfoList, err := c.listPartitionInfos(ctx)
		if err != nil {
			return err
		}
		(*objList) = (*partitionInfoList)
	default:
		return errors.New(http.StatusText(http.StatusNotImplemented))
	}

	return nil
}

// Update implements Client.
func (c *client) Update(
	ctx context.Context,
	obj object.Object,
	opts ...UpdateOption,
) error {
	// Apply options
	options := &UpdateOptions{}
	options.ApplyOptions(opts)

	switch o := obj.(type) {
	case *slurmtypes.Node:
		// Get original node to get deltas from
		originalNode := &slurmtypes.Node{}
		if err := c.Get(ctx, object.ObjectKey(o.Name), originalNode); err != nil {
			return err
		}
		if err := c.updateNode(ctx, o, originalNode); err != nil {
			return err
		}
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
	c.mu.RLock()
	if c.started {
		defer c.mu.RUnlock()
		return
	}
	c.mu.RUnlock()

	c.mu.Lock()
	c.ctx, c.cancel = context.WithCancel(ctx)
	c.mu.Unlock()

	ticker := time.NewTicker(defaultSyncPeriod)
	defer ticker.Stop()
	stopCh := make(chan struct{})
	for {
		c.mu.Lock()
		for objectType, ic := range c.informers {
			if c.uncached.Has(objectType) {
				continue
			}
			go ic.Run(stopCh)
		}
		c.started = true
		c.mu.Unlock()

		select {
		case <-ticker.C:
			// wait for tick
		case <-ctx.Done():
			c.mu.Lock()
			defer c.mu.Unlock()
			c.started = false
			close(stopCh)
			return
		}
	}
}

// Stop implements Client.
func (c *client) Stop() {
	c.mu.RLock()
	if !c.started {
		defer c.mu.RUnlock()
		return
	}
	c.mu.RUnlock()

	c.mu.Lock()
	c.cancel()
	c.started = false
	c.mu.Unlock()
}

func tolerateError(err error) bool {
	errText := err.Error()
	if errText == http.StatusText(http.StatusNotFound) ||
		errText == http.StatusText(http.StatusNoContent) {
		return true
	}
	return false
}
