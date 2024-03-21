// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"jmergepatchgroup"
	"net/http"
	"reflect"
)

// JsonMergePatchServer is a fake server for instances of the jmergepatchgroup.JsonMergePatchClient type.
type JsonMergePatchServer struct {
	// CreateResource is the fake for method JsonMergePatchClient.CreateResource
	// HTTP status codes to indicate success: http.StatusOK
	CreateResource func(ctx context.Context, body jmergepatchgroup.Resource, options *jmergepatchgroup.JsonMergePatchClientCreateResourceOptions) (resp azfake.Responder[jmergepatchgroup.JsonMergePatchClientCreateResourceResponse], errResp azfake.ErrorResponder)

	// UpdateOptionalResource is the fake for method JsonMergePatchClient.UpdateOptionalResource
	// HTTP status codes to indicate success: http.StatusOK
	UpdateOptionalResource func(ctx context.Context, options *jmergepatchgroup.JsonMergePatchClientUpdateOptionalResourceOptions) (resp azfake.Responder[jmergepatchgroup.JsonMergePatchClientUpdateOptionalResourceResponse], errResp azfake.ErrorResponder)

	// UpdateResource is the fake for method JsonMergePatchClient.UpdateResource
	// HTTP status codes to indicate success: http.StatusOK
	UpdateResource func(ctx context.Context, body jmergepatchgroup.ResourcePatch, options *jmergepatchgroup.JsonMergePatchClientUpdateResourceOptions) (resp azfake.Responder[jmergepatchgroup.JsonMergePatchClientUpdateResourceResponse], errResp azfake.ErrorResponder)
}

// NewJsonMergePatchServerTransport creates a new instance of JsonMergePatchServerTransport with the provided implementation.
// The returned JsonMergePatchServerTransport instance is connected to an instance of jmergepatchgroup.JsonMergePatchClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewJsonMergePatchServerTransport(srv *JsonMergePatchServer) *JsonMergePatchServerTransport {
	return &JsonMergePatchServerTransport{srv: srv}
}

// JsonMergePatchServerTransport connects instances of jmergepatchgroup.JsonMergePatchClient to instances of JsonMergePatchServer.
// Don't use this type directly, use NewJsonMergePatchServerTransport instead.
type JsonMergePatchServerTransport struct {
	srv *JsonMergePatchServer
}

// Do implements the policy.Transporter interface for JsonMergePatchServerTransport.
func (j *JsonMergePatchServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "JsonMergePatchClient.CreateResource":
		resp, err = j.dispatchCreateResource(req)
	case "JsonMergePatchClient.UpdateOptionalResource":
		resp, err = j.dispatchUpdateOptionalResource(req)
	case "JsonMergePatchClient.UpdateResource":
		resp, err = j.dispatchUpdateResource(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (j *JsonMergePatchServerTransport) dispatchCreateResource(req *http.Request) (*http.Response, error) {
	if j.srv.CreateResource == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateResource not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[jmergepatchgroup.Resource](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := j.srv.CreateResource(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Resource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (j *JsonMergePatchServerTransport) dispatchUpdateOptionalResource(req *http.Request) (*http.Response, error) {
	if j.srv.UpdateOptionalResource == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateOptionalResource not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[jmergepatchgroup.ResourcePatch](req)
	if err != nil {
		return nil, err
	}
	var options *jmergepatchgroup.JsonMergePatchClientUpdateOptionalResourceOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &jmergepatchgroup.JsonMergePatchClientUpdateOptionalResourceOptions{
			Body: &body,
		}
	}
	respr, errRespr := j.srv.UpdateOptionalResource(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Resource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (j *JsonMergePatchServerTransport) dispatchUpdateResource(req *http.Request) (*http.Response, error) {
	if j.srv.UpdateResource == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateResource not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[jmergepatchgroup.ResourcePatch](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := j.srv.UpdateResource(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Resource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
