// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package srvdrivennewgroup_test

import (
	"context"
	"srvdrivennewgroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestResiliencyServiceDrivenClient_AddOperation(t *testing.T) {
	client, err := srvdrivennewgroup.NewResiliencyServiceDrivenClient(nil)
	require.NoError(t, err)
	resp, err := client.AddOperation(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestResiliencyServiceDrivenClient_FromNone(t *testing.T) {
	client, err := srvdrivennewgroup.NewResiliencyServiceDrivenClient(nil)
	require.NoError(t, err)
	resp, err := client.FromNone(context.Background(), &srvdrivennewgroup.ResiliencyServiceDrivenClientFromNoneOptions{
		NewParameter: to.Ptr("new"),
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestResiliencyServiceDrivenClient_FromOneOptional(t *testing.T) {
	client, err := srvdrivennewgroup.NewResiliencyServiceDrivenClient(nil)
	require.NoError(t, err)
	resp, err := client.FromOneOptional(context.Background(), &srvdrivennewgroup.ResiliencyServiceDrivenClientFromOneOptionalOptions{
		NewParameter: to.Ptr("new"),
		Parameter:    to.Ptr("optional"),
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestResiliencyServiceDrivenClient_FromOneRequired(t *testing.T) {
	client, err := srvdrivennewgroup.NewResiliencyServiceDrivenClient(nil)
	require.NoError(t, err)
	resp, err := client.FromOneRequired(context.Background(), "required", &srvdrivennewgroup.ResiliencyServiceDrivenClientFromOneRequiredOptions{
		NewParameter: to.Ptr("new"),
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}