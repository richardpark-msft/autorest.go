//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azalias

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strconv"
	"strings"
)

type client struct {
	endpoint      string
	clientVersion string
	clientIndex   int32
	pl            runtime.Pipeline
}

// newClient creates a new instance of client with the specified values.
// geography - This parameter specifies where the Azure Maps Creator resource is located. Valid values are us and eu.
// clientVersion - Version number of Azure Maps API.
// clientIndex - Index number of Azure Maps API.
// pl - the pipeline used for sending requests and handling responses.
func newClient(geography *Geography, clientVersion *string, clientIndex *int32, pl runtime.Pipeline) *client {
	hostURL := "https://{geography}.atlas.microsoft.com"
	if geography == nil {
		defaultValue := GeographyUs
		geography = &defaultValue
	}
	hostURL = strings.ReplaceAll(hostURL, "{geography}", string(*geography))
	if clientVersion == nil {
		clientVersionDefault := "2.0"
		clientVersion = &clientVersionDefault
	}
	if clientIndex == nil {
		clientIndexDefault := int32(567)
		clientIndex = &clientIndexDefault
	}
	client := &client{
		endpoint:      hostURL,
		clientVersion: *clientVersion,
		clientIndex:   *clientIndex,
		pl:            pl,
	}
	return client
}

// Create - Applies to: see pricing tiers [https://aka.ms/AzureMapsPricingTier].
// Creator makes it possible to develop applications based on your private indoor map data using Azure Maps API and SDK. This
// [https://docs.microsoft.com/azure/azure-maps/creator-indoor-maps] article
// introduces concepts and tools that apply to Azure Maps Creator.
// This API allows the caller to create an alias. You can also assign the alias during the create request. An alias can reference
// an ID generated by a creator service, but cannot reference another alias
// ID.
// SUBMIT CREATE REQUEST To create your alias, you will use a POST request. If you would like to assign the alias during the
// creation, you will pass the resourceId query parameter.
// CREATE ALIAS RESPONSE The Create API returns a HTTP 201 Created response with the alias resource in the body.
// A sample response from creating an alias:
// { "createdTimestamp": "2020-02-13T21:19:11.123Z", "aliasId": "a8a4b8bb-ecf4-fb27-a618-f41721552766", "creatorDataItemId":
// "e89aebb9-70a3-8fe1-32bb-1fbd0c725f14", "lastUpdatedTimestamp":
// "2020-02-13T21:19:22.123Z" }
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2.0
// options - clientCreateOptions contains the optional parameters for the client.Create method.
func (client *client) Create(ctx context.Context, options *clientCreateOptions) (clientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, options)
	if err != nil {
		return clientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return clientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return clientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *client) createCreateRequest(ctx context.Context, options *clientCreateOptions) (*policy.Request, error) {
	urlPath := "/aliases"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("client-version", client.clientVersion)
	creatorIDDefault := int32(123)
	if options != nil && options.CreatorID != nil {
		creatorIDDefault = *options.CreatorID
	}
	reqQP.Set("creator-id", strconv.FormatInt(int64(creatorIDDefault), 10))
	if options != nil && options.GroupBy != nil {
		for _, qv := range options.GroupBy {
			reqQP.Add("groupBy", fmt.Sprintf("%d", qv))
		}
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("client-index", strconv.FormatInt(int64(client.clientIndex), 10))
	assignedIDDefault := float32(8989)
	if options != nil && options.AssignedID != nil {
		assignedIDDefault = *options.AssignedID
	}
	req.Raw().Header.Set("assigned-id", strconv.FormatFloat(float64(assignedIDDefault), 'f', -1, 32))
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *client) createHandleResponse(resp *http.Response) (clientCreateResponse, error) {
	result := clientCreateResponse{}
	if val := resp.Header.Get("Access-Control-Expose-Headers"); val != "" {
		result.AccessControlExposeHeaders = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AliasesCreateResponse); err != nil {
		return clientCreateResponse{}, err
	}
	return result, nil
}

// GetScript - Retrieve the configuration script identified by configuration name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2.0
// options - clientGetScriptOptions contains the optional parameters for the client.GetScript method.
func (client *client) GetScript(ctx context.Context, props GeoJSONObjectNamedCollection, options *clientGetScriptOptions) (clientGetScriptResponse, error) {
	req, err := client.getScriptCreateRequest(ctx, props, options)
	if err != nil {
		return clientGetScriptResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return clientGetScriptResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return clientGetScriptResponse{}, runtime.NewResponseError(resp)
	}
	return client.getScriptHandleResponse(resp)
}

// getScriptCreateRequest creates the GetScript request.
func (client *client) getScriptCreateRequest(ctx context.Context, props GeoJSONObjectNamedCollection, options *clientGetScriptOptions) (*policy.Request, error) {
	urlPath := "/scripts"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "text/powershell")
	return req, runtime.MarshalAsJSON(req, props)
}

