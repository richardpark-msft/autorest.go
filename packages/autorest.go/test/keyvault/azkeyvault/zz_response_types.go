//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azkeyvault

// ClientBackupCertificateResponse contains the response from method Client.BackupCertificate.
type ClientBackupCertificateResponse struct {
	BackupCertificateResult
}

// ClientBackupKeyResponse contains the response from method Client.BackupKey.
type ClientBackupKeyResponse struct {
	BackupKeyResult
}

// ClientBackupSecretResponse contains the response from method Client.BackupSecret.
type ClientBackupSecretResponse struct {
	BackupSecretResult
}

// ClientBackupStorageAccountResponse contains the response from method Client.BackupStorageAccount.
type ClientBackupStorageAccountResponse struct {
	BackupStorageResult
}

// ClientCreateCertificateResponse contains the response from method Client.CreateCertificate.
type ClientCreateCertificateResponse struct {
	CertificateOperation
}

// ClientCreateKeyResponse contains the response from method Client.CreateKey.
type ClientCreateKeyResponse struct {
	KeyBundle
}

// ClientDecryptResponse contains the response from method Client.Decrypt.
type ClientDecryptResponse struct {
	KeyOperationResult
}

// ClientDeleteCertificateContactsResponse contains the response from method Client.DeleteCertificateContacts.
type ClientDeleteCertificateContactsResponse struct {
	Contacts
}

// ClientDeleteCertificateIssuerResponse contains the response from method Client.DeleteCertificateIssuer.
type ClientDeleteCertificateIssuerResponse struct {
	IssuerBundle
}

// ClientDeleteCertificateOperationResponse contains the response from method Client.DeleteCertificateOperation.
type ClientDeleteCertificateOperationResponse struct {
	CertificateOperation
}

// ClientDeleteCertificateResponse contains the response from method Client.DeleteCertificate.
type ClientDeleteCertificateResponse struct {
	DeletedCertificateBundle
}

// ClientDeleteKeyResponse contains the response from method Client.DeleteKey.
type ClientDeleteKeyResponse struct {
	DeletedKeyBundle
}

// ClientDeleteSasDefinitionResponse contains the response from method Client.DeleteSasDefinition.
type ClientDeleteSasDefinitionResponse struct {
	DeletedSasDefinitionBundle
}

// ClientDeleteSecretResponse contains the response from method Client.DeleteSecret.
type ClientDeleteSecretResponse struct {
	DeletedSecretBundle
}

// ClientDeleteStorageAccountResponse contains the response from method Client.DeleteStorageAccount.
type ClientDeleteStorageAccountResponse struct {
	DeletedStorageBundle
}

// ClientEncryptResponse contains the response from method Client.Encrypt.
type ClientEncryptResponse struct {
	KeyOperationResult
}

// ClientFullBackupResponse contains the response from method Client.BeginFullBackup.
type ClientFullBackupResponse struct {
	FullBackupOperation
}

// ClientFullBackupStatusResponse contains the response from method Client.FullBackupStatus.
type ClientFullBackupStatusResponse struct {
	FullBackupOperation
}

// ClientFullRestoreOperationResponse contains the response from method Client.BeginFullRestoreOperation.
type ClientFullRestoreOperationResponse struct {
	RestoreOperation
}

// ClientGetCertificateContactsResponse contains the response from method Client.GetCertificateContacts.
type ClientGetCertificateContactsResponse struct {
	Contacts
}

// ClientGetCertificateIssuerResponse contains the response from method Client.GetCertificateIssuer.
type ClientGetCertificateIssuerResponse struct {
	IssuerBundle
}

// ClientGetCertificateIssuersResponse contains the response from method Client.NewGetCertificateIssuersPager.
type ClientGetCertificateIssuersResponse struct {
	CertificateIssuerListResult
}

// ClientGetCertificateOperationResponse contains the response from method Client.GetCertificateOperation.
type ClientGetCertificateOperationResponse struct {
	CertificateOperation
}

// ClientGetCertificatePolicyResponse contains the response from method Client.GetCertificatePolicy.
type ClientGetCertificatePolicyResponse struct {
	CertificatePolicy
}

// ClientGetCertificateResponse contains the response from method Client.GetCertificate.
type ClientGetCertificateResponse struct {
	CertificateBundle
}

