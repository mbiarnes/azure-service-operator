// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20210101p "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20210101preview"
	v20210101ps "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20210101previewstorage"
	v20211101 "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20211101"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20211101storage"
	v20221001p "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20221001preview"
	v20221001ps "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20221001previewstorage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type NamespacesQueueExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *NamespacesQueueExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20210101p.NamespacesQueue{},
		&v20210101ps.NamespacesQueue{},
		&v20211101.NamespacesQueue{},
		&v20211101s.NamespacesQueue{},
		&v20221001p.NamespacesQueue{},
		&v20221001ps.NamespacesQueue{}}
}
