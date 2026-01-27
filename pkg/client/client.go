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

	"k8s.io/utils/ptr"
	"k8s.io/utils/set"

	v0041 "github.com/SlinkyProject/slurm-client/pkg/client/api/v0041"
	v0042 "github.com/SlinkyProject/slurm-client/pkg/client/api/v0042"
	v0043 "github.com/SlinkyProject/slurm-client/pkg/client/api/v0043"
	v0044 "github.com/SlinkyProject/slurm-client/pkg/client/api/v0044"
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

	v0041Client v0041.ClientInterface
	v0042Client v0042.ClientInterface
	v0043Client v0043.ClientInterface
	v0044Client v0044.ClientInterface

	config Config

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
		DisableFor: []object.Object{
			&types.V0044NodeResourceLayout{},
			&types.V0044Reconfigure{},
			&types.V0043Reconfigure{},
			&types.V0042Reconfigure{},
			&types.V0041Reconfigure{},
		},
	}
	options.ApplyOptions(opts)

	// create return client object
	c := &client{
		informers:       make(map[object.ObjectType]InformerCache),
		uncached:        make(set.Set[object.ObjectType]),
		config:          ptr.Deref(config, Config{}),
		cacheSyncPeriod: options.CacheSyncPeriod,
	}

	c.ctx, c.cancel = context.WithCancel(context.Background())

	if err := c.createApiClients(); err != nil {
		return nil, fmt.Errorf("unable to create client: %w", err)
	}

	for _, obj := range options.EnableFor {
		c.GetInformer(obj.GetType())
	}
	for _, obj := range options.DisableFor {
		c.uncached.Insert(obj.GetType())
	}

	return c, nil
}

