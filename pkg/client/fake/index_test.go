// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"testing"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestFakeInformerIndex(t *testing.T) {
	job := &api.V0044JobInfo{JobId: ptr.To(int32(1)), AdminComment: ptr.To("a")}
	c := NewClientBuilder().WithObjects(job).Build()
	i := c.GetInformer(job.GetType())
	extract := func(o object.Object) []string { return []string{*o.(*api.V0044JobInfo).AdminComment} }
	if err := i.AddIndex("owner", extract); err != nil {
		t.Fatal(err)
	}
	if c.GetInformer(job.GetType()) != i {
		t.Fatal("GetInformer did not preserve registered index")
	}
	if err := i.AddIndex("owner", extract); err == nil {
		t.Fatal("duplicate accepted")
	}
	if err := i.AddIndex("", extract); err == nil {
		t.Fatal("empty name accepted")
	}
	if err := i.AddIndex("nil", nil); err == nil {
		t.Fatal("nil extractor accepted")
	}
	if err := i.ByIndex(t.Context(), "unknown", "a", &api.V0044JobInfoObjectList{}); err == nil {
		t.Fatal("unknown index accepted")
	}
	if err := i.ByIndex(t.Context(), "owner", "a", &api.V0044NodeObjectList{}); err == nil {
		t.Fatal("wrong type accepted")
	}
	check := func(value string, want int) {
		t.Helper()
		jobs := &api.V0044JobInfoObjectList{}
		if err := i.ByIndex(t.Context(), "owner", value, jobs); err != nil {
			t.Fatal(err)
		}
		if len(jobs.Items) != want {
			t.Fatalf("%q: got %d want %d", value, len(jobs.Items), want)
		}
		if want > 0 {
			*jobs.Items[0].AdminComment = "mutated result"
		}
	}
	check("a", 1)
	check("a", 1)
	*job.AdminComment = "b"
	if err := c.Update(t.Context(), job, nil); err != nil {
		t.Fatal(err)
	}
	check("a", 0)
	check("b", 1)
	if err := c.Delete(t.Context(), job); err != nil {
		t.Fatal(err)
	}
	check("b", 0)
}
