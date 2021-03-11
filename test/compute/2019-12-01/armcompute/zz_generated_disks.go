// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// DisksClient contains the methods for the Disks group.
// Don't use this type directly, use NewDisksClient() instead.
type DisksClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewDisksClient creates a new instance of DisksClient with the specified values.
func NewDisksClient(con *armcore.Connection, subscriptionID string) *DisksClient {
	return &DisksClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a disk.
func (client *DisksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, diskName string, disk Disk, options *DisksBeginCreateOrUpdateOptions) (DiskPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, diskName, disk, options)
	if err != nil {
		return DiskPollerResponse{}, err
	}
	result := DiskPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("DisksClient.CreateOrUpdate", "", resp, client.createOrUpdateHandleError)
	if err != nil {
		return DiskPollerResponse{}, err
	}
	poller := &diskPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (DiskResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new DiskPoller from the specified resume token.
// token - The value must come from a previous call to DiskPoller.ResumeToken().
func (client *DisksClient) ResumeCreateOrUpdate(token string) (DiskPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("DisksClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &diskPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a disk.
func (client *DisksClient) createOrUpdate(ctx context.Context, resourceGroupName string, diskName string, disk Disk, options *DisksBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, diskName, disk, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DisksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, diskName string, disk Disk, options *DisksBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-11-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(disk)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DisksClient) createOrUpdateHandleResponse(resp *azcore.Response) (DiskResponse, error) {
	var val *Disk
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DiskResponse{}, err
	}
	return DiskResponse{RawResponse: resp.Response, Disk: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DisksClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginDelete - Deletes a disk.
func (client *DisksClient) BeginDelete(ctx context.Context, resourceGroupName string, diskName string, options *DisksBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.delete(ctx, resourceGroupName, diskName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("DisksClient.Delete", "", resp, client.deleteHandleError)
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
func (client *DisksClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("DisksClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes a disk.
func (client *DisksClient) delete(ctx context.Context, resourceGroupName string, diskName string, options *DisksBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, diskName, options)
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
func (client *DisksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, diskName string, options *DisksBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-11-01")
	req.URL.RawQuery = query.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *DisksClient) deleteHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Get - Gets information about a disk.
func (client *DisksClient) Get(ctx context.Context, resourceGroupName string, diskName string, options *DisksGetOptions) (DiskResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, diskName, options)
	if err != nil {
		return DiskResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DiskResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DiskResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DisksClient) getCreateRequest(ctx context.Context, resourceGroupName string, diskName string, options *DisksGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-11-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DisksClient) getHandleResponse(resp *azcore.Response) (DiskResponse, error) {
	var val *Disk
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DiskResponse{}, err
	}
	return DiskResponse{RawResponse: resp.Response, Disk: val}, nil
}

// getHandleError handles the Get error response.
func (client *DisksClient) getHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginGrantAccess - Grants access to a disk.
func (client *DisksClient) BeginGrantAccess(ctx context.Context, resourceGroupName string, diskName string, grantAccessData GrantAccessData, options *DisksBeginGrantAccessOptions) (AccessURIPollerResponse, error) {
	resp, err := client.grantAccess(ctx, resourceGroupName, diskName, grantAccessData, options)
	if err != nil {
		return AccessURIPollerResponse{}, err
	}
	result := AccessURIPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("DisksClient.GrantAccess", "location", resp, client.grantAccessHandleError)
	if err != nil {
		return AccessURIPollerResponse{}, err
	}
	poller := &accessURIPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (AccessURIResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeGrantAccess creates a new AccessURIPoller from the specified resume token.
// token - The value must come from a previous call to AccessURIPoller.ResumeToken().
func (client *DisksClient) ResumeGrantAccess(token string) (AccessURIPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("DisksClient.GrantAccess", token, client.grantAccessHandleError)
	if err != nil {
		return nil, err
	}
	return &accessURIPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// GrantAccess - Grants access to a disk.
func (client *DisksClient) grantAccess(ctx context.Context, resourceGroupName string, diskName string, grantAccessData GrantAccessData, options *DisksBeginGrantAccessOptions) (*azcore.Response, error) {
	req, err := client.grantAccessCreateRequest(ctx, resourceGroupName, diskName, grantAccessData, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.grantAccessHandleError(resp)
	}
	return resp, nil
}

// grantAccessCreateRequest creates the GrantAccess request.
func (client *DisksClient) grantAccessCreateRequest(ctx context.Context, resourceGroupName string, diskName string, grantAccessData GrantAccessData, options *DisksBeginGrantAccessOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}/beginGetAccess"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-11-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(grantAccessData)
}

// grantAccessHandleResponse handles the GrantAccess response.
func (client *DisksClient) grantAccessHandleResponse(resp *azcore.Response) (AccessURIResponse, error) {
	var val *AccessURI
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return AccessURIResponse{}, err
	}
	return AccessURIResponse{RawResponse: resp.Response, AccessURI: val}, nil
}

// grantAccessHandleError handles the GrantAccess error response.
func (client *DisksClient) grantAccessHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// List - Lists all the disks under a subscription.
func (client *DisksClient) List(options *DisksListOptions) DiskListPager {
	return &diskListPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp DiskListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.DiskList.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *DisksClient) listCreateRequest(ctx context.Context, options *DisksListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/disks"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-11-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DisksClient) listHandleResponse(resp *azcore.Response) (DiskListResponse, error) {
	var val *DiskList
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DiskListResponse{}, err
	}
	return DiskListResponse{RawResponse: resp.Response, DiskList: val}, nil
}

// listHandleError handles the List error response.
func (client *DisksClient) listHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListByResourceGroup - Lists all the disks under a resource group.
func (client *DisksClient) ListByResourceGroup(resourceGroupName string, options *DisksListByResourceGroupOptions) DiskListPager {
	return &diskListPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listByResourceGroupHandleResponse,
		errorer:   client.listByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp DiskListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.DiskList.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DisksClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DisksListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-11-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DisksClient) listByResourceGroupHandleResponse(resp *azcore.Response) (DiskListResponse, error) {
	var val *DiskList
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DiskListResponse{}, err
	}
	return DiskListResponse{RawResponse: resp.Response, DiskList: val}, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *DisksClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginRevokeAccess - Revokes access to a disk.
func (client *DisksClient) BeginRevokeAccess(ctx context.Context, resourceGroupName string, diskName string, options *DisksBeginRevokeAccessOptions) (HTTPPollerResponse, error) {
	resp, err := client.revokeAccess(ctx, resourceGroupName, diskName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("DisksClient.RevokeAccess", "location", resp, client.revokeAccessHandleError)
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

// ResumeRevokeAccess creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *DisksClient) ResumeRevokeAccess(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("DisksClient.RevokeAccess", token, client.revokeAccessHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// RevokeAccess - Revokes access to a disk.
func (client *DisksClient) revokeAccess(ctx context.Context, resourceGroupName string, diskName string, options *DisksBeginRevokeAccessOptions) (*azcore.Response, error) {
	req, err := client.revokeAccessCreateRequest(ctx, resourceGroupName, diskName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.revokeAccessHandleError(resp)
	}
	return resp, nil
}

// revokeAccessCreateRequest creates the RevokeAccess request.
func (client *DisksClient) revokeAccessCreateRequest(ctx context.Context, resourceGroupName string, diskName string, options *DisksBeginRevokeAccessOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}/endGetAccess"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-11-01")
	req.URL.RawQuery = query.Encode()
	return req, nil
}

// revokeAccessHandleError handles the RevokeAccess error response.
func (client *DisksClient) revokeAccessHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginUpdate - Updates (patches) a disk.
func (client *DisksClient) BeginUpdate(ctx context.Context, resourceGroupName string, diskName string, disk DiskUpdate, options *DisksBeginUpdateOptions) (DiskPollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, diskName, disk, options)
	if err != nil {
		return DiskPollerResponse{}, err
	}
	result := DiskPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("DisksClient.Update", "", resp, client.updateHandleError)
	if err != nil {
		return DiskPollerResponse{}, err
	}
	poller := &diskPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (DiskResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdate creates a new DiskPoller from the specified resume token.
// token - The value must come from a previous call to DiskPoller.ResumeToken().
func (client *DisksClient) ResumeUpdate(token string) (DiskPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("DisksClient.Update", token, client.updateHandleError)
	if err != nil {
		return nil, err
	}
	return &diskPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Update - Updates (patches) a disk.
func (client *DisksClient) update(ctx context.Context, resourceGroupName string, diskName string, disk DiskUpdate, options *DisksBeginUpdateOptions) (*azcore.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, diskName, disk, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *DisksClient) updateCreateRequest(ctx context.Context, resourceGroupName string, diskName string, disk DiskUpdate, options *DisksBeginUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-11-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(disk)
}

// updateHandleResponse handles the Update response.
func (client *DisksClient) updateHandleResponse(resp *azcore.Response) (DiskResponse, error) {
	var val *Disk
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DiskResponse{}, err
	}
	return DiskResponse{RawResponse: resp.Response, Disk: val}, nil
}

// updateHandleError handles the Update error response.
func (client *DisksClient) updateHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
