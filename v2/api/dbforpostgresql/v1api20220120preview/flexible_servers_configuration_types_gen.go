// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220120preview

import (
	"fmt"
	v20220120ps "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20220120previewstorage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Generator information:
// - Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2022-01-20-preview/postgresql.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/configurations/{configurationName}
type FlexibleServersConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServers_Configuration_Spec   `json:"spec,omitempty"`
	Status            FlexibleServers_Configuration_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &FlexibleServersConfiguration{}

// GetConditions returns the conditions of the resource
func (configuration *FlexibleServersConfiguration) GetConditions() conditions.Conditions {
	return configuration.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (configuration *FlexibleServersConfiguration) SetConditions(conditions conditions.Conditions) {
	configuration.Status.Conditions = conditions
}

var _ conversion.Convertible = &FlexibleServersConfiguration{}

// ConvertFrom populates our FlexibleServersConfiguration from the provided hub FlexibleServersConfiguration
func (configuration *FlexibleServersConfiguration) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source v20220120ps.FlexibleServersConfiguration

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = configuration.AssignProperties_From_FlexibleServersConfiguration(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to configuration")
	}

	return nil
}

// ConvertTo populates the provided hub FlexibleServersConfiguration from our FlexibleServersConfiguration
func (configuration *FlexibleServersConfiguration) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination v20220120ps.FlexibleServersConfiguration
	err := configuration.AssignProperties_To_FlexibleServersConfiguration(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from configuration")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-dbforpostgresql-azure-com-v1api20220120preview-flexibleserversconfiguration,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbforpostgresql.azure.com,resources=flexibleserversconfigurations,verbs=create;update,versions=v1api20220120preview,name=default.v1api20220120preview.flexibleserversconfigurations.dbforpostgresql.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &FlexibleServersConfiguration{}

// Default applies defaults to the FlexibleServersConfiguration resource
func (configuration *FlexibleServersConfiguration) Default() {
	configuration.defaultImpl()
	var temp any = configuration
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (configuration *FlexibleServersConfiguration) defaultAzureName() {
	if configuration.Spec.AzureName == "" {
		configuration.Spec.AzureName = configuration.Name
	}
}

// defaultImpl applies the code generated defaults to the FlexibleServersConfiguration resource
func (configuration *FlexibleServersConfiguration) defaultImpl() { configuration.defaultAzureName() }

var _ genruntime.KubernetesResource = &FlexibleServersConfiguration{}

// AzureName returns the Azure name of the resource
func (configuration *FlexibleServersConfiguration) AzureName() string {
	return configuration.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-01-20-preview"
func (configuration FlexibleServersConfiguration) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (configuration *FlexibleServersConfiguration) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (configuration *FlexibleServersConfiguration) GetSpec() genruntime.ConvertibleSpec {
	return &configuration.Spec
}

// GetStatus returns the status of this resource
func (configuration *FlexibleServersConfiguration) GetStatus() genruntime.ConvertibleStatus {
	return &configuration.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers/configurations"
func (configuration *FlexibleServersConfiguration) GetType() string {
	return "Microsoft.DBforPostgreSQL/flexibleServers/configurations"
}

// NewEmptyStatus returns a new empty (blank) status
func (configuration *FlexibleServersConfiguration) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &FlexibleServers_Configuration_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (configuration *FlexibleServersConfiguration) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(configuration.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  configuration.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (configuration *FlexibleServersConfiguration) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*FlexibleServers_Configuration_STATUS); ok {
		configuration.Status = *st
		return nil
	}

	// Convert status to required version
	var st FlexibleServers_Configuration_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	configuration.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-dbforpostgresql-azure-com-v1api20220120preview-flexibleserversconfiguration,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbforpostgresql.azure.com,resources=flexibleserversconfigurations,verbs=create;update,versions=v1api20220120preview,name=validate.v1api20220120preview.flexibleserversconfigurations.dbforpostgresql.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &FlexibleServersConfiguration{}

// ValidateCreate validates the creation of the resource
func (configuration *FlexibleServersConfiguration) ValidateCreate() (admission.Warnings, error) {
	validations := configuration.createValidations()
	var temp any = configuration
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (configuration *FlexibleServersConfiguration) ValidateDelete() (admission.Warnings, error) {
	validations := configuration.deleteValidations()
	var temp any = configuration
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (configuration *FlexibleServersConfiguration) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := configuration.updateValidations()
	var temp any = configuration
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (configuration *FlexibleServersConfiguration) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){configuration.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (configuration *FlexibleServersConfiguration) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (configuration *FlexibleServersConfiguration) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return configuration.validateResourceReferences()
		},
		configuration.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (configuration *FlexibleServersConfiguration) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&configuration.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (configuration *FlexibleServersConfiguration) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*FlexibleServersConfiguration)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, configuration)
}

// AssignProperties_From_FlexibleServersConfiguration populates our FlexibleServersConfiguration from the provided source FlexibleServersConfiguration
func (configuration *FlexibleServersConfiguration) AssignProperties_From_FlexibleServersConfiguration(source *v20220120ps.FlexibleServersConfiguration) error {

	// ObjectMeta
	configuration.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec FlexibleServers_Configuration_Spec
	err := spec.AssignProperties_From_FlexibleServers_Configuration_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_FlexibleServers_Configuration_Spec() to populate field Spec")
	}
	configuration.Spec = spec

	// Status
	var status FlexibleServers_Configuration_STATUS
	err = status.AssignProperties_From_FlexibleServers_Configuration_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_FlexibleServers_Configuration_STATUS() to populate field Status")
	}
	configuration.Status = status

	// No error
	return nil
}

// AssignProperties_To_FlexibleServersConfiguration populates the provided destination FlexibleServersConfiguration from our FlexibleServersConfiguration
func (configuration *FlexibleServersConfiguration) AssignProperties_To_FlexibleServersConfiguration(destination *v20220120ps.FlexibleServersConfiguration) error {

	// ObjectMeta
	destination.ObjectMeta = *configuration.ObjectMeta.DeepCopy()

	// Spec
	var spec v20220120ps.FlexibleServers_Configuration_Spec
	err := configuration.Spec.AssignProperties_To_FlexibleServers_Configuration_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_FlexibleServers_Configuration_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20220120ps.FlexibleServers_Configuration_STATUS
	err = configuration.Status.AssignProperties_To_FlexibleServers_Configuration_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_FlexibleServers_Configuration_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (configuration *FlexibleServersConfiguration) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: configuration.Spec.OriginalVersion(),
		Kind:    "FlexibleServersConfiguration",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2022-01-20-preview/postgresql.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/configurations/{configurationName}
type FlexibleServersConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServersConfiguration `json:"items"`
}

type FlexibleServers_Configuration_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a dbforpostgresql.azure.com/FlexibleServer resource
	Owner *genruntime.KnownResourceReference `group:"dbforpostgresql.azure.com" json:"owner,omitempty" kind:"FlexibleServer"`

	// Source: Source of the configuration.
	Source *string `json:"source,omitempty"`

	// Value: Value of the configuration.
	Value *string `json:"value,omitempty"`
}