// ClientGetCertificateVersionsResponse contains the response from method Client.NewGetCertificateVersionsPager.
type ClientGetCertificateVersionsResponse struct {
	CertificateListResult
}

// ClientGetCertificatesResponse contains the response from method Client.NewGetCertificatesPager.
type ClientGetCertificatesResponse struct {
	CertificateListResult
}

// ClientGetDeletedCertificateResponse contains the response from method Client.GetDeletedCertificate.
type ClientGetDeletedCertificateResponse struct {
	DeletedCertificateBundle
}

// ClientGetDeletedCertificatesResponse contains the response from method Client.NewGetDeletedCertificatesPager.
type ClientGetDeletedCertificatesResponse struct {
	DeletedCertificateListResult
}

// ClientGetDeletedKeyResponse contains the response from method Client.GetDeletedKey.
type ClientGetDeletedKeyResponse struct {
	DeletedKeyBundle
}

// ClientGetDeletedKeysResponse contains the response from method Client.NewGetDeletedKeysPager.
type ClientGetDeletedKeysResponse struct {
	DeletedKeyListResult
}

// ClientGetDeletedSasDefinitionResponse contains the response from method Client.GetDeletedSasDefinition.
type ClientGetDeletedSasDefinitionResponse struct {
	DeletedSasDefinitionBundle
}

// ClientGetDeletedSasDefinitionsResponse contains the response from method Client.NewGetDeletedSasDefinitionsPager.
type ClientGetDeletedSasDefinitionsResponse struct {
	DeletedSasDefinitionListResult
}

// ClientGetDeletedSecretResponse contains the response from method Client.GetDeletedSecret.
type ClientGetDeletedSecretResponse struct {
	DeletedSecretBundle
}

// ClientGetDeletedSecretsResponse contains the response from method Client.NewGetDeletedSecretsPager.
type ClientGetDeletedSecretsResponse struct {
	DeletedSecretListResult
}

// ClientGetDeletedStorageAccountResponse contains the response from method Client.GetDeletedStorageAccount.
type ClientGetDeletedStorageAccountResponse struct {
	DeletedStorageBundle
}

// ClientGetDeletedStorageAccountsResponse contains the response from method Client.NewGetDeletedStorageAccountsPager.
type ClientGetDeletedStorageAccountsResponse struct {
	DeletedStorageListResult
}

// ClientGetKeyResponse contains the response from method Client.GetKey.
type ClientGetKeyResponse struct {
	KeyBundle
}

// ClientGetKeyVersionsResponse contains the response from method Client.NewGetKeyVersionsPager.
type ClientGetKeyVersionsResponse struct {
	KeyListResult
}

// ClientGetKeysResponse contains the response from method Client.NewGetKeysPager.
type ClientGetKeysResponse struct {
	KeyListResult
}

// ClientGetSasDefinitionResponse contains the response from method Client.GetSasDefinition.
type ClientGetSasDefinitionResponse struct {
	SasDefinitionBundle
}

// ClientGetSasDefinitionsResponse contains the response from method Client.NewGetSasDefinitionsPager.
type ClientGetSasDefinitionsResponse struct {
	SasDefinitionListResult
}

// ClientGetSecretResponse contains the response from method Client.GetSecret.
type ClientGetSecretResponse struct {
	SecretBundle
}

// ClientGetSecretVersionsResponse contains the response from method Client.NewGetSecretVersionsPager.
type ClientGetSecretVersionsResponse struct {
	SecretListResult
}

// ClientGetSecretsResponse contains the response from method Client.NewGetSecretsPager.
type ClientGetSecretsResponse struct {
	SecretListResult
}

// ClientGetStorageAccountResponse contains the response from method Client.GetStorageAccount.
type ClientGetStorageAccountResponse struct {
	StorageBundle
}

// ClientGetStorageAccountsResponse contains the response from method Client.NewGetStorageAccountsPager.
type ClientGetStorageAccountsResponse struct {
	StorageListResult
}

// ClientImportCertificateResponse contains the response from method Client.ImportCertificate.
type ClientImportCertificateResponse struct {
	CertificateBundle
}

// ClientImportKeyResponse contains the response from method Client.ImportKey.
type ClientImportKeyResponse struct {
	KeyBundle
}

