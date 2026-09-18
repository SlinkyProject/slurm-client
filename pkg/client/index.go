// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"fmt"

	"github.com/SlinkyProject/slurm-client/pkg/object"
)

// secondaryIndex is protected by the informer's cache mutex. Recording the keys
// per object avoids re-running extractors when removing an old object version.
type secondaryIndex struct {
	extract IndexFunc
	values  map[string]map[object.ObjectKey]struct{}
	keys    map[object.ObjectKey][]string
}

func (i *informerCache) AddIndex(name string, extract IndexFunc) error {
	i.mu.Lock()
	defer i.mu.Unlock()
	if name == "" || extract == nil {
		return fmt.Errorf("index name and extractor must be set")
	}
	if _, exists := i.indexes[name]; exists {
		return fmt.Errorf("index %q already exists", name)
	}
	index := &secondaryIndex{
		extract: extract,
		values:  make(map[string]map[object.ObjectKey]struct{}),
		keys:    make(map[object.ObjectKey][]string),
	}
	for _, entry := range i.cache {
		if entry != nil && entry.object != nil {
			index.update(entry.object)
		}
	}
	if i.indexes == nil {
		i.indexes = make(map[string]*secondaryIndex)
	}
	i.indexes[name] = index
	return nil
}

func (i *informerCache) ByIndex(ctx context.Context, name, value string, list object.ObjectList) error {
	if list.GetType() != i.objectType {
		return fmt.Errorf("index on %s cannot populate %s", i.objectType, list.GetType())
	}
	i.mu.RLock()
	_, exists := i.indexes[name]
	i.mu.RUnlock()
	if !exists {
		return fmt.Errorf("index %q is not registered", name)
	}
	if err := i.waitForSyncList(ctx); err != nil {
		return fmt.Errorf("failed to wait on type %s cache sync: %w", i.objectType, err)
	}
	i.mu.RLock()
	defer i.mu.RUnlock()
	for key := range i.indexes[name].values[value] {
		list.AppendItem(i.cache[key].object.DeepCopyObject().(object.Object))
	}
	return nil
}

func (i *secondaryIndex) remove(key object.ObjectKey) {
	for _, value := range i.keys[key] {
		delete(i.values[value], key)
		if len(i.values[value]) == 0 {
			delete(i.values, value)
		}
	}
	delete(i.keys, key)
}

func (i *secondaryIndex) update(obj object.Object) {
	key := obj.GetKey()
	i.remove(key)
	for _, value := range i.extract(obj) {
		if i.values[value] == nil {
			i.values[value] = make(map[object.ObjectKey]struct{})
		}
		if _, exists := i.values[value][key]; exists {
			continue
		}
		i.values[value][key] = struct{}{}
		i.keys[key] = append(i.keys[key], value)
	}
}
