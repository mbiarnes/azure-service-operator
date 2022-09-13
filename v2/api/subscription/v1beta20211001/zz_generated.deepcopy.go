//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta20211001

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Alias) DeepCopyInto(out *Alias) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Alias.
func (in *Alias) DeepCopy() *Alias {
	if in == nil {
		return nil
	}
	out := new(Alias)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Alias) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasList) DeepCopyInto(out *AliasList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Alias, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasList.
func (in *AliasList) DeepCopy() *AliasList {
	if in == nil {
		return nil
	}
	out := new(AliasList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AliasList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Alias_Spec) DeepCopyInto(out *Alias_Spec) {
	*out = *in
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(PutAliasRequestProperties)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Alias_Spec.
func (in *Alias_Spec) DeepCopy() *Alias_Spec {
	if in == nil {
		return nil
	}
	out := new(Alias_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Alias_Spec_ARM) DeepCopyInto(out *Alias_Spec_ARM) {
	*out = *in
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(PutAliasRequestProperties_ARM)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Alias_Spec_ARM.
func (in *Alias_Spec_ARM) DeepCopy() *Alias_Spec_ARM {
	if in == nil {
		return nil
	}
	out := new(Alias_Spec_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PutAliasRequestAdditionalProperties) DeepCopyInto(out *PutAliasRequestAdditionalProperties) {
	*out = *in
	if in.ManagementGroupId != nil {
		in, out := &in.ManagementGroupId, &out.ManagementGroupId
		*out = new(string)
		**out = **in
	}
	if in.SubscriptionOwnerId != nil {
		in, out := &in.SubscriptionOwnerId, &out.SubscriptionOwnerId
		*out = new(string)
		**out = **in
	}
	if in.SubscriptionTenantId != nil {
		in, out := &in.SubscriptionTenantId, &out.SubscriptionTenantId
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PutAliasRequestAdditionalProperties.
func (in *PutAliasRequestAdditionalProperties) DeepCopy() *PutAliasRequestAdditionalProperties {
	if in == nil {
		return nil
	}
	out := new(PutAliasRequestAdditionalProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PutAliasRequestAdditionalProperties_ARM) DeepCopyInto(out *PutAliasRequestAdditionalProperties_ARM) {
	*out = *in
	if in.ManagementGroupId != nil {
		in, out := &in.ManagementGroupId, &out.ManagementGroupId
		*out = new(string)
		**out = **in
	}
	if in.SubscriptionOwnerId != nil {
		in, out := &in.SubscriptionOwnerId, &out.SubscriptionOwnerId
		*out = new(string)
		**out = **in
	}
	if in.SubscriptionTenantId != nil {
		in, out := &in.SubscriptionTenantId, &out.SubscriptionTenantId
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PutAliasRequestAdditionalProperties_ARM.
func (in *PutAliasRequestAdditionalProperties_ARM) DeepCopy() *PutAliasRequestAdditionalProperties_ARM {
	if in == nil {
		return nil
	}
	out := new(PutAliasRequestAdditionalProperties_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PutAliasRequestProperties) DeepCopyInto(out *PutAliasRequestProperties) {
	*out = *in
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = new(PutAliasRequestAdditionalProperties)
		(*in).DeepCopyInto(*out)
	}
	if in.BillingScope != nil {
		in, out := &in.BillingScope, &out.BillingScope
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.ResellerId != nil {
		in, out := &in.ResellerId, &out.ResellerId
		*out = new(string)
		**out = **in
	}
	if in.SubscriptionId != nil {
		in, out := &in.SubscriptionId, &out.SubscriptionId
		*out = new(string)
		**out = **in
	}
	if in.Workload != nil {
		in, out := &in.Workload, &out.Workload
		*out = new(PutAliasRequestProperties_Workload)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PutAliasRequestProperties.
func (in *PutAliasRequestProperties) DeepCopy() *PutAliasRequestProperties {
	if in == nil {
		return nil
	}
	out := new(PutAliasRequestProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PutAliasRequestProperties_ARM) DeepCopyInto(out *PutAliasRequestProperties_ARM) {
	*out = *in
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = new(PutAliasRequestAdditionalProperties_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.BillingScope != nil {
		in, out := &in.BillingScope, &out.BillingScope
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.ResellerId != nil {
		in, out := &in.ResellerId, &out.ResellerId
		*out = new(string)
		**out = **in
	}
	if in.SubscriptionId != nil {
		in, out := &in.SubscriptionId, &out.SubscriptionId
		*out = new(string)
		**out = **in
	}
	if in.Workload != nil {
		in, out := &in.Workload, &out.Workload
		*out = new(PutAliasRequestProperties_Workload)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PutAliasRequestProperties_ARM.
func (in *PutAliasRequestProperties_ARM) DeepCopy() *PutAliasRequestProperties_ARM {
	if in == nil {
		return nil
	}
	out := new(PutAliasRequestProperties_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionAliasResponseProperties_STATUS) DeepCopyInto(out *SubscriptionAliasResponseProperties_STATUS) {
	*out = *in
	if in.AcceptOwnershipState != nil {
		in, out := &in.AcceptOwnershipState, &out.AcceptOwnershipState
		*out = new(AcceptOwnershipState_STATUS)
		**out = **in
	}
	if in.AcceptOwnershipUrl != nil {
		in, out := &in.AcceptOwnershipUrl, &out.AcceptOwnershipUrl
		*out = new(string)
		**out = **in
	}
	if in.BillingScope != nil {
		in, out := &in.BillingScope, &out.BillingScope
		*out = new(string)
		**out = **in
	}
	if in.CreatedTime != nil {
		in, out := &in.CreatedTime, &out.CreatedTime
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.ManagementGroupId != nil {
		in, out := &in.ManagementGroupId, &out.ManagementGroupId
		*out = new(string)
		**out = **in
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(SubscriptionAliasResponseProperties_ProvisioningState_STATUS)
		**out = **in
	}
	if in.ResellerId != nil {
		in, out := &in.ResellerId, &out.ResellerId
		*out = new(string)
		**out = **in
	}
	if in.SubscriptionId != nil {
		in, out := &in.SubscriptionId, &out.SubscriptionId
		*out = new(string)
		**out = **in
	}
	if in.SubscriptionOwnerId != nil {
		in, out := &in.SubscriptionOwnerId, &out.SubscriptionOwnerId
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Workload != nil {
		in, out := &in.Workload, &out.Workload
		*out = new(Workload_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionAliasResponseProperties_STATUS.
func (in *SubscriptionAliasResponseProperties_STATUS) DeepCopy() *SubscriptionAliasResponseProperties_STATUS {
	if in == nil {
		return nil
	}
	out := new(SubscriptionAliasResponseProperties_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionAliasResponseProperties_STATUS_ARM) DeepCopyInto(out *SubscriptionAliasResponseProperties_STATUS_ARM) {
	*out = *in
	if in.AcceptOwnershipState != nil {
		in, out := &in.AcceptOwnershipState, &out.AcceptOwnershipState
		*out = new(AcceptOwnershipState_STATUS)
		**out = **in
	}
	if in.AcceptOwnershipUrl != nil {
		in, out := &in.AcceptOwnershipUrl, &out.AcceptOwnershipUrl
		*out = new(string)
		**out = **in
	}
	if in.BillingScope != nil {
		in, out := &in.BillingScope, &out.BillingScope
		*out = new(string)
		**out = **in
	}
	if in.CreatedTime != nil {
		in, out := &in.CreatedTime, &out.CreatedTime
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.ManagementGroupId != nil {
		in, out := &in.ManagementGroupId, &out.ManagementGroupId
		*out = new(string)
		**out = **in
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(SubscriptionAliasResponseProperties_ProvisioningState_STATUS)
		**out = **in
	}
	if in.ResellerId != nil {
		in, out := &in.ResellerId, &out.ResellerId
		*out = new(string)
		**out = **in
	}
	if in.SubscriptionId != nil {
		in, out := &in.SubscriptionId, &out.SubscriptionId
		*out = new(string)
		**out = **in
	}
	if in.SubscriptionOwnerId != nil {
		in, out := &in.SubscriptionOwnerId, &out.SubscriptionOwnerId
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Workload != nil {
		in, out := &in.Workload, &out.Workload
		*out = new(Workload_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionAliasResponseProperties_STATUS_ARM.
func (in *SubscriptionAliasResponseProperties_STATUS_ARM) DeepCopy() *SubscriptionAliasResponseProperties_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(SubscriptionAliasResponseProperties_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionAliasResponse_STATUS) DeepCopyInto(out *SubscriptionAliasResponse_STATUS) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(SubscriptionAliasResponseProperties_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.SystemData != nil {
		in, out := &in.SystemData, &out.SystemData
		*out = new(SystemData_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionAliasResponse_STATUS.
func (in *SubscriptionAliasResponse_STATUS) DeepCopy() *SubscriptionAliasResponse_STATUS {
	if in == nil {
		return nil
	}
	out := new(SubscriptionAliasResponse_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionAliasResponse_STATUS_ARM) DeepCopyInto(out *SubscriptionAliasResponse_STATUS_ARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(SubscriptionAliasResponseProperties_STATUS_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.SystemData != nil {
		in, out := &in.SystemData, &out.SystemData
		*out = new(SystemData_STATUS_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionAliasResponse_STATUS_ARM.
func (in *SubscriptionAliasResponse_STATUS_ARM) DeepCopy() *SubscriptionAliasResponse_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(SubscriptionAliasResponse_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemData_STATUS) DeepCopyInto(out *SystemData_STATUS) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.CreatedBy != nil {
		in, out := &in.CreatedBy, &out.CreatedBy
		*out = new(string)
		**out = **in
	}
	if in.CreatedByType != nil {
		in, out := &in.CreatedByType, &out.CreatedByType
		*out = new(SystemData_CreatedByType_STATUS)
		**out = **in
	}
	if in.LastModifiedAt != nil {
		in, out := &in.LastModifiedAt, &out.LastModifiedAt
		*out = new(string)
		**out = **in
	}
	if in.LastModifiedBy != nil {
		in, out := &in.LastModifiedBy, &out.LastModifiedBy
		*out = new(string)
		**out = **in
	}
	if in.LastModifiedByType != nil {
		in, out := &in.LastModifiedByType, &out.LastModifiedByType
		*out = new(SystemData_LastModifiedByType_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemData_STATUS.
func (in *SystemData_STATUS) DeepCopy() *SystemData_STATUS {
	if in == nil {
		return nil
	}
	out := new(SystemData_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemData_STATUS_ARM) DeepCopyInto(out *SystemData_STATUS_ARM) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.CreatedBy != nil {
		in, out := &in.CreatedBy, &out.CreatedBy
		*out = new(string)
		**out = **in
	}
	if in.CreatedByType != nil {
		in, out := &in.CreatedByType, &out.CreatedByType
		*out = new(SystemData_CreatedByType_STATUS)
		**out = **in
	}
	if in.LastModifiedAt != nil {
		in, out := &in.LastModifiedAt, &out.LastModifiedAt
		*out = new(string)
		**out = **in
	}
	if in.LastModifiedBy != nil {
		in, out := &in.LastModifiedBy, &out.LastModifiedBy
		*out = new(string)
		**out = **in
	}
	if in.LastModifiedByType != nil {
		in, out := &in.LastModifiedByType, &out.LastModifiedByType
		*out = new(SystemData_LastModifiedByType_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemData_STATUS_ARM.
func (in *SystemData_STATUS_ARM) DeepCopy() *SystemData_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(SystemData_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}