// ClientMergeCertificateResponse contains the response from method Client.MergeCertificate.
type ClientMergeCertificateResponse struct {
	CertificateBundle
}

// ClientPurgeDeletedCertificateResponse contains the response from method Client.PurgeDeletedCertificate.
type ClientPurgeDeletedCertificateResponse struct {
	// placeholder for future response values
}

// ClientPurgeDeletedKeyResponse contains the response from method Client.PurgeDeletedKey.
type ClientPurgeDeletedKeyResponse struct {
	// placeholder for future response values
}

// ClientPurgeDeletedSecretResponse contains the response from method Client.PurgeDeletedSecret.
type ClientPurgeDeletedSecretResponse struct {
	// placeholder for future response values
}

// ClientPurgeDeletedStorageAccountResponse contains the response from method Client.PurgeDeletedStorageAccount.
type ClientPurgeDeletedStorageAccountResponse struct {
	// placeholder for future response values
}

// ClientRecoverDeletedCertificateResponse contains the response from method Client.RecoverDeletedCertificate.
type ClientRecoverDeletedCertificateResponse struct {
	CertificateBundle
}

// ClientRecoverDeletedKeyResponse contains the response from method Client.RecoverDeletedKey.
type ClientRecoverDeletedKeyResponse struct {
	KeyBundle
}

// ClientRecoverDeletedSasDefinitionResponse contains the response from method Client.RecoverDeletedSasDefinition.
type ClientRecoverDeletedSasDefinitionResponse struct {
	SasDefinitionBundle
}

// ClientRecoverDeletedSecretResponse contains the response from method Client.RecoverDeletedSecret.
type ClientRecoverDeletedSecretResponse struct {
	SecretBundle
}

// ClientRecoverDeletedStorageAccountResponse contains the response from method Client.RecoverDeletedStorageAccount.
type ClientRecoverDeletedStorageAccountResponse struct {
	StorageBundle
}

// ClientRegenerateStorageAccountKeyResponse contains the response from method Client.RegenerateStorageAccountKey.
type ClientRegenerateStorageAccountKeyResponse struct {
	StorageBundle
}

// ClientRestoreCertificateResponse contains the response from method Client.RestoreCertificate.
type ClientRestoreCertificateResponse struct {
	CertificateBundle
}

// ClientRestoreKeyResponse contains the response from method Client.RestoreKey.
type ClientRestoreKeyResponse struct {
	KeyBundle
}

// ClientRestoreSecretResponse contains the response from method Client.RestoreSecret.
type ClientRestoreSecretResponse struct {
	SecretBundle
}

// ClientRestoreStatusResponse contains the response from method Client.RestoreStatus.
type ClientRestoreStatusResponse struct {
	RestoreOperation
}

// ClientRestoreStorageAccountResponse contains the response from method Client.RestoreStorageAccount.
type ClientRestoreStorageAccountResponse struct {
	StorageBundle
}

// ClientSelectiveKeyRestoreOperationResponse contains the response from method Client.BeginSelectiveKeyRestoreOperation.
type ClientSelectiveKeyRestoreOperationResponse struct {
	SelectiveKeyRestoreOperation
}

// ClientSetCertificateContactsResponse contains the response from method Client.SetCertificateContacts.
type ClientSetCertificateContactsResponse struct {
	Contacts
}

// ClientSetCertificateIssuerResponse contains the response from method Client.SetCertificateIssuer.
type ClientSetCertificateIssuerResponse struct {
	IssuerBundle
}

// ClientSetSasDefinitionResponse contains the response from method Client.SetSasDefinition.
type ClientSetSasDefinitionResponse struct {
	SasDefinitionBundle
}

// ClientSetSecretResponse contains the response from method Client.SetSecret.
type ClientSetSecretResponse struct {
	SecretBundle
}

// ClientSetStorageAccountResponse contains the response from method Client.SetStorageAccount.
type ClientSetStorageAccountResponse struct {
	StorageBundle
}

// ClientSignResponse contains the response from method Client.Sign.
type ClientSignResponse struct {
	KeyOperationResult
}

// ClientUnwrapKeyResponse contains the response from method Client.UnwrapKey.
type ClientUnwrapKeyResponse struct {
	KeyOperationResult
}

