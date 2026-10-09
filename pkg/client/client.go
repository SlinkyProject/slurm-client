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

	apiv0042 "github.com/SlinkyProject/slurm-client/api/v0042"
	apiv0043 "github.com/SlinkyProject/slurm-client/api/v0043"
	apiv0044 "github.com/SlinkyProject/slurm-client/api/v0044"
	apiv0045 "github.com/SlinkyProject/slurm-client/api/v0045"
	clientapi "github.com/SlinkyProject/slurm-client/pkg/client/api"
	v0042 "github.com/SlinkyProject/slurm-client/pkg/client/api/v0042"
	v0043 "github.com/SlinkyProject/slurm-client/pkg/client/api/v0043"
	v0044 "github.com/SlinkyProject/slurm-client/pkg/client/api/v0044"
	v0045 "github.com/SlinkyProject/slurm-client/pkg/client/api/v0045"
	tokenprovider "github.com/SlinkyProject/slurm-client/pkg/client/token"
	apierrors "github.com/SlinkyProject/slurm-client/pkg/errors"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

// Config holds the common attributes that can be passed to a Slurm client on
// initialization.
type Config struct {
	// The server with port (e.g. `http://localhost:8080/`).
	// +required
	Server string

	// TokenProvider supplies the Slurm JWT token for every request.
	TokenProvider tokenprovider.Provider

	// HTTPClient is the HTTP client to use for requests.
	HTTPClient *http.Client
}

func validate(config *Config) error {
	switch {
	case config == nil:
		return fmt.Errorf("config cannot be nil")
	case config.Server == "":
		return fmt.Errorf("server cannot be empty")
	case config.TokenProvider == nil:
		return fmt.Errorf("tokenProvider cannot be nil")
	}
	return nil
}

type tokenProviderValue struct {
	tokenprovider.Provider
}

type client struct {
	mu        sync.RWMutex
	informers map[object.ObjectType]InformerCache
	uncached  set.Set[object.ObjectType]
	started   bool
	stopped   bool
	ctx       context.Context
	cancel    context.CancelFunc

	v0042Client v0042.ClientInterface
	v0043Client v0043.ClientInterface
	v0044Client v0044.ClientInterface
	v0045Client v0045.ClientInterface
	resources   map[object.ObjectType]resource

	config Config

	cacheSyncPeriod time.Duration

	tokenMu       sync.RWMutex
	authToken     string
	tokenProvider *tokenProviderValue
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

	// create return client object
	c := &client{
		informers:       make(map[object.ObjectType]InformerCache),
		uncached:        make(set.Set[object.ObjectType]),
		config:          ptr.Deref(config, Config{}),
		cacheSyncPeriod: options.CacheSyncPeriod,
	}
	c.tokenProvider = &tokenProviderValue{Provider: c.config.TokenProvider}

	c.ctx, c.cancel = context.WithCancel(context.Background())

	if err := c.createApiClients(); err != nil {
		return nil, fmt.Errorf("unable to create client: %w", err)
	}

	for _, obj := range options.EnableFor {
		c.GetInformer(obj.GetType())
	}
	for _, res := range c.resources {
		if res.cacheable {
			continue
		}
		objectType := normalizeObjectType(res.newObject().GetType())
		c.uncached.Insert(objectType)
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

	tokenProviderOption := clientapi.WithTokenProvider(tokenprovider.ProviderFunc(c.resolveToken))

	c.v0042Client, err = v0042.NewSlurmClient(c.config.Server, "", c.config.HTTPClient, tokenProviderOption)
	if err != nil {
		return fmt.Errorf("unable to create client: %w", err)
	}

	c.v0043Client, err = v0043.NewSlurmClient(c.config.Server, "", c.config.HTTPClient, tokenProviderOption)
	if err != nil {
		return fmt.Errorf("unable to create client: %w", err)
	}

	c.v0044Client, err = v0044.NewSlurmClient(c.config.Server, "", c.config.HTTPClient, tokenProviderOption)
	if err != nil {
		return fmt.Errorf("unable to create client: %w", err)
	}

	c.v0045Client, err = v0045.NewSlurmClient(c.config.Server, "", c.config.HTTPClient, tokenProviderOption)
	if err != nil {
		return fmt.Errorf("unable to create client: %w", err)
	}

	c.resources = cloneCatalog()
	bindV0042(c.resources, c.v0042Client)
	bindV0043(c.resources, c.v0043Client)
	bindV0044(c.resources, c.v0044Client)
	bindV0045(c.resources, c.v0045Client)

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

	c.mu.RLock()
	r, ok := c.resources[obj.GetType()]
	c.mu.RUnlock()
	if !ok || r.create == nil {
		return apierrors.ErrNotImplemented
	}
	key, err := r.create(ctx, req)
	if err != nil {
		return err
	}
	if err := obj.SetKey(key); err != nil {
		return fmt.Errorf("set key for created object %q: %w", key, err)
	}
	if options.SkipReadAfterCreate {
		return nil
	}
	return c.Get(ctx, key, obj, &GetOptions{RefreshCache: true})
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

	c.mu.RLock()
	r, ok := c.resources[obj.GetType()]
	c.mu.RUnlock()
	if !ok || r.delete == nil {
		return apierrors.ErrNotImplemented
	}
	if err := r.delete(ctx, string(obj.GetKey())); err != nil {
		return err
	}

	err := c.Get(ctx, obj.GetKey(), obj, &GetOptions{RefreshCache: true})
	if err != nil {
		// We expect the error to always be NotFound because we deleted the
		// object from Slurm then attempted to Get the deleted object with
		// refreshed cache.
		if !errors.Is(err, apierrors.ErrNotFound) {
			return err
		}
	}

	return nil
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

	c.mu.RLock()
	r, ok := c.resources[obj.GetType()]
	c.mu.RUnlock()
	if !ok || r.update == nil {
		return apierrors.ErrNotImplemented
	}
	if err := r.update(ctx, string(obj.GetKey()), req); err != nil {
		return err
	}

	return c.Get(ctx, obj.GetKey(), obj, &GetOptions{RefreshCache: true})
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
		objectType = normalizeObjectType(objectType)
		informerCache := c.GetInformer(objectType)
		if informerCache != nil && informerCache.HasStarted() {
			return informerCache.Get(ctx, key, obj, opts...)
		}
	}

	c.mu.RLock()
	r, ok := c.resources[obj.GetType()]
	c.mu.RUnlock()
	if !ok || r.get == nil {
		return apierrors.ErrNotImplemented
	}
	return r.get(ctx, key, obj)
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
		objectType = normalizeObjectType(objectType)
		informerCache := c.GetInformer(objectType)
		if informerCache != nil && informerCache.HasStarted() {
			return informerCache.List(ctx, list, opts...)
		}
	}

	c.mu.RLock()
	r, ok := c.resources[list.GetType()]
	c.mu.RUnlock()
	if !ok || r.list == nil {
		return apierrors.ErrNotImplemented
	}
	return r.list(ctx, list)
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

