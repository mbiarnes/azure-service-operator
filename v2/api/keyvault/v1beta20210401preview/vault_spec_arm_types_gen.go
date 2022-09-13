// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210401preview

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Vault_Spec_ARM struct {
	// Location: The supported Azure location where the key vault should be created.
	Location *string `json:"location,omitempty"`

	// Name: Name of the resource
	Name string `json:"name,omitempty"`

	// Properties: Properties of the vault
	Properties *VaultProperties_ARM `json:"properties,omitempty"`

	// Tags: The tags that will be assigned to the key vault.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Vault_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01-preview"
func (vault Vault_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (vault *Vault_Spec_ARM) GetName() string {
	return vault.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.KeyVault/vaults"
func (vault *Vault_Spec_ARM) GetType() string {
	return "Microsoft.KeyVault/vaults"
}

// Generated from: https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/VaultProperties
type VaultProperties_ARM struct {
	// AccessPolicies: An array of 0 to 1024 identities that have access to the key vault. All identities in the array must use
	// the same tenant ID as the key vault's tenant ID. When `createMode` is set to `recover`, access policies are not
	// required. Otherwise, access policies are required.
	AccessPolicies []AccessPolicyEntry_ARM `json:"accessPolicies,omitempty"`

	// CreateMode: The vault's create mode to indicate whether the vault need to be recovered or not.
	CreateMode *VaultProperties_CreateMode `json:"createMode,omitempty"`

	// EnablePurgeProtection: Property specifying whether protection against purge is enabled for this vault. Setting this
	// property to true activates protection against purge for this vault and its content - only the Key Vault service may
	// initiate a hard, irrecoverable deletion. The setting is effective only if soft delete is also enabled. Enabling this
	// functionality is irreversible - that is, the property does not accept false as its value.
	EnablePurgeProtection *bool `json:"enablePurgeProtection,omitempty"`

	// EnableRbacAuthorization: Property that controls how data actions are authorized. When true, the key vault will use Role
	// Based Access Control (RBAC) for authorization of data actions, and the access policies specified in vault properties
	// will be  ignored. When false, the key vault will use the access policies specified in vault properties, and any policy
	// stored on Azure Resource Manager will be ignored. If null or not specified, the vault is created with the default value
	// of false. Note that management actions are always authorized with RBAC.
	EnableRbacAuthorization *bool `json:"enableRbacAuthorization,omitempty"`

	// EnableSoftDelete: Property to specify whether the 'soft delete' functionality is enabled for this key vault. If it's not
	// set to any value(true or false) when creating new key vault, it will be set to true by default. Once set to true, it
	// cannot be reverted to false.
	EnableSoftDelete *bool `json:"enableSoftDelete,omitempty"`

	// EnabledForDeployment: Property to specify whether Azure Virtual Machines are permitted to retrieve certificates stored
	// as secrets from the key vault.
	EnabledForDeployment *bool `json:"enabledForDeployment,omitempty"`

	// EnabledForDiskEncryption: Property to specify whether Azure Disk Encryption is permitted to retrieve secrets from the
	// vault and unwrap keys.
	EnabledForDiskEncryption *bool `json:"enabledForDiskEncryption,omitempty"`

	// EnabledForTemplateDeployment: Property to specify whether Azure Resource Manager is permitted to retrieve secrets from
	// the key vault.
	EnabledForTemplateDeployment *bool `json:"enabledForTemplateDeployment,omitempty"`

	// NetworkAcls: A set of rules governing the network accessibility of a vault.
	NetworkAcls *NetworkRuleSet_ARM `json:"networkAcls,omitempty"`

	// ProvisioningState: Provisioning state of the vault.
	ProvisioningState *VaultProperties_ProvisioningState `json:"provisioningState,omitempty"`

	// Sku: SKU details
	Sku *Sku_ARM `json:"sku,omitempty"`

	// SoftDeleteRetentionInDays: softDelete data retention days. It accepts >=7 and <=90.
	SoftDeleteRetentionInDays *int `json:"softDeleteRetentionInDays,omitempty"`

	// TenantId: The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantId *string `json:"tenantId,omitempty"`

	// VaultUri: The URI of the vault for performing operations on keys and secrets.
	VaultUri *string `json:"vaultUri,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/AccessPolicyEntry
type AccessPolicyEntry_ARM struct {
	// ApplicationId:  Application ID of the client making request on behalf of a principal
	ApplicationId *string `json:"applicationId,omitempty"`

	// ObjectId: The object ID of a user, service principal or security group in the Azure Active Directory tenant for the
	// vault. The object ID must be unique for the list of access policies.
	ObjectId *string `json:"objectId,omitempty"`

	// Permissions: Permissions the identity has for keys, secrets, certificates and storage.
	Permissions *Permissions_ARM `json:"permissions,omitempty"`

	// TenantId: The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantId *string `json:"tenantId,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/NetworkRuleSet
type NetworkRuleSet_ARM struct {
	// Bypass: Tells what traffic can bypass network rules. This can be 'AzureServices' or 'None'.  If not specified the
	// default is 'AzureServices'.
	Bypass *NetworkRuleSet_Bypass `json:"bypass,omitempty"`

	// DefaultAction: The default action when no rule from ipRules and from virtualNetworkRules match. This is only used after
	// the bypass property has been evaluated.
	DefaultAction *NetworkRuleSet_DefaultAction `json:"defaultAction,omitempty"`

	// IpRules: The list of IP address rules.
	IpRules []IPRule_ARM `json:"ipRules,omitempty"`

	// VirtualNetworkRules: The list of virtual network rules.
	VirtualNetworkRules []VirtualNetworkRule_ARM `json:"virtualNetworkRules,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/Sku
type Sku_ARM struct {
	// Family: SKU family name
	Family *Sku_Family `json:"family,omitempty"`

	// Name: SKU name to specify whether the key vault is a standard vault or a premium vault.
	Name *Sku_Name `json:"name,omitempty"`
}

// +kubebuilder:validation:Enum={"default","recover"}
type VaultProperties_CreateMode string

const (
	VaultProperties_CreateMode_Default = VaultProperties_CreateMode("default")
	VaultProperties_CreateMode_Recover = VaultProperties_CreateMode("recover")
)

// +kubebuilder:validation:Enum={"RegisteringDns","Succeeded"}
type VaultProperties_ProvisioningState string

const (
	VaultProperties_ProvisioningState_RegisteringDns = VaultProperties_ProvisioningState("RegisteringDns")
	VaultProperties_ProvisioningState_Succeeded      = VaultProperties_ProvisioningState("Succeeded")
)

// Generated from: https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/IPRule
type IPRule_ARM struct {
	// Value: An IPv4 address range in CIDR notation, such as '124.56.78.91' (simple IP address) or '124.56.78.0/24' (all
	// addresses that start with 124.56.78).
	Value *string `json:"value,omitempty"`
}

// +kubebuilder:validation:Enum={"AzureServices","None"}
type NetworkRuleSet_Bypass string

const (
	NetworkRuleSet_Bypass_AzureServices = NetworkRuleSet_Bypass("AzureServices")
	NetworkRuleSet_Bypass_None          = NetworkRuleSet_Bypass("None")
)

// +kubebuilder:validation:Enum={"Allow","Deny"}
type NetworkRuleSet_DefaultAction string

const (
	NetworkRuleSet_DefaultAction_Allow = NetworkRuleSet_DefaultAction("Allow")
	NetworkRuleSet_DefaultAction_Deny  = NetworkRuleSet_DefaultAction("Deny")
)

// Generated from: https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/Permissions
type Permissions_ARM struct {
	// Certificates: Permissions to certificates
	Certificates []Permissions_Certificates `json:"certificates,omitempty"`

	// Keys: Permissions to keys
	Keys []Permissions_Keys `json:"keys,omitempty"`

	// Secrets: Permissions to secrets
	Secrets []Permissions_Secrets `json:"secrets,omitempty"`

	// Storage: Permissions to storage accounts
	Storage []Permissions_Storage `json:"storage,omitempty"`
}

// +kubebuilder:validation:Enum={"A"}
type Sku_Family string

const Sku_Family_A = Sku_Family("A")

// +kubebuilder:validation:Enum={"premium","standard"}
type Sku_Name string

const (
	Sku_Name_Premium  = Sku_Name("premium")
	Sku_Name_Standard = Sku_Name("standard")
)

// Generated from: https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/VirtualNetworkRule
type VirtualNetworkRule_ARM struct {
	Id *string `json:"id,omitempty"`

	// IgnoreMissingVnetServiceEndpoint: Property to specify whether NRP will ignore the check if parent subnet has
	// serviceEndpoints configured.
	IgnoreMissingVnetServiceEndpoint *bool `json:"ignoreMissingVnetServiceEndpoint,omitempty"`
}

// +kubebuilder:validation:Enum={"backup","create","delete","deleteissuers","get","getissuers","import","list","listissuers","managecontacts","manageissuers","purge","recover","restore","setissuers","update"}
type Permissions_Certificates string

const (
	Permissions_Certificates_Backup         = Permissions_Certificates("backup")
	Permissions_Certificates_Create         = Permissions_Certificates("create")
	Permissions_Certificates_Delete         = Permissions_Certificates("delete")
	Permissions_Certificates_Deleteissuers  = Permissions_Certificates("deleteissuers")
	Permissions_Certificates_Get            = Permissions_Certificates("get")
	Permissions_Certificates_Getissuers     = Permissions_Certificates("getissuers")
	Permissions_Certificates_Import         = Permissions_Certificates("import")
	Permissions_Certificates_List           = Permissions_Certificates("list")
	Permissions_Certificates_Listissuers    = Permissions_Certificates("listissuers")
	Permissions_Certificates_Managecontacts = Permissions_Certificates("managecontacts")
	Permissions_Certificates_Manageissuers  = Permissions_Certificates("manageissuers")
	Permissions_Certificates_Purge          = Permissions_Certificates("purge")
	Permissions_Certificates_Recover        = Permissions_Certificates("recover")
	Permissions_Certificates_Restore        = Permissions_Certificates("restore")
	Permissions_Certificates_Setissuers     = Permissions_Certificates("setissuers")
	Permissions_Certificates_Update         = Permissions_Certificates("update")
)

// +kubebuilder:validation:Enum={"backup","create","decrypt","delete","encrypt","get","import","list","purge","recover","release","restore","sign","unwrapKey","update","verify","wrapKey"}
type Permissions_Keys string

const (
	Permissions_Keys_Backup    = Permissions_Keys("backup")
	Permissions_Keys_Create    = Permissions_Keys("create")
	Permissions_Keys_Decrypt   = Permissions_Keys("decrypt")
	Permissions_Keys_Delete    = Permissions_Keys("delete")
	Permissions_Keys_Encrypt   = Permissions_Keys("encrypt")
	Permissions_Keys_Get       = Permissions_Keys("get")
	Permissions_Keys_Import    = Permissions_Keys("import")
	Permissions_Keys_List      = Permissions_Keys("list")
	Permissions_Keys_Purge     = Permissions_Keys("purge")
	Permissions_Keys_Recover   = Permissions_Keys("recover")
	Permissions_Keys_Release   = Permissions_Keys("release")
	Permissions_Keys_Restore   = Permissions_Keys("restore")
	Permissions_Keys_Sign      = Permissions_Keys("sign")
	Permissions_Keys_UnwrapKey = Permissions_Keys("unwrapKey")
	Permissions_Keys_Update    = Permissions_Keys("update")
	Permissions_Keys_Verify    = Permissions_Keys("verify")
	Permissions_Keys_WrapKey   = Permissions_Keys("wrapKey")
)

// +kubebuilder:validation:Enum={"backup","delete","get","list","purge","recover","restore","set"}
type Permissions_Secrets string

const (
	Permissions_Secrets_Backup  = Permissions_Secrets("backup")
	Permissions_Secrets_Delete  = Permissions_Secrets("delete")
	Permissions_Secrets_Get     = Permissions_Secrets("get")
	Permissions_Secrets_List    = Permissions_Secrets("list")
	Permissions_Secrets_Purge   = Permissions_Secrets("purge")
	Permissions_Secrets_Recover = Permissions_Secrets("recover")
	Permissions_Secrets_Restore = Permissions_Secrets("restore")
	Permissions_Secrets_Set     = Permissions_Secrets("set")
)

// +kubebuilder:validation:Enum={"backup","delete","deletesas","get","getsas","list","listsas","purge","recover","regeneratekey","restore","set","setsas","update"}
type Permissions_Storage string

const (
	Permissions_Storage_Backup        = Permissions_Storage("backup")
	Permissions_Storage_Delete        = Permissions_Storage("delete")
	Permissions_Storage_Deletesas     = Permissions_Storage("deletesas")
	Permissions_Storage_Get           = Permissions_Storage("get")
	Permissions_Storage_Getsas        = Permissions_Storage("getsas")
	Permissions_Storage_List          = Permissions_Storage("list")
	Permissions_Storage_Listsas       = Permissions_Storage("listsas")
	Permissions_Storage_Purge         = Permissions_Storage("purge")
	Permissions_Storage_Recover       = Permissions_Storage("recover")
	Permissions_Storage_Regeneratekey = Permissions_Storage("regeneratekey")
	Permissions_Storage_Restore       = Permissions_Storage("restore")
	Permissions_Storage_Set           = Permissions_Storage("set")
	Permissions_Storage_Setsas        = Permissions_Storage("setsas")
	Permissions_Storage_Update        = Permissions_Storage("update")
)
