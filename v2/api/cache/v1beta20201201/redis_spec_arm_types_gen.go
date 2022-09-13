// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201201

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Redis_Spec_ARM struct {
	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the Redis cache.
	Name string `json:"name,omitempty"`

	// Properties: Properties supplied to Create Redis operation.
	Properties *RedisCreateProperties_ARM `json:"properties,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`

	// Zones: A list of availability zones denoting where the resource needs to come from.
	Zones []string `json:"zones,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Redis_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (redis Redis_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (redis *Redis_Spec_ARM) GetName() string {
	return redis.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis"
func (redis *Redis_Spec_ARM) GetType() string {
	return "Microsoft.Cache/redis"
}

// Generated from: https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/definitions/RedisCreateProperties
type RedisCreateProperties_ARM struct {
	// EnableNonSslPort: Specifies whether the non-ssl Redis server port (6379) is enabled.
	EnableNonSslPort *bool `json:"enableNonSslPort,omitempty"`

	// MinimumTlsVersion: Optional: requires clients to use a specified TLS version (or higher) to connect (e,g, '1.0', '1.1',
	// '1.2').
	MinimumTlsVersion *RedisCreateProperties_MinimumTlsVersion `json:"minimumTlsVersion,omitempty"`

	// PublicNetworkAccess: Whether or not public endpoint access is allowed for this cache.  Value is optional but if passed
	// in, must be 'Enabled' or 'Disabled'. If 'Disabled', private endpoints are the exclusive access method. Default value is
	// 'Enabled'.
	PublicNetworkAccess *RedisCreateProperties_PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// RedisConfiguration: All Redis Settings. Few possible keys:
	// rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value
	// etc.
	RedisConfiguration map[string]string `json:"redisConfiguration,omitempty"`

	// RedisVersion: Redis version. This should be in the form 'major[.minor[.build]]' (only 'major' is required) or the value
	// 'latest' which refers to the latest stable Redis version that is available. Only the major and minor version are used in
	// a PUT/PATCH request. Supported versions: 4.0, 6.0.
	RedisVersion *string `json:"redisVersion,omitempty"`

	// ReplicasPerMaster: The number of replicas to be created per primary.
	ReplicasPerMaster *int `json:"replicasPerMaster,omitempty"`

	// ReplicasPerPrimary: The number of replicas to be created per primary.
	ReplicasPerPrimary *int `json:"replicasPerPrimary,omitempty"`

	// ShardCount: The number of shards to be created on a Premium Cluster Cache.
	ShardCount *int `json:"shardCount,omitempty"`

	// Sku: SKU parameters supplied to the create Redis operation.
	Sku *Sku_ARM `json:"sku,omitempty"`

	// StaticIP: Static IP address. Optionally, may be specified when deploying a Redis cache inside an existing Azure Virtual
	// Network; auto assigned by default.
	StaticIP *string `json:"staticIP,omitempty"`
	SubnetId *string `json:"subnetId,omitempty"`

	// TenantSettings: A dictionary of tenant settings
	TenantSettings map[string]string `json:"tenantSettings,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/definitions/Sku
type Sku_ARM struct {
	// Capacity: The size of the Redis cache to deploy. Valid values: for C (Basic/Standard) family (0, 1, 2, 3, 4, 5, 6), for
	// P (Premium) family (1, 2, 3, 4).
	Capacity *int `json:"capacity,omitempty"`

	// Family: The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).
	Family *Sku_Family `json:"family,omitempty"`

	// Name: The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium).
	Name *Sku_Name `json:"name,omitempty"`
}
