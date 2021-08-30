//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package morecustombaseurigroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ConnectionOptions contains configuration settings for the connection's pipeline.
// All zero-value fields will be initialized with their default values.
type ConnectionOptions struct {
	// HTTPClient sets the transport for making HTTP requests.
	HTTPClient policy.Transporter
	// Retry configures the built-in retry policy behavior.
	Retry policy.RetryOptions
	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry policy.TelemetryOptions
	// Logging configures the built-in logging policy behavior.
	Logging policy.LogOptions
	// PerCallPolicies contains custom policies to inject into the pipeline.
	// Each policy is executed once per request.
	PerCallPolicies []policy.Policy
	// PerRetryPolicies contains custom policies to inject into the pipeline.
	// Each policy is executed once per request, and for each retry request.
	PerRetryPolicies []policy.Policy
}

// Connection - Test Infrastructure for AutoRest
type Connection struct {
	dnsSuffix string
	p         runtime.Pipeline
}

// NewConnection creates an instance of the Connection type with the specified endpoint.
// Pass nil to accept the default options; this is the same as passing a zero-value options.
func NewConnection(dnsSuffix *string, options *ConnectionOptions) *Connection {
	if options == nil {
		options = &ConnectionOptions{}
	}
	policies := []policy.Policy{}
	if !options.Telemetry.Disabled {
		policies = append(policies, runtime.NewTelemetryPolicy(module, version, &options.Telemetry))
	}
	policies = append(policies, options.PerCallPolicies...)
	policies = append(policies, runtime.NewRetryPolicy(&options.Retry))
	policies = append(policies, options.PerRetryPolicies...)
	policies = append(policies, runtime.NewLogPolicy(&options.Logging))
	client := &Connection{
		p:         runtime.NewPipeline(options.HTTPClient, policies...),
		dnsSuffix: "host",
	}
	if dnsSuffix != nil {
		client.dnsSuffix = *dnsSuffix
	}
	return client
}

// DnsSuffix returns part of the parameterized host.
func (c *Connection) DnsSuffix() string {
	return c.dnsSuffix
}

// Pipeline returns the connection's pipeline.
func (c *Connection) Pipeline() runtime.Pipeline {
	return c.p
}
