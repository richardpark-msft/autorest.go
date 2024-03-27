// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"basicgroup"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

// BasicServer is a fake server for instances of the basicgroup.BasicClient type.
type BasicServer struct {
	// TwoModelsAsPageItemServer contains the fakes for client TwoModelsAsPageItemClient
	TwoModelsAsPageItemServer TwoModelsAsPageItemServer

	// CreateOrReplace is the fake for method BasicClient.CreateOrReplace
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrReplace func(ctx context.Context, id int32, resource basicgroup.User, options *basicgroup.BasicClientCreateOrReplaceOptions) (resp azfake.Responder[basicgroup.BasicClientCreateOrReplaceResponse], errResp azfake.ErrorResponder)

	// CreateOrUpdate is the fake for method BasicClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, id int32, resource basicgroup.User, options *basicgroup.BasicClientCreateOrUpdateOptions) (resp azfake.Responder[basicgroup.BasicClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method BasicClient.Delete
	// HTTP status codes to indicate success: http.StatusNoContent
	Delete func(ctx context.Context, id int32, options *basicgroup.BasicClientDeleteOptions) (resp azfake.Responder[basicgroup.BasicClientDeleteResponse], errResp azfake.ErrorResponder)

	// Export is the fake for method BasicClient.Export
	// HTTP status codes to indicate success: http.StatusOK
	Export func(ctx context.Context, id int32, formatParam string, options *basicgroup.BasicClientExportOptions) (resp azfake.Responder[basicgroup.BasicClientExportResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method BasicClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, id int32, options *basicgroup.BasicClientGetOptions) (resp azfake.Responder[basicgroup.BasicClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method BasicClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *basicgroup.BasicClientListOptions) (resp azfake.PagerResponder[basicgroup.BasicClientListResponse])

	// NewListWithCustomPageModelPager is the fake for method BasicClient.NewListWithCustomPageModelPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListWithCustomPageModelPager func(options *basicgroup.BasicClientListWithCustomPageModelOptions) (resp azfake.PagerResponder[basicgroup.BasicClientListWithCustomPageModelResponse])

	// NewListWithPagePager is the fake for method BasicClient.NewListWithPagePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListWithPagePager func(options *basicgroup.BasicClientListWithPageOptions) (resp azfake.PagerResponder[basicgroup.BasicClientListWithPageResponse])

	// NewListWithParametersPager is the fake for method BasicClient.NewListWithParametersPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListWithParametersPager func(bodyInput basicgroup.ListItemInputBody, options *basicgroup.BasicClientListWithParametersOptions) (resp azfake.PagerResponder[basicgroup.BasicClientListWithParametersResponse])
}

// NewBasicServerTransport creates a new instance of BasicServerTransport with the provided implementation.
// The returned BasicServerTransport instance is connected to an instance of basicgroup.BasicClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBasicServerTransport(srv *BasicServer) *BasicServerTransport {
	return &BasicServerTransport{
		srv:                             srv,
		newListPager:                    newTracker[azfake.PagerResponder[basicgroup.BasicClientListResponse]](),
		newListWithCustomPageModelPager: newTracker[azfake.PagerResponder[basicgroup.BasicClientListWithCustomPageModelResponse]](),
		newListWithPagePager:            newTracker[azfake.PagerResponder[basicgroup.BasicClientListWithPageResponse]](),
		newListWithParametersPager:      newTracker[azfake.PagerResponder[basicgroup.BasicClientListWithParametersResponse]](),
	}
}

// BasicServerTransport connects instances of basicgroup.BasicClient to instances of BasicServer.
// Don't use this type directly, use NewBasicServerTransport instead.
type BasicServerTransport struct {
	srv                             *BasicServer
	trMu                            sync.Mutex
	trTwoModelsAsPageItemServer     *TwoModelsAsPageItemServerTransport
	newListPager                    *tracker[azfake.PagerResponder[basicgroup.BasicClientListResponse]]
	newListWithCustomPageModelPager *tracker[azfake.PagerResponder[basicgroup.BasicClientListWithCustomPageModelResponse]]
	newListWithPagePager            *tracker[azfake.PagerResponder[basicgroup.BasicClientListWithPageResponse]]
	newListWithParametersPager      *tracker[azfake.PagerResponder[basicgroup.BasicClientListWithParametersResponse]]
}

// Do implements the policy.Transporter interface for BasicServerTransport.
func (b *BasicServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	if client := method[:strings.Index(method, ".")]; client != "BasicClient" {
		return b.dispatchToClientFake(req, client)
	}
	return b.dispatchToMethodFake(req, method)
}

func (b *BasicServerTransport) dispatchToClientFake(req *http.Request, client string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch client {
	case "TwoModelsAsPageItemClient":
		initServer(&b.trMu, &b.trTwoModelsAsPageItemServer, func() *TwoModelsAsPageItemServerTransport {
			return NewTwoModelsAsPageItemServerTransport(&b.srv.TwoModelsAsPageItemServer)
		})
		resp, err = b.trTwoModelsAsPageItemServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	return resp, err
}

func (b *BasicServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "BasicClient.CreateOrReplace":
		resp, err = b.dispatchCreateOrReplace(req)
	case "BasicClient.CreateOrUpdate":
		resp, err = b.dispatchCreateOrUpdate(req)
	case "BasicClient.Delete":
		resp, err = b.dispatchDelete(req)
	case "BasicClient.Export":
		resp, err = b.dispatchExport(req)
	case "BasicClient.Get":
		resp, err = b.dispatchGet(req)
	case "BasicClient.NewListPager":
		resp, err = b.dispatchNewListPager(req)
	case "BasicClient.NewListWithCustomPageModelPager":
		resp, err = b.dispatchNewListWithCustomPageModelPager(req)
	case "BasicClient.NewListWithPagePager":
		resp, err = b.dispatchNewListWithPagePager(req)
	case "BasicClient.NewListWithParametersPager":
		resp, err = b.dispatchNewListWithParametersPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (b *BasicServerTransport) dispatchCreateOrReplace(req *http.Request) (*http.Response, error) {
	if b.srv.CreateOrReplace == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrReplace not implemented")}
	}
	const regexStr = `/azure/core/basic/users/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[basicgroup.User](req)
	if err != nil {
		return nil, err
	}
	idUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
	if err != nil {
		return nil, err
	}
	idParam, err := parseWithCast(idUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.CreateOrReplace(req.Context(), idParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).User, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BasicServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if b.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/azure/core/basic/users/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[basicgroup.User](req)
	if err != nil {
		return nil, err
	}
	idUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
	if err != nil {
		return nil, err
	}
	idParam, err := parseWithCast(idUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.CreateOrUpdate(req.Context(), idParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).User, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BasicServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if b.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/azure/core/basic/users/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	idUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
	if err != nil {
		return nil, err
	}
	idParam, err := parseWithCast(idUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Delete(req.Context(), idParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BasicServerTransport) dispatchExport(req *http.Request) (*http.Response, error) {
	if b.srv.Export == nil {
		return nil, &nonRetriableError{errors.New("fake for method Export not implemented")}
	}
	const regexStr = `/azure/core/basic/users/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+):export`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	idUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
	if err != nil {
		return nil, err
	}
	idParam, err := parseWithCast(idUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	formatParamParam, err := url.QueryUnescape(qp.Get("format"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Export(req.Context(), idParam, formatParamParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).User, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BasicServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if b.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/azure/core/basic/users/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	idUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
	if err != nil {
		return nil, err
	}
	idParam, err := parseWithCast(idUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Get(req.Context(), idParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).User, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BasicServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := b.newListPager.get(req)
	if newListPager == nil {
		qp := req.URL.Query()
		topUnescaped, err := url.QueryUnescape(qp.Get("top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		maxpagesizeUnescaped, err := url.QueryUnescape(qp.Get("maxpagesize"))
		if err != nil {
			return nil, err
		}
		maxpagesizeParam, err := parseOptional(maxpagesizeUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		orderbyEscaped := qp["orderby"]
		orderbyParam := make([]string, len(orderbyEscaped))
		for i, v := range orderbyEscaped {
			u, unescapeErr := url.QueryUnescape(v)
			if unescapeErr != nil {
				return nil, unescapeErr
			}
			orderbyParam[i] = u
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		selectParamEscaped := qp["select"]
		selectParamParam := make([]string, len(selectParamEscaped))
		for i, v := range selectParamEscaped {
			u, unescapeErr := url.QueryUnescape(v)
			if unescapeErr != nil {
				return nil, unescapeErr
			}
			selectParamParam[i] = u
		}
		expandEscaped := qp["expand"]
		expandParam := make([]string, len(expandEscaped))
		for i, v := range expandEscaped {
			u, unescapeErr := url.QueryUnescape(v)
			if unescapeErr != nil {
				return nil, unescapeErr
			}
			expandParam[i] = u
		}
		var options *basicgroup.BasicClientListOptions
		if topParam != nil || skipParam != nil || maxpagesizeParam != nil || len(orderbyParam) > 0 || filterParam != nil || len(selectParamParam) > 0 || len(expandParam) > 0 {
			options = &basicgroup.BasicClientListOptions{
				Top:         topParam,
				Skip:        skipParam,
				Maxpagesize: maxpagesizeParam,
				Orderby:     orderbyParam,
				Filter:      filterParam,
				SelectParam: selectParamParam,
				Expand:      expandParam,
			}
		}
		resp := b.srv.NewListPager(options)
		newListPager = &resp
		b.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *basicgroup.BasicClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		b.newListPager.remove(req)
	}
	return resp, nil
}

func (b *BasicServerTransport) dispatchNewListWithCustomPageModelPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListWithCustomPageModelPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListWithCustomPageModelPager not implemented")}
	}
	newListWithCustomPageModelPager := b.newListWithCustomPageModelPager.get(req)
	if newListWithCustomPageModelPager == nil {
		resp := b.srv.NewListWithCustomPageModelPager(nil)
		newListWithCustomPageModelPager = &resp
		b.newListWithCustomPageModelPager.add(req, newListWithCustomPageModelPager)
		server.PagerResponderInjectNextLinks(newListWithCustomPageModelPager, req, func(page *basicgroup.BasicClientListWithCustomPageModelResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListWithCustomPageModelPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListWithCustomPageModelPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListWithCustomPageModelPager) {
		b.newListWithCustomPageModelPager.remove(req)
	}
	return resp, nil
}

func (b *BasicServerTransport) dispatchNewListWithPagePager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListWithPagePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListWithPagePager not implemented")}
	}
	newListWithPagePager := b.newListWithPagePager.get(req)
	if newListWithPagePager == nil {
		resp := b.srv.NewListWithPagePager(nil)
		newListWithPagePager = &resp
		b.newListWithPagePager.add(req, newListWithPagePager)
		server.PagerResponderInjectNextLinks(newListWithPagePager, req, func(page *basicgroup.BasicClientListWithPageResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListWithPagePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListWithPagePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListWithPagePager) {
		b.newListWithPagePager.remove(req)
	}
	return resp, nil
}

func (b *BasicServerTransport) dispatchNewListWithParametersPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListWithParametersPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListWithParametersPager not implemented")}
	}
	newListWithParametersPager := b.newListWithParametersPager.get(req)
	if newListWithParametersPager == nil {
		qp := req.URL.Query()
		body, err := server.UnmarshalRequestAsJSON[basicgroup.ListItemInputBody](req)
		if err != nil {
			return nil, err
		}
		anotherUnescaped, err := url.QueryUnescape(qp.Get("another"))
		if err != nil {
			return nil, err
		}
		anotherParam := getOptional(basicgroup.ListItemInputExtensibleEnum(anotherUnescaped))
		var options *basicgroup.BasicClientListWithParametersOptions
		if anotherParam != nil {
			options = &basicgroup.BasicClientListWithParametersOptions{
				Another: anotherParam,
			}
		}
		resp := b.srv.NewListWithParametersPager(body, options)
		newListWithParametersPager = &resp
		b.newListWithParametersPager.add(req, newListWithParametersPager)
		server.PagerResponderInjectNextLinks(newListWithParametersPager, req, func(page *basicgroup.BasicClientListWithParametersResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListWithParametersPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListWithParametersPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListWithParametersPager) {
		b.newListWithParametersPager.remove(req)
	}
	return resp, nil
}
