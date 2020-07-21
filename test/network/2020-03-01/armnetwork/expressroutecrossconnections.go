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

// ExpressRouteCrossConnectionsOperations contains the methods for the ExpressRouteCrossConnections group.
type ExpressRouteCrossConnectionsOperations interface {
	// BeginCreateOrUpdate - Update the specified ExpressRouteCrossConnection.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, parameters ExpressRouteCrossConnection) (*ExpressRouteCrossConnectionPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (ExpressRouteCrossConnectionPoller, error)
	// Get - Gets details about the specified ExpressRouteCrossConnection.
	Get(ctx context.Context, resourceGroupName string, crossConnectionName string) (*ExpressRouteCrossConnectionResponse, error)
	// List - Retrieves all the ExpressRouteCrossConnections in a subscription.
	List() (ExpressRouteCrossConnectionListResultPager, error)
	// BeginListArpTable - Gets the currently advertised ARP table associated with the express route cross connection in a resource group.
	BeginListArpTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string) (*ExpressRouteCircuitsArpTableListResultPollerResponse, error)
	// ResumeListArpTable - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeListArpTable(token string) (ExpressRouteCircuitsArpTableListResultPoller, error)
	// ListByResourceGroup - Retrieves all the ExpressRouteCrossConnections in a resource group.
	ListByResourceGroup(resourceGroupName string) (ExpressRouteCrossConnectionListResultPager, error)
	// BeginListRoutesTable - Gets the currently advertised routes table associated with the express route cross connection in a resource group.
	BeginListRoutesTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string) (*ExpressRouteCircuitsRoutesTableListResultPollerResponse, error)
	// ResumeListRoutesTable - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeListRoutesTable(token string) (ExpressRouteCircuitsRoutesTableListResultPoller, error)
	// BeginListRoutesTableSummary - Gets the route table summary associated with the express route cross connection in a resource group.
	BeginListRoutesTableSummary(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string) (*ExpressRouteCrossConnectionsRoutesTableSummaryListResultPollerResponse, error)
	// ResumeListRoutesTableSummary - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeListRoutesTableSummary(token string) (ExpressRouteCrossConnectionsRoutesTableSummaryListResultPoller, error)
	// UpdateTags - Updates an express route cross connection tags.
	UpdateTags(ctx context.Context, resourceGroupName string, crossConnectionName string, crossConnectionParameters TagsObject) (*ExpressRouteCrossConnectionResponse, error)
}

