package durationgroupapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"tests/generated/durationgroup"
)

// DurationClientAPI contains the set of methods on the DurationClient type.
type DurationClientAPI interface {
	GetInvalid(ctx context.Context) (result durationgroup.TimeSpan, err error)
	GetNull(ctx context.Context) (result durationgroup.TimeSpan, err error)
	GetPositiveDuration(ctx context.Context) (result durationgroup.TimeSpan, err error)
	PutPositiveDuration(ctx context.Context, durationBody string) (result autorest.Response, err error)
}

var _ DurationClientAPI = (*durationgroup.DurationClient)(nil)