// SPDX-FileCopyrightText: Copyright (C) NVIDIA CORPORATION & AFFILIATES. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package client

import "github.com/SlinkyProject/slurm-client/pkg/object"

// Fetch generations order cache publication by request start, not response arrival.
// All methods here require mu to be held; HTTP requests remain concurrent.
func (i *informerCache) beginFetch() uint64 {
	i.generation++
	if i.inflight == nil {
		i.inflight = make(map[uint64]struct{})
	}
	i.inflight[i.generation] = struct{}{}
	return i.generation
}

func (i *informerCache) observedGeneration(key object.ObjectKey) uint64 {
	generation := max(i.listedGeneration, i.deleted[key])
	if entry := i.cache[key]; entry != nil {
		generation = max(generation, entry.generation)
	}
	return generation
}

func (i *informerCache) publishObject(obj object.Object, generation uint64) {
	key := obj.GetKey()
	if generation < i.observedGeneration(key) {
		return
	}
	i.processObject(obj)
	i.cache[key].generation = generation
	delete(i.deleted, key)
}

func (i *informerCache) endFetch(generation uint64) {
	delete(i.inflight, generation)
	if len(i.deleted) == 0 {
		return
	}
	oldest := i.generation + 1
	for active := range i.inflight {
		oldest = min(oldest, active)
	}
	// A not-found result needs a tombstone only while an older response could
	// resurrect it. A newer full list supplies the same fence for absent keys.
	for key, observed := range i.deleted {
		if observed < oldest || observed <= i.listedGeneration {
			delete(i.deleted, key)
		}
	}
}