// expressRouteCrossConnectionsOperations implements the ExpressRouteCrossConnectionsOperations interface.
type expressRouteCrossConnectionsOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - Update the specified ExpressRouteCrossConnection.
func (client *expressRouteCrossConnectionsOperations) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, parameters ExpressRouteCrossConnection) (*ExpressRouteCrossConnectionPollerResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, crossConnectionName, parameters)
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
	pt, err := createPollingTracker("expressRouteCrossConnectionsOperations.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteCrossConnectionPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteCrossConnectionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *expressRouteCrossConnectionsOperations) ResumeCreateOrUpdate(token string) (ExpressRouteCrossConnectionPoller, error) {
	pt, err := resumePollingTracker("expressRouteCrossConnectionsOperations.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCrossConnectionPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *expressRouteCrossConnectionsOperations) createOrUpdateCreateRequest(resourceGroupName string, crossConnectionName string, parameters ExpressRouteCrossConnection) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
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
func (client *expressRouteCrossConnectionsOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*ExpressRouteCrossConnectionPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return &ExpressRouteCrossConnectionPollerResponse{RawResponse: resp.Response}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *expressRouteCrossConnectionsOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets details about the specified ExpressRouteCrossConnection.
func (client *expressRouteCrossConnectionsOperations) Get(ctx context.Context, resourceGroupName string, crossConnectionName string) (*ExpressRouteCrossConnectionResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, crossConnectionName)
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
func (client *expressRouteCrossConnectionsOperations) getCreateRequest(resourceGroupName string, crossConnectionName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
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
func (client *expressRouteCrossConnectionsOperations) getHandleResponse(resp *azcore.Response) (*ExpressRouteCrossConnectionResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := ExpressRouteCrossConnectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCrossConnection)
}

// getHandleError handles the Get error response.
func (client *expressRouteCrossConnectionsOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Retrieves all the ExpressRouteCrossConnections in a subscription.
func (client *expressRouteCrossConnectionsOperations) List() (ExpressRouteCrossConnectionListResultPager, error) {
	req, err := client.listCreateRequest()
	if err != nil {
		return nil, err
	}
	return &expressRouteCrossConnectionListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *ExpressRouteCrossConnectionListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.ExpressRouteCrossConnectionListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.ExpressRouteCrossConnectionListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *expressRouteCrossConnectionsOperations) listCreateRequest() (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteCrossConnections"
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
func (client *expressRouteCrossConnectionsOperations) listHandleResponse(resp *azcore.Response) (*ExpressRouteCrossConnectionListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := ExpressRouteCrossConnectionListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCrossConnectionListResult)
}

// listHandleError handles the List error response.
func (client *expressRouteCrossConnectionsOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListArpTable - Gets the currently advertised ARP table associated with the express route cross connection in a resource group.
func (client *expressRouteCrossConnectionsOperations) BeginListArpTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string) (*ExpressRouteCircuitsArpTableListResultPollerResponse, error) {
	req, err := client.listArpTableCreateRequest(resourceGroupName, crossConnectionName, peeringName, devicePath)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.listArpTableHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("expressRouteCrossConnectionsOperations.ListArpTable", "location", resp, client.listArpTableHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteCircuitsArpTableListResultPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteCircuitsArpTableListResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *expressRouteCrossConnectionsOperations) ResumeListArpTable(token string) (ExpressRouteCircuitsArpTableListResultPoller, error) {
	pt, err := resumePollingTracker("expressRouteCrossConnectionsOperations.ListArpTable", token, client.listArpTableHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCircuitsArpTableListResultPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// listArpTableCreateRequest creates the ListArpTable request.
func (client *expressRouteCrossConnectionsOperations) listArpTableCreateRequest(resourceGroupName string, crossConnectionName string, peeringName string, devicePath string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}/arpTables/{devicePath}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, nil
}

// listArpTableHandleResponse handles the ListArpTable response.
func (client *expressRouteCrossConnectionsOperations) listArpTableHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitsArpTableListResultPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.listArpTableHandleError(resp)
	}
	return &ExpressRouteCircuitsArpTableListResultPollerResponse{RawResponse: resp.Response}, nil
}

// listArpTableHandleError handles the ListArpTable error response.
func (client *expressRouteCrossConnectionsOperations) listArpTableHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListByResourceGroup - Retrieves all the ExpressRouteCrossConnections in a resource group.
func (client *expressRouteCrossConnectionsOperations) ListByResourceGroup(resourceGroupName string) (ExpressRouteCrossConnectionListResultPager, error) {
	req, err := client.listByResourceGroupCreateRequest(resourceGroupName)
	if err != nil {
		return nil, err
	}
	return &expressRouteCrossConnectionListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listByResourceGroupHandleResponse,
		advancer: func(resp *ExpressRouteCrossConnectionListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.ExpressRouteCrossConnectionListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.ExpressRouteCrossConnectionListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *expressRouteCrossConnectionsOperations) listByResourceGroupCreateRequest(resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections"
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
func (client *expressRouteCrossConnectionsOperations) listByResourceGroupHandleResponse(resp *azcore.Response) (*ExpressRouteCrossConnectionListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listByResourceGroupHandleError(resp)
	}
	result := ExpressRouteCrossConnectionListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCrossConnectionListResult)
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *expressRouteCrossConnectionsOperations) listByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListRoutesTable - Gets the currently advertised routes table associated with the express route cross connection in a resource group.
func (client *expressRouteCrossConnectionsOperations) BeginListRoutesTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string) (*ExpressRouteCircuitsRoutesTableListResultPollerResponse, error) {
	req, err := client.listRoutesTableCreateRequest(resourceGroupName, crossConnectionName, peeringName, devicePath)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.listRoutesTableHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("expressRouteCrossConnectionsOperations.ListRoutesTable", "location", resp, client.listRoutesTableHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteCircuitsRoutesTableListResultPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteCircuitsRoutesTableListResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *expressRouteCrossConnectionsOperations) ResumeListRoutesTable(token string) (ExpressRouteCircuitsRoutesTableListResultPoller, error) {
	pt, err := resumePollingTracker("expressRouteCrossConnectionsOperations.ListRoutesTable", token, client.listRoutesTableHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCircuitsRoutesTableListResultPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// listRoutesTableCreateRequest creates the ListRoutesTable request.
func (client *expressRouteCrossConnectionsOperations) listRoutesTableCreateRequest(resourceGroupName string, crossConnectionName string, peeringName string, devicePath string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}/routeTables/{devicePath}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, nil
}

// listRoutesTableHandleResponse handles the ListRoutesTable response.
func (client *expressRouteCrossConnectionsOperations) listRoutesTableHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitsRoutesTableListResultPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.listRoutesTableHandleError(resp)
	}
	return &ExpressRouteCircuitsRoutesTableListResultPollerResponse{RawResponse: resp.Response}, nil
}

// listRoutesTableHandleError handles the ListRoutesTable error response.
func (client *expressRouteCrossConnectionsOperations) listRoutesTableHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListRoutesTableSummary - Gets the route table summary associated with the express route cross connection in a resource group.
func (client *expressRouteCrossConnectionsOperations) BeginListRoutesTableSummary(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string) (*ExpressRouteCrossConnectionsRoutesTableSummaryListResultPollerResponse, error) {
	req, err := client.listRoutesTableSummaryCreateRequest(resourceGroupName, crossConnectionName, peeringName, devicePath)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.listRoutesTableSummaryHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("expressRouteCrossConnectionsOperations.ListRoutesTableSummary", "location", resp, client.listRoutesTableSummaryHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteCrossConnectionsRoutesTableSummaryListResultPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteCrossConnectionsRoutesTableSummaryListResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *expressRouteCrossConnectionsOperations) ResumeListRoutesTableSummary(token string) (ExpressRouteCrossConnectionsRoutesTableSummaryListResultPoller, error) {
	pt, err := resumePollingTracker("expressRouteCrossConnectionsOperations.ListRoutesTableSummary", token, client.listRoutesTableSummaryHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCrossConnectionsRoutesTableSummaryListResultPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// listRoutesTableSummaryCreateRequest creates the ListRoutesTableSummary request.
func (client *expressRouteCrossConnectionsOperations) listRoutesTableSummaryCreateRequest(resourceGroupName string, crossConnectionName string, peeringName string, devicePath string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}/routeTablesSummary/{devicePath}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, nil
}

// listRoutesTableSummaryHandleResponse handles the ListRoutesTableSummary response.
func (client *expressRouteCrossConnectionsOperations) listRoutesTableSummaryHandleResponse(resp *azcore.Response) (*ExpressRouteCrossConnectionsRoutesTableSummaryListResultPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.listRoutesTableSummaryHandleError(resp)
	}
	return &ExpressRouteCrossConnectionsRoutesTableSummaryListResultPollerResponse{RawResponse: resp.Response}, nil
}

// listRoutesTableSummaryHandleError handles the ListRoutesTableSummary error response.
func (client *expressRouteCrossConnectionsOperations) listRoutesTableSummaryHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// UpdateTags - Updates an express route cross connection tags.
func (client *expressRouteCrossConnectionsOperations) UpdateTags(ctx context.Context, resourceGroupName string, crossConnectionName string, crossConnectionParameters TagsObject) (*ExpressRouteCrossConnectionResponse, error) {
	req, err := client.updateTagsCreateRequest(resourceGroupName, crossConnectionName, crossConnectionParameters)
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
func (client *expressRouteCrossConnectionsOperations) updateTagsCreateRequest(resourceGroupName string, crossConnectionName string, crossConnectionParameters TagsObject) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPatch, *u)
	return req, req.MarshalAsJSON(crossConnectionParameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *expressRouteCrossConnectionsOperations) updateTagsHandleResponse(resp *azcore.Response) (*ExpressRouteCrossConnectionResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.updateTagsHandleError(resp)
	}
	result := ExpressRouteCrossConnectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCrossConnection)
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *expressRouteCrossConnectionsOperations) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
