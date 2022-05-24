//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta20200601

import (
	"github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointConnection_Status) DeepCopyInto(out *PrivateEndpointConnection_Status) {
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
	if in.PrivateEndpoint != nil {
		in, out := &in.PrivateEndpoint, &out.PrivateEndpoint
		*out = new(PrivateEndpoint_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.PrivateLinkServiceConnectionState != nil {
		in, out := &in.PrivateLinkServiceConnectionState, &out.PrivateLinkServiceConnectionState
		*out = new(PrivateLinkServiceConnectionState_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(PrivateEndpointConnectionProvisioningState_Status)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointConnection_Status.
func (in *PrivateEndpointConnection_Status) DeepCopy() *PrivateEndpointConnection_Status {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointConnection_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpoint_Status) DeepCopyInto(out *PrivateEndpoint_Status) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpoint_Status.
func (in *PrivateEndpoint_Status) DeepCopy() *PrivateEndpoint_Status {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpoint_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateLinkServiceConnectionState_Status) DeepCopyInto(out *PrivateLinkServiceConnectionState_Status) {
	*out = *in
	if in.ActionsRequired != nil {
		in, out := &in.ActionsRequired, &out.ActionsRequired
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(PrivateEndpointServiceConnectionStatus_Status)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateLinkServiceConnectionState_Status.
func (in *PrivateLinkServiceConnectionState_Status) DeepCopy() *PrivateLinkServiceConnectionState_Status {
	if in == nil {
		return nil
	}
	out := new(PrivateLinkServiceConnectionState_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Redis) DeepCopyInto(out *Redis) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Redis.
func (in *Redis) DeepCopy() *Redis {
	if in == nil {
		return nil
	}
	out := new(Redis)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Redis) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisAccessKeys_Status) DeepCopyInto(out *RedisAccessKeys_Status) {
	*out = *in
	if in.PrimaryKey != nil {
		in, out := &in.PrimaryKey, &out.PrimaryKey
		*out = new(string)
		**out = **in
	}
	if in.SecondaryKey != nil {
		in, out := &in.SecondaryKey, &out.SecondaryKey
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisAccessKeys_Status.
func (in *RedisAccessKeys_Status) DeepCopy() *RedisAccessKeys_Status {
	if in == nil {
		return nil
	}
	out := new(RedisAccessKeys_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisCommonPropertiesRedisConfiguration) DeepCopyInto(out *RedisCommonPropertiesRedisConfiguration) {
	*out = *in
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AofStorageConnectionString0 != nil {
		in, out := &in.AofStorageConnectionString0, &out.AofStorageConnectionString0
		*out = new(string)
		**out = **in
	}
	if in.AofStorageConnectionString1 != nil {
		in, out := &in.AofStorageConnectionString1, &out.AofStorageConnectionString1
		*out = new(string)
		**out = **in
	}
	if in.MaxfragmentationmemoryReserved != nil {
		in, out := &in.MaxfragmentationmemoryReserved, &out.MaxfragmentationmemoryReserved
		*out = new(string)
		**out = **in
	}
	if in.MaxmemoryDelta != nil {
		in, out := &in.MaxmemoryDelta, &out.MaxmemoryDelta
		*out = new(string)
		**out = **in
	}
	if in.MaxmemoryPolicy != nil {
		in, out := &in.MaxmemoryPolicy, &out.MaxmemoryPolicy
		*out = new(string)
		**out = **in
	}
	if in.MaxmemoryReserved != nil {
		in, out := &in.MaxmemoryReserved, &out.MaxmemoryReserved
		*out = new(string)
		**out = **in
	}
	if in.RdbBackupEnabled != nil {
		in, out := &in.RdbBackupEnabled, &out.RdbBackupEnabled
		*out = new(string)
		**out = **in
	}
	if in.RdbBackupFrequency != nil {
		in, out := &in.RdbBackupFrequency, &out.RdbBackupFrequency
		*out = new(string)
		**out = **in
	}
	if in.RdbBackupMaxSnapshotCount != nil {
		in, out := &in.RdbBackupMaxSnapshotCount, &out.RdbBackupMaxSnapshotCount
		*out = new(string)
		**out = **in
	}
	if in.RdbStorageConnectionString != nil {
		in, out := &in.RdbStorageConnectionString, &out.RdbStorageConnectionString
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisCommonPropertiesRedisConfiguration.
func (in *RedisCommonPropertiesRedisConfiguration) DeepCopy() *RedisCommonPropertiesRedisConfiguration {
	if in == nil {
		return nil
	}
	out := new(RedisCommonPropertiesRedisConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisInstanceDetails_Status) DeepCopyInto(out *RedisInstanceDetails_Status) {
	*out = *in
	if in.IsMaster != nil {
		in, out := &in.IsMaster, &out.IsMaster
		*out = new(bool)
		**out = **in
	}
	if in.NonSslPort != nil {
		in, out := &in.NonSslPort, &out.NonSslPort
		*out = new(int)
		**out = **in
	}
	if in.ShardId != nil {
		in, out := &in.ShardId, &out.ShardId
		*out = new(int)
		**out = **in
	}
	if in.SslPort != nil {
		in, out := &in.SslPort, &out.SslPort
		*out = new(int)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisInstanceDetails_Status.
func (in *RedisInstanceDetails_Status) DeepCopy() *RedisInstanceDetails_Status {
	if in == nil {
		return nil
	}
	out := new(RedisInstanceDetails_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisLinkedServer_Status) DeepCopyInto(out *RedisLinkedServer_Status) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisLinkedServer_Status.
func (in *RedisLinkedServer_Status) DeepCopy() *RedisLinkedServer_Status {
	if in == nil {
		return nil
	}
	out := new(RedisLinkedServer_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisList) DeepCopyInto(out *RedisList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Redis, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisList.
func (in *RedisList) DeepCopy() *RedisList {
	if in == nil {
		return nil
	}
	out := new(RedisList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisParameters) DeepCopyInto(out *RedisParameters) {
	*out = *in
	if in.EnableNonSslPort != nil {
		in, out := &in.EnableNonSslPort, &out.EnableNonSslPort
		*out = new(bool)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MinimumTlsVersion != nil {
		in, out := &in.MinimumTlsVersion, &out.MinimumTlsVersion
		*out = new(RedisCreatePropertiesMinimumTlsVersion)
		**out = **in
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(RedisCreatePropertiesPublicNetworkAccess)
		**out = **in
	}
	if in.RedisConfiguration != nil {
		in, out := &in.RedisConfiguration, &out.RedisConfiguration
		*out = new(RedisCommonPropertiesRedisConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ReplicasPerMaster != nil {
		in, out := &in.ReplicasPerMaster, &out.ReplicasPerMaster
		*out = new(int)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1alpha1.Reference)
		**out = **in
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ShardCount != nil {
		in, out := &in.ShardCount, &out.ShardCount
		*out = new(int)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku)
		(*in).DeepCopyInto(*out)
	}
	if in.StaticIP != nil {
		in, out := &in.StaticIP, &out.StaticIP
		*out = new(string)
		**out = **in
	}
	if in.SubnetId != nil {
		in, out := &in.SubnetId, &out.SubnetId
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
	if in.TenantSettings != nil {
		in, out := &in.TenantSettings, &out.TenantSettings
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisParameters.
func (in *RedisParameters) DeepCopy() *RedisParameters {
	if in == nil {
		return nil
	}
	out := new(RedisParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisProperties_Status_RedisConfiguration) DeepCopyInto(out *RedisProperties_Status_RedisConfiguration) {
	*out = *in
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AofStorageConnectionString0 != nil {
		in, out := &in.AofStorageConnectionString0, &out.AofStorageConnectionString0
		*out = new(string)
		**out = **in
	}
	if in.AofStorageConnectionString1 != nil {
		in, out := &in.AofStorageConnectionString1, &out.AofStorageConnectionString1
		*out = new(string)
		**out = **in
	}
	if in.Maxclients != nil {
		in, out := &in.Maxclients, &out.Maxclients
		*out = new(string)
		**out = **in
	}
	if in.MaxfragmentationmemoryReserved != nil {
		in, out := &in.MaxfragmentationmemoryReserved, &out.MaxfragmentationmemoryReserved
		*out = new(string)
		**out = **in
	}
	if in.MaxmemoryDelta != nil {
		in, out := &in.MaxmemoryDelta, &out.MaxmemoryDelta
		*out = new(string)
		**out = **in
	}
	if in.MaxmemoryPolicy != nil {
		in, out := &in.MaxmemoryPolicy, &out.MaxmemoryPolicy
		*out = new(string)
		**out = **in
	}
	if in.MaxmemoryReserved != nil {
		in, out := &in.MaxmemoryReserved, &out.MaxmemoryReserved
		*out = new(string)
		**out = **in
	}
	if in.RdbBackupEnabled != nil {
		in, out := &in.RdbBackupEnabled, &out.RdbBackupEnabled
		*out = new(string)
		**out = **in
	}
	if in.RdbBackupFrequency != nil {
		in, out := &in.RdbBackupFrequency, &out.RdbBackupFrequency
		*out = new(string)
		**out = **in
	}
	if in.RdbBackupMaxSnapshotCount != nil {
		in, out := &in.RdbBackupMaxSnapshotCount, &out.RdbBackupMaxSnapshotCount
		*out = new(string)
		**out = **in
	}
	if in.RdbStorageConnectionString != nil {
		in, out := &in.RdbStorageConnectionString, &out.RdbStorageConnectionString
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisProperties_Status_RedisConfiguration.
func (in *RedisProperties_Status_RedisConfiguration) DeepCopy() *RedisProperties_Status_RedisConfiguration {
	if in == nil {
		return nil
	}
	out := new(RedisProperties_Status_RedisConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisResourceObservation) DeepCopyInto(out *RedisResourceObservation) {
	*out = *in
	if in.AccessKeys != nil {
		in, out := &in.AccessKeys, &out.AccessKeys
		*out = new(RedisAccessKeys_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.EnableNonSslPort != nil {
		in, out := &in.EnableNonSslPort, &out.EnableNonSslPort
		*out = new(bool)
		**out = **in
	}
	if in.HostName != nil {
		in, out := &in.HostName, &out.HostName
		*out = new(string)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = make([]RedisInstanceDetails_Status, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LinkedServers != nil {
		in, out := &in.LinkedServers, &out.LinkedServers
		*out = make([]RedisLinkedServer_Status, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MinimumTlsVersion != nil {
		in, out := &in.MinimumTlsVersion, &out.MinimumTlsVersion
		*out = new(RedisPropertiesStatusMinimumTlsVersion)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
	if in.PrivateEndpointConnections != nil {
		in, out := &in.PrivateEndpointConnections, &out.PrivateEndpointConnections
		*out = make([]PrivateEndpointConnection_Status, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(RedisPropertiesStatusProvisioningState)
		**out = **in
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(RedisPropertiesStatusPublicNetworkAccess)
		**out = **in
	}
	if in.RedisConfiguration != nil {
		in, out := &in.RedisConfiguration, &out.RedisConfiguration
		*out = new(RedisProperties_Status_RedisConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.RedisVersion != nil {
		in, out := &in.RedisVersion, &out.RedisVersion
		*out = new(string)
		**out = **in
	}
	if in.ReplicasPerMaster != nil {
		in, out := &in.ReplicasPerMaster, &out.ReplicasPerMaster
		*out = new(int)
		**out = **in
	}
	if in.ShardCount != nil {
		in, out := &in.ShardCount, &out.ShardCount
		*out = new(int)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.SslPort != nil {
		in, out := &in.SslPort, &out.SslPort
		*out = new(int)
		**out = **in
	}
	if in.StaticIP != nil {
		in, out := &in.StaticIP, &out.StaticIP
		*out = new(string)
		**out = **in
	}
	if in.SubnetId != nil {
		in, out := &in.SubnetId, &out.SubnetId
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
	if in.TenantSettings != nil {
		in, out := &in.TenantSettings, &out.TenantSettings
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisResourceObservation.
func (in *RedisResourceObservation) DeepCopy() *RedisResourceObservation {
	if in == nil {
		return nil
	}
	out := new(RedisResourceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisResource_Status) DeepCopyInto(out *RedisResource_Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisResource_Status.
func (in *RedisResource_Status) DeepCopy() *RedisResource_Status {
	if in == nil {
		return nil
	}
	out := new(RedisResource_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Redis_Spec) DeepCopyInto(out *Redis_Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Redis_Spec.
func (in *Redis_Spec) DeepCopy() *Redis_Spec {
	if in == nil {
		return nil
	}
	out := new(Redis_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sku) DeepCopyInto(out *Sku) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(int)
		**out = **in
	}
	if in.Family != nil {
		in, out := &in.Family, &out.Family
		*out = new(SkuFamily)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(SkuName)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sku.
func (in *Sku) DeepCopy() *Sku {
	if in == nil {
		return nil
	}
	out := new(Sku)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sku_Status) DeepCopyInto(out *Sku_Status) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(int)
		**out = **in
	}
	if in.Family != nil {
		in, out := &in.Family, &out.Family
		*out = new(SkuStatusFamily)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(SkuStatusName)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sku_Status.
func (in *Sku_Status) DeepCopy() *Sku_Status {
	if in == nil {
		return nil
	}
	out := new(Sku_Status)
	in.DeepCopyInto(out)
	return out
}
