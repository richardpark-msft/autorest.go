//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"azartifacts"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"net/http"
	"regexp"
)

// TriggerServer is a fake server for instances of the azartifacts.TriggerClient type.
type TriggerServer struct {
	// BeginCreateOrUpdateTrigger is the fake for method TriggerClient.BeginCreateOrUpdateTrigger
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdateTrigger func(ctx context.Context, triggerName string, trigger azartifacts.TriggerResource, options *azartifacts.TriggerClientBeginCreateOrUpdateTriggerOptions) (resp azfake.PollerResponder[azartifacts.TriggerClientCreateOrUpdateTriggerResponse], errResp azfake.ErrorResponder)

	// BeginDeleteTrigger is the fake for method TriggerClient.BeginDeleteTrigger
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDeleteTrigger func(ctx context.Context, triggerName string, options *azartifacts.TriggerClientBeginDeleteTriggerOptions) (resp azfake.PollerResponder[azartifacts.TriggerClientDeleteTriggerResponse], errResp azfake.ErrorResponder)

	// GetEventSubscriptionStatus is the fake for method TriggerClient.GetEventSubscriptionStatus
	// HTTP status codes to indicate success: http.StatusOK
	GetEventSubscriptionStatus func(ctx context.Context, triggerName string, options *azartifacts.TriggerClientGetEventSubscriptionStatusOptions) (resp azfake.Responder[azartifacts.TriggerClientGetEventSubscriptionStatusResponse], errResp azfake.ErrorResponder)

	// GetTrigger is the fake for method TriggerClient.GetTrigger
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNotModified
	GetTrigger func(ctx context.Context, triggerName string, options *azartifacts.TriggerClientGetTriggerOptions) (resp azfake.Responder[azartifacts.TriggerClientGetTriggerResponse], errResp azfake.ErrorResponder)

	// NewGetTriggersByWorkspacePager is the fake for method TriggerClient.NewGetTriggersByWorkspacePager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetTriggersByWorkspacePager func(options *azartifacts.TriggerClientGetTriggersByWorkspaceOptions) (resp azfake.PagerResponder[azartifacts.TriggerClientGetTriggersByWorkspaceResponse])

	// BeginStartTrigger is the fake for method TriggerClient.BeginStartTrigger
	// HTTP status codes to indicate success: http.StatusOK
	BeginStartTrigger func(ctx context.Context, triggerName string, options *azartifacts.TriggerClientBeginStartTriggerOptions) (resp azfake.PollerResponder[azartifacts.TriggerClientStartTriggerResponse], errResp azfake.ErrorResponder)

	// BeginStopTrigger is the fake for method TriggerClient.BeginStopTrigger
	// HTTP status codes to indicate success: http.StatusOK
	BeginStopTrigger func(ctx context.Context, triggerName string, options *azartifacts.TriggerClientBeginStopTriggerOptions) (resp azfake.PollerResponder[azartifacts.TriggerClientStopTriggerResponse], errResp azfake.ErrorResponder)

	// BeginSubscribeTriggerToEvents is the fake for method TriggerClient.BeginSubscribeTriggerToEvents
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginSubscribeTriggerToEvents func(ctx context.Context, triggerName string, options *azartifacts.TriggerClientBeginSubscribeTriggerToEventsOptions) (resp azfake.PollerResponder[azartifacts.TriggerClientSubscribeTriggerToEventsResponse], errResp azfake.ErrorResponder)

	// BeginUnsubscribeTriggerFromEvents is the fake for method TriggerClient.BeginUnsubscribeTriggerFromEvents
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUnsubscribeTriggerFromEvents func(ctx context.Context, triggerName string, options *azartifacts.TriggerClientBeginUnsubscribeTriggerFromEventsOptions) (resp azfake.PollerResponder[azartifacts.TriggerClientUnsubscribeTriggerFromEventsResponse], errResp azfake.ErrorResponder)
}

// NewTriggerServerTransport creates a new instance of TriggerServerTransport with the provided implementation.
// The returned TriggerServerTransport instance is connected to an instance of azartifacts.TriggerClient by way of the
// undefined.Transporter field.
func NewTriggerServerTransport(srv *TriggerServer) *TriggerServerTransport {
	return &TriggerServerTransport{srv: srv}
}

