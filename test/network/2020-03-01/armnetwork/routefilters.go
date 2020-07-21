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

// RouteFiltersOperations contains the methods for the RouteFilters group.
type RouteFiltersOperations interface {
	// BeginCreateOrUpdate - Creates or updates a route filter in a specified resource group.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, routeFilterName string, routeFilterParameters RouteFilter) (*RouteFilterPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (RouteFilterPoller, error)
	// BeginDelete - Deletes the specified route filter.
	BeginDelete(ctx context.Context, resourceGroupName string, routeFilterName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified route filter.
	Get(ctx context.Context, resourceGroupName string, routeFilterName string, routeFiltersGetOptions *RouteFiltersGetOptions) (*RouteFilterResponse, error)
	// List - Gets all route filters in a subscription.
	List() (RouteFilterListResultPager, error)
	// ListByResourceGroup - Gets all route filters in a resource group.
	ListByResourceGroup(resourceGroupName string) (RouteFilterListResultPager, error)
	// UpdateTags - Updates tags of a route filter.
	UpdateTags(ctx context.Context, resourceGroupName string, routeFilterName string, parameters TagsObject) (*RouteFilterResponse, error)
}

// routeFiltersOperations implements the RouteFiltersOperations interface.
type routeFiltersOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - Creates or updates a route filter in a specified resource group.
func (client *routeFiltersOperations) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, routeFilterName string, routeFilterParameters RouteFilter) (*RouteFilterPollerResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, routeFilterName, routeFilterParameters)
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
	pt, err := createPollingTracker("routeFiltersOperations.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &routeFilterPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*RouteFilterResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *routeFiltersOperations) ResumeCreateOrUpdate(token string) (RouteFilterPoller, error) {
	pt, err := resumePollingTracker("routeFiltersOperations.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &routeFilterPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *routeFiltersOperations) createOrUpdateCreateRequest(resourceGroupName string, routeFilterName string, routeFilterParameters RouteFilter) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(routeFilterParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *routeFiltersOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*RouteFilterPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return &RouteFilterPollerResponse{RawResponse: resp.Response}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *routeFiltersOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the specified route filter.
func (client *routeFiltersOperations) BeginDelete(ctx context.Context, resourceGroupName string, routeFilterName string) (*HTTPPollerResponse, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, routeFilterName)
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
	pt, err := createPollingTracker("routeFiltersOperations.Delete", "location", resp, client.deleteHandleError)
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

func (client *routeFiltersOperations) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("routeFiltersOperations.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *routeFiltersOperations) deleteCreateRequest(resourceGroupName string, routeFilterName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
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
func (client *routeFiltersOperations) deleteHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// deleteHandleError handles the Delete error response.
func (client *routeFiltersOperations) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets the specified route filter.
func (client *routeFiltersOperations) Get(ctx context.Context, resourceGroupName string, routeFilterName string, routeFiltersGetOptions *RouteFiltersGetOptions) (*RouteFilterResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, routeFilterName, routeFiltersGetOptions)
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
func (client *routeFiltersOperations) getCreateRequest(resourceGroupName string, routeFilterName string, routeFiltersGetOptions *RouteFiltersGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	if routeFiltersGetOptions != nil && routeFiltersGetOptions.Expand != nil {
		query.Set("$expand", *routeFiltersGetOptions.Expand)
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *routeFiltersOperations) getHandleResponse(resp *azcore.Response) (*RouteFilterResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := RouteFilterResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.RouteFilter)
}

// getHandleError handles the Get error response.
func (client *routeFiltersOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Gets all route filters in a subscription.
func (client *routeFiltersOperations) List() (RouteFilterListResultPager, error) {
	req, err := client.listCreateRequest()
	if err != nil {
		return nil, err
	}
	return &routeFilterListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *RouteFilterListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.RouteFilterListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.RouteFilterListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *routeFiltersOperations) listCreateRequest() (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/routeFilters"
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

// listHandleResponse handles the List response.
func (client *routeFiltersOperations) listHandleResponse(resp *azcore.Response) (*RouteFilterListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := RouteFilterListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.RouteFilterListResult)
}

// listHandleError handles the List error response.
func (client *routeFiltersOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListByResourceGroup - Gets all route filters in a resource group.
func (client *routeFiltersOperations) ListByResourceGroup(resourceGroupName string) (RouteFilterListResultPager, error) {
	req, err := client.listByResourceGroupCreateRequest(resourceGroupName)
	if err != nil {
		return nil, err
	}
	return &routeFilterListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listByResourceGroupHandleResponse,
		advancer: func(resp *RouteFilterListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.RouteFilterListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.RouteFilterListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *routeFiltersOperations) listByResourceGroupCreateRequest(resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters"
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
func (client *routeFiltersOperations) listByResourceGroupHandleResponse(resp *azcore.Response) (*RouteFilterListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listByResourceGroupHandleError(resp)
	}
	result := RouteFilterListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.RouteFilterListResult)
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *routeFiltersOperations) listByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// UpdateTags - Updates tags of a route filter.
func (client *routeFiltersOperations) UpdateTags(ctx context.Context, resourceGroupName string, routeFilterName string, parameters TagsObject) (*RouteFilterResponse, error) {
	req, err := client.updateTagsCreateRequest(resourceGroupName, routeFilterName, parameters)
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
func (client *routeFiltersOperations) updateTagsCreateRequest(resourceGroupName string, routeFilterName string, parameters TagsObject) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPatch, *u)
	return req, req.MarshalAsJSON(parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *routeFiltersOperations) updateTagsHandleResponse(resp *azcore.Response) (*RouteFilterResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.updateTagsHandleError(resp)
	}
	result := RouteFilterResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.RouteFilter)
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *routeFiltersOperations) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
