//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azalias

// ClientCreateResponse contains the response from method Client.Create.
type ClientCreateResponse struct {
	// The response model for the Alias Create API for the case when the alias was successfully created.
	AliasesCreateResponse

	// AccessControlExposeHeaders contains the information returned from the Access-Control-Expose-Headers header response.
	AccessControlExposeHeaders *string
}

// ClientGetScriptResponse contains the response from method Client.GetScript.
type ClientGetScriptResponse struct {
	Value *string
}

// ClientListLROResponse contains the response from method Client.BeginListLRO.
type ClientListLROResponse struct {
	PagesOfThings
}

// ClientListResponse contains the response from method Client.NewListPager.
type ClientListResponse struct {
	// The response model for the List API. Returns a list of all the previously created aliases.
	ListResponse
}

// ClientListWithSharedNextOneResponse contains the response from method Client.NewListWithSharedNextOnePager.
type ClientListWithSharedNextOneResponse struct {
	// The response model for the List API. Returns a list of all the previously created aliases.
	ListResponse
}

// ClientListWithSharedNextTwoResponse contains the response from method Client.NewListWithSharedNextTwoPager.
type ClientListWithSharedNextTwoResponse struct {
	// The response model for the List API. Returns a list of all the previously created aliases.
	ListResponse
}

// ClientPolicyAssignmentResponse contains the response from method Client.PolicyAssignment.
type ClientPolicyAssignmentResponse struct {
	PolicyAssignmentProperties
}