// getScriptHandleResponse handles the GetScript response.
func (client *client) getScriptHandleResponse(resp *http.Response) (clientGetScriptResponse, error) {
	result := clientGetScriptResponse{}
	body, err := runtime.Payload(resp)
	if err != nil {
		return clientGetScriptResponse{}, err
	}
	txt := string(body)
	result.Value = &txt
	return result, nil
}

// NewListPager - Applies to: see pricing tiers [https://aka.ms/AzureMapsPricingTier].
// Creator makes it possible to develop applications based on your private indoor map data using Azure Maps API and SDK. This
// [https://docs.microsoft.com/azure/azure-maps/creator-indoor-maps] article
// introduces concepts and tools that apply to Azure Maps Creator.
// This API allows the caller to fetch a list of all previously successfully created aliases.
// SUBMIT LIST REQUEST To list all your aliases, you will issue a GET request with no additional parameters.
// LIST DATA RESPONSE The List API returns the complete list of all aliases in json format. The response contains the following
// details for each alias resource:
// > createdTimestamp - The timestamp that the alias was created. Format yyyy-MM-ddTHH:mm:ss.sssZ aliasId - The id for the
// alias. creatorDataItemId - The id for the creator data item that this alias
// references (could be null if the alias has not been assigned). lastUpdatedTimestamp - The last time the alias was assigned
// to a resource. Format yyyy-MM-ddTHH:mm:ss.sssZ
// A sample response returning 2 alias resources:
// { "aliases": [ { "createdTimestamp": "2020-02-13T21:19:11.123Z", "aliasId": "a8a4b8bb-ecf4-fb27-a618-f41721552766", "creatorDataItemId":
// "e89aebb9-70a3-8fe1-32bb-1fbd0c725f14", "lastUpdatedTimestamp":
// "2020-02-13T21:19:22.123Z" }, { "createdTimestamp": "2020-02-18T19:53:33.123Z", "aliasId": "1856dbfc-7a66-ee5a-bf8d-51dbfe1906f6",
// "creatorDataItemId": null, "lastUpdatedTimestamp":
// "2020-02-18T19:53:33.123Z" } ] }
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2.0
// options - clientListOptions contains the optional parameters for the client.List method.
func (client *client) NewListPager(options *clientListOptions) *runtime.Pager[clientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[clientListResponse]{
		More: func(page clientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *clientListResponse) (clientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return clientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return clientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return clientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *client) listCreateRequest(ctx context.Context, options *clientListOptions) (*policy.Request, error) {
	urlPath := "/aliases"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("client-version", client.clientVersion)
	if options != nil && options.GroupBy != nil {
		for _, qv := range options.GroupBy {
			reqQP.Add("groupBy", string(qv))
		}
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *client) listHandleResponse(resp *http.Response) (clientListResponse, error) {
	result := clientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResponse); err != nil {
		return clientListResponse{}, err
	}
	return result, nil
}

// PolicyAssignment -
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2.0
// options - clientPolicyAssignmentOptions contains the optional parameters for the client.PolicyAssignment method.
func (client *client) PolicyAssignment(ctx context.Context, props ScheduleCreateOrUpdateProperties, options *clientPolicyAssignmentOptions) (clientPolicyAssignmentResponse, error) {
	req, err := client.policyAssignmentCreateRequest(ctx, props, options)
	if err != nil {
		return clientPolicyAssignmentResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return clientPolicyAssignmentResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return clientPolicyAssignmentResponse{}, runtime.NewResponseError(resp)
	}
	return client.policyAssignmentHandleResponse(resp)
}

// policyAssignmentCreateRequest creates the PolicyAssignment request.
func (client *client) policyAssignmentCreateRequest(ctx context.Context, props ScheduleCreateOrUpdateProperties, options *clientPolicyAssignmentOptions) (*policy.Request, error) {
	urlPath := "/policy"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, props)
}

// policyAssignmentHandleResponse handles the PolicyAssignment response.
func (client *client) policyAssignmentHandleResponse(resp *http.Response) (clientPolicyAssignmentResponse, error) {
	result := clientPolicyAssignmentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyAssignmentProperties); err != nil {
		return clientPolicyAssignmentResponse{}, err
	}
	return result, nil
}