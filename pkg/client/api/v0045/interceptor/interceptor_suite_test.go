// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package interceptor

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestInterceptor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Interceptor Suite")
}
