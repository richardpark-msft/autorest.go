// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armcodesigning

import "time"

// Trusted signing account resource.
type Account struct {
	// REQUIRED; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string

	// The resource-specific properties for this resource.
	Properties *AccountProperties

	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Trusted Signing account name.
	Name *string
}

// The response of a CodeSigningAccount list operation.
type AccountListResult struct {
	// REQUIRED; The CodeSigningAccount items on this page
	Value []*Account

	// The link to the next page of items
	NextLink *string
}

// Parameters for creating or updating a trusted signing account.
type AccountPatch struct {
	// Properties of the trusted signing account.
	Properties *AccountPatchProperties

	// Resource tags.
	Tags map[string]*string
}

// Properties of the trusted signing account.
type AccountPatchProperties struct {
	// SKU of the trusted signing account.
	SKU *AccountSKU
}

// Properties of the trusted signing account.
type AccountProperties struct {
	// The URI of the trusted signing account which is used during signing files.
	AccountURI *string

	// Status of the current operation on trusted signing account.
	ProvisioningState *ProvisioningState

	// SKU of the trusted signing account.
	SKU *AccountSKU
}

// SKU of the trusted signing account.
type AccountSKU struct {
	// REQUIRED; Name of the SKU.
	Name *SKUName
}

// Common properties for all Azure Resource Manager resources.
type ArmResource struct {
	// REQUIRED; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// REQUIRED; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string

	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData
}

// Base class used for type definitions
type ArmResourceBase struct {
}

// Properties of the certificate.
type Certificate struct {
	// Certificate created date.
	CreatedDate *string

	// Certificate expiry date.
	ExpiryDate *string

	// Revocations history of a certificate.
	Revocation *Revocation

	// Serial number of the certificate.
	SerialNumber *string

	// Status of the certificate.
	Status *CertificateStatus

	// Subject name of the certificate.
	SubjectName *string

	// Thumbprint of the certificate.
	Thumbprint *string
}

// Certificate profile resource.
type CertificateProfile struct {
	// REQUIRED; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// REQUIRED; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string

	// The resource-specific properties for this resource.
	Properties *CertificateProfileProperties

	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; Certificate profile name.
	Name *string
}

// The response of a CertificateProfile list operation.
type CertificateProfileListResult struct {
	// REQUIRED; The CertificateProfile items on this page
	Value []*CertificateProfile

	// The link to the next page of items
	NextLink *string
}

// Properties of the certificate profile.
type CertificateProfileProperties struct {
	// REQUIRED; Profile type of the certificate.
	ProfileType *ProfileType

	// List of renewed certificates.
	Certificates []*Certificate

	// Used as L in the certificate subject name.
	City *string

	// Used as CN in the certificate subject name.
	CommonName *string

	// Used as C in the certificate subject name.
	Country *string

	// Enhanced key usage of the certificate.
	EnhancedKeyUsage *string

	// Identity validation id used for the certificate subject name.
	IdentityValidationID *string

	// Whether to include L in the certificate subject name. Applicable only for private trust, private trust ci profile types
	IncludeCity *bool

	// Whether to include C in the certificate subject name. Applicable only for private trust, private trust ci profile types
	IncludeCountry *bool

	// Whether to include PC in the certificate subject name.
	IncludePostalCode *bool

	// Whether to include S in the certificate subject name. Applicable only for private trust, private trust ci profile types
	IncludeState *bool

	// Whether to include STREET in the certificate subject name.
	IncludeStreetAddress *bool

	// Used as O in the certificate subject name.
	Organization *string

	// Used as OU in the private trust certificate subject name.
	OrganizationUnit *string

	// Used as PC in the certificate subject name.
	PostalCode *string

	// Status of the current operation on certificate profile.
	ProvisioningState *ProvisioningState

	// Used as S in the certificate subject name.
	State *string

	// Status of the certificate profile.
	Status *CertificateProfileStatus

	// Used as STREET in the certificate subject name.
	StreetAddress *string
}

// The parameters used to check the availability of the trusted signing account name.
type CheckNameAvailability struct {
	// REQUIRED; Trusted signing account name.
	Name *string
}

// The CheckNameAvailability operation response.
type CheckNameAvailabilityResult struct {
	// An error message explaining the Reason value in more detail.
	Message *string

	// A boolean value that indicates whether the name is available for you to use. If true, the name is available. If false,
	// the name has already been taken or is invalid and cannot be used.
	NameAvailable *bool

	// The reason that a trusted signing account name could not be used. The Reason element is only returned if nameAvailable
	// is false.
	Reason *NameUnavailabilityReason
}

// Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// Localized display information for this particular operation.
	Display *OperationDisplay

	// Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for Azure Resource Manager/control-plane
	// operations.
	IsDataAction *bool

	// The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default value is
	// "user,system"
	Origin *Origin
}

// Localized display information for and operation.
type OperationDisplay struct {
	// The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual Machine",
	// "Restart Virtual Machine".
	Operation *string

	// The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft Compute".
	Provider *string

	// The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job Schedule Collections".
	Resource *string
}

// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
type PagedOperation struct {
	// REQUIRED; The Operation items on this page
	Value []*Operation

	// The link to the next page of items
	NextLink *string
}

// The base proxy resource.
type ProxyResourceBase struct {
	// REQUIRED; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// REQUIRED; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string

	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData
}

// Revocation details of the certificate.
type Revocation struct {
	// The timestamp when the revocation is effective.
	EffectiveAt *time.Time

	// Reason for the revocation failure.
	FailureReason *string

	// Reason for revocation.
	Reason *string

	// Remarks for the revocation.
	Remarks *string

	// The timestamp when the revocation is requested.
	RequestedAt *time.Time

	// Status of the revocation.
	Status *RevocationStatus
}

// Defines the certificate revocation properties.
type RevokeCertificate struct {
	// REQUIRED; The timestamp when the revocation is effective.
	EffectiveAt *time.Time

	// REQUIRED; Reason for the revocation.
	Reason *string

	// REQUIRED; Serial number of the certificate.
	SerialNumber *string

	// REQUIRED; Thumbprint of the certificate.
	Thumbprint *string

	// Remarks for the revocation.
	Remarks *string
}

// Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The type of identity that created the resource.
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// The resource model definition for an Azure Resource Manager tracked top level resource
type TrackedResourceBase struct {
	// REQUIRED; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string

	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// Resource tags.
	Tags map[string]*string
}