// ClientUpdateCertificateIssuerResponse contains the response from method Client.UpdateCertificateIssuer.
type ClientUpdateCertificateIssuerResponse struct {
	IssuerBundle
}

// ClientUpdateCertificateOperationResponse contains the response from method Client.UpdateCertificateOperation.
type ClientUpdateCertificateOperationResponse struct {
	CertificateOperation
}

// ClientUpdateCertificatePolicyResponse contains the response from method Client.UpdateCertificatePolicy.
type ClientUpdateCertificatePolicyResponse struct {
	CertificatePolicy
}

// ClientUpdateCertificateResponse contains the response from method Client.UpdateCertificate.
type ClientUpdateCertificateResponse struct {
	CertificateBundle
}

// ClientUpdateKeyResponse contains the response from method Client.UpdateKey.
type ClientUpdateKeyResponse struct {
	KeyBundle
}

// ClientUpdateSasDefinitionResponse contains the response from method Client.UpdateSasDefinition.
type ClientUpdateSasDefinitionResponse struct {
	SasDefinitionBundle
}

// ClientUpdateSecretResponse contains the response from method Client.UpdateSecret.
type ClientUpdateSecretResponse struct {
	SecretBundle
}

// ClientUpdateStorageAccountResponse contains the response from method Client.UpdateStorageAccount.
type ClientUpdateStorageAccountResponse struct {
	StorageBundle
}

// ClientVerifyResponse contains the response from method Client.Verify.
type ClientVerifyResponse struct {
	KeyVerifyResult
}

// ClientWrapKeyResponse contains the response from method Client.WrapKey.
type ClientWrapKeyResponse struct {
	KeyOperationResult
}

// HSMSecurityDomainClientDownloadPendingResponse contains the response from method HSMSecurityDomainClient.DownloadPending.
type HSMSecurityDomainClientDownloadPendingResponse struct {
	SecurityDomainOperationStatus
}

// HSMSecurityDomainClientDownloadResponse contains the response from method HSMSecurityDomainClient.BeginDownload.
type HSMSecurityDomainClientDownloadResponse struct {
	SecurityDomainObject
}

// HSMSecurityDomainClientTransferKeyResponse contains the response from method HSMSecurityDomainClient.TransferKey.
type HSMSecurityDomainClientTransferKeyResponse struct {
	TransferKey
}

// HSMSecurityDomainClientUploadPendingResponse contains the response from method HSMSecurityDomainClient.UploadPending.
type HSMSecurityDomainClientUploadPendingResponse struct {
	SecurityDomainOperationStatus
}

// HSMSecurityDomainClientUploadResponse contains the response from method HSMSecurityDomainClient.BeginUpload.
type HSMSecurityDomainClientUploadResponse struct {
	SecurityDomainOperationStatus
}

// RoleAssignmentsClientCreateResponse contains the response from method RoleAssignmentsClient.Create.
type RoleAssignmentsClientCreateResponse struct {
	RoleAssignment
}

// RoleAssignmentsClientDeleteResponse contains the response from method RoleAssignmentsClient.Delete.
type RoleAssignmentsClientDeleteResponse struct {
	RoleAssignment
}

// RoleAssignmentsClientGetResponse contains the response from method RoleAssignmentsClient.Get.
type RoleAssignmentsClientGetResponse struct {
	RoleAssignment
}

// RoleAssignmentsClientListForScopeResponse contains the response from method RoleAssignmentsClient.NewListForScopePager.
type RoleAssignmentsClientListForScopeResponse struct {
	RoleAssignmentListResult
}

// RoleDefinitionsClientCreateOrUpdateResponse contains the response from method RoleDefinitionsClient.CreateOrUpdate.
type RoleDefinitionsClientCreateOrUpdateResponse struct {
	RoleDefinition
}

// RoleDefinitionsClientDeleteResponse contains the response from method RoleDefinitionsClient.Delete.
type RoleDefinitionsClientDeleteResponse struct {
	RoleDefinition
}

// RoleDefinitionsClientGetResponse contains the response from method RoleDefinitionsClient.Get.
type RoleDefinitionsClientGetResponse struct {
	RoleDefinition
}

// RoleDefinitionsClientListResponse contains the response from method RoleDefinitionsClient.NewListPager.
type RoleDefinitionsClientListResponse struct {
	RoleDefinitionListResult
}
