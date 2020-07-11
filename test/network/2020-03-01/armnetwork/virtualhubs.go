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
	"strings"
	"time"
)

// VirtualHubsOperations contains the methods for the VirtualHubs group.
type VirtualHubsOperations interface {
	// BeginCreateOrUpdate - Creates a VirtualHub resource if it doesn't exist else updates the existing VirtualHub.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, virtualHubParameters VirtualHub) (*VirtualHubPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (VirtualHubPoller, error)
	// BeginDelete - Deletes a VirtualHub.
	BeginDelete(ctx context.Context, resourceGroupName string, virtualHubName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Retrieves the details of a VirtualHub.
	Get(ctx context.Context, resourceGroupName string, virtualHubName string) (*VirtualHubResponse, error)
	// List - Lists all the VirtualHubs in a subscription.
	List() (ListVirtualHubsResultPager, error)
	// ListByResourceGroup - Lists all the VirtualHubs in a resource group.
	ListByResourceGroup(resourceGroupName string) (ListVirtualHubsResultPager, error)
	// UpdateTags - Updates VirtualHub tags.
	UpdateTags(ctx context.Context, resourceGroupName string, virtualHubName string, virtualHubParameters TagsObject) (*VirtualHubResponse, error)
}

// virtualHubsOperations implements the VirtualHubsOperations interface.
type virtualHubsOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - Creates a VirtualHub resource if it doesn't exist else updates the existing VirtualHub.
func (client *virtualHubsOperations) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, virtualHubParameters VirtualHub) (*VirtualHubPollerResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, virtualHubName, virtualHubParameters)
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
	pt, err := createPollingTracker("virtualHubsOperations.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &virtualHubPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VirtualHubResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *virtualHubsOperations) ResumeCreateOrUpdate(token string) (VirtualHubPoller, error) {
	pt, err := resumePollingTracker("virtualHubsOperations.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualHubPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *virtualHubsOperations) createOrUpdateCreateRequest(resourceGroupName string, virtualHubName string, virtualHubParameters VirtualHub) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(virtualHubParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *virtualHubsOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*VirtualHubPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return &VirtualHubPollerResponse{RawResponse: resp.Response}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *virtualHubsOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes a VirtualHub.
func (client *virtualHubsOperations) BeginDelete(ctx context.Context, resourceGroupName string, virtualHubName string) (*HTTPPollerResponse, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, virtualHubName)
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
	pt, err := createPollingTracker("virtualHubsOperations.Delete", "location", resp, client.deleteHandleError)
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

func (client *virtualHubsOperations) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("virtualHubsOperations.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *virtualHubsOperations) deleteCreateRequest(resourceGroupName string, virtualHubName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	u, err := client.u.Parse(urlPath)
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
func (client *virtualHubsOperations) deleteHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// deleteHandleError handles the Delete error response.
func (client *virtualHubsOperations) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Retrieves the details of a VirtualHub.
func (client *virtualHubsOperations) Get(ctx context.Context, resourceGroupName string, virtualHubName string) (*VirtualHubResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, virtualHubName)
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
func (client *virtualHubsOperations) getCreateRequest(resourceGroupName string, virtualHubName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	u, err := client.u.Parse(urlPath)
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
func (client *virtualHubsOperations) getHandleResponse(resp *azcore.Response) (*VirtualHubResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := VirtualHubResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualHub)
}

// getHandleError handles the Get error response.
func (client *virtualHubsOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Lists all the VirtualHubs in a subscription.
func (client *virtualHubsOperations) List() (ListVirtualHubsResultPager, error) {
	req, err := client.listCreateRequest()
	if err != nil {
		return nil, err
	}
	return &listVirtualHubsResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *ListVirtualHubsResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.ListVirtualHubsResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.ListVirtualHubsResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *virtualHubsOperations) listCreateRequest() (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualHubs"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listHandleResponse handles the List response.
func (client *virtualHubsOperations) listHandleResponse(resp *azcore.Response) (*ListVirtualHubsResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := ListVirtualHubsResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ListVirtualHubsResult)
}

// listHandleError handles the List error response.
func (client *virtualHubsOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListByResourceGroup - Lists all the VirtualHubs in a resource group.
func (client *virtualHubsOperations) ListByResourceGroup(resourceGroupName string) (ListVirtualHubsResultPager, error) {
	req, err := client.listByResourceGroupCreateRequest(resourceGroupName)
	if err != nil {
		return nil, err
	}
	return &listVirtualHubsResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listByResourceGroupHandleResponse,
		advancer: func(resp *ListVirtualHubsResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.ListVirtualHubsResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.ListVirtualHubsResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *virtualHubsOperations) listByResourceGroupCreateRequest(resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	u, err := client.u.Parse(urlPath)
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
func (client *virtualHubsOperations) listByResourceGroupHandleResponse(resp *azcore.Response) (*ListVirtualHubsResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listByResourceGroupHandleError(resp)
	}
	result := ListVirtualHubsResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ListVirtualHubsResult)
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *virtualHubsOperations) listByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// UpdateTags - Updates VirtualHub tags.
func (client *virtualHubsOperations) UpdateTags(ctx context.Context, resourceGroupName string, virtualHubName string, virtualHubParameters TagsObject) (*VirtualHubResponse, error) {
	req, err := client.updateTagsCreateRequest(resourceGroupName, virtualHubName, virtualHubParameters)
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
func (client *virtualHubsOperations) updateTagsCreateRequest(resourceGroupName string, virtualHubName string, virtualHubParameters TagsObject) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPatch, *u)
	return req, req.MarshalAsJSON(virtualHubParameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *virtualHubsOperations) updateTagsHandleResponse(resp *azcore.Response) (*VirtualHubResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.updateTagsHandleError(resp)
	}
	result := VirtualHubResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualHub)
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *virtualHubsOperations) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
