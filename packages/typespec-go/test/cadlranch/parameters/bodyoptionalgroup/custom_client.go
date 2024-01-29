//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package bodyoptionalgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewBodyOptionalityClient(options *azcore.ClientOptions) (*BodyOptionalityClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &BodyOptionalityClient{
		internal: internal,
	}, nil
}

func NewOptionalExplicitClient(options *azcore.ClientOptions) (*OptionalExplicitClient, error) {
	internal, err := newClient(options)
	if err != nil {
		return nil, err
	}
	return &OptionalExplicitClient{
		internal: internal,
	}, nil
}

func newClient(options *azcore.ClientOptions) (*azcore.Client, error) {
	return azcore.NewClient("bodyoptionalgroup", "v0.1.0", runtime.PipelineOptions{}, options)
}
