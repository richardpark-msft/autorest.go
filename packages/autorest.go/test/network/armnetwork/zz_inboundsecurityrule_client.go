//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// InboundSecurityRuleClient contains the methods for the InboundSecurityRule group.
// Don't use this type directly, use NewInboundSecurityRuleClient() instead.
type InboundSecurityRuleClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewInboundSecurityRuleClient creates a new instance of InboundSecurityRuleClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewInboundSecurityRuleClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*InboundSecurityRuleClient, error) {
	cl, err := arm.NewClient(moduleName+".InboundSecurityRuleClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &InboundSecurityRuleClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates the specified Network Virtual Appliance Inbound Security Rules.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - networkVirtualApplianceName - The name of the Network Virtual Appliance.
//   - ruleCollectionName - The name of security rule collection.
//   - parameters - Parameters supplied to the create or update Network Virtual Appliance Inbound Security Rules operation.
//   - options - InboundSecurityRuleClientBeginCreateOrUpdateOptions contains the optional parameters for the InboundSecurityRuleClient.BeginCreateOrUpdate
//     method.
func (client *InboundSecurityRuleClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, ruleCollectionName string, parameters InboundSecurityRule, options *InboundSecurityRuleClientBeginCreateOrUpdateOptions) (*runtime.Poller[InboundSecurityRuleClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, networkVirtualApplianceName, ruleCollectionName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[InboundSecurityRuleClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[InboundSecurityRuleClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates the specified Network Virtual Appliance Inbound Security Rules.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *InboundSecurityRuleClient) createOrUpdate(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, ruleCollectionName string, parameters InboundSecurityRule, options *InboundSecurityRuleClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "InboundSecurityRuleClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, ruleCollectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *InboundSecurityRuleClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, ruleCollectionName string, parameters InboundSecurityRule, options *InboundSecurityRuleClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}/inboundSecurityRules/{ruleCollectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkVirtualApplianceName == "" {
		return nil, errors.New("parameter networkVirtualApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	if ruleCollectionName == "" {
		return nil, errors.New("parameter ruleCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleCollectionName}", url.PathEscape(ruleCollectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
