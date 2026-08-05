// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-FileCopyrightText: Copyright 2015 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package cache

// ResourceEventHandler can handle notifications for events that
// happen to a resource. The events are informational only, so you
// can't return an error.  The handlers MUST NOT modify the objects
// received; this concerns not only the top level of structure but all
// the data structures reachable from it.
//   - OnAdd is called when an object is added.
//   - OnUpdate is called when an object is modified. Note that oldObj is the
//     last known state of the object-- it is possible that several changes
//     were combined together, so you can't use this to see every single
//     change. OnUpdate is also called when a re-list happens, and it will
//     get called even if nothing changed. This is useful for periodically
//     evaluating or syncing something.
//   - OnDelete will get the final state of the item if it is known, otherwise
//     it will get an object of type DeletedFinalStateUnknown. This can
//     happen if the watch is closed and misses the delete event and we don't
//     notice the deletion until the subsequent re-list.
type ResourceEventHandler interface {
	OnAdd(obj interface{}, isInInitialList bool)
	OnUpdate(oldObj, newObj interface{})
	OnDelete(obj interface{})
}

// ResourceEventHandlerFuncs is an adaptor to let you easily specify as many or
// as few of the notification functions as you want while still implementing
// ResourceEventHandler.  This adapter does not remove the prohibition against
// modifying the objects.
//
// See ResourceEventHandlerDetailedFuncs if your use needs to propagate
// HasSynced.
type ResourceEventHandlerFuncs struct {
	AddFunc    func(obj interface{})
	UpdateFunc func(oldObj, newObj interface{})
	DeleteFunc func(obj interface{})
}

// OnAdd calls AddFunc if it's not nil.
func (r ResourceEventHandlerFuncs) OnAdd(obj interface{}, isInInitialList bool) {
	if r.AddFunc != nil {
		r.AddFunc(obj)
	}
}

// OnUpdate calls UpdateFunc if it's not nil.
func (r ResourceEventHandlerFuncs) OnUpdate(oldObj, newObj interface{}) {
	if r.UpdateFunc != nil {
		r.UpdateFunc(oldObj, newObj)
	}
}

// OnDelete calls DeleteFunc if it's not nil.
func (r ResourceEventHandlerFuncs) OnDelete(obj interface{}) {
	if r.DeleteFunc != nil {
		r.DeleteFunc(obj)
	}
}

// ResourceEventHandlerDetailedFuncs is exactly like ResourceEventHandlerFuncs
// except its AddFunc accepts the isInInitialList parameter, for propagating
// HasSynced.
type ResourceEventHandlerDetailedFuncs struct {
	AddFunc    func(obj interface{}, isInInitialList bool)
	UpdateFunc func(oldObj, newObj interface{})
	DeleteFunc func(obj interface{})
}

// OnAdd calls AddFunc if it's not nil.
func (r ResourceEventHandlerDetailedFuncs) OnAdd(obj interface{}, isInInitialList bool) {
	if r.AddFunc != nil {
		r.AddFunc(obj, isInInitialList)
	}
}

// OnUpdate calls UpdateFunc if it's not nil.
func (r ResourceEventHandlerDetailedFuncs) OnUpdate(oldObj, newObj interface{}) {
	if r.UpdateFunc != nil {
		r.UpdateFunc(oldObj, newObj)
	}
}

// OnDelete calls DeleteFunc if it's not nil.
func (r ResourceEventHandlerDetailedFuncs) OnDelete(obj interface{}) {
	if r.DeleteFunc != nil {
		r.DeleteFunc(obj)
	}
}
