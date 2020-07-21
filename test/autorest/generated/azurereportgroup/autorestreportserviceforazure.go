// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurereportgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"path"
)

// AutoRestReportServiceForAzureOperations contains the methods for the AutoRestReportServiceForAzure group.
type AutoRestReportServiceForAzureOperations interface {
	// GetReport - Get test coverage report
	GetReport(ctx context.Context, autoRestReportServiceForAzureGetReportOptions *AutoRestReportServiceForAzureGetReportOptions) (*MapOfInt32Response, error)
}

// autoRestReportServiceForAzureOperations implements the AutoRestReportServiceForAzureOperations interface.
type autoRestReportServiceForAzureOperations struct {
	*Client
}

// GetReport - Get test coverage report
func (client *autoRestReportServiceForAzureOperations) GetReport(ctx context.Context, autoRestReportServiceForAzureGetReportOptions *AutoRestReportServiceForAzureGetReportOptions) (*MapOfInt32Response, error) {
	req, err := client.getReportCreateRequest(autoRestReportServiceForAzureGetReportOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getReportHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getReportCreateRequest creates the GetReport request.
func (client *autoRestReportServiceForAzureOperations) getReportCreateRequest(autoRestReportServiceForAzureGetReportOptions *AutoRestReportServiceForAzureGetReportOptions) (*azcore.Request, error) {
	urlPath := "/report/azure"
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	if autoRestReportServiceForAzureGetReportOptions != nil && autoRestReportServiceForAzureGetReportOptions.Qualifier != nil {
		query.Set("qualifier", *autoRestReportServiceForAzureGetReportOptions.Qualifier)
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getReportHandleResponse handles the GetReport response.
func (client *autoRestReportServiceForAzureOperations) getReportHandleResponse(resp *azcore.Response) (*MapOfInt32Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getReportHandleError(resp)
	}
	result := MapOfInt32Response{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// getReportHandleError handles the GetReport error response.
func (client *autoRestReportServiceForAzureOperations) getReportHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
