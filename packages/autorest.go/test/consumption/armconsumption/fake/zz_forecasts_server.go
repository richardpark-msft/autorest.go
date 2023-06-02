//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package fake

import (
	"armconsumption"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"regexp"
)

// ForecastsServer is a fake server for instances of the armconsumption.ForecastsClient type.
type ForecastsServer struct {
	// NewListPager is the fake for method ForecastsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armconsumption.ForecastsClientListOptions) (resp azfake.PagerResponder[armconsumption.ForecastsClientListResponse])
}

// NewForecastsServerTransport creates a new instance of ForecastsServerTransport with the provided implementation.
// The returned ForecastsServerTransport instance is connected to an instance of armconsumption.ForecastsClient by way of the
// undefined.Transporter field.
func NewForecastsServerTransport(srv *ForecastsServer) *ForecastsServerTransport {
	return &ForecastsServerTransport{srv: srv}
}

// ForecastsServerTransport connects instances of armconsumption.ForecastsClient to instances of ForecastsServer.
// Don't use this type directly, use NewForecastsServerTransport instead.
type ForecastsServerTransport struct {
	srv          *ForecastsServer
	newListPager *azfake.PagerResponder[armconsumption.ForecastsClientListResponse]
}

// Do implements the policy.Transporter interface for ForecastsServerTransport.
func (f *ForecastsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ForecastsClient.NewListPager":
		resp, err = f.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (f *ForecastsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if f.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Consumption/forecasts"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		filterParam := getOptional(qp.Get("$filter"))
		var options *armconsumption.ForecastsClientListOptions
		if filterParam != nil {
			options = &armconsumption.ForecastsClientListOptions{
				Filter: filterParam,
			}
		}
		resp := f.srv.NewListPager(options)
		f.newListPager = &resp
	}
	resp, err := server.PagerResponderNext(f.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(f.newListPager) {
		f.newListPager = nil
	}
	return resp, nil
}