// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

const scope = "https://dev.azuresynapse.net/.default"
const telemetryInfo = "azsdk-go-azartifacts/<version>"

// ClientOptions contains configuration settings for the default client's pipeline.
type ClientOptions struct {
	// HTTPClient sets the transport for making HTTP requests.
	HTTPClient azcore.Transport
	// LogOptions configures the built-in request logging policy behavior.
	LogOptions azcore.RequestLogOptions
	// Retry configures the built-in retry policy behavior.
	Retry azcore.RetryOptions
	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry azcore.TelemetryOptions
}

// DefaultClientOptions creates a ClientOptions type initialized with default values.
func DefaultClientOptions() ClientOptions {
	return ClientOptions{
		HTTPClient: azcore.DefaultHTTPClientTransport(),
		Retry:      azcore.DefaultRetryOptions(),
	}
}

func (c *ClientOptions) telemetryOptions() azcore.TelemetryOptions {
	to := c.Telemetry
	if to.Value == "" {
		to.Value = telemetryInfo
	} else {
		to.Value = fmt.Sprintf("%s %s", telemetryInfo, to.Value)
	}
	return to
}

type Client struct {
	u string
	p azcore.Pipeline
}

// NewClient creates an instance of the Client type with the specified endpoint.
func NewClient(endpoint string, cred azcore.Credential, options *ClientOptions) *Client {
	if options == nil {
		o := DefaultClientOptions()
		options = &o
	}
	p := azcore.NewPipeline(options.HTTPClient,
		azcore.NewTelemetryPolicy(options.telemetryOptions()),
		azcore.NewUniqueRequestIDPolicy(),
		azcore.NewRetryPolicy(&options.Retry),
		cred.AuthenticationPolicy(azcore.AuthenticationPolicyOptions{Options: azcore.TokenRequestOptions{Scopes: []string{scope}}}),
		azcore.NewRequestLogPolicy(options.LogOptions))
	return NewClientWithPipeline(endpoint, p)
}

// NewClientWithPipeline creates an instance of the Client type with the specified endpoint and pipeline.
func NewClientWithPipeline(endpoint string, p azcore.Pipeline) *Client {
	return &Client{u: endpoint, p: p}
}
