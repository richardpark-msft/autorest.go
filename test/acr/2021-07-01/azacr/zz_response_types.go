//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azacr

import "io"

// AuthenticationClientExchangeAADAccessTokenForAcrRefreshTokenResponse contains the response from method authenticationClient.ExchangeAADAccessTokenForAcrRefreshToken.
type AuthenticationClientExchangeAADAccessTokenForAcrRefreshTokenResponse struct {
	RefreshToken
}

// AuthenticationClientExchangeAcrRefreshTokenForAcrAccessTokenResponse contains the response from method authenticationClient.ExchangeAcrRefreshTokenForAcrAccessToken.
type AuthenticationClientExchangeAcrRefreshTokenForAcrAccessTokenResponse struct {
	AccessToken
}

// AuthenticationClientGetAcrAccessTokenFromLoginResponse contains the response from method authenticationClient.GetAcrAccessTokenFromLogin.
type AuthenticationClientGetAcrAccessTokenFromLoginResponse struct {
	AccessToken
}

// ContainerRegistryBlobClientCancelUploadResponse contains the response from method containerRegistryBlobClient.CancelUpload.
type ContainerRegistryBlobClientCancelUploadResponse struct {
	// placeholder for future response values
}

// ContainerRegistryBlobClientCheckBlobExistsResponse contains the response from method containerRegistryBlobClient.CheckBlobExists.
type ContainerRegistryBlobClientCheckBlobExistsResponse struct {
	// ContentLength contains the information returned from the Content-Length header response.
	ContentLength *int64

	// DockerContentDigest contains the information returned from the Docker-Content-Digest header response.
	DockerContentDigest *string
}

// ContainerRegistryBlobClientCheckChunkExistsResponse contains the response from method containerRegistryBlobClient.CheckChunkExists.
type ContainerRegistryBlobClientCheckChunkExistsResponse struct {
	// ContentLength contains the information returned from the Content-Length header response.
	ContentLength *int64

	// ContentRange contains the information returned from the Content-Range header response.
	ContentRange *string
}

// ContainerRegistryBlobClientCompleteUploadResponse contains the response from method containerRegistryBlobClient.CompleteUpload.
type ContainerRegistryBlobClientCompleteUploadResponse struct {
	// DockerContentDigest contains the information returned from the Docker-Content-Digest header response.
	DockerContentDigest *string

	// Location contains the information returned from the Location header response.
	Location *string

	// Range contains the information returned from the Range header response.
	Range *string
}

// ContainerRegistryBlobClientDeleteBlobResponse contains the response from method containerRegistryBlobClient.DeleteBlob.
type ContainerRegistryBlobClientDeleteBlobResponse struct {
	// Body contains the streaming response.
	Body io.ReadCloser

	// DockerContentDigest contains the information returned from the Docker-Content-Digest header response.
	DockerContentDigest *string
}

// ContainerRegistryBlobClientGetBlobResponse contains the response from method containerRegistryBlobClient.GetBlob.
type ContainerRegistryBlobClientGetBlobResponse struct {
	// Body contains the streaming response.
	Body io.ReadCloser

	// ContentLength contains the information returned from the Content-Length header response.
	ContentLength *int64

	// DockerContentDigest contains the information returned from the Docker-Content-Digest header response.
	DockerContentDigest *string
}

// ContainerRegistryBlobClientGetChunkResponse contains the response from method containerRegistryBlobClient.GetChunk.
type ContainerRegistryBlobClientGetChunkResponse struct {
	// Body contains the streaming response.
	Body io.ReadCloser

	// ContentLength contains the information returned from the Content-Length header response.
	ContentLength *int64

	// ContentRange contains the information returned from the Content-Range header response.
	ContentRange *string
}

// ContainerRegistryBlobClientGetUploadStatusResponse contains the response from method containerRegistryBlobClient.GetUploadStatus.
type ContainerRegistryBlobClientGetUploadStatusResponse struct {
	// DockerUploadUUID contains the information returned from the Docker-Upload-UUID header response.
	DockerUploadUUID *string

	// Range contains the information returned from the Range header response.
	Range *string
}

// ContainerRegistryBlobClientMountBlobResponse contains the response from method containerRegistryBlobClient.MountBlob.
type ContainerRegistryBlobClientMountBlobResponse struct {
	// DockerContentDigest contains the information returned from the Docker-Content-Digest header response.
	DockerContentDigest *string

	// DockerUploadUUID contains the information returned from the Docker-Upload-UUID header response.
	DockerUploadUUID *string

	// Location contains the information returned from the Location header response.
	Location *string
}