func (c *client) createApiClients() error {
	var err error

	c.mu.Lock()
	defer c.mu.Unlock()

	c.v0041Client, err = v0041.NewSlurmClient(c.config.Server, c.config.AuthToken, c.config.HTTPClient)
	if err != nil {
		return fmt.Errorf("unable to create client: %w", err)
	}

	c.v0042Client, err = v0042.NewSlurmClient(c.config.Server, c.config.AuthToken, c.config.HTTPClient)
	if err != nil {
		return fmt.Errorf("unable to create client: %w", err)
	}

	c.v0043Client, err = v0043.NewSlurmClient(c.config.Server, c.config.AuthToken, c.config.HTTPClient)
	if err != nil {
		return fmt.Errorf("unable to create client: %w", err)
	}

	c.v0044Client, err = v0044.NewSlurmClient(c.config.Server, c.config.AuthToken, c.config.HTTPClient)
	if err != nil {
		return fmt.Errorf("unable to create client: %w", err)
	}

	return nil
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

	switch o := obj.(type) {
	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0041JobInfo:
		var jobId *int32
		var err error
		if options.Allocation {
			jobId, err = c.v0041Client.CreateJobInfoAlloc(ctx, req)
		} else {
			jobId, err = c.v0041Client.CreateJobInfo(ctx, req)
		}
		if err != nil {
			return err
		}
		key := object.ObjectKey(fmt.Sprintf("%d", *jobId))
		return c.Get(ctx, key, o)

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0042JobInfo:
		var jobId *int32
		var err error
		if options.Allocation {
			jobId, err = c.v0042Client.CreateJobInfoAlloc(ctx, req)
		} else {
			jobId, err = c.v0042Client.CreateJobInfo(ctx, req)
		}
		if err != nil {
			return err
		}
		key := object.ObjectKey(fmt.Sprintf("%d", *jobId))
		return c.Get(ctx, key, o)

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0043JobInfo:
		var jobId *int32
		var err error
		if options.Allocation {
			jobId, err = c.v0043Client.CreateJobInfoAlloc(ctx, req)
		} else {
			jobId, err = c.v0043Client.CreateJobInfo(ctx, req)
		}
		if err != nil {
			return err
		}
		key := object.ObjectKey(fmt.Sprintf("%d", *jobId))
		return c.Get(ctx, key, o)

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0044JobInfo:
		var jobId *int32
		var err error
		if options.Allocation {
			jobId, err = c.v0044Client.CreateJobInfoAlloc(ctx, req)
		} else {
			jobId, err = c.v0044Client.CreateJobInfo(ctx, req)
		}
		if err != nil {
			return err
		}
		key := object.ObjectKey(fmt.Sprintf("%d", *jobId))
		return c.Get(ctx, key, o)
	case *types.V0044ReservationInfo:
		reservationName, err := c.v0044Client.CreateReservationInfo(ctx, req)
		if err != nil {
			return err
		}
		key := object.ObjectKey(reservationName)
		return c.Get(ctx, key, o)

	case *types.V0044Node:
		nodeName, err := c.v0044Client.CreateNewNode(ctx, req)
		if err != nil {
			return err
		}
		key := object.ObjectKey(*nodeName)
		return c.Get(ctx, key, o, &GetOptions{SkipCache: true})

	/////////////////////////////////////////////////////////////////////////////////

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
	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0041JobInfo:
		return c.v0041Client.DeleteJobInfo(ctx, key)
	case *types.V0041Node:
		return c.v0041Client.DeleteNode(ctx, key)

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0042JobInfo:
		return c.v0042Client.DeleteJobInfo(ctx, key)
	case *types.V0042Node:
		return c.v0042Client.DeleteNode(ctx, key)

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0043JobInfo:
		return c.v0043Client.DeleteJobInfo(ctx, key)
	case *types.V0043Node:
		return c.v0043Client.DeleteNode(ctx, key)

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0044JobInfo:
		return c.v0044Client.DeleteJobInfo(ctx, key)
	case *types.V0044Node:
		return c.v0044Client.DeleteNode(ctx, key)
	case *types.V0044ReservationInfo:
		return c.v0044Client.DeleteReservationInfo(ctx, key)

	/////////////////////////////////////////////////////////////////////////////////

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
	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0041JobInfo:
		err := c.v0041Client.UpdateJobInfo(ctx, key, req)
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

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0042JobInfo:
		err := c.v0042Client.UpdateJobInfo(ctx, key, req)
		if err != nil {
			return err
		}
		return c.Get(ctx, obj.GetKey(), o)
	case *types.V0042Node:
		err := c.v0042Client.UpdateNode(ctx, key, req)
		if err != nil {
			return err
		}
		return c.Get(ctx, obj.GetKey(), o)

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0043JobInfo:
		err := c.v0043Client.UpdateJobInfo(ctx, key, req)
		if err != nil {
			return err
		}
		return c.Get(ctx, obj.GetKey(), o)
	case *types.V0043Node:
		err := c.v0043Client.UpdateNode(ctx, key, req)
		if err != nil {
			return err
		}
		return c.Get(ctx, obj.GetKey(), o)

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0044JobInfo:
		err := c.v0044Client.UpdateJobInfo(ctx, key, req)
		if err != nil {
			return err
		}
		return c.Get(ctx, obj.GetKey(), o)
	case *types.V0044Node:
		err := c.v0044Client.UpdateNode(ctx, key, req)
		if err != nil {
			return err
		}
		return c.Get(ctx, obj.GetKey(), o)
	case *types.V0044ReservationInfo:
		err := c.v0044Client.UpdateReservationInfo(ctx, req)
		if err != nil {
			return err
		}
		return c.Get(ctx, obj.GetKey(), o)

	/////////////////////////////////////////////////////////////////////////////////

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
	/////////////////////////////////////////////////////////////////////////////////

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
	case *types.V0041Reconfigure:
		out, err := c.v0041Client.GetReconfigure(ctx)
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0041Stats:
		out, err := c.v0041Client.GetStats(ctx)
		if err != nil {
			return err
		}
		*o = *out

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0042ControllerPing:
		out, err := c.v0042Client.GetControllerPing(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0042JobInfo:
		out, err := c.v0042Client.GetJobInfo(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0042Node:
		out, err := c.v0042Client.GetNode(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0042PartitionInfo:
		out, err := c.v0042Client.GetPartitionInfo(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0042Reconfigure:
		out, err := c.v0042Client.GetReconfigure(ctx)
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0042Stats:
		out, err := c.v0042Client.GetStats(ctx)
		if err != nil {
			return err
		}
		*o = *out

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0043ControllerPing:
		out, err := c.v0043Client.GetControllerPing(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0043JobInfo:
		out, err := c.v0043Client.GetJobInfo(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0043Node:
		out, err := c.v0043Client.GetNode(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0043PartitionInfo:
		out, err := c.v0043Client.GetPartitionInfo(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0043Reconfigure:
		out, err := c.v0043Client.GetReconfigure(ctx)
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0043Stats:
		out, err := c.v0043Client.GetStats(ctx)
		if err != nil {
			return err
		}
		*o = *out

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0044ControllerPing:
		out, err := c.v0044Client.GetControllerPing(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0044JobInfo:
		out, err := c.v0044Client.GetJobInfo(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0044NodeResourceLayout:
		out, err := c.v0044Client.GetNodeResourceLayout(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0044Node:
		out, err := c.v0044Client.GetNode(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0044PartitionInfo:
		out, err := c.v0044Client.GetPartitionInfo(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0044Reconfigure:
		out, err := c.v0044Client.GetReconfigure(ctx)
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0044ReservationInfo:
		out, err := c.v0044Client.GetReservationInfo(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0044Stats:
		out, err := c.v0044Client.GetStats(ctx)
		if err != nil {
			return err
		}
		*o = *out

	/////////////////////////////////////////////////////////////////////////////////

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
	/////////////////////////////////////////////////////////////////////////////////

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
	case *types.V0041ReconfigureList:
		out, err := c.v0041Client.ListReconfigure(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0041StatsList:
		out, err := c.v0041Client.ListStats(ctx)
		if err != nil {
			return err
		}
		*objList = *out

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0042ControllerPingList:
		out, err := c.v0042Client.ListControllerPing(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0042JobInfoList:
		out, err := c.v0042Client.ListJobInfo(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0042NodeList:
		out, err := c.v0042Client.ListNodes(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0042PartitionInfoList:
		out, err := c.v0042Client.ListPartitionInfo(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0042ReconfigureList:
		out, err := c.v0042Client.ListReconfigure(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0042StatsList:
		out, err := c.v0042Client.ListStats(ctx)
		if err != nil {
			return err
		}
		*objList = *out

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0043ControllerPingList:
		out, err := c.v0043Client.ListControllerPing(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0043JobInfoList:
		out, err := c.v0043Client.ListJobInfo(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0043NodeList:
		out, err := c.v0043Client.ListNodes(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0043PartitionInfoList:
		out, err := c.v0043Client.ListPartitionInfo(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0043ReconfigureList:
		out, err := c.v0043Client.ListReconfigure(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0043StatsList:
		out, err := c.v0043Client.ListStats(ctx)
		if err != nil {
			return err
		}
		*objList = *out

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0044ControllerPingList:
		out, err := c.v0044Client.ListControllerPing(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0044JobInfoList:
		out, err := c.v0044Client.ListJobInfo(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0044NodeList:
		out, err := c.v0044Client.ListNodes(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0044PartitionInfoList:
		out, err := c.v0044Client.ListPartitionInfo(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0044ReconfigureList:
		out, err := c.v0044Client.ListReconfigure(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0044ReservationInfoList:
		out, err := c.v0044Client.ListReservationInfo(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0044StatsList:
		out, err := c.v0044Client.ListStats(ctx)
		if err != nil {
			return err
		}
		*objList = *out

	/////////////////////////////////////////////////////////////////////////////////

	default:
		return errors.New(http.StatusText(http.StatusNotImplemented))
	}

	return nil
}

// GetServer returns the client server.
func (c *client) GetServer() string {
	return c.config.Server
}

// GetServer returns the client server.
func (c *client) SetServer(server string) {
	c.config.Server = server
	if err := c.createApiClients(); err != nil {
		panic(fmt.Errorf("unable to create client: %w", err))
	}
}

// GetToken returns the client token.
func (c *client) GetToken() string {
	return c.config.AuthToken
}

// GetToken returns the client token.
func (c *client) SetToken(token string) {
	c.config.AuthToken = token
	if err := c.createApiClients(); err != nil {
		panic(fmt.Errorf("unable to create client: %w", err))
	}
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
	// Reverify existence after acquiring lock
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
