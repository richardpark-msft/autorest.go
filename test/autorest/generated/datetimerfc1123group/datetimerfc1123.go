// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package datetimerfc1123group

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"path"
	"time"
)

// Datetimerfc1123Operations contains the methods for the Datetimerfc1123 group.
type Datetimerfc1123Operations interface {
	// GetInvalid - Get invalid datetime value
	GetInvalid(ctx context.Context) (*TimeResponse, error)
	// GetNull - Get null datetime value
	GetNull(ctx context.Context) (*TimeResponse, error)
	// GetOverflow - Get overflow datetime value
	GetOverflow(ctx context.Context) (*TimeResponse, error)
	// GetUTCLowercaseMaxDateTime - Get max datetime value fri, 31 dec 9999 23:59:59 gmt
	GetUTCLowercaseMaxDateTime(ctx context.Context) (*TimeResponse, error)
	// GetUTCMinDateTime - Get min datetime value Mon, 1 Jan 0001 00:00:00 GMT
	GetUTCMinDateTime(ctx context.Context) (*TimeResponse, error)
	// GetUTCUppercaseMaxDateTime - Get max datetime value FRI, 31 DEC 9999 23:59:59 GMT
	GetUTCUppercaseMaxDateTime(ctx context.Context) (*TimeResponse, error)
	// GetUnderflow - Get underflow datetime value
	GetUnderflow(ctx context.Context) (*TimeResponse, error)
	// PutUTCMaxDateTime - Put max datetime value Fri, 31 Dec 9999 23:59:59 GMT
	PutUTCMaxDateTime(ctx context.Context, datetimeBody time.Time) (*http.Response, error)
	// PutUTCMinDateTime - Put min datetime value Mon, 1 Jan 0001 00:00:00 GMT
	PutUTCMinDateTime(ctx context.Context, datetimeBody time.Time) (*http.Response, error)
}

// datetimerfc1123Operations implements the Datetimerfc1123Operations interface.
type datetimerfc1123Operations struct {
	*Client
}

