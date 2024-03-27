// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armcodesigning

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// CodeSigningClient - Code Signing resource provider api.
// Don't use this type directly, use NewCodeSigningClient() instead.
type CodeSigningClient struct {
	internal *arm.Client
}

// NewCodeSigningClient creates a new instance of CodeSigningClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCodeSigningClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*CodeSigningClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CodeSigningClient{
		internal: cl,
	}
	return client, nil
}

// NewCertificateProfilesClient creates a new instance of [CertificateProfilesClient].
func (client *CodeSigningClient) NewCertificateProfilesClient() *CertificateProfilesClient {
	return &CertificateProfilesClient{
		internal: client.internal,
	}
}

// NewCodeSigningAccountsClient creates a new instance of [CodeSigningAccountsClient].
func (client *CodeSigningClient) NewCodeSigningAccountsClient() *CodeSigningAccountsClient {
	return &CodeSigningAccountsClient{
		internal: client.internal,
	}
}

// NewOperationsClient creates a new instance of [OperationsClient].
func (client *CodeSigningClient) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: client.internal,
	}
}