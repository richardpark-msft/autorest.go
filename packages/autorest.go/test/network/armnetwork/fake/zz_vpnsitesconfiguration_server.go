//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"armnetwork"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"regexp"
)

// VPNSitesConfigurationServer is a fake server for instances of the armnetwork.VPNSitesConfigurationClient type.
type VPNSitesConfigurationServer struct {
	// BeginDownload is the fake for method VPNSitesConfigurationClient.BeginDownload
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginDownload func(ctx context.Context, resourceGroupName string, virtualWANName string, request armnetwork.GetVPNSitesConfigurationRequest, options *armnetwork.VPNSitesConfigurationClientBeginDownloadOptions) (resp azfake.PollerResponder[armnetwork.VPNSitesConfigurationClientDownloadResponse], errResp azfake.ErrorResponder)
}

// NewVPNSitesConfigurationServerTransport creates a new instance of VPNSitesConfigurationServerTransport with the provided implementation.
// The returned VPNSitesConfigurationServerTransport instance is connected to an instance of armnetwork.VPNSitesConfigurationClient by way of the
// undefined.Transporter field.
func NewVPNSitesConfigurationServerTransport(srv *VPNSitesConfigurationServer) *VPNSitesConfigurationServerTransport {
	return &VPNSitesConfigurationServerTransport{srv: srv}
}

// VPNSitesConfigurationServerTransport connects instances of armnetwork.VPNSitesConfigurationClient to instances of VPNSitesConfigurationServer.
// Don't use this type directly, use NewVPNSitesConfigurationServerTransport instead.
type VPNSitesConfigurationServerTransport struct {
	srv           *VPNSitesConfigurationServer
	beginDownload *azfake.PollerResponder[armnetwork.VPNSitesConfigurationClientDownloadResponse]
}

// Do implements the policy.Transporter interface for VPNSitesConfigurationServerTransport.
func (v *VPNSitesConfigurationServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VPNSitesConfigurationClient.BeginDownload":
		resp, err = v.dispatchBeginDownload(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VPNSitesConfigurationServerTransport) dispatchBeginDownload(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDownload == nil {
		return nil, &nonRetriableError{errors.New("method BeginDownload not implemented")}
	}
	if v.beginDownload == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/virtualWans/(?P<virtualWANName>[a-zA-Z0-9-_]+)/vpnConfiguration"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.GetVPNSitesConfigurationRequest](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginDownload(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("virtualWANName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		v.beginDownload = &respr
	}

	resp, err := server.PollerResponderNext(v.beginDownload, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(v.beginDownload) {
		v.beginDownload = nil
	}

	return resp, nil
}
