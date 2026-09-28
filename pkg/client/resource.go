// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"fmt"
	"maps"
	"slices"
	"strings"

	"k8s.io/utils/ptr"

	apierrors "github.com/SlinkyProject/slurm-client/pkg/errors"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

// resource is one Slurm object the versioned clients already implement.
// A nil operation means that call is unsupported.
type resource struct {
	newObject func() object.Object
	newList   func() object.ObjectList
	cacheable bool
	get       func(context.Context, object.ObjectKey, object.Object) error
	list      func(context.Context, object.ObjectList) error
	create    func(context.Context, any) (object.ObjectKey, error)
	update    func(context.Context, string, any) error
	delete    func(context.Context, string) error
}

// resourceCatalog holds constructors and cacheability. Operation closures are
// filled per client in createApiClients so two clients do not share them.
var resourceCatalog = map[object.ObjectType]resource{}

func addResource(typ object.ObjectType, r resource) {
	if _, ok := resourceCatalog[typ]; ok {
		panic("duplicate resource " + typ)
	}
	resourceCatalog[typ] = r
}

func cloneCatalog() map[object.ObjectType]resource {
	out := make(map[object.ObjectType]resource, len(resourceCatalog))
	maps.Copy(out, resourceCatalog)
	return out
}

func setOps(dst map[object.ObjectType]resource, typ object.ObjectType, ops resource) {
	base, ok := dst[typ]
	if !ok {
		panic("resource " + typ + " is not registered")
	}
	base.get = ops.get
	base.list = ops.list
	base.create = ops.create
	base.update = ops.update
	base.delete = ops.delete
	dst[typ] = base
}

func uncachedObjects() []object.Object {
	out := make([]object.Object, 0)
	for _, r := range resourceCatalog {
		if !r.cacheable {
			out = append(out, r.newObject())
		}
	}
	slices.SortFunc(out, func(a, b object.Object) int {
		return strings.Compare(string(a.GetType()), string(b.GetType()))
	})
	return out
}

func getInto[T any](obj object.Object, fetch func() (*T, error)) error {
	dst, ok := any(obj).(*T)
	if !ok {
		return apierrors.ErrNotImplemented
	}
	out, err := fetch()
	if err != nil {
		return err
	}
	*dst = *out
	return nil
}

func listInto[T any](list object.ObjectList, fetch func() (*T, error)) error {
	dst, ok := any(list).(*T)
	if !ok {
		return apierrors.ErrNotImplemented
	}
	out, err := fetch()
	if err != nil {
		return err
	}
	*dst = *out
	return nil
}

func jobKey(id *int32, err error) (object.ObjectKey, error) {
	if err != nil {
		return "", err
	}
	return object.ObjectKey(fmt.Sprintf("%d", ptr.Deref(id, 0))), nil
}

func nodeKey(name *string, err error) (object.ObjectKey, error) {
	if err != nil {
		return "", err
	}
	return object.ObjectKey(ptr.Deref(name, "")), nil
}

func nameKey(name string, err error) (object.ObjectKey, error) {
	if err != nil {
		return "", err
	}
	return object.ObjectKey(name), nil
}
