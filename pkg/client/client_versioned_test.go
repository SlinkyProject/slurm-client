// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"net/http"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/SlinkyProject/slurm-client/pkg/client/token"
)

var _ = Describe("VersionedClients", func() {
	const testTimeout = 30 * time.Second
	var cl Client

	BeforeEach(func() {
		var err error
		cl, err = NewClient(&Config{
			Server:        restapiServer,
			TokenProvider: token.StaticProvider(slurmJwt),
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(cl).NotTo(BeNil())
	})

	It("should ping the cluster through the v0042 client", func(ctx SpecContext) {
		versionedClient := cl.Versioned().V0042()
		Expect(versionedClient).NotTo(BeNil())

		res, err := versionedClient.SlurmV0042GetPingWithResponse(ctx)
		Expect(err).NotTo(HaveOccurred())
		Expect(res.StatusCode()).To(Equal(http.StatusOK))
		Expect(res.JSON200).NotTo(BeNil())
		Expect(res.JSON200.Pings).To(ContainElement(HaveField("Hostname", new("slurmctld"))))
	}, SpecTimeout(testTimeout))

	It("should ping the cluster through the v0043 client", func(ctx SpecContext) {
		versionedClient := cl.Versioned().V0043()
		Expect(versionedClient).NotTo(BeNil())

		res, err := versionedClient.SlurmV0043GetPingWithResponse(ctx)
		Expect(err).NotTo(HaveOccurred())
		Expect(res.StatusCode()).To(Equal(http.StatusOK))
		Expect(res.JSON200).NotTo(BeNil())
		Expect(res.JSON200.Pings).To(ContainElement(HaveField("Hostname", new("slurmctld"))))
	}, SpecTimeout(testTimeout))

	It("should ping the cluster through the v0044 client", func(ctx SpecContext) {
		versionedClient := cl.Versioned().V0044()
		Expect(versionedClient).NotTo(BeNil())

		res, err := versionedClient.SlurmV0044GetPingWithResponse(ctx)
		Expect(err).NotTo(HaveOccurred())
		Expect(res.StatusCode()).To(Equal(http.StatusOK))
		Expect(res.JSON200).NotTo(BeNil())
		Expect(res.JSON200.Pings).To(ContainElement(HaveField("Hostname", new("slurmctld"))))
	}, SpecTimeout(testTimeout))

	It("should ping the cluster through the v0045 client", func(ctx SpecContext) {
		versionedClient := cl.Versioned().V0045()
		Expect(versionedClient).NotTo(BeNil())

		res, err := versionedClient.SlurmV0045GetPingWithResponse(ctx)
		Expect(err).NotTo(HaveOccurred())
		Expect(res.StatusCode()).To(Equal(http.StatusOK))
		Expect(res.JSON200).NotTo(BeNil())
		Expect(res.JSON200.Pings).To(ContainElement(HaveField("Hostname", new("slurmctld"))))
	}, SpecTimeout(testTimeout))
})
