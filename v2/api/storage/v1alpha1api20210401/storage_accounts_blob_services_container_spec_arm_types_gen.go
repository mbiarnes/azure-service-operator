// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of StorageAccounts_BlobServices_Container_Spec. Use v1beta20210401.StorageAccounts_BlobServices_Container_Spec instead
type StorageAccounts_BlobServices_Container_Spec_ARM struct {
	Location   *string                  `json:"location,omitempty"`
	Name       string                   `json:"name,omitempty"`
	Properties *ContainerProperties_ARM `json:"properties,omitempty"`
	Tags       map[string]string        `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &StorageAccounts_BlobServices_Container_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (container StorageAccounts_BlobServices_Container_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (container *StorageAccounts_BlobServices_Container_Spec_ARM) GetName() string {
	return container.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/blobServices/containers"
func (container *StorageAccounts_BlobServices_Container_Spec_ARM) GetType() string {
	return "Microsoft.Storage/storageAccounts/blobServices/containers"
}

// Deprecated version of ContainerProperties. Use v1beta20210401.ContainerProperties instead
type ContainerProperties_ARM struct {
	DefaultEncryptionScope         *string                             `json:"defaultEncryptionScope,omitempty"`
	DenyEncryptionScopeOverride    *bool                               `json:"denyEncryptionScopeOverride,omitempty"`
	ImmutableStorageWithVersioning *ImmutableStorageWithVersioning_ARM `json:"immutableStorageWithVersioning,omitempty"`
	Metadata                       map[string]string                   `json:"metadata,omitempty"`
	PublicAccess                   *ContainerProperties_PublicAccess   `json:"publicAccess,omitempty"`
}

// Deprecated version of ImmutableStorageWithVersioning. Use v1beta20210401.ImmutableStorageWithVersioning instead
type ImmutableStorageWithVersioning_ARM struct {
	Enabled *bool `json:"enabled,omitempty"`
}
