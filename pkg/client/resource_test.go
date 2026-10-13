// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"k8s.io/utils/ptr"

	apiv0042 "github.com/SlinkyProject/slurm-client/api/v0042"
	apiv0045 "github.com/SlinkyProject/slurm-client/api/v0045"
	apierrors "github.com/SlinkyProject/slurm-client/pkg/errors"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestResourceCatalog(t *testing.T) {
	require.NotEmpty(t, resourceCatalog)
	for typ, r := range resourceCatalog {
		require.NotNil(t, r.newObject)
		require.Equal(t, typ, r.newObject().GetType())
		uncached := strings.HasSuffix(string(typ), "NodeResourceLayoutList")
		require.Equal(t, !uncached, r.cacheable)
		if uncached {
			require.Nil(t, r.newList)
			continue
		}
		require.NotNil(t, r.newList)
		require.Equal(t, typ, r.newList().GetType())
	}

	uncached := uncachedObjects()
	require.NotEmpty(t, uncached)
	for _, obj := range uncached {
		require.False(t, resourceCatalog[obj.GetType()].cacheable)
	}

	bound := cloneCatalog()
	bindV0042(bound, nil)
	bindV0043(bound, nil)
	bindV0044(bound, nil)
	bindV0045(bound, nil)
	for typ, r := range bound {
		require.NotNil(t, r.get)
		if strings.HasSuffix(string(typ), "NodeResourceLayoutList") {
			require.Nil(t, r.list)
			continue
		}
		require.NotNil(t, r.list)
	}
	require.Nil(t, bound[apiv0042.ObjectTypeV0042Node].create)
	require.NotNil(t, bound[apiv0042.ObjectTypeV0042JobInfo].create)
	require.NotNil(t, bound[apiv0045.ObjectTypeV0045Node].create)
	require.NotNil(t, bound[apiv0045.ObjectTypeV0045ReservationInfo].create)
}

func TestCopyObject(t *testing.T) {
	src := &apiv0045.V0045Node{Name: ptr.To("n1")}
	dst := &apiv0045.V0045Node{}
	require.NoError(t, copyObject(dst, src))
	require.Equal(t, "n1", ptr.Deref(dst.Name, ""))
	err := copyObject(&apiv0045.V0045JobInfo{}, src)
	require.EqualError(t, err, "cannot copy object of type *v0045.V0045Node into *v0045.V0045JobInfo")
	require.NotErrorIs(t, err, apierrors.ErrNotImplemented)
}

func TestClientResourceErrors(t *testing.T) {
	obj := &apiv0045.V0045Node{}
	list := &apiv0045.V0045NodeObjectList{}
	for _, tc := range []struct {
		operation string
		call      func(*client) error
	}{
		{"create", func(c *client) error { return c.Create(t.Context(), obj, nil) }},
		{"delete", func(c *client) error { return c.Delete(t.Context(), obj) }},
		{"update", func(c *client) error { return c.Update(t.Context(), obj, nil) }},
		{"get", func(c *client) error { return c.Get(t.Context(), "n1", obj, &GetOptions{SkipCache: true}) }},
		{"list", func(c *client) error { return c.List(t.Context(), list, &ListOptions{SkipCache: true}) }},
	} {
		t.Run(tc.operation, func(t *testing.T) {
			c := &client{resources: map[object.ObjectType]resource{}}
			require.ErrorIs(t, tc.call(c), apierrors.ErrNotImplemented)

			c.resources[obj.GetType()] = resource{}
			require.ErrorIs(t, tc.call(c), apierrors.ErrNotImplemented)
		})
	}
}

func TestGetIntoTypeMismatch(t *testing.T) {
	fetched := false
	err := getInto(&apiv0045.V0045JobInfo{}, func() (*apiv0045.V0045Node, error) {
		fetched = true
		return &apiv0045.V0045Node{}, nil
	})
	require.EqualError(t, err, "expected object of type *v0045.V0045Node, got *v0045.V0045JobInfo")
	require.NotErrorIs(t, err, apierrors.ErrNotImplemented)
	require.False(t, fetched, "invalid destination must be rejected before fetching")
}

func TestListIntoTypeMismatch(t *testing.T) {
	fetched := false
	err := listInto(&apiv0045.V0045JobInfoObjectList{}, func() (*apiv0045.V0045NodeObjectList, error) {
		fetched = true
		return &apiv0045.V0045NodeObjectList{}, nil
	})
	require.EqualError(t, err, "expected list of type *v0045.V0045NodeObjectList, got *v0045.V0045JobInfoObjectList")
	require.NotErrorIs(t, err, apierrors.ErrNotImplemented)
	require.False(t, fetched, "invalid destination must be rejected before fetching")
}