// GetToken returns the current client token.
func (c *client) GetToken() string {
	c.tokenMu.RLock()
	defer c.tokenMu.RUnlock()

	return c.authToken
}

// SetTokenProvider updates the token provider used by subsequent requests.
func (c *client) SetTokenProvider(tokenProvider tokenprovider.Provider) {
	c.tokenMu.Lock()
	defer c.tokenMu.Unlock()

	c.tokenProvider = &tokenProviderValue{Provider: tokenProvider}
	c.authToken = ""
}

func (c *client) resolveToken(ctx context.Context) (string, error) {
	c.tokenMu.RLock()
	tokenProvider := c.tokenProvider
	c.tokenMu.RUnlock()

	token, err := tokenProvider.Token(ctx)
	if err != nil {
		return "", err
	}

	c.tokenMu.Lock()
	if tokenProvider == c.tokenProvider {
		c.authToken = token
	}
	c.tokenMu.Unlock()

	return token, nil
}

// GetInformer implements Client.
func (c *client) GetInformer(objectType object.ObjectType) InformerCache {
	objectType = normalizeObjectType(objectType)

	c.mu.RLock()
	if r, ok := c.resources[objectType]; !ok || !r.cacheable {
		defer c.mu.RUnlock()
		return nil
	}

	if c.uncached.Has(objectType) {
		defer c.mu.RUnlock()
		return nil
	}

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
	c.informers[objectType] = newInformer(objectType, c, c.cacheSyncPeriod)
	return c.informers[objectType]
}

func normalizeObjectType(objectType object.ObjectType) object.ObjectType {
	return object.ObjectType(strings.TrimSuffix(string(objectType), "ObjectList"))
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
			if c.uncached.Has(objectType) || ic.HasStarted() {
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

// Versioned implements Client.
func (c *client) Versioned() VersionedClients {
	c.tokenMu.RLock()
	defer c.tokenMu.RUnlock()
	return VersionedClients{
		c: c,
	}
}

type VersionedClients struct {
	c *client
}

func (v VersionedClients) V0042() apiv0042.ClientWithResponsesInterface {
	if v.c == nil {
		return nil
	}
	v.c.mu.RLock()
	defer v.c.mu.RUnlock()
	return v.c.v0042Client
}

func (v VersionedClients) V0043() apiv0043.ClientWithResponsesInterface {
	if v.c == nil {
		return nil
	}
	v.c.mu.RLock()
	defer v.c.mu.RUnlock()
	return v.c.v0043Client
}

func (v VersionedClients) V0044() apiv0044.ClientWithResponsesInterface {
	if v.c == nil {
		return nil
	}
	v.c.mu.RLock()
	defer v.c.mu.RUnlock()
	return v.c.v0044Client
}

func (v VersionedClients) V0045() apiv0045.ClientWithResponsesInterface {
	if v.c == nil {
		return nil
	}
	v.c.mu.RLock()
	defer v.c.mu.RUnlock()
	return v.c.v0045Client
}
