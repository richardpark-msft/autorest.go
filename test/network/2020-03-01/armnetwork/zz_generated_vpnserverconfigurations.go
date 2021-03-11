// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// VPNServerConfigurationsClient contains the methods for the VPNServerConfigurations group.
// Don't use this type directly, use NewVPNServerConfigurationsClient() instead.
type VPNServerConfigurationsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVPNServerConfigurationsClient creates a new instance of VPNServerConfigurationsClient with the specified values.
func NewVPNServerConfigurationsClient(con *armcore.Connection, subscriptionID string) *VPNServerConfigurationsClient {
	return &VPNServerConfigurationsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates a VpnServerConfiguration resource if it doesn't exist else updates the existing VpnServerConfiguration.
func (client *VPNServerConfigurationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, vpnServerConfigurationParameters VPNServerConfiguration, options *VPNServerConfigurationsBeginCreateOrUpdateOptions) (VPNServerConfigurationPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, vpnServerConfigurationName, vpnServerConfigurationParameters, options)
	if err != nil {
		return VPNServerConfigurationPollerResponse{}, err
	}
	result := VPNServerConfigurationPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VPNServerConfigurationsClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return VPNServerConfigurationPollerResponse{}, err
	}
	poller := &vpnServerConfigurationPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VPNServerConfigurationResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new VPNServerConfigurationPoller from the specified resume token.
// token - The value must come from a previous call to VPNServerConfigurationPoller.ResumeToken().
func (client *VPNServerConfigurationsClient) ResumeCreateOrUpdate(token string) (VPNServerConfigurationPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VPNServerConfigurationsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &vpnServerConfigurationPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates a VpnServerConfiguration resource if it doesn't exist else updates the existing VpnServerConfiguration.
func (client *VPNServerConfigurationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, vpnServerConfigurationParameters VPNServerConfiguration, options *VPNServerConfigurationsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vpnServerConfigurationName, vpnServerConfigurationParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VPNServerConfigurationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, vpnServerConfigurationParameters VPNServerConfiguration, options *VPNServerConfigurationsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vpnServerConfigurationName}", url.PathEscape(vpnServerConfigurationName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(vpnServerConfigurationParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *VPNServerConfigurationsClient) createOrUpdateHandleResponse(resp *azcore.Response) (VPNServerConfigurationResponse, error) {
	var val *VPNServerConfiguration
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VPNServerConfigurationResponse{}, err
	}
	return VPNServerConfigurationResponse{RawResponse: resp.Response, VPNServerConfiguration: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VPNServerConfigurationsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes a VpnServerConfiguration.
func (client *VPNServerConfigurationsClient) BeginDelete(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, options *VPNServerConfigurationsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.delete(ctx, resourceGroupName, vpnServerConfigurationName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VPNServerConfigurationsClient.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *VPNServerConfigurationsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VPNServerConfigurationsClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes a VpnServerConfiguration.
func (client *VPNServerConfigurationsClient) delete(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, options *VPNServerConfigurationsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vpnServerConfigurationName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VPNServerConfigurationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, options *VPNServerConfigurationsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vpnServerConfigurationName}", url.PathEscape(vpnServerConfigurationName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *VPNServerConfigurationsClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Retrieves the details of a VpnServerConfiguration.
func (client *VPNServerConfigurationsClient) Get(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, options *VPNServerConfigurationsGetOptions) (VPNServerConfigurationResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vpnServerConfigurationName, options)
	if err != nil {
		return VPNServerConfigurationResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VPNServerConfigurationResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VPNServerConfigurationResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VPNServerConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, options *VPNServerConfigurationsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vpnServerConfigurationName}", url.PathEscape(vpnServerConfigurationName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VPNServerConfigurationsClient) getHandleResponse(resp *azcore.Response) (VPNServerConfigurationResponse, error) {
	var val *VPNServerConfiguration
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VPNServerConfigurationResponse{}, err
	}
	return VPNServerConfigurationResponse{RawResponse: resp.Response, VPNServerConfiguration: val}, nil
}

// getHandleError handles the Get error response.
func (client *VPNServerConfigurationsClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Lists all the VpnServerConfigurations in a subscription.
func (client *VPNServerConfigurationsClient) List(options *VPNServerConfigurationsListOptions) ListVPNServerConfigurationsResultPager {
	return &listVPNServerConfigurationsResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp ListVPNServerConfigurationsResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListVPNServerConfigurationsResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *VPNServerConfigurationsClient) listCreateRequest(ctx context.Context, options *VPNServerConfigurationsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/vpnServerConfigurations"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VPNServerConfigurationsClient) listHandleResponse(resp *azcore.Response) (ListVPNServerConfigurationsResultResponse, error) {
	var val *ListVPNServerConfigurationsResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ListVPNServerConfigurationsResultResponse{}, err
	}
	return ListVPNServerConfigurationsResultResponse{RawResponse: resp.Response, ListVPNServerConfigurationsResult: val}, nil
}

// listHandleError handles the List error response.
func (client *VPNServerConfigurationsClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByResourceGroup - Lists all the vpnServerConfigurations in a resource group.
func (client *VPNServerConfigurationsClient) ListByResourceGroup(resourceGroupName string, options *VPNServerConfigurationsListByResourceGroupOptions) ListVPNServerConfigurationsResultPager {
	return &listVPNServerConfigurationsResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listByResourceGroupHandleResponse,
		errorer:   client.listByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp ListVPNServerConfigurationsResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListVPNServerConfigurationsResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VPNServerConfigurationsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VPNServerConfigurationsListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *VPNServerConfigurationsClient) listByResourceGroupHandleResponse(resp *azcore.Response) (ListVPNServerConfigurationsResultResponse, error) {
	var val *ListVPNServerConfigurationsResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ListVPNServerConfigurationsResultResponse{}, err
	}
	return ListVPNServerConfigurationsResultResponse{RawResponse: resp.Response, ListVPNServerConfigurationsResult: val}, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *VPNServerConfigurationsClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates VpnServerConfiguration tags.
func (client *VPNServerConfigurationsClient) UpdateTags(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, vpnServerConfigurationParameters TagsObject, options *VPNServerConfigurationsUpdateTagsOptions) (VPNServerConfigurationResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, vpnServerConfigurationName, vpnServerConfigurationParameters, options)
	if err != nil {
		return VPNServerConfigurationResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VPNServerConfigurationResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VPNServerConfigurationResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *VPNServerConfigurationsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, vpnServerConfigurationParameters TagsObject, options *VPNServerConfigurationsUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vpnServerConfigurationName}", url.PathEscape(vpnServerConfigurationName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(vpnServerConfigurationParameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *VPNServerConfigurationsClient) updateTagsHandleResponse(resp *azcore.Response) (VPNServerConfigurationResponse, error) {
	var val *VPNServerConfiguration
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VPNServerConfigurationResponse{}, err
	}
	return VPNServerConfigurationResponse{RawResponse: resp.Response, VPNServerConfiguration: val}, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *VPNServerConfigurationsClient) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
