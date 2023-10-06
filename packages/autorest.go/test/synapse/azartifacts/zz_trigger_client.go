//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// TriggerClient contains the methods for the Trigger group.
// Don't use this type directly, use a constructor function instead.
type TriggerClient struct {
	internal *azcore.Client
	endpoint string
}

// BeginCreateOrUpdateTrigger - Creates or updates a trigger.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - triggerName - The trigger name.
//   - trigger - Trigger resource definition.
//   - options - TriggerClientBeginCreateOrUpdateTriggerOptions contains the optional parameters for the TriggerClient.BeginCreateOrUpdateTrigger
//     method.
func (client *TriggerClient) BeginCreateOrUpdateTrigger(ctx context.Context, triggerName string, trigger TriggerResource, options *TriggerClientBeginCreateOrUpdateTriggerOptions) (*runtime.Poller[TriggerClientCreateOrUpdateTriggerResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdateTrigger(ctx, triggerName, trigger, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[TriggerClientCreateOrUpdateTriggerResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[TriggerClientCreateOrUpdateTriggerResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdateTrigger - Creates or updates a trigger.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
func (client *TriggerClient) createOrUpdateTrigger(ctx context.Context, triggerName string, trigger TriggerResource, options *TriggerClientBeginCreateOrUpdateTriggerOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateTriggerCreateRequest(ctx, triggerName, trigger, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateTriggerCreateRequest creates the CreateOrUpdateTrigger request.
func (client *TriggerClient) createOrUpdateTriggerCreateRequest(ctx context.Context, triggerName string, trigger TriggerResource, options *TriggerClientBeginCreateOrUpdateTriggerOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, trigger); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDeleteTrigger - Deletes a trigger.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - triggerName - The trigger name.
//   - options - TriggerClientBeginDeleteTriggerOptions contains the optional parameters for the TriggerClient.BeginDeleteTrigger
//     method.
func (client *TriggerClient) BeginDeleteTrigger(ctx context.Context, triggerName string, options *TriggerClientBeginDeleteTriggerOptions) (*runtime.Poller[TriggerClientDeleteTriggerResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteTrigger(ctx, triggerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[TriggerClientDeleteTriggerResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[TriggerClientDeleteTriggerResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// DeleteTrigger - Deletes a trigger.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
func (client *TriggerClient) deleteTrigger(ctx context.Context, triggerName string, options *TriggerClientBeginDeleteTriggerOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteTriggerCreateRequest(ctx, triggerName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteTriggerCreateRequest creates the DeleteTrigger request.
func (client *TriggerClient) deleteTriggerCreateRequest(ctx context.Context, triggerName string, options *TriggerClientBeginDeleteTriggerOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetEventSubscriptionStatus - Get a trigger's event subscription status.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - triggerName - The trigger name.
//   - options - TriggerClientGetEventSubscriptionStatusOptions contains the optional parameters for the TriggerClient.GetEventSubscriptionStatus
//     method.
func (client *TriggerClient) GetEventSubscriptionStatus(ctx context.Context, triggerName string, options *TriggerClientGetEventSubscriptionStatusOptions) (TriggerClientGetEventSubscriptionStatusResponse, error) {
	var err error
	req, err := client.getEventSubscriptionStatusCreateRequest(ctx, triggerName, options)
	if err != nil {
		return TriggerClientGetEventSubscriptionStatusResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TriggerClientGetEventSubscriptionStatusResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TriggerClientGetEventSubscriptionStatusResponse{}, err
	}
	resp, err := client.getEventSubscriptionStatusHandleResponse(httpResp)
	return resp, err
}

// getEventSubscriptionStatusCreateRequest creates the GetEventSubscriptionStatus request.
func (client *TriggerClient) getEventSubscriptionStatusCreateRequest(ctx context.Context, triggerName string, options *TriggerClientGetEventSubscriptionStatusOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}/getEventSubscriptionStatus"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEventSubscriptionStatusHandleResponse handles the GetEventSubscriptionStatus response.
func (client *TriggerClient) getEventSubscriptionStatusHandleResponse(resp *http.Response) (TriggerClientGetEventSubscriptionStatusResponse, error) {
	result := TriggerClientGetEventSubscriptionStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TriggerSubscriptionOperationStatus); err != nil {
		return TriggerClientGetEventSubscriptionStatusResponse{}, err
	}
	return result, nil
}

// GetTrigger - Gets a trigger.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - triggerName - The trigger name.
//   - options - TriggerClientGetTriggerOptions contains the optional parameters for the TriggerClient.GetTrigger method.
func (client *TriggerClient) GetTrigger(ctx context.Context, triggerName string, options *TriggerClientGetTriggerOptions) (TriggerClientGetTriggerResponse, error) {
	var err error
	req, err := client.getTriggerCreateRequest(ctx, triggerName, options)
	if err != nil {
		return TriggerClientGetTriggerResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TriggerClientGetTriggerResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNotModified) {
		err = runtime.NewResponseError(httpResp)
		return TriggerClientGetTriggerResponse{}, err
	}
	resp, err := client.getTriggerHandleResponse(httpResp)
	return resp, err
}

// getTriggerCreateRequest creates the GetTrigger request.
func (client *TriggerClient) getTriggerCreateRequest(ctx context.Context, triggerName string, options *TriggerClientGetTriggerOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getTriggerHandleResponse handles the GetTrigger response.
func (client *TriggerClient) getTriggerHandleResponse(resp *http.Response) (TriggerClientGetTriggerResponse, error) {
	result := TriggerClientGetTriggerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TriggerResource); err != nil {
		return TriggerClientGetTriggerResponse{}, err
	}
	return result, nil
}

// NewGetTriggersByWorkspacePager - Lists triggers.
//
// Generated from API version 2020-12-01
//   - options - TriggerClientGetTriggersByWorkspaceOptions contains the optional parameters for the TriggerClient.NewGetTriggersByWorkspacePager
//     method.
func (client *TriggerClient) NewGetTriggersByWorkspacePager(options *TriggerClientGetTriggersByWorkspaceOptions) *runtime.Pager[TriggerClientGetTriggersByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[TriggerClientGetTriggersByWorkspaceResponse]{
		More: func(page TriggerClientGetTriggersByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TriggerClientGetTriggersByWorkspaceResponse) (TriggerClientGetTriggersByWorkspaceResponse, error) {
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.getTriggersByWorkspaceCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return TriggerClientGetTriggersByWorkspaceResponse{}, err
			}
			return client.getTriggersByWorkspaceHandleResponse(resp)
		},
	})
}

// getTriggersByWorkspaceCreateRequest creates the GetTriggersByWorkspace request.
func (client *TriggerClient) getTriggersByWorkspaceCreateRequest(ctx context.Context, options *TriggerClientGetTriggersByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/triggers"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getTriggersByWorkspaceHandleResponse handles the GetTriggersByWorkspace response.
func (client *TriggerClient) getTriggersByWorkspaceHandleResponse(resp *http.Response) (TriggerClientGetTriggersByWorkspaceResponse, error) {
	result := TriggerClientGetTriggersByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TriggerListResponse); err != nil {
		return TriggerClientGetTriggersByWorkspaceResponse{}, err
	}
	return result, nil
}

// BeginStartTrigger - Starts a trigger.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - triggerName - The trigger name.
//   - options - TriggerClientBeginStartTriggerOptions contains the optional parameters for the TriggerClient.BeginStartTrigger
//     method.
func (client *TriggerClient) BeginStartTrigger(ctx context.Context, triggerName string, options *TriggerClientBeginStartTriggerOptions) (*runtime.Poller[TriggerClientStartTriggerResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.startTrigger(ctx, triggerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[TriggerClientStartTriggerResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[TriggerClientStartTriggerResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// StartTrigger - Starts a trigger.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
func (client *TriggerClient) startTrigger(ctx context.Context, triggerName string, options *TriggerClientBeginStartTriggerOptions) (*http.Response, error) {
	var err error
	req, err := client.startTriggerCreateRequest(ctx, triggerName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// startTriggerCreateRequest creates the StartTrigger request.
func (client *TriggerClient) startTriggerCreateRequest(ctx context.Context, triggerName string, options *TriggerClientBeginStartTriggerOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}/start"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginStopTrigger - Stops a trigger.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - triggerName - The trigger name.
//   - options - TriggerClientBeginStopTriggerOptions contains the optional parameters for the TriggerClient.BeginStopTrigger
//     method.
func (client *TriggerClient) BeginStopTrigger(ctx context.Context, triggerName string, options *TriggerClientBeginStopTriggerOptions) (*runtime.Poller[TriggerClientStopTriggerResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.stopTrigger(ctx, triggerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[TriggerClientStopTriggerResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[TriggerClientStopTriggerResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// StopTrigger - Stops a trigger.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
func (client *TriggerClient) stopTrigger(ctx context.Context, triggerName string, options *TriggerClientBeginStopTriggerOptions) (*http.Response, error) {
	var err error
	req, err := client.stopTriggerCreateRequest(ctx, triggerName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// stopTriggerCreateRequest creates the StopTrigger request.
func (client *TriggerClient) stopTriggerCreateRequest(ctx context.Context, triggerName string, options *TriggerClientBeginStopTriggerOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}/stop"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginSubscribeTriggerToEvents - Subscribe event trigger to events.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - triggerName - The trigger name.
//   - options - TriggerClientBeginSubscribeTriggerToEventsOptions contains the optional parameters for the TriggerClient.BeginSubscribeTriggerToEvents
//     method.
func (client *TriggerClient) BeginSubscribeTriggerToEvents(ctx context.Context, triggerName string, options *TriggerClientBeginSubscribeTriggerToEventsOptions) (*runtime.Poller[TriggerClientSubscribeTriggerToEventsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.subscribeTriggerToEvents(ctx, triggerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[TriggerClientSubscribeTriggerToEventsResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[TriggerClientSubscribeTriggerToEventsResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// SubscribeTriggerToEvents - Subscribe event trigger to events.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
func (client *TriggerClient) subscribeTriggerToEvents(ctx context.Context, triggerName string, options *TriggerClientBeginSubscribeTriggerToEventsOptions) (*http.Response, error) {
	var err error
	req, err := client.subscribeTriggerToEventsCreateRequest(ctx, triggerName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// subscribeTriggerToEventsCreateRequest creates the SubscribeTriggerToEvents request.
func (client *TriggerClient) subscribeTriggerToEventsCreateRequest(ctx context.Context, triggerName string, options *TriggerClientBeginSubscribeTriggerToEventsOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}/subscribeToEvents"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginUnsubscribeTriggerFromEvents - Unsubscribe event trigger from events.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - triggerName - The trigger name.
//   - options - TriggerClientBeginUnsubscribeTriggerFromEventsOptions contains the optional parameters for the TriggerClient.BeginUnsubscribeTriggerFromEvents
//     method.
func (client *TriggerClient) BeginUnsubscribeTriggerFromEvents(ctx context.Context, triggerName string, options *TriggerClientBeginUnsubscribeTriggerFromEventsOptions) (*runtime.Poller[TriggerClientUnsubscribeTriggerFromEventsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.unsubscribeTriggerFromEvents(ctx, triggerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[TriggerClientUnsubscribeTriggerFromEventsResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[TriggerClientUnsubscribeTriggerFromEventsResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// UnsubscribeTriggerFromEvents - Unsubscribe event trigger from events.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
func (client *TriggerClient) unsubscribeTriggerFromEvents(ctx context.Context, triggerName string, options *TriggerClientBeginUnsubscribeTriggerFromEventsOptions) (*http.Response, error) {
	var err error
	req, err := client.unsubscribeTriggerFromEventsCreateRequest(ctx, triggerName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// unsubscribeTriggerFromEventsCreateRequest creates the UnsubscribeTriggerFromEvents request.
func (client *TriggerClient) unsubscribeTriggerFromEventsCreateRequest(ctx context.Context, triggerName string, options *TriggerClientBeginUnsubscribeTriggerFromEventsOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}/unsubscribeFromEvents"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
