//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package nonstringenumgroup

// FloatClientGetResponse contains the response from method FloatClient.Get.
type FloatClientGetResponse struct {
	// List of float enums
	Value *FloatEnum
}

// FloatClientPutResponse contains the response from method FloatClient.Put.
type FloatClientPutResponse struct {
	Value *string
}

// IntClientGetResponse contains the response from method IntClient.Get.
type IntClientGetResponse struct {
	// List of integer enums
	Value *IntEnum
}

// IntClientPutResponse contains the response from method IntClient.Put.
type IntClientPutResponse struct {
	Value *string
}