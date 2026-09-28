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
	require.ErrorIs(t, copyObject(&apiv0045.V0045JobInfo{}, src), apierrors.ErrNotImplemented)
}
