// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	const testTimeout = 30 * time.Second
	var cfg *Config

	BeforeEach(func() {
		cfg = &Config{
			Server:    restapiServer,
			AuthToken: slurmJwt,
		}
	})

	Describe("Client", func() {
		var cl Client

		BeforeEach(func() {
			var err error
			cl, err = NewClient(cfg)
			Expect(err).NotTo(HaveOccurred())
			Expect(cl).NotTo(BeNil())
		})

		Context("New", func() {
			It("should return a new Client", func() {
				cl, err := NewClient(cfg)
				Expect(err).NotTo(HaveOccurred())
				Expect(cl).NotTo(BeNil())
			})
			It("should fail if config is nil", func() {
				cl, err := NewClient(nil)
				Expect(err).To(HaveOccurred())
				Expect(cl).To(BeNil())
			})
			It("should fail if Server is empty", func() {
				cfg2 := &Config{
					AuthToken: cfg.AuthToken,
				}
				cl, err := NewClient(cfg2)
				Expect(err).To(HaveOccurred())
				Expect(cl).To(BeNil())
			})
			It("should fail if AuthToken is empty", func() {
				cfg2 := &Config{
					Server: cfg.Server,
				}
				cl, err := NewClient(cfg2)
				Expect(err).To(HaveOccurred())
				Expect(cl).To(BeNil())
			})
		})

		Context("Stop", func() {
			var cl Client

			BeforeEach(func() {
				var err error
				cl, err = NewClient(cfg)
				Expect(err).NotTo(HaveOccurred())
				Expect(cl).NotTo(BeNil())
			})

			It("should stop by Stop function", func() {
				By("waiting for the client cache to Start then invoking Stop")
				ctx := context.TODO()
				startCh := make(chan bool)
				stopCh := make(chan bool)
				go func() {
					startCh <- true
					cl.Start(ctx)
					stopCh <- true
				}()
				<-startCh
				time.Sleep(2 * time.Second)
				cl.Stop()
				Eventually(stopCh).Within(testTimeout).Should(Receive())
			})
			It("should stop by context cancel", func() {
				By("waiting for the client cache to Start then invoking Stop")
				ctx, cancel := context.WithCancel(context.TODO())
				startCh := make(chan bool)
				stopCh := make(chan bool)
				go func() {
					startCh <- true
					cl.Start(ctx)
					stopCh <- true
				}()
				<-startCh
				time.Sleep(2 * time.Second)
				cancel()
				Eventually(stopCh).Within(testTimeout).Should(Receive())
			})
		})
	})
})