// TriggerServerTransport connects instances of azartifacts.TriggerClient to instances of TriggerServer.
// Don't use this type directly, use NewTriggerServerTransport instead.
type TriggerServerTransport struct {
	srv                               *TriggerServer
	beginCreateOrUpdateTrigger        *azfake.PollerResponder[azartifacts.TriggerClientCreateOrUpdateTriggerResponse]
	beginDeleteTrigger                *azfake.PollerResponder[azartifacts.TriggerClientDeleteTriggerResponse]
	newGetTriggersByWorkspacePager    *azfake.PagerResponder[azartifacts.TriggerClientGetTriggersByWorkspaceResponse]
	beginStartTrigger                 *azfake.PollerResponder[azartifacts.TriggerClientStartTriggerResponse]
	beginStopTrigger                  *azfake.PollerResponder[azartifacts.TriggerClientStopTriggerResponse]
	beginSubscribeTriggerToEvents     *azfake.PollerResponder[azartifacts.TriggerClientSubscribeTriggerToEventsResponse]
	beginUnsubscribeTriggerFromEvents *azfake.PollerResponder[azartifacts.TriggerClientUnsubscribeTriggerFromEventsResponse]
}

// Do implements the policy.Transporter interface for TriggerServerTransport.
func (t *TriggerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "TriggerClient.BeginCreateOrUpdateTrigger":
		resp, err = t.dispatchBeginCreateOrUpdateTrigger(req)
	case "TriggerClient.BeginDeleteTrigger":
		resp, err = t.dispatchBeginDeleteTrigger(req)
	case "TriggerClient.GetEventSubscriptionStatus":
		resp, err = t.dispatchGetEventSubscriptionStatus(req)
	case "TriggerClient.GetTrigger":
		resp, err = t.dispatchGetTrigger(req)
	case "TriggerClient.NewGetTriggersByWorkspacePager":
		resp, err = t.dispatchNewGetTriggersByWorkspacePager(req)
	case "TriggerClient.BeginStartTrigger":
		resp, err = t.dispatchBeginStartTrigger(req)
	case "TriggerClient.BeginStopTrigger":
		resp, err = t.dispatchBeginStopTrigger(req)
	case "TriggerClient.BeginSubscribeTriggerToEvents":
		resp, err = t.dispatchBeginSubscribeTriggerToEvents(req)
	case "TriggerClient.BeginUnsubscribeTriggerFromEvents":
		resp, err = t.dispatchBeginUnsubscribeTriggerFromEvents(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *TriggerServerTransport) dispatchBeginCreateOrUpdateTrigger(req *http.Request) (*http.Response, error) {
	if t.srv.BeginCreateOrUpdateTrigger == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdateTrigger not implemented")}
	}
	if t.beginCreateOrUpdateTrigger == nil {
		const regexStr = "/triggers/(?P<triggerName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[azartifacts.TriggerResource](req)
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		var options *azartifacts.TriggerClientBeginCreateOrUpdateTriggerOptions
		if ifMatchParam != nil {
			options = &azartifacts.TriggerClientBeginCreateOrUpdateTriggerOptions{
				IfMatch: ifMatchParam,
			}
		}
		respr, errRespr := t.srv.BeginCreateOrUpdateTrigger(req.Context(), matches[regex.SubexpIndex("triggerName")], body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		t.beginCreateOrUpdateTrigger = &respr
	}

	resp, err := server.PollerResponderNext(t.beginCreateOrUpdateTrigger, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(t.beginCreateOrUpdateTrigger) {
		t.beginCreateOrUpdateTrigger = nil
	}

	return resp, nil
}

func (t *TriggerServerTransport) dispatchBeginDeleteTrigger(req *http.Request) (*http.Response, error) {
	if t.srv.BeginDeleteTrigger == nil {
		return nil, &nonRetriableError{errors.New("method BeginDeleteTrigger not implemented")}
	}
	if t.beginDeleteTrigger == nil {
		const regexStr = "/triggers/(?P<triggerName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := t.srv.BeginDeleteTrigger(req.Context(), matches[regex.SubexpIndex("triggerName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		t.beginDeleteTrigger = &respr
	}

	resp, err := server.PollerResponderNext(t.beginDeleteTrigger, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(t.beginDeleteTrigger) {
		t.beginDeleteTrigger = nil
	}

	return resp, nil
}

func (t *TriggerServerTransport) dispatchGetEventSubscriptionStatus(req *http.Request) (*http.Response, error) {
	if t.srv.GetEventSubscriptionStatus == nil {
		return nil, &nonRetriableError{errors.New("method GetEventSubscriptionStatus not implemented")}
	}
	const regexStr = "/triggers/(?P<triggerName>[a-zA-Z0-9-_]+)/getEventSubscriptionStatus"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := t.srv.GetEventSubscriptionStatus(req.Context(), matches[regex.SubexpIndex("triggerName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TriggerSubscriptionOperationStatus, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TriggerServerTransport) dispatchGetTrigger(req *http.Request) (*http.Response, error) {
	if t.srv.GetTrigger == nil {
		return nil, &nonRetriableError{errors.New("method GetTrigger not implemented")}
	}
	const regexStr = "/triggers/(?P<triggerName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
	var options *azartifacts.TriggerClientGetTriggerOptions
	if ifNoneMatchParam != nil {
		options = &azartifacts.TriggerClientGetTriggerOptions{
			IfNoneMatch: ifNoneMatchParam,
		}
	}
	respr, errRespr := t.srv.GetTrigger(req.Context(), matches[regex.SubexpIndex("triggerName")], options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNotModified}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNotModified", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TriggerResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TriggerServerTransport) dispatchNewGetTriggersByWorkspacePager(req *http.Request) (*http.Response, error) {
	if t.srv.NewGetTriggersByWorkspacePager == nil {
		return nil, &nonRetriableError{errors.New("method NewGetTriggersByWorkspacePager not implemented")}
	}
	if t.newGetTriggersByWorkspacePager == nil {
		resp := t.srv.NewGetTriggersByWorkspacePager(nil)
		t.newGetTriggersByWorkspacePager = &resp
		server.PagerResponderInjectNextLinks(t.newGetTriggersByWorkspacePager, req, func(page *azartifacts.TriggerClientGetTriggersByWorkspaceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(t.newGetTriggersByWorkspacePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(t.newGetTriggersByWorkspacePager) {
		t.newGetTriggersByWorkspacePager = nil
	}
	return resp, nil
}

func (t *TriggerServerTransport) dispatchBeginStartTrigger(req *http.Request) (*http.Response, error) {
	if t.srv.BeginStartTrigger == nil {
		return nil, &nonRetriableError{errors.New("method BeginStartTrigger not implemented")}
	}
	if t.beginStartTrigger == nil {
		const regexStr = "/triggers/(?P<triggerName>[a-zA-Z0-9-_]+)/start"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := t.srv.BeginStartTrigger(req.Context(), matches[regex.SubexpIndex("triggerName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		t.beginStartTrigger = &respr
	}

	resp, err := server.PollerResponderNext(t.beginStartTrigger, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(t.beginStartTrigger) {
		t.beginStartTrigger = nil
	}

	return resp, nil
}

func (t *TriggerServerTransport) dispatchBeginStopTrigger(req *http.Request) (*http.Response, error) {
	if t.srv.BeginStopTrigger == nil {
		return nil, &nonRetriableError{errors.New("method BeginStopTrigger not implemented")}
	}
	if t.beginStopTrigger == nil {
		const regexStr = "/triggers/(?P<triggerName>[a-zA-Z0-9-_]+)/stop"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := t.srv.BeginStopTrigger(req.Context(), matches[regex.SubexpIndex("triggerName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		t.beginStopTrigger = &respr
	}

	resp, err := server.PollerResponderNext(t.beginStopTrigger, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(t.beginStopTrigger) {
		t.beginStopTrigger = nil
	}

	return resp, nil
}

func (t *TriggerServerTransport) dispatchBeginSubscribeTriggerToEvents(req *http.Request) (*http.Response, error) {
	if t.srv.BeginSubscribeTriggerToEvents == nil {
		return nil, &nonRetriableError{errors.New("method BeginSubscribeTriggerToEvents not implemented")}
	}
	if t.beginSubscribeTriggerToEvents == nil {
		const regexStr = "/triggers/(?P<triggerName>[a-zA-Z0-9-_]+)/subscribeToEvents"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := t.srv.BeginSubscribeTriggerToEvents(req.Context(), matches[regex.SubexpIndex("triggerName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		t.beginSubscribeTriggerToEvents = &respr
	}

	resp, err := server.PollerResponderNext(t.beginSubscribeTriggerToEvents, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(t.beginSubscribeTriggerToEvents) {
		t.beginSubscribeTriggerToEvents = nil
	}

	return resp, nil
}

func (t *TriggerServerTransport) dispatchBeginUnsubscribeTriggerFromEvents(req *http.Request) (*http.Response, error) {
	if t.srv.BeginUnsubscribeTriggerFromEvents == nil {
		return nil, &nonRetriableError{errors.New("method BeginUnsubscribeTriggerFromEvents not implemented")}
	}
	if t.beginUnsubscribeTriggerFromEvents == nil {
		const regexStr = "/triggers/(?P<triggerName>[a-zA-Z0-9-_]+)/unsubscribeFromEvents"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := t.srv.BeginUnsubscribeTriggerFromEvents(req.Context(), matches[regex.SubexpIndex("triggerName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		t.beginUnsubscribeTriggerFromEvents = &respr
	}

	resp, err := server.PollerResponderNext(t.beginUnsubscribeTriggerFromEvents, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(t.beginUnsubscribeTriggerFromEvents) {
		t.beginUnsubscribeTriggerFromEvents = nil
	}

	return resp, nil
}
