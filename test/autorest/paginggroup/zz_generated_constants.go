// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package paginggroup

// OperationResultStatus - The status of the request
type OperationResultStatus string

const (
	OperationResultStatusAccepted  OperationResultStatus = "Accepted"
	OperationResultStatusCanceled  OperationResultStatus = "canceled"
	OperationResultStatusCreated   OperationResultStatus = "Created"
	OperationResultStatusCreating  OperationResultStatus = "Creating"
	OperationResultStatusDeleted   OperationResultStatus = "Deleted"
	OperationResultStatusDeleting  OperationResultStatus = "Deleting"
	OperationResultStatusFailed    OperationResultStatus = "Failed"
	OperationResultStatusOK        OperationResultStatus = "OK"
	OperationResultStatusSucceeded OperationResultStatus = "Succeeded"
	OperationResultStatusUpdated   OperationResultStatus = "Updated"
	OperationResultStatusUpdating  OperationResultStatus = "Updating"
)

// PossibleOperationResultStatusValues returns the possible values for the OperationResultStatus const type.
func PossibleOperationResultStatusValues() []OperationResultStatus {
	return []OperationResultStatus{
		OperationResultStatusAccepted,
		OperationResultStatusCanceled,
		OperationResultStatusCreated,
		OperationResultStatusCreating,
		OperationResultStatusDeleted,
		OperationResultStatusDeleting,
		OperationResultStatusFailed,
		OperationResultStatusOK,
		OperationResultStatusSucceeded,
		OperationResultStatusUpdated,
		OperationResultStatusUpdating,
	}
}

// ToPtr() returns a *OperationResultStatus pointing to the current value.
func (c OperationResultStatus) ToPtr() *OperationResultStatus {
	return &c
}
