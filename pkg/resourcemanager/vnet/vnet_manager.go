// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package vnet

import (
	"context"

	vnetwork "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/go-autorest/autorest"
)

// NewAzureVNetManager creates a new instance of AzureVNetManager
func NewAzureVNetManager(creds config.Credentials) *AzureVNetManager {
	return &AzureVNetManager{Creds: creds}
}

// VNetManager manages VNet service components
type VNetManager interface {
	CreateVNet(ctx context.Context,
		location string,
		resourceGroupName string,
		resourceName string,
		addressSpace string,
		subnets []azurev1alpha1.VNetSubnets) (vnetwork.VirtualNetwork, error)

	DeleteVNet(ctx context.Context,
		resourceGroupName string,
		resourceName string) (autorest.Response, error)

	GetVNet(ctx context.Context,
		resourceGroupName string,
		resourceName string) (vnetwork.VirtualNetwork, error)

	GetAvailableIP(ctc context.Context, resourceGroup, vnet, subnet string) (string, error)

	// also embed async client methods
	resourcemanager.ARMClient
}
