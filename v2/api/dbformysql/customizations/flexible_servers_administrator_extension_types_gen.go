// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20220101 "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1api20220101"
	v20220101s "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1api20220101storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type FlexibleServersAdministratorExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *FlexibleServersAdministratorExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20220101.FlexibleServersAdministrator{},
		&v20220101s.FlexibleServersAdministrator{}}
}
