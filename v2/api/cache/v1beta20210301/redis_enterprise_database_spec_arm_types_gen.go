// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210301

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type RedisEnterprise_Database_Spec_ARM struct {
	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// Name: The name of the database.
	Name string `json:"name,omitempty"`

	// Properties: Properties of RedisEnterprise databases, as opposed to general resource properties like location, tags
	Properties *DatabaseProperties_ARM `json:"properties,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &RedisEnterprise_Database_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-03-01"
func (database RedisEnterprise_Database_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (database *RedisEnterprise_Database_Spec_ARM) GetName() string {
	return database.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redisEnterprise/databases"
func (database *RedisEnterprise_Database_Spec_ARM) GetType() string {
	return "Microsoft.Cache/redisEnterprise/databases"
}

// Generated from: https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/definitions/DatabaseProperties
type DatabaseProperties_ARM struct {
	// ClientProtocol: Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is
	// TLS-encrypted.
	ClientProtocol *DatabaseProperties_ClientProtocol `json:"clientProtocol,omitempty"`

	// ClusteringPolicy: Clustering policy - default is OSSCluster. Specified at create time.
	ClusteringPolicy *DatabaseProperties_ClusteringPolicy `json:"clusteringPolicy,omitempty"`

	// EvictionPolicy: Redis eviction policy - default is VolatileLRU.
	EvictionPolicy *DatabaseProperties_EvictionPolicy `json:"evictionPolicy,omitempty"`

	// Modules: Optional set of redis modules to enable in this database - modules can only be added at creation time.
	Modules []Module_ARM `json:"modules,omitempty"`

	// Persistence: Persistence-related configuration for the RedisEnterprise database
	Persistence *Persistence_ARM `json:"persistence,omitempty"`

	// Port: TCP port of the database endpoint. Specified at create time. Defaults to an available port.
	Port *int `json:"port,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/definitions/Module
type Module_ARM struct {
	// Args: Configuration options for the module, e.g. 'ERROR_RATE 0.00 INITIAL_SIZE 400'.
	Args *string `json:"args,omitempty"`

	// Name: The name of the module, e.g. 'RedisBloom', 'RediSearch', 'RedisTimeSeries'
	Name *string `json:"name,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/definitions/Persistence
type Persistence_ARM struct {
	// AofEnabled: Sets whether AOF is enabled.
	AofEnabled *bool `json:"aofEnabled,omitempty"`

	// AofFrequency: Sets the frequency at which data is written to disk.
	AofFrequency *Persistence_AofFrequency `json:"aofFrequency,omitempty"`

	// RdbEnabled: Sets whether RDB is enabled.
	RdbEnabled *bool `json:"rdbEnabled,omitempty"`

	// RdbFrequency: Sets the frequency at which a snapshot of the database is created.
	RdbFrequency *Persistence_RdbFrequency `json:"rdbFrequency,omitempty"`
}
