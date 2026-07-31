// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"os"
	"path/filepath"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/SlinkyProject/slurm-client/pkg/client/token"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

var _ = Describe("Token provider", func() {
	const testTimeout = 30 * time.Second

	It("reads a rotated token file for every request", func(specCtx SpecContext) {
		tokenPath := filepath.Join(GinkgoT().TempDir(), "token")
		replaceTokenFile(tokenPath, slurmJwt)

		slurmClient, err := NewClient(&Config{
			Server:        restapiServer,
			TokenProvider: token.FileProvider{Path: tokenPath},
		})
		Expect(err).NotTo(HaveOccurred())

		list := &types.V0045ControllerPingList{}
		Expect(slurmClient.List(specCtx, list, &ListOptions{SkipCache: true})).To(Succeed())
		Expect(list.Items).NotTo(BeEmpty())

		replaceTokenFile(tokenPath, "")
		Expect(slurmClient.List(specCtx, list, &ListOptions{SkipCache: true})).To(MatchError(ContainSubstring("token file")))

		rotatedToken := mintSlurmToken("3600")
		Expect(rotatedToken).NotTo(Equal(slurmJwt))
		replaceTokenFile(tokenPath, rotatedToken)

		Expect(slurmClient.List(specCtx, list, &ListOptions{SkipCache: true})).To(Succeed())
		Expect(list.Items).NotTo(BeEmpty())
		Expect(slurmClient.GetToken()).To(Equal(rotatedToken))
	}, SpecTimeout(testTimeout))
})

func replaceTokenFile(path, token string) {
	GinkgoHelper()
	nextPath := path + ".next"
	Expect(os.WriteFile(nextPath, []byte(token+"\n"), 0o600)).To(Succeed())
	Expect(os.Rename(nextPath, path)).To(Succeed())
}