var _ genruntime.ARMTransformer = &FlexibleServers_Configuration_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (configuration *FlexibleServers_Configuration_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if configuration == nil {
		return nil, nil
	}
	result := &FlexibleServers_Configuration_Spec_ARM{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if configuration.Source != nil || configuration.Value != nil {
		result.Properties = &ConfigurationProperties_ARM{}
	}
	if configuration.Source != nil {
		source := *configuration.Source
		result.Properties.Source = &source
	}
	if configuration.Value != nil {
		value := *configuration.Value
		result.Properties.Value = &value
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (configuration *FlexibleServers_Configuration_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &FlexibleServers_Configuration_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (configuration *FlexibleServers_Configuration_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(FlexibleServers_Configuration_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected FlexibleServers_Configuration_Spec_ARM, got %T", armInput)
	}

	// Set property "AzureName":
	configuration.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property "Owner":
	configuration.Owner = &genruntime.KnownResourceReference{Name: owner.Name}

	// Set property "Source":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Source != nil {
			source := *typedInput.Properties.Source
			configuration.Source = &source
		}
	}

	// Set property "Value":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Value != nil {
			value := *typedInput.Properties.Value
			configuration.Value = &value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &FlexibleServers_Configuration_Spec{}

// ConvertSpecFrom populates our FlexibleServers_Configuration_Spec from the provided source
func (configuration *FlexibleServers_Configuration_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20220120ps.FlexibleServers_Configuration_Spec)
	if ok {
		// Populate our instance from source
		return configuration.AssignProperties_From_FlexibleServers_Configuration_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20220120ps.FlexibleServers_Configuration_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = configuration.AssignProperties_From_FlexibleServers_Configuration_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our FlexibleServers_Configuration_Spec
func (configuration *FlexibleServers_Configuration_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20220120ps.FlexibleServers_Configuration_Spec)
	if ok {
		// Populate destination from our instance
		return configuration.AssignProperties_To_FlexibleServers_Configuration_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20220120ps.FlexibleServers_Configuration_Spec{}
	err := configuration.AssignProperties_To_FlexibleServers_Configuration_Spec(dst)
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

// AssignProperties_From_FlexibleServers_Configuration_Spec populates our FlexibleServers_Configuration_Spec from the provided source FlexibleServers_Configuration_Spec
func (configuration *FlexibleServers_Configuration_Spec) AssignProperties_From_FlexibleServers_Configuration_Spec(source *v20220120ps.FlexibleServers_Configuration_Spec) error {

	// AzureName
	configuration.AzureName = source.AzureName

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		configuration.Owner = &owner
	} else {
		configuration.Owner = nil
	}

	// Source
	configuration.Source = genruntime.ClonePointerToString(source.Source)

	// Value
	configuration.Value = genruntime.ClonePointerToString(source.Value)

	// No error
	return nil
}

// AssignProperties_To_FlexibleServers_Configuration_Spec populates the provided destination FlexibleServers_Configuration_Spec from our FlexibleServers_Configuration_Spec
func (configuration *FlexibleServers_Configuration_Spec) AssignProperties_To_FlexibleServers_Configuration_Spec(destination *v20220120ps.FlexibleServers_Configuration_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = configuration.AzureName

	// OriginalVersion
	destination.OriginalVersion = configuration.OriginalVersion()

	// Owner
	if configuration.Owner != nil {
		owner := configuration.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Source
	destination.Source = genruntime.ClonePointerToString(configuration.Source)

	// Value
	destination.Value = genruntime.ClonePointerToString(configuration.Value)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (configuration *FlexibleServers_Configuration_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (configuration *FlexibleServers_Configuration_Spec) SetAzureName(azureName string) {
	configuration.AzureName = azureName
}

type FlexibleServers_Configuration_STATUS struct {
	// AllowedValues: Allowed values of the configuration.
	AllowedValues *string `json:"allowedValues,omitempty"`

	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// DataType: Data type of the configuration.
	DataType *ConfigurationProperties_DataType_STATUS `json:"dataType,omitempty"`

	// DefaultValue: Default value of the configuration.
	DefaultValue *string `json:"defaultValue,omitempty"`

	// Description: Description of the configuration.
	Description *string `json:"description,omitempty"`

	// DocumentationLink: Configuration documentation link.
	DocumentationLink *string `json:"documentationLink,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// IsConfigPendingRestart: Configuration is pending restart or not.
	IsConfigPendingRestart *bool `json:"isConfigPendingRestart,omitempty"`

	// IsDynamicConfig: Configuration dynamic or static.
	IsDynamicConfig *bool `json:"isDynamicConfig,omitempty"`

	// IsReadOnly: Configuration read-only or not.
	IsReadOnly *bool `json:"isReadOnly,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Source: Source of the configuration.
	Source *string `json:"source,omitempty"`

	// SystemData: The system metadata relating to this resource.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`

	// Unit: Configuration unit.
	Unit *string `json:"unit,omitempty"`

	// Value: Value of the configuration.
	Value *string `json:"value,omitempty"`
}

var _ genruntime.ConvertibleStatus = &FlexibleServers_Configuration_STATUS{}

// ConvertStatusFrom populates our FlexibleServers_Configuration_STATUS from the provided source
func (configuration *FlexibleServers_Configuration_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20220120ps.FlexibleServers_Configuration_STATUS)
	if ok {
		// Populate our instance from source
		return configuration.AssignProperties_From_FlexibleServers_Configuration_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20220120ps.FlexibleServers_Configuration_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = configuration.AssignProperties_From_FlexibleServers_Configuration_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our FlexibleServers_Configuration_STATUS
func (configuration *FlexibleServers_Configuration_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20220120ps.FlexibleServers_Configuration_STATUS)
	if ok {
		// Populate destination from our instance
		return configuration.AssignProperties_To_FlexibleServers_Configuration_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20220120ps.FlexibleServers_Configuration_STATUS{}
	err := configuration.AssignProperties_To_FlexibleServers_Configuration_STATUS(dst)
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

var _ genruntime.FromARMConverter = &FlexibleServers_Configuration_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (configuration *FlexibleServers_Configuration_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &FlexibleServers_Configuration_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (configuration *FlexibleServers_Configuration_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(FlexibleServers_Configuration_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected FlexibleServers_Configuration_STATUS_ARM, got %T", armInput)
	}

	// Set property "AllowedValues":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.AllowedValues != nil {
			allowedValues := *typedInput.Properties.AllowedValues
			configuration.AllowedValues = &allowedValues
		}
	}

	// no assignment for property "Conditions"

	// Set property "DataType":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.DataType != nil {
			dataType := *typedInput.Properties.DataType
			configuration.DataType = &dataType
		}
	}

	// Set property "DefaultValue":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.DefaultValue != nil {
			defaultValue := *typedInput.Properties.DefaultValue
			configuration.DefaultValue = &defaultValue
		}
	}

	// Set property "Description":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Description != nil {
			description := *typedInput.Properties.Description
			configuration.Description = &description
		}
	}

	// Set property "DocumentationLink":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.DocumentationLink != nil {
			documentationLink := *typedInput.Properties.DocumentationLink
			configuration.DocumentationLink = &documentationLink
		}
	}

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		configuration.Id = &id
	}

	// Set property "IsConfigPendingRestart":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.IsConfigPendingRestart != nil {
			isConfigPendingRestart := *typedInput.Properties.IsConfigPendingRestart
			configuration.IsConfigPendingRestart = &isConfigPendingRestart
		}
	}

	// Set property "IsDynamicConfig":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.IsDynamicConfig != nil {
			isDynamicConfig := *typedInput.Properties.IsDynamicConfig
			configuration.IsDynamicConfig = &isDynamicConfig
		}
	}

	// Set property "IsReadOnly":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.IsReadOnly != nil {
			isReadOnly := *typedInput.Properties.IsReadOnly
			configuration.IsReadOnly = &isReadOnly
		}
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		configuration.Name = &name
	}

	// Set property "Source":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Source != nil {
			source := *typedInput.Properties.Source
			configuration.Source = &source
		}
	}

	// Set property "SystemData":
	if typedInput.SystemData != nil {
		var systemData1 SystemData_STATUS
		err := systemData1.PopulateFromARM(owner, *typedInput.SystemData)
		if err != nil {
			return err
		}
		systemData := systemData1
		configuration.SystemData = &systemData
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		configuration.Type = &typeVar
	}

	// Set property "Unit":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Unit != nil {
			unit := *typedInput.Properties.Unit
			configuration.Unit = &unit
		}
	}

	// Set property "Value":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Value != nil {
			value := *typedInput.Properties.Value
			configuration.Value = &value
		}
	}

	// No error
	return nil
}

