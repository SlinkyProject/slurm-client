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

	v0042 "github.com/SlinkyProject/slurm-client/pkg/client/api/v0042"
	v0043 "github.com/SlinkyProject/slurm-client/pkg/client/api/v0043"
	v0044 "github.com/SlinkyProject/slurm-client/pkg/client/api/v0044"
	v0045 "github.com/SlinkyProject/slurm-client/pkg/client/api/v0045"
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

	v0042Client v0042.ClientInterface
	v0043Client v0043.ClientInterface
	v0044Client v0044.ClientInterface
	v0045Client v0045.ClientInterface

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
			&types.V0045NodeResourceLayout{},
			&types.V0045Reconfigure{},
			&types.V0044NodeResourceLayout{},
			&types.V0044Reconfigure{},
			&types.V0043Reconfigure{},
			&types.V0042Reconfigure{},
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

	c.v0045Client, err = v0045.NewSlurmClient(c.config.Server, c.config.AuthToken, c.config.HTTPClient)
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

	var err error
	var key object.ObjectKey
	switch obj.(type) {
	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0042JobInfo:
		var jobId *int32
		jobId, err = c.v0042Client.CreateJobInfo(ctx, req)
		key = object.ObjectKey(fmt.Sprintf("%d", ptr.Deref(jobId, 0)))

	case *types.V0042Account:
		var name string
		name, err = c.v0042Client.CreateAccount(ctx, req)
		key = object.ObjectKey(name)

	case *types.V0042User:
		var name string
		name, err = c.v0042Client.CreateUser(ctx, req)
		key = object.ObjectKey(name)

	case *types.V0042Assoc:
		var assocKey string
		assocKey, err = c.v0042Client.CreateAssoc(ctx, req)
		key = object.ObjectKey(assocKey)

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0043JobInfo:
		var jobId *int32
		jobId, err = c.v0043Client.CreateJobInfo(ctx, req)
		key = object.ObjectKey(fmt.Sprintf("%d", ptr.Deref(jobId, 0)))

	case *types.V0043Account:
		var name string
		name, err = c.v0043Client.CreateAccount(ctx, req)
		key = object.ObjectKey(name)

	case *types.V0043User:
		var name string
		name, err = c.v0043Client.CreateUser(ctx, req)
		key = object.ObjectKey(name)

	case *types.V0043Assoc:
		var assocKey string
		assocKey, err = c.v0043Client.CreateAssoc(ctx, req)
		key = object.ObjectKey(assocKey)

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0044JobInfo:
		var jobId *int32
		jobId, err = c.v0044Client.CreateJobInfo(ctx, req)
		key = object.ObjectKey(fmt.Sprintf("%d", ptr.Deref(jobId, 0)))

	case *types.V0044ReservationInfo:
		var reservationName string
		reservationName, err = c.v0044Client.CreateReservationInfo(ctx, req)
		key = object.ObjectKey(reservationName)

	case *types.V0044Node:
		var nodeName *string
		nodeName, err = c.v0044Client.CreateNewNode(ctx, req)
		key = object.ObjectKey(ptr.Deref(nodeName, ""))

	case *types.V0044Account:
		var name string
		name, err = c.v0044Client.CreateAccount(ctx, req)
		key = object.ObjectKey(name)

	case *types.V0044User:
		var name string
		name, err = c.v0044Client.CreateUser(ctx, req)
		key = object.ObjectKey(name)

	case *types.V0044Assoc:
		var assocKey string
		assocKey, err = c.v0044Client.CreateAssoc(ctx, req)
		key = object.ObjectKey(assocKey)

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0045JobInfo:
		var jobId *int32
		jobId, err = c.v0045Client.CreateJobInfo(ctx, req)
		key = object.ObjectKey(fmt.Sprintf("%d", ptr.Deref(jobId, 0)))

	case *types.V0045ReservationInfo:
		var reservationName string
		reservationName, err = c.v0045Client.CreateReservationInfo(ctx, req)
		key = object.ObjectKey(reservationName)

	case *types.V0045Node:
		var nodeName *string
		nodeName, err = c.v0045Client.CreateNewNode(ctx, req)
		key = object.ObjectKey(ptr.Deref(nodeName, ""))

	case *types.V0045Account:
		var name string
		name, err = c.v0045Client.CreateAccount(ctx, req)
		key = object.ObjectKey(name)

	case *types.V0045User:
		var name string
		name, err = c.v0045Client.CreateUser(ctx, req)
		key = object.ObjectKey(name)

	case *types.V0045Assoc:
		var assocKey string
		assocKey, err = c.v0045Client.CreateAssoc(ctx, req)
		key = object.ObjectKey(assocKey)

	/////////////////////////////////////////////////////////////////////////////////

	default:
		return errors.New(http.StatusText(http.StatusNotImplemented))
	}

	if err != nil {
		return err
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

	var err error
	key := string(obj.GetKey())
	switch o := obj.(type) {
	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0042JobInfo:
		err = c.v0042Client.DeleteJobInfo(ctx, key)
	case *types.V0042Node:
		err = c.v0042Client.DeleteNode(ctx, key)
	case *types.V0042Account:
		err = c.v0042Client.DeleteAccount(ctx, key)
	case *types.V0042User:
		err = c.v0042Client.DeleteUser(ctx, key)
	case *types.V0042Assoc:
		err = c.v0042Client.DeleteAssoc(ctx, o.V0042Assoc)

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0043JobInfo:
		err = c.v0043Client.DeleteJobInfo(ctx, key)
	case *types.V0043Node:
		err = c.v0043Client.DeleteNode(ctx, key)
	case *types.V0043Account:
		err = c.v0043Client.DeleteAccount(ctx, key)
	case *types.V0043User:
		err = c.v0043Client.DeleteUser(ctx, key)
	case *types.V0043Assoc:
		err = c.v0043Client.DeleteAssoc(ctx, o.V0043Assoc)

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0044JobInfo:
		err = c.v0044Client.DeleteJobInfo(ctx, key)
	case *types.V0044Node:
		err = c.v0044Client.DeleteNode(ctx, key)
	case *types.V0044ReservationInfo:
		err = c.v0044Client.DeleteReservationInfo(ctx, key)
	case *types.V0044Account:
		err = c.v0044Client.DeleteAccount(ctx, key)
	case *types.V0044User:
		err = c.v0044Client.DeleteUser(ctx, key)
	case *types.V0044Assoc:
		err = c.v0044Client.DeleteAssoc(ctx, o.V0044Assoc)

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0045JobInfo:
		err = c.v0045Client.DeleteJobInfo(ctx, key)
	case *types.V0045Node:
		err = c.v0045Client.DeleteNode(ctx, key)
	case *types.V0045ReservationInfo:
		err = c.v0045Client.DeleteReservationInfo(ctx, key)
	case *types.V0045Account:
		err = c.v0045Client.DeleteAccount(ctx, key)
	case *types.V0045User:
		err = c.v0045Client.DeleteUser(ctx, key)
	case *types.V0045Assoc:
		err = c.v0045Client.DeleteAssoc(ctx, o.V0045Assoc)

	/////////////////////////////////////////////////////////////////////////////////

	default:
		return errors.New(http.StatusText(http.StatusNotImplemented))
	}

	if err != nil {
		return err
	}

	err = c.Get(ctx, obj.GetKey(), obj, &GetOptions{RefreshCache: true})
	if err != nil {
		// We expect the error to always be NotFound because we deleted the
		// object from Slurm then attempted to Get the deleted object with
		// refreshed cache.
		if err.Error() != http.StatusText(http.StatusNotFound) {
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

	var err error
	key := string(obj.GetKey())
	switch obj.(type) {
	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0042JobInfo:
		err = c.v0042Client.UpdateJobInfo(ctx, key, req)
	case *types.V0042Node:
		err = c.v0042Client.UpdateNode(ctx, key, req)
	case *types.V0042Account:
		err = c.v0042Client.UpdateAccount(ctx, key, req)
	case *types.V0042User:
		err = c.v0042Client.UpdateUser(ctx, key, req)
	case *types.V0042Assoc:
		err = c.v0042Client.UpdateAssoc(ctx, key, req)

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0043JobInfo:
		err = c.v0043Client.UpdateJobInfo(ctx, key, req)
	case *types.V0043Node:
		err = c.v0043Client.UpdateNode(ctx, key, req)
	case *types.V0043Account:
		err = c.v0043Client.UpdateAccount(ctx, key, req)
	case *types.V0043User:
		err = c.v0043Client.UpdateUser(ctx, key, req)
	case *types.V0043Assoc:
		err = c.v0043Client.UpdateAssoc(ctx, key, req)

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0044JobInfo:
		err = c.v0044Client.UpdateJobInfo(ctx, key, req)
	case *types.V0044Node:
		err = c.v0044Client.UpdateNode(ctx, key, req)
	case *types.V0044ReservationInfo:
		err = c.v0044Client.UpdateReservationInfo(ctx, key, req)
	case *types.V0044Account:
		err = c.v0044Client.UpdateAccount(ctx, key, req)
	case *types.V0044User:
		err = c.v0044Client.UpdateUser(ctx, key, req)
	case *types.V0044Assoc:
		err = c.v0044Client.UpdateAssoc(ctx, key, req)

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0045JobInfo:
		err = c.v0045Client.UpdateJobInfo(ctx, key, req)
	case *types.V0045Node:
		err = c.v0045Client.UpdateNode(ctx, key, req)
	case *types.V0045ReservationInfo:
		err = c.v0045Client.UpdateReservationInfo(ctx, key, req)
	case *types.V0045Account:
		err = c.v0045Client.UpdateAccount(ctx, key, req)
	case *types.V0045User:
		err = c.v0045Client.UpdateUser(ctx, key, req)
	case *types.V0045Assoc:
		err = c.v0045Client.UpdateAssoc(ctx, key, req)

	/////////////////////////////////////////////////////////////////////////////////

	default:
		return errors.New(http.StatusText(http.StatusNotImplemented))
	}

	if err != nil {
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
		objectType = object.ObjectType(strings.TrimSuffix(string(objectType), "List"))
		informerCache := c.GetInformer(objectType)
		if informerCache.HasStarted() && !c.uncached.Has(objectType) {
			return informerCache.Get(ctx, key, obj, opts...)
		}
	}

	switch o := obj.(type) {
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
	case *types.V0042Account:
		out, err := c.v0042Client.GetAccount(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0042User:
		out, err := c.v0042Client.GetUser(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0042Assoc:
		out, err := c.v0042Client.GetAssoc(ctx, string(key))
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
	case *types.V0043Account:
		out, err := c.v0043Client.GetAccount(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0043User:
		out, err := c.v0043Client.GetUser(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0043Assoc:
		out, err := c.v0043Client.GetAssoc(ctx, string(key))
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
	case *types.V0044Account:
		out, err := c.v0044Client.GetAccount(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0044User:
		out, err := c.v0044Client.GetUser(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0044Assoc:
		out, err := c.v0044Client.GetAssoc(ctx, string(key))
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

	case *types.V0045ControllerPing:
		out, err := c.v0045Client.GetControllerPing(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0045JobInfo:
		out, err := c.v0045Client.GetJobInfo(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0045NodeResourceLayout:
		out, err := c.v0045Client.GetNodeResourceLayout(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0045Node:
		out, err := c.v0045Client.GetNode(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0045PartitionInfo:
		out, err := c.v0045Client.GetPartitionInfo(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0045Reconfigure:
		out, err := c.v0045Client.GetReconfigure(ctx)
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0045ReservationInfo:
		out, err := c.v0045Client.GetReservationInfo(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0045Account:
		out, err := c.v0045Client.GetAccount(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0045User:
		out, err := c.v0045Client.GetUser(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0045Assoc:
		out, err := c.v0045Client.GetAssoc(ctx, string(key))
		if err != nil {
			return err
		}
		*o = *out
	case *types.V0045Stats:
		out, err := c.v0045Client.GetStats(ctx)
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
	case *types.V0042AccountList:
		out, err := c.v0042Client.ListAccount(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0042UserList:
		out, err := c.v0042Client.ListUser(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0042AssocList:
		out, err := c.v0042Client.ListAssoc(ctx)
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
	case *types.V0043AccountList:
		out, err := c.v0043Client.ListAccount(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0043UserList:
		out, err := c.v0043Client.ListUser(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0043AssocList:
		out, err := c.v0043Client.ListAssoc(ctx)
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
	case *types.V0044AccountList:
		out, err := c.v0044Client.ListAccount(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0044UserList:
		out, err := c.v0044Client.ListUser(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0044AssocList:
		out, err := c.v0044Client.ListAssoc(ctx)
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

	case *types.V0045ControllerPingList:
		out, err := c.v0045Client.ListControllerPing(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0045JobInfoList:
		out, err := c.v0045Client.ListJobInfo(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0045NodeList:
		out, err := c.v0045Client.ListNodes(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0045PartitionInfoList:
		out, err := c.v0045Client.ListPartitionInfo(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0045ReconfigureList:
		out, err := c.v0045Client.ListReconfigure(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0045ReservationInfoList:
		out, err := c.v0045Client.ListReservationInfo(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0045AccountList:
		out, err := c.v0045Client.ListAccount(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0045UserList:
		out, err := c.v0045Client.ListUser(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0045AssocList:
		out, err := c.v0045Client.ListAssoc(ctx)
		if err != nil {
			return err
		}
		*objList = *out
	case *types.V0045StatsList:
		out, err := c.v0045Client.ListStats(ctx)
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
	c.informers[objectType] = newInformer(objectType, c, c.cacheSyncPeriod)
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
