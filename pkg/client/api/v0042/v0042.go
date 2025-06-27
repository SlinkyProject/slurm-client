// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0042

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"k8s.io/utils/ptr"
)

const (
	headerSlurmUserName  = "X-SLURM-USER-NAME"
	headerSlurmUserToken = "X-SLURM-USER-TOKEN" //nolint:gosec // disable G101

	headerContentType     = "Content-Type"
	headerApplicationJson = "application/json"
)

type ClientInterface interface {
	api.ClientWithResponsesInterface
	ControllerPingInfoInterface
	JobInfoInterface
	NodeInterface
	PartitionInterface
	ReconfigureInterface
	StatsInterface
}

type SlurmClient struct {
	api.ClientWithResponsesInterface
}

var _ ClientInterface = &SlurmClient{}

func NewSlurmClient(server, token string, httpServer *http.Client) (ClientInterface, error) {
	httpClient := http.DefaultClient
	if httpServer != nil {
		httpClient = httpServer
	}

	// Create header injection function
	headerFunc := func(ctx context.Context, req *http.Request) error {
		req.Header.Add(headerSlurmUserToken, token)
		req.Header.Add(headerContentType, headerApplicationJson)
		return nil
	}

	// Create wrapper client
	client, err := api.NewClientWithResponses(
		server,
		api.WithHTTPClient(httpClient),
		api.WithRequestEditorFn(headerFunc),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to create client: %w", err)
	}

	return &SlurmClient{client}, nil
}

func getOpenapiErrors(oapierrors *api.V0042OpenapiErrors) []error {
	errs := []error{}
	for _, err := range ptr.Deref(oapierrors, []api.V0042OpenapiError{}) {
		if err.Error != nil {
			errs = append(errs, errors.New(*err.Error))
		}
	}
	return errs
}
