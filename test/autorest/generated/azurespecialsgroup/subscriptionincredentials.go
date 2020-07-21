// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurespecialsgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// SubscriptionInCredentialsOperations contains the methods for the SubscriptionInCredentials group.
type SubscriptionInCredentialsOperations interface {
	// PostMethodGlobalNotProvidedValid - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
	PostMethodGlobalNotProvidedValid(ctx context.Context) (*http.Response, error)
	// PostMethodGlobalNull - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to null, and client-side validation should prevent you from making this call
	PostMethodGlobalNull(ctx context.Context) (*http.Response, error)
	// PostMethodGlobalValid - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
	PostMethodGlobalValid(ctx context.Context) (*http.Response, error)
	// PostPathGlobalValid - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
	PostPathGlobalValid(ctx context.Context) (*http.Response, error)
	// PostSwaggerGlobalValid - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
	PostSwaggerGlobalValid(ctx context.Context) (*http.Response, error)
}

// subscriptionInCredentialsOperations implements the SubscriptionInCredentialsOperations interface.
type subscriptionInCredentialsOperations struct {
	*Client
	subscriptionID string
}

// PostMethodGlobalNotProvidedValid - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
func (client *subscriptionInCredentialsOperations) PostMethodGlobalNotProvidedValid(ctx context.Context) (*http.Response, error) {
	req, err := client.postMethodGlobalNotProvidedValidCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.postMethodGlobalNotProvidedValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// postMethodGlobalNotProvidedValidCreateRequest creates the PostMethodGlobalNotProvidedValid request.
func (client *subscriptionInCredentialsOperations) postMethodGlobalNotProvidedValidCreateRequest() (*azcore.Request, error) {
	urlPath := "/azurespecials/subscriptionId/method/string/none/path/globalNotProvided/1234-5678-9012-3456/{subscriptionId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2015-07-01-preview")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, nil
}

// postMethodGlobalNotProvidedValidHandleResponse handles the PostMethodGlobalNotProvidedValid response.
func (client *subscriptionInCredentialsOperations) postMethodGlobalNotProvidedValidHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postMethodGlobalNotProvidedValidHandleError(resp)
	}
	return resp.Response, nil
}

// postMethodGlobalNotProvidedValidHandleError handles the PostMethodGlobalNotProvidedValid error response.
func (client *subscriptionInCredentialsOperations) postMethodGlobalNotProvidedValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PostMethodGlobalNull - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to null, and client-side validation should prevent you from making this call
func (client *subscriptionInCredentialsOperations) PostMethodGlobalNull(ctx context.Context) (*http.Response, error) {
	req, err := client.postMethodGlobalNullCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.postMethodGlobalNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// postMethodGlobalNullCreateRequest creates the PostMethodGlobalNull request.
func (client *subscriptionInCredentialsOperations) postMethodGlobalNullCreateRequest() (*azcore.Request, error) {
	urlPath := "/azurespecials/subscriptionId/method/string/none/path/global/null/{subscriptionId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, nil
}

// postMethodGlobalNullHandleResponse handles the PostMethodGlobalNull response.
func (client *subscriptionInCredentialsOperations) postMethodGlobalNullHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postMethodGlobalNullHandleError(resp)
	}
	return resp.Response, nil
}

// postMethodGlobalNullHandleError handles the PostMethodGlobalNull error response.
func (client *subscriptionInCredentialsOperations) postMethodGlobalNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PostMethodGlobalValid - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
func (client *subscriptionInCredentialsOperations) PostMethodGlobalValid(ctx context.Context) (*http.Response, error) {
	req, err := client.postMethodGlobalValidCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.postMethodGlobalValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// postMethodGlobalValidCreateRequest creates the PostMethodGlobalValid request.
func (client *subscriptionInCredentialsOperations) postMethodGlobalValidCreateRequest() (*azcore.Request, error) {
	urlPath := "/azurespecials/subscriptionId/method/string/none/path/global/1234-5678-9012-3456/{subscriptionId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, nil
}

// postMethodGlobalValidHandleResponse handles the PostMethodGlobalValid response.
func (client *subscriptionInCredentialsOperations) postMethodGlobalValidHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postMethodGlobalValidHandleError(resp)
	}
	return resp.Response, nil
}

// postMethodGlobalValidHandleError handles the PostMethodGlobalValid error response.
func (client *subscriptionInCredentialsOperations) postMethodGlobalValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PostPathGlobalValid - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
func (client *subscriptionInCredentialsOperations) PostPathGlobalValid(ctx context.Context) (*http.Response, error) {
	req, err := client.postPathGlobalValidCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.postPathGlobalValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// postPathGlobalValidCreateRequest creates the PostPathGlobalValid request.
func (client *subscriptionInCredentialsOperations) postPathGlobalValidCreateRequest() (*azcore.Request, error) {
	urlPath := "/azurespecials/subscriptionId/path/string/none/path/global/1234-5678-9012-3456/{subscriptionId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, nil
}

// postPathGlobalValidHandleResponse handles the PostPathGlobalValid response.
func (client *subscriptionInCredentialsOperations) postPathGlobalValidHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postPathGlobalValidHandleError(resp)
	}
	return resp.Response, nil
}

// postPathGlobalValidHandleError handles the PostPathGlobalValid error response.
func (client *subscriptionInCredentialsOperations) postPathGlobalValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PostSwaggerGlobalValid - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
func (client *subscriptionInCredentialsOperations) PostSwaggerGlobalValid(ctx context.Context) (*http.Response, error) {
	req, err := client.postSwaggerGlobalValidCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.postSwaggerGlobalValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// postSwaggerGlobalValidCreateRequest creates the PostSwaggerGlobalValid request.
func (client *subscriptionInCredentialsOperations) postSwaggerGlobalValidCreateRequest() (*azcore.Request, error) {
	urlPath := "/azurespecials/subscriptionId/swagger/string/none/path/global/1234-5678-9012-3456/{subscriptionId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(path.Join(client.u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, nil
}

// postSwaggerGlobalValidHandleResponse handles the PostSwaggerGlobalValid response.
func (client *subscriptionInCredentialsOperations) postSwaggerGlobalValidHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postSwaggerGlobalValidHandleError(resp)
	}
	return resp.Response, nil
}

// postSwaggerGlobalValidHandleError handles the PostSwaggerGlobalValid error response.
func (client *subscriptionInCredentialsOperations) postSwaggerGlobalValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