// ContainerRegistryBlobClientStartUploadResponse contains the response from method containerRegistryBlobClient.StartUpload.
type ContainerRegistryBlobClientStartUploadResponse struct {
	// DockerUploadUUID contains the information returned from the Docker-Upload-UUID header response.
	DockerUploadUUID *string

	// Location contains the information returned from the Location header response.
	Location *string

	// Range contains the information returned from the Range header response.
	Range *string
}

// ContainerRegistryBlobClientUploadChunkResponse contains the response from method containerRegistryBlobClient.UploadChunk.
type ContainerRegistryBlobClientUploadChunkResponse struct {
	// DockerUploadUUID contains the information returned from the Docker-Upload-UUID header response.
	DockerUploadUUID *string

	// Location contains the information returned from the Location header response.
	Location *string

	// Range contains the information returned from the Range header response.
	Range *string
}

// ContainerRegistryClientCheckDockerV2SupportResponse contains the response from method containerRegistryClient.CheckDockerV2Support.
type ContainerRegistryClientCheckDockerV2SupportResponse struct {
	// placeholder for future response values
}

// ContainerRegistryClientCreateManifestResponse contains the response from method containerRegistryClient.CreateManifest.
type ContainerRegistryClientCreateManifestResponse struct {
	// ContentLength contains the information returned from the Content-Length header response.
	ContentLength *int64

	// DockerContentDigest contains the information returned from the Docker-Content-Digest header response.
	DockerContentDigest *string

	// Anything
	Interface any

	// Location contains the information returned from the Location header response.
	Location *string
}

// ContainerRegistryClientDeleteManifestResponse contains the response from method containerRegistryClient.DeleteManifest.
type ContainerRegistryClientDeleteManifestResponse struct {
	// placeholder for future response values
}

// ContainerRegistryClientDeleteRepositoryResponse contains the response from method containerRegistryClient.DeleteRepository.
type ContainerRegistryClientDeleteRepositoryResponse struct {
	DeleteRepositoryResult
}

// ContainerRegistryClientDeleteTagResponse contains the response from method containerRegistryClient.DeleteTag.
type ContainerRegistryClientDeleteTagResponse struct {
	// placeholder for future response values
}

// ContainerRegistryClientGetManifestPropertiesResponse contains the response from method containerRegistryClient.GetManifestProperties.
type ContainerRegistryClientGetManifestPropertiesResponse struct {
	ArtifactManifestProperties
}

// ContainerRegistryClientGetManifestResponse contains the response from method containerRegistryClient.GetManifest.
type ContainerRegistryClientGetManifestResponse struct {
	ManifestWrapper
}

// ContainerRegistryClientGetManifestsResponse contains the response from method containerRegistryClient.GetManifests.
type ContainerRegistryClientGetManifestsResponse struct {
	Manifests
	// Link contains the information returned from the Link header response.
	Link *string
}

// ContainerRegistryClientGetPropertiesResponse contains the response from method containerRegistryClient.GetProperties.
type ContainerRegistryClientGetPropertiesResponse struct {
	ContainerRepositoryProperties
}

// ContainerRegistryClientGetRepositoriesResponse contains the response from method containerRegistryClient.GetRepositories.
type ContainerRegistryClientGetRepositoriesResponse struct {
	Repositories
	// Link contains the information returned from the Link header response.
	Link *string
}

// ContainerRegistryClientGetTagPropertiesResponse contains the response from method containerRegistryClient.GetTagProperties.
type ContainerRegistryClientGetTagPropertiesResponse struct {
	ArtifactTagProperties
}

// ContainerRegistryClientGetTagsResponse contains the response from method containerRegistryClient.GetTags.
type ContainerRegistryClientGetTagsResponse struct {
	TagList
	// Link contains the information returned from the Link header response.
	Link *string
}

// ContainerRegistryClientUpdateManifestPropertiesResponse contains the response from method containerRegistryClient.UpdateManifestProperties.
type ContainerRegistryClientUpdateManifestPropertiesResponse struct {
	ArtifactManifestProperties
}

// ContainerRegistryClientUpdatePropertiesResponse contains the response from method containerRegistryClient.UpdateProperties.
type ContainerRegistryClientUpdatePropertiesResponse struct {
	ContainerRepositoryProperties
}

// ContainerRegistryClientUpdateTagAttributesResponse contains the response from method containerRegistryClient.UpdateTagAttributes.
type ContainerRegistryClientUpdateTagAttributesResponse struct {
	ArtifactTagProperties
}
