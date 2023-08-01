// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101storage

import (
	"fmt"
	v20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20201101.RouteTable
// Deprecated version of RouteTable. Use v1api20201101.RouteTable instead
type RouteTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteTable_Spec   `json:"spec,omitempty"`
	Status            RouteTable_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RouteTable{}

// GetConditions returns the conditions of the resource
func (table *RouteTable) GetConditions() conditions.Conditions {
	return table.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (table *RouteTable) SetConditions(conditions conditions.Conditions) {
	table.Status.Conditions = conditions
}

var _ conversion.Convertible = &RouteTable{}

// ConvertFrom populates our RouteTable from the provided hub RouteTable
func (table *RouteTable) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20201101s.RouteTable)
	if !ok {
		return fmt.Errorf("expected network/v1api20201101storage/RouteTable but received %T instead", hub)
	}

	return table.AssignProperties_From_RouteTable(source)
}

// ConvertTo populates the provided hub RouteTable from our RouteTable
func (table *RouteTable) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20201101s.RouteTable)
	if !ok {
		return fmt.Errorf("expected network/v1api20201101storage/RouteTable but received %T instead", hub)
	}

	return table.AssignProperties_To_RouteTable(destination)
}

var _ genruntime.KubernetesResource = &RouteTable{}

// AzureName returns the Azure name of the resource
func (table *RouteTable) AzureName() string {
	return table.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (table RouteTable) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (table *RouteTable) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (table *RouteTable) GetSpec() genruntime.ConvertibleSpec {
	return &table.Spec
}

// GetStatus returns the status of this resource
func (table *RouteTable) GetStatus() genruntime.ConvertibleStatus {
	return &table.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/routeTables"
func (table *RouteTable) GetType() string {
	return "Microsoft.Network/routeTables"
}

// NewEmptyStatus returns a new empty (blank) status
func (table *RouteTable) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RouteTable_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (table *RouteTable) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(table.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  table.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (table *RouteTable) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RouteTable_STATUS); ok {
		table.Status = *st
		return nil
	}

	// Convert status to required version
	var st RouteTable_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	table.Status = st
	return nil
}

// AssignProperties_From_RouteTable populates our RouteTable from the provided source RouteTable
func (table *RouteTable) AssignProperties_From_RouteTable(source *v20201101s.RouteTable) error {

	// ObjectMeta
	table.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec RouteTable_Spec
	err := spec.AssignProperties_From_RouteTable_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_RouteTable_Spec() to populate field Spec")
	}
	table.Spec = spec

	// Status
	var status RouteTable_STATUS
	err = status.AssignProperties_From_RouteTable_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_RouteTable_STATUS() to populate field Status")
	}
	table.Status = status

	// Invoke the augmentConversionForRouteTable interface (if implemented) to customize the conversion
	var tableAsAny any = table
	if augmentedTable, ok := tableAsAny.(augmentConversionForRouteTable); ok {
		err := augmentedTable.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_RouteTable populates the provided destination RouteTable from our RouteTable
func (table *RouteTable) AssignProperties_To_RouteTable(destination *v20201101s.RouteTable) error {

	// ObjectMeta
	destination.ObjectMeta = *table.ObjectMeta.DeepCopy()

	// Spec
	var spec v20201101s.RouteTable_Spec
	err := table.Spec.AssignProperties_To_RouteTable_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_RouteTable_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20201101s.RouteTable_STATUS
	err = table.Status.AssignProperties_To_RouteTable_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_RouteTable_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForRouteTable interface (if implemented) to customize the conversion
	var tableAsAny any = table
	if augmentedTable, ok := tableAsAny.(augmentConversionForRouteTable); ok {
		err := augmentedTable.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (table *RouteTable) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: table.Spec.OriginalVersion,
		Kind:    "RouteTable",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20201101.RouteTable
// Deprecated version of RouteTable. Use v1api20201101.RouteTable instead
type RouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTable `json:"items"`
}

type augmentConversionForRouteTable interface {
	AssignPropertiesFrom(src *v20201101s.RouteTable) error
	AssignPropertiesTo(dst *v20201101s.RouteTable) error
}

// Storage version of v1beta20201101.RouteTable_Spec
type RouteTable_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                  string  `json:"azureName,omitempty"`
	DisableBgpRoutePropagation *bool   `json:"disableBgpRoutePropagation,omitempty"`
	Location                   *string `json:"location,omitempty"`
	OriginalVersion            string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner       *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &RouteTable_Spec{}

// ConvertSpecFrom populates our RouteTable_Spec from the provided source
func (table *RouteTable_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20201101s.RouteTable_Spec)
	if ok {
		// Populate our instance from source
		return table.AssignProperties_From_RouteTable_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20201101s.RouteTable_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = table.AssignProperties_From_RouteTable_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RouteTable_Spec
func (table *RouteTable_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20201101s.RouteTable_Spec)
	if ok {
		// Populate destination from our instance
		return table.AssignProperties_To_RouteTable_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20201101s.RouteTable_Spec{}
	err := table.AssignProperties_To_RouteTable_Spec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_RouteTable_Spec populates our RouteTable_Spec from the provided source RouteTable_Spec
func (table *RouteTable_Spec) AssignProperties_From_RouteTable_Spec(source *v20201101s.RouteTable_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	table.AzureName = source.AzureName

	// DisableBgpRoutePropagation
	if source.DisableBgpRoutePropagation != nil {
		disableBgpRoutePropagation := *source.DisableBgpRoutePropagation
		table.DisableBgpRoutePropagation = &disableBgpRoutePropagation
	} else {
		table.DisableBgpRoutePropagation = nil
	}

	// Location
	table.Location = genruntime.ClonePointerToString(source.Location)

	// OriginalVersion
	table.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		table.Owner = &owner
	} else {
		table.Owner = nil
	}

	// Tags
	table.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		table.PropertyBag = propertyBag
	} else {
		table.PropertyBag = nil
	}

	// Invoke the augmentConversionForRouteTable_Spec interface (if implemented) to customize the conversion
	var tableAsAny any = table
	if augmentedTable, ok := tableAsAny.(augmentConversionForRouteTable_Spec); ok {
		err := augmentedTable.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_RouteTable_Spec populates the provided destination RouteTable_Spec from our RouteTable_Spec
func (table *RouteTable_Spec) AssignProperties_To_RouteTable_Spec(destination *v20201101s.RouteTable_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(table.PropertyBag)

	// AzureName
	destination.AzureName = table.AzureName

	// DisableBgpRoutePropagation
	if table.DisableBgpRoutePropagation != nil {
		disableBgpRoutePropagation := *table.DisableBgpRoutePropagation
		destination.DisableBgpRoutePropagation = &disableBgpRoutePropagation
	} else {
		destination.DisableBgpRoutePropagation = nil
	}

	// Location
	destination.Location = genruntime.ClonePointerToString(table.Location)

	// OriginalVersion
	destination.OriginalVersion = table.OriginalVersion

	// Owner
	if table.Owner != nil {
		owner := table.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(table.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForRouteTable_Spec interface (if implemented) to customize the conversion
	var tableAsAny any = table
	if augmentedTable, ok := tableAsAny.(augmentConversionForRouteTable_Spec); ok {
		err := augmentedTable.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1beta20201101.RouteTable_STATUS
// Deprecated version of RouteTable_STATUS. Use v1api20201101.RouteTable_STATUS instead
type RouteTable_STATUS struct {
	Conditions                 []conditions.Condition `json:"conditions,omitempty"`
	DisableBgpRoutePropagation *bool                  `json:"disableBgpRoutePropagation,omitempty"`
	Etag                       *string                `json:"etag,omitempty"`
	Id                         *string                `json:"id,omitempty"`
	Location                   *string                `json:"location,omitempty"`
	Name                       *string                `json:"name,omitempty"`
	PropertyBag                genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState          *string                `json:"provisioningState,omitempty"`
	ResourceGuid               *string                `json:"resourceGuid,omitempty"`
	Tags                       map[string]string      `json:"tags,omitempty"`
	Type                       *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RouteTable_STATUS{}

// ConvertStatusFrom populates our RouteTable_STATUS from the provided source
func (table *RouteTable_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20201101s.RouteTable_STATUS)
	if ok {
		// Populate our instance from source
		return table.AssignProperties_From_RouteTable_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20201101s.RouteTable_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = table.AssignProperties_From_RouteTable_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our RouteTable_STATUS
func (table *RouteTable_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20201101s.RouteTable_STATUS)
	if ok {
		// Populate destination from our instance
		return table.AssignProperties_To_RouteTable_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20201101s.RouteTable_STATUS{}
	err := table.AssignProperties_To_RouteTable_STATUS(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

// AssignProperties_From_RouteTable_STATUS populates our RouteTable_STATUS from the provided source RouteTable_STATUS
func (table *RouteTable_STATUS) AssignProperties_From_RouteTable_STATUS(source *v20201101s.RouteTable_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	table.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// DisableBgpRoutePropagation
	if source.DisableBgpRoutePropagation != nil {
		disableBgpRoutePropagation := *source.DisableBgpRoutePropagation
		table.DisableBgpRoutePropagation = &disableBgpRoutePropagation
	} else {
		table.DisableBgpRoutePropagation = nil
	}

	// Etag
	table.Etag = genruntime.ClonePointerToString(source.Etag)

	// Id
	table.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	table.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	table.Name = genruntime.ClonePointerToString(source.Name)

	// ProvisioningState
	table.ProvisioningState = genruntime.ClonePointerToString(source.ProvisioningState)

	// ResourceGuid
	table.ResourceGuid = genruntime.ClonePointerToString(source.ResourceGuid)

	// Tags
	table.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	table.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		table.PropertyBag = propertyBag
	} else {
		table.PropertyBag = nil
	}

	// Invoke the augmentConversionForRouteTable_STATUS interface (if implemented) to customize the conversion
	var tableAsAny any = table
	if augmentedTable, ok := tableAsAny.(augmentConversionForRouteTable_STATUS); ok {
		err := augmentedTable.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_RouteTable_STATUS populates the provided destination RouteTable_STATUS from our RouteTable_STATUS
func (table *RouteTable_STATUS) AssignProperties_To_RouteTable_STATUS(destination *v20201101s.RouteTable_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(table.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(table.Conditions)

	// DisableBgpRoutePropagation
	if table.DisableBgpRoutePropagation != nil {
		disableBgpRoutePropagation := *table.DisableBgpRoutePropagation
		destination.DisableBgpRoutePropagation = &disableBgpRoutePropagation
	} else {
		destination.DisableBgpRoutePropagation = nil
	}

	// Etag
	destination.Etag = genruntime.ClonePointerToString(table.Etag)

	// Id
	destination.Id = genruntime.ClonePointerToString(table.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(table.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(table.Name)

	// ProvisioningState
	destination.ProvisioningState = genruntime.ClonePointerToString(table.ProvisioningState)

	// ResourceGuid
	destination.ResourceGuid = genruntime.ClonePointerToString(table.ResourceGuid)

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(table.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(table.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForRouteTable_STATUS interface (if implemented) to customize the conversion
	var tableAsAny any = table
	if augmentedTable, ok := tableAsAny.(augmentConversionForRouteTable_STATUS); ok {
		err := augmentedTable.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForRouteTable_Spec interface {
	AssignPropertiesFrom(src *v20201101s.RouteTable_Spec) error
	AssignPropertiesTo(dst *v20201101s.RouteTable_Spec) error
}

type augmentConversionForRouteTable_STATUS interface {
	AssignPropertiesFrom(src *v20201101s.RouteTable_STATUS) error
	AssignPropertiesTo(dst *v20201101s.RouteTable_STATUS) error
}

func init() {
	SchemeBuilder.Register(&RouteTable{}, &RouteTableList{})
}