// AssignProperties_From_FlexibleServers_Configuration_STATUS populates our FlexibleServers_Configuration_STATUS from the provided source FlexibleServers_Configuration_STATUS
func (configuration *FlexibleServers_Configuration_STATUS) AssignProperties_From_FlexibleServers_Configuration_STATUS(source *v20220120ps.FlexibleServers_Configuration_STATUS) error {

	// AllowedValues
	configuration.AllowedValues = genruntime.ClonePointerToString(source.AllowedValues)

	// Conditions
	configuration.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// DataType
	if source.DataType != nil {
		dataType := ConfigurationProperties_DataType_STATUS(*source.DataType)
		configuration.DataType = &dataType
	} else {
		configuration.DataType = nil
	}

	// DefaultValue
	configuration.DefaultValue = genruntime.ClonePointerToString(source.DefaultValue)

	// Description
	configuration.Description = genruntime.ClonePointerToString(source.Description)

	// DocumentationLink
	configuration.DocumentationLink = genruntime.ClonePointerToString(source.DocumentationLink)

	// Id
	configuration.Id = genruntime.ClonePointerToString(source.Id)

	// IsConfigPendingRestart
	if source.IsConfigPendingRestart != nil {
		isConfigPendingRestart := *source.IsConfigPendingRestart
		configuration.IsConfigPendingRestart = &isConfigPendingRestart
	} else {
		configuration.IsConfigPendingRestart = nil
	}

	// IsDynamicConfig
	if source.IsDynamicConfig != nil {
		isDynamicConfig := *source.IsDynamicConfig
		configuration.IsDynamicConfig = &isDynamicConfig
	} else {
		configuration.IsDynamicConfig = nil
	}

	// IsReadOnly
	if source.IsReadOnly != nil {
		isReadOnly := *source.IsReadOnly
		configuration.IsReadOnly = &isReadOnly
	} else {
		configuration.IsReadOnly = nil
	}

	// Name
	configuration.Name = genruntime.ClonePointerToString(source.Name)

	// Source
	configuration.Source = genruntime.ClonePointerToString(source.Source)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
		}
		configuration.SystemData = &systemDatum
	} else {
		configuration.SystemData = nil
	}

	// Type
	configuration.Type = genruntime.ClonePointerToString(source.Type)

	// Unit
	configuration.Unit = genruntime.ClonePointerToString(source.Unit)

	// Value
	configuration.Value = genruntime.ClonePointerToString(source.Value)

	// No error
	return nil
}

