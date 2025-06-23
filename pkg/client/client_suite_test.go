// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/docker/docker/pkg/stdcopy"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	tcc "github.com/testcontainers/testcontainers-go/modules/compose"
)

var (
	ctx     context.Context
	cancel  context.CancelFunc
	compose tcc.ComposeStack

	restapiServer string
	slurmJwt      string
)

func TestClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Client Suite")
}

var _ = BeforeSuite(func() {
	var err error
	ctx, cancel = context.WithCancel(context.TODO())

	composeFilePaths := []string{"./.testdata/compose.yaml"}
	compose, err = tcc.NewDockerComposeWith(tcc.WithStackFiles(composeFilePaths...), tcc.StackIdentifier("testdata"))
	Expect(err).NotTo(HaveOccurred())

	containerCleanup()

	err = compose.Up(ctx, tcc.Wait(true))
	Expect(err).NotTo(HaveOccurred())

	restapiC, err := compose.ServiceContainer(ctx, "restapi")
	Expect(err).NotTo(HaveOccurred())

	restapiServer, err = restapiC.Endpoint(ctx, "http")
	Expect(err).NotTo(HaveOccurred())
	Expect(restapiServer).NotTo(BeEmpty())

	authcredC, err := compose.ServiceContainer(ctx, "authcred")
	Expect(err).NotTo(HaveOccurred())

	cmd := "scontrol token username=slurm lifespan=infinite"
	rc, reader, err := authcredC.Exec(ctx, strings.Split(cmd, " "))
	Expect(err).NotTo(HaveOccurred())
	stdout, _ := demultiplexReader(reader)
	Expect(rc).To(Equal(0))
	token := strings.TrimPrefix(stdout, "SLURM_JWT=")
	token = strings.TrimSuffix(token, "\n")
	slurmJwt = token
	Expect(slurmJwt).NotTo(BeEmpty())
})

func demultiplexReader(multiplexedReader io.Reader) (string, string) {
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	_, err := stdcopy.StdCopy(stdout, stderr, multiplexedReader)
	Expect(err).NotTo(HaveOccurred())
	return stdout.String(), stderr.String()
}

var _ = AfterSuite(func() {
	containerCleanup()
	cancel()
})

func containerCleanup() {
	err := compose.Down(context.Background(), tcc.RemoveOrphans(true), tcc.RemoveVolumes(true), tcc.RemoveImagesLocal)
	Expect(err).NotTo(HaveOccurred())
}
