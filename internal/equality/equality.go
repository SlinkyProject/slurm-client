// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package equality

import (
	"time"
)

// Semantic can do semantic deep equality checks for objects.
// Example: equality.Semantic.DeepEqual(obj, objWithNonNilButEmptyMaps) == true
var Semantic = EqualitiesOrDie(
	func(a, b time.Time) bool {
		return a.UTC().Equal(b.UTC())
	},
)