// AssignProperties_To_FlexibleServers_Configuration_STATUS populates the provided destination FlexibleServers_Configuration_STATUS from our FlexibleServers_Configuration_STATUS
func (configuration *FlexibleServers_Configuration_STATUS) AssignProperties_To_FlexibleServers_Configuration_STATUS(destination *v20220120ps.FlexibleServers_Configuration_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AllowedValues
	destination.AllowedValues = genruntime.ClonePointerToString(configuration.AllowedValues)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(configuration.Conditions)

	// DataType
	if configuration.DataType != nil {
		dataType := string(*configuration.DataType)
		destination.DataType = &dataType
	} else {
		destination.DataType = nil
	}

	// DefaultValue
	destination.DefaultValue = genruntime.ClonePointerToString(configuration.DefaultValue)

	// Description
	destination.Description = genruntime.ClonePointerToString(configuration.Description)

	// DocumentationLink
	destination.DocumentationLink = genruntime.ClonePointerToString(configuration.DocumentationLink)

	// Id
	destination.Id = genruntime.ClonePointerToString(configuration.Id)

	// IsConfigPendingRestart
	if configuration.IsConfigPendingRestart != nil {
		isConfigPendingRestart := *configuration.IsConfigPendingRestart
		destination.IsConfigPendingRestart = &isConfigPendingRestart
	} else {
		destination.IsConfigPendingRestart = nil
	}

	// IsDynamicConfig
	if configuration.IsDynamicConfig != nil {
		isDynamicConfig := *configuration.IsDynamicConfig
		destination.IsDynamicConfig = &isDynamicConfig
	} else {
		destination.IsDynamicConfig = nil
	}

	// IsReadOnly
	if configuration.IsReadOnly != nil {
		isReadOnly := *configuration.IsReadOnly
		destination.IsReadOnly = &isReadOnly
	} else {
		destination.IsReadOnly = nil
	}

	// Name
	destination.Name = genruntime.ClonePointerToString(configuration.Name)

	// Source
	destination.Source = genruntime.ClonePointerToString(configuration.Source)

	// SystemData
	if configuration.SystemData != nil {
		var systemDatum v20220120ps.SystemData_STATUS
		err := configuration.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(configuration.Type)

	// Unit
	destination.Unit = genruntime.ClonePointerToString(configuration.Unit)

	// Value
	destination.Value = genruntime.ClonePointerToString(configuration.Value)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type ConfigurationProperties_DataType_STATUS string

const (
	ConfigurationProperties_DataType_STATUS_Boolean     = ConfigurationProperties_DataType_STATUS("Boolean")
	ConfigurationProperties_DataType_STATUS_Enumeration = ConfigurationProperties_DataType_STATUS("Enumeration")
	ConfigurationProperties_DataType_STATUS_Integer     = ConfigurationProperties_DataType_STATUS("Integer")
	ConfigurationProperties_DataType_STATUS_Numeric     = ConfigurationProperties_DataType_STATUS("Numeric")
)

func init() {
	SchemeBuilder.Register(&FlexibleServersConfiguration{}, &FlexibleServersConfigurationList{})
}
