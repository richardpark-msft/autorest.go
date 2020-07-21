// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

// VirtualNetworkTapsOperations contains the methods for the VirtualNetworkTaps group.
type VirtualNetworkTapsOperations interface {
	// BeginCreateOrUpdate - Creates or updates a Virtual Network Tap.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, tapName string, parameters VirtualNetworkTap) (*VirtualNetworkTapPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (VirtualNetworkTapPoller, error)
	// BeginDelete - Deletes the specified virtual network tap.
	BeginDelete(ctx context.Context, resourceGroupName string, tapName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets information about the specified virtual network tap.
	Get(ctx context.Context, resourceGroupName string, tapName string) (*VirtualNetworkTapResponse, error)
	// ListAll - Gets all the VirtualNetworkTaps in a subscription.
	ListAll() (VirtualNetworkTapListResultPager, error)
	// ListByResourceGroup - Gets all the VirtualNetworkTaps in a subscription.
	ListByResourceGroup(resourceGroupName string) (VirtualNetworkTapListResultPager, error)
	// UpdateTags - Updates an VirtualNetworkTap tags.
	UpdateTags(ctx context.Context, resourceGroupName string, tapName string, tapParameters TagsObject) (*VirtualNetworkTapResponse, error)
}

// virtualNetworkTapsOperations implements the VirtualNetworkTapsOperations interface.
type virtualNetworkTapsOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - Creates or updates a Virtual Network Tap.
func (client *virtualNetworkTapsOperations) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, tapName string, parameters VirtualNetworkTap) (*VirtualNetworkTapPollerResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, tapName, parameters)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.createOrUpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("virtualNetworkTapsOperations.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &virtualNetworkTapPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VirtualNetworkTapResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *virtualNetworkTapsOperations) ResumeCreateOrUpdate(token string) (VirtualNetworkTapPoller, error) {
	pt, err := resumePollingTracker("virtualNetworkTapsOperations.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualNetworkTapPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *virtualNetworkTapsOperations) createOrUpdateCreateRequest(resourceGroupName string, tapName string, parameters VirtualNetworkTap) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{tapName}", url.PathEscape(tapName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *virtualNetworkTapsOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*VirtualNetworkTapPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return &VirtualNetworkTapPollerResponse{RawResponse: resp.Response}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *virtualNetworkTapsOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the specified virtual network tap.
func (client *virtualNetworkTapsOperations) BeginDelete(ctx context.Context, resourceGroupName string, tapName string) (*HTTPPollerResponse, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, tapName)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.deleteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("virtualNetworkTapsOperations.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *virtualNetworkTapsOperations) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("virtualNetworkTapsOperations.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *virtualNetworkTapsOperations) deleteCreateRequest(resourceGroupName string, tapName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{tapName}", url.PathEscape(tapName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *virtualNetworkTapsOperations) deleteHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// deleteHandleError handles the Delete error response.
func (client *virtualNetworkTapsOperations) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets information about the specified virtual network tap.
func (client *virtualNetworkTapsOperations) Get(ctx context.Context, resourceGroupName string, tapName string) (*VirtualNetworkTapResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, tapName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client *virtualNetworkTapsOperations) getCreateRequest(resourceGroupName string, tapName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{tapName}", url.PathEscape(tapName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *virtualNetworkTapsOperations) getHandleResponse(resp *azcore.Response) (*VirtualNetworkTapResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := VirtualNetworkTapResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualNetworkTap)
}

// getHandleError handles the Get error response.
func (client *virtualNetworkTapsOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListAll - Gets all the VirtualNetworkTaps in a subscription.
func (client *virtualNetworkTapsOperations) ListAll() (VirtualNetworkTapListResultPager, error) {
	req, err := client.listAllCreateRequest()
	if err != nil {
		return nil, err
	}
	return &virtualNetworkTapListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listAllHandleResponse,
		advancer: func(resp *VirtualNetworkTapListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.VirtualNetworkTapListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.VirtualNetworkTapListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listAllCreateRequest creates the ListAll request.
func (client *virtualNetworkTapsOperations) listAllCreateRequest() (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualNetworkTaps"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *virtualNetworkTapsOperations) listAllHandleResponse(resp *azcore.Response) (*VirtualNetworkTapListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listAllHandleError(resp)
	}
	result := VirtualNetworkTapListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualNetworkTapListResult)
}

// listAllHandleError handles the ListAll error response.
func (client *virtualNetworkTapsOperations) listAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListByResourceGroup - Gets all the VirtualNetworkTaps in a subscription.
func (client *virtualNetworkTapsOperations) ListByResourceGroup(resourceGroupName string) (VirtualNetworkTapListResultPager, error) {
	req, err := client.listByResourceGroupCreateRequest(resourceGroupName)
	if err != nil {
		return nil, err
	}
	return &virtualNetworkTapListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listByResourceGroupHandleResponse,
		advancer: func(resp *VirtualNetworkTapListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.VirtualNetworkTapListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.VirtualNetworkTapListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *virtualNetworkTapsOperations) listByResourceGroupCreateRequest(resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *virtualNetworkTapsOperations) listByResourceGroupHandleResponse(resp *azcore.Response) (*VirtualNetworkTapListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listByResourceGroupHandleError(resp)
	}
	result := VirtualNetworkTapListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualNetworkTapListResult)
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *virtualNetworkTapsOperations) listByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// UpdateTags - Updates an VirtualNetworkTap tags.
func (client *virtualNetworkTapsOperations) UpdateTags(ctx context.Context, resourceGroupName string, tapName string, tapParameters TagsObject) (*VirtualNetworkTapResponse, error) {
	req, err := client.updateTagsCreateRequest(resourceGroupName, tapName, tapParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.updateTagsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *virtualNetworkTapsOperations) updateTagsCreateRequest(resourceGroupName string, tapName string, tapParameters TagsObject) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{tapName}", url.PathEscape(tapName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPatch, *u)
	return req, req.MarshalAsJSON(tapParameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *virtualNetworkTapsOperations) updateTagsHandleResponse(resp *azcore.Response) (*VirtualNetworkTapResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.updateTagsHandleError(resp)
	}
	result := VirtualNetworkTapResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualNetworkTap)
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *virtualNetworkTapsOperations) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
