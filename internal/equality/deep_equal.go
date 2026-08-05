// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-FileCopyrightText: Copyright 2015 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package equality

import (
	"maps"

	"github.com/SlinkyProject/slurm-client/third_party/forked/golang/reflect"
)

// The code for this type must be located in third_party, since it forks from
// go std lib. But for convenience, we expose the type here, too.
type Equalities struct {
	reflect.Equalities
}

// For convenience, panics on errors
func EqualitiesOrDie(funcs ...any) Equalities {
	e := Equalities{reflect.Equalities{}}
	if err := e.AddFuncs(funcs...); err != nil {
		panic(err)
	}
	return e
}

// Performs a shallow copy of the equalities map
func (e Equalities) Copy() Equalities {
	result := Equalities{reflect.Equalities{}}

	maps.Copy(result.Equalities, e.Equalities)

	return result
}