// GetInvalid - Get invalid datetime value
func (client *datetimerfc1123Operations) GetInvalid(ctx context.Context) (*TimeResponse, error) {
	req, err := client.getInvalidCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getInvalidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getInvalidCreateRequest creates the GetInvalid request.
func (client *datetimerfc1123Operations) getInvalidCreateRequest() (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/invalid"
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getInvalidHandleResponse handles the GetInvalid response.
func (client *datetimerfc1123Operations) getInvalidHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getInvalidHandleError(resp)
	}
	var aux *timeRFC1123
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// getInvalidHandleError handles the GetInvalid error response.
func (client *datetimerfc1123Operations) getInvalidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetNull - Get null datetime value
func (client *datetimerfc1123Operations) GetNull(ctx context.Context) (*TimeResponse, error) {
	req, err := client.getNullCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getNullCreateRequest creates the GetNull request.
func (client *datetimerfc1123Operations) getNullCreateRequest() (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/null"
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client *datetimerfc1123Operations) getNullHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getNullHandleError(resp)
	}
	var aux *timeRFC1123
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// getNullHandleError handles the GetNull error response.
func (client *datetimerfc1123Operations) getNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetOverflow - Get overflow datetime value
func (client *datetimerfc1123Operations) GetOverflow(ctx context.Context) (*TimeResponse, error) {
	req, err := client.getOverflowCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getOverflowHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getOverflowCreateRequest creates the GetOverflow request.
func (client *datetimerfc1123Operations) getOverflowCreateRequest() (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/overflow"
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getOverflowHandleResponse handles the GetOverflow response.
func (client *datetimerfc1123Operations) getOverflowHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getOverflowHandleError(resp)
	}
	var aux *timeRFC1123
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// getOverflowHandleError handles the GetOverflow error response.
func (client *datetimerfc1123Operations) getOverflowHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetUTCLowercaseMaxDateTime - Get max datetime value fri, 31 dec 9999 23:59:59 gmt
func (client *datetimerfc1123Operations) GetUTCLowercaseMaxDateTime(ctx context.Context) (*TimeResponse, error) {
	req, err := client.getUtcLowercaseMaxDateTimeCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getUtcLowercaseMaxDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getUtcLowercaseMaxDateTimeCreateRequest creates the GetUTCLowercaseMaxDateTime request.
func (client *datetimerfc1123Operations) getUtcLowercaseMaxDateTimeCreateRequest() (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/max/lowercase"
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getUtcLowercaseMaxDateTimeHandleResponse handles the GetUTCLowercaseMaxDateTime response.
func (client *datetimerfc1123Operations) getUtcLowercaseMaxDateTimeHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getUtcLowercaseMaxDateTimeHandleError(resp)
	}
	var aux *timeRFC1123
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// getUtcLowercaseMaxDateTimeHandleError handles the GetUTCLowercaseMaxDateTime error response.
func (client *datetimerfc1123Operations) getUtcLowercaseMaxDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetUTCMinDateTime - Get min datetime value Mon, 1 Jan 0001 00:00:00 GMT
func (client *datetimerfc1123Operations) GetUTCMinDateTime(ctx context.Context) (*TimeResponse, error) {
	req, err := client.getUtcMinDateTimeCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getUtcMinDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getUtcMinDateTimeCreateRequest creates the GetUTCMinDateTime request.
func (client *datetimerfc1123Operations) getUtcMinDateTimeCreateRequest() (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/min"
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getUtcMinDateTimeHandleResponse handles the GetUTCMinDateTime response.
func (client *datetimerfc1123Operations) getUtcMinDateTimeHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getUtcMinDateTimeHandleError(resp)
	}
	var aux *timeRFC1123
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// getUtcMinDateTimeHandleError handles the GetUTCMinDateTime error response.
func (client *datetimerfc1123Operations) getUtcMinDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetUTCUppercaseMaxDateTime - Get max datetime value FRI, 31 DEC 9999 23:59:59 GMT
func (client *datetimerfc1123Operations) GetUTCUppercaseMaxDateTime(ctx context.Context) (*TimeResponse, error) {
	req, err := client.getUtcUppercaseMaxDateTimeCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getUtcUppercaseMaxDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getUtcUppercaseMaxDateTimeCreateRequest creates the GetUTCUppercaseMaxDateTime request.
func (client *datetimerfc1123Operations) getUtcUppercaseMaxDateTimeCreateRequest() (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/max/uppercase"
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getUtcUppercaseMaxDateTimeHandleResponse handles the GetUTCUppercaseMaxDateTime response.
func (client *datetimerfc1123Operations) getUtcUppercaseMaxDateTimeHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getUtcUppercaseMaxDateTimeHandleError(resp)
	}
	var aux *timeRFC1123
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// getUtcUppercaseMaxDateTimeHandleError handles the GetUTCUppercaseMaxDateTime error response.
func (client *datetimerfc1123Operations) getUtcUppercaseMaxDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetUnderflow - Get underflow datetime value
func (client *datetimerfc1123Operations) GetUnderflow(ctx context.Context) (*TimeResponse, error) {
	req, err := client.getUnderflowCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getUnderflowHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getUnderflowCreateRequest creates the GetUnderflow request.
func (client *datetimerfc1123Operations) getUnderflowCreateRequest() (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/underflow"
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getUnderflowHandleResponse handles the GetUnderflow response.
func (client *datetimerfc1123Operations) getUnderflowHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getUnderflowHandleError(resp)
	}
	var aux *timeRFC1123
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// getUnderflowHandleError handles the GetUnderflow error response.
func (client *datetimerfc1123Operations) getUnderflowHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutUTCMaxDateTime - Put max datetime value Fri, 31 Dec 9999 23:59:59 GMT
func (client *datetimerfc1123Operations) PutUTCMaxDateTime(ctx context.Context, datetimeBody time.Time) (*http.Response, error) {
	req, err := client.putUtcMaxDateTimeCreateRequest(datetimeBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putUtcMaxDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putUtcMaxDateTimeCreateRequest creates the PutUTCMaxDateTime request.
func (client *datetimerfc1123Operations) putUtcMaxDateTimeCreateRequest(datetimeBody time.Time) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/max"
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	aux := timeRFC1123(datetimeBody)
	return req, req.MarshalAsJSON(aux)
}

// putUtcMaxDateTimeHandleResponse handles the PutUTCMaxDateTime response.
func (client *datetimerfc1123Operations) putUtcMaxDateTimeHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putUtcMaxDateTimeHandleError(resp)
	}
	return resp.Response, nil
}

// putUtcMaxDateTimeHandleError handles the PutUTCMaxDateTime error response.
func (client *datetimerfc1123Operations) putUtcMaxDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutUTCMinDateTime - Put min datetime value Mon, 1 Jan 0001 00:00:00 GMT
func (client *datetimerfc1123Operations) PutUTCMinDateTime(ctx context.Context, datetimeBody time.Time) (*http.Response, error) {
	req, err := client.putUtcMinDateTimeCreateRequest(datetimeBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putUtcMinDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putUtcMinDateTimeCreateRequest creates the PutUTCMinDateTime request.
func (client *datetimerfc1123Operations) putUtcMinDateTimeCreateRequest(datetimeBody time.Time) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/min"
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	aux := timeRFC1123(datetimeBody)
	return req, req.MarshalAsJSON(aux)
}

// putUtcMinDateTimeHandleResponse handles the PutUTCMinDateTime response.
func (client *datetimerfc1123Operations) putUtcMinDateTimeHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putUtcMinDateTimeHandleError(resp)
	}
	return resp.Response, nil
}

// putUtcMinDateTimeHandleError handles the PutUTCMinDateTime error response.
func (client *datetimerfc1123Operations) putUtcMinDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
