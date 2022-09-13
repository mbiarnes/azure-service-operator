// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210501

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// Deprecated version of ManagedCluster_Spec. Use v1beta20210501.ManagedCluster_Spec instead
type ManagedCluster_Spec_ARM struct {
	ExtendedLocation *ExtendedLocation_ARM         `json:"extendedLocation,omitempty"`
	Identity         *ManagedClusterIdentity_ARM   `json:"identity,omitempty"`
	Location         *string                       `json:"location,omitempty"`
	Name             string                        `json:"name,omitempty"`
	Properties       *ManagedClusterProperties_ARM `json:"properties,omitempty"`
	Sku              *ManagedClusterSKU_ARM        `json:"sku,omitempty"`
	Tags             map[string]string             `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &ManagedCluster_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-01"
func (cluster ManagedCluster_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (cluster *ManagedCluster_Spec_ARM) GetName() string {
	return cluster.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ContainerService/managedClusters"
func (cluster *ManagedCluster_Spec_ARM) GetType() string {
	return "Microsoft.ContainerService/managedClusters"
}

// Deprecated version of ExtendedLocation. Use v1beta20210501.ExtendedLocation instead
type ExtendedLocation_ARM struct {
	Name *string                `json:"name,omitempty"`
	Type *ExtendedLocation_Type `json:"type,omitempty"`
}

// Deprecated version of ManagedClusterIdentity. Use v1beta20210501.ManagedClusterIdentity instead
type ManagedClusterIdentity_ARM struct {
	Type                   *ManagedClusterIdentity_Type `json:"type,omitempty"`
	UserAssignedIdentities map[string]v1.JSON           `json:"userAssignedIdentities,omitempty"`
}

// Deprecated version of ManagedClusterProperties. Use v1beta20210501.ManagedClusterProperties instead
type ManagedClusterProperties_ARM struct {
	AadProfile              *ManagedClusterAADProfile_ARM                                                                               `json:"aadProfile,omitempty"`
	AddonProfiles           map[string]ManagedClusterAddonProfile_ARM                                                                   `json:"addonProfiles,omitempty"`
	AgentPoolProfiles       []ManagedClusterAgentPoolProfile_ARM                                                                        `json:"agentPoolProfiles,omitempty"`
	ApiServerAccessProfile  *ManagedClusterAPIServerAccessProfile_ARM                                                                   `json:"apiServerAccessProfile,omitempty"`
	AutoScalerProfile       *ManagedClusterPropertiesAutoScalerProfile_ARM                                                              `json:"autoScalerProfile,omitempty"`
	AutoUpgradeProfile      *ManagedClusterAutoUpgradeProfile_ARM                                                                       `json:"autoUpgradeProfile,omitempty"`
	DisableLocalAccounts    *bool                                                                                                       `json:"disableLocalAccounts,omitempty"`
	DiskEncryptionSetID     *string                                                                                                     `json:"diskEncryptionSetID,omitempty"`
	DnsPrefix               *string                                                                                                     `json:"dnsPrefix,omitempty"`
	EnablePodSecurityPolicy *bool                                                                                                       `json:"enablePodSecurityPolicy,omitempty"`
	EnableRBAC              *bool                                                                                                       `json:"enableRBAC,omitempty"`
	FqdnSubdomain           *string                                                                                                     `json:"fqdnSubdomain,omitempty"`
	HttpProxyConfig         *ManagedClusterHTTPProxyConfig_ARM                                                                          `json:"httpProxyConfig,omitempty"`
	IdentityProfile         map[string]Componentsqit0Etschemasmanagedclusterpropertiespropertiesidentityprofileadditionalproperties_ARM `json:"identityProfile,omitempty"`
	KubernetesVersion       *string                                                                                                     `json:"kubernetesVersion,omitempty"`
	LinuxProfile            *ContainerServiceLinuxProfile_ARM                                                                           `json:"linuxProfile,omitempty"`
	NetworkProfile          *ContainerServiceNetworkProfile_ARM                                                                         `json:"networkProfile,omitempty"`
	NodeResourceGroup       *string                                                                                                     `json:"nodeResourceGroup,omitempty"`
	PodIdentityProfile      *ManagedClusterPodIdentityProfile_ARM                                                                       `json:"podIdentityProfile,omitempty"`
	PrivateLinkResources    []PrivateLinkResource_ARM                                                                                   `json:"privateLinkResources,omitempty"`
	ServicePrincipalProfile *ManagedClusterServicePrincipalProfile_ARM                                                                  `json:"servicePrincipalProfile,omitempty"`
	WindowsProfile          *ManagedClusterWindowsProfile_ARM                                                                           `json:"windowsProfile,omitempty"`
}

// Deprecated version of ManagedClusterSKU. Use v1beta20210501.ManagedClusterSKU instead
type ManagedClusterSKU_ARM struct {
	Name *ManagedClusterSKU_Name `json:"name,omitempty"`
	Tier *ManagedClusterSKU_Tier `json:"tier,omitempty"`
}

// Deprecated version of Componentsqit0Etschemasmanagedclusterpropertiespropertiesidentityprofileadditionalproperties. Use v1beta20210501.Componentsqit0Etschemasmanagedclusterpropertiespropertiesidentityprofileadditionalproperties instead
type Componentsqit0Etschemasmanagedclusterpropertiespropertiesidentityprofileadditionalproperties_ARM struct {
	ClientId   *string `json:"clientId,omitempty"`
	ObjectId   *string `json:"objectId,omitempty"`
	ResourceId *string `json:"resourceId,omitempty"`
}

// Deprecated version of ContainerServiceLinuxProfile. Use v1beta20210501.ContainerServiceLinuxProfile instead
type ContainerServiceLinuxProfile_ARM struct {
	AdminUsername *string                               `json:"adminUsername,omitempty"`
	Ssh           *ContainerServiceSshConfiguration_ARM `json:"ssh,omitempty"`
}

// Deprecated version of ContainerServiceNetworkProfile. Use v1beta20210501.ContainerServiceNetworkProfile instead
type ContainerServiceNetworkProfile_ARM struct {
	DnsServiceIP        *string                                         `json:"dnsServiceIP,omitempty"`
	DockerBridgeCidr    *string                                         `json:"dockerBridgeCidr,omitempty"`
	LoadBalancerProfile *ManagedClusterLoadBalancerProfile_ARM          `json:"loadBalancerProfile,omitempty"`
	LoadBalancerSku     *ContainerServiceNetworkProfile_LoadBalancerSku `json:"loadBalancerSku,omitempty"`
	NetworkMode         *ContainerServiceNetworkProfile_NetworkMode     `json:"networkMode,omitempty"`
	NetworkPlugin       *ContainerServiceNetworkProfile_NetworkPlugin   `json:"networkPlugin,omitempty"`
	NetworkPolicy       *ContainerServiceNetworkProfile_NetworkPolicy   `json:"networkPolicy,omitempty"`
	OutboundType        *ContainerServiceNetworkProfile_OutboundType    `json:"outboundType,omitempty"`
	PodCidr             *string                                         `json:"podCidr,omitempty"`
	ServiceCidr         *string                                         `json:"serviceCidr,omitempty"`
}

// Deprecated version of ExtendedLocation_Type. Use v1beta20210501.ExtendedLocation_Type instead
// +kubebuilder:validation:Enum={"EdgeZone"}
type ExtendedLocation_Type string

const ExtendedLocation_Type_EdgeZone = ExtendedLocation_Type("EdgeZone")

// Deprecated version of ManagedClusterAADProfile. Use v1beta20210501.ManagedClusterAADProfile instead
type ManagedClusterAADProfile_ARM struct {
	AdminGroupObjectIDs []string `json:"adminGroupObjectIDs,omitempty"`
	ClientAppID         *string  `json:"clientAppID,omitempty"`
	EnableAzureRBAC     *bool    `json:"enableAzureRBAC,omitempty"`
	Managed             *bool    `json:"managed,omitempty"`
	ServerAppID         *string  `json:"serverAppID,omitempty"`
	ServerAppSecret     *string  `json:"serverAppSecret,omitempty"`
	TenantID            *string  `json:"tenantID,omitempty"`
}

// Deprecated version of ManagedClusterAddonProfile. Use v1beta20210501.ManagedClusterAddonProfile instead
type ManagedClusterAddonProfile_ARM struct {
	Config  map[string]string `json:"config,omitempty"`
	Enabled *bool             `json:"enabled,omitempty"`
}

// Deprecated version of ManagedClusterAgentPoolProfile. Use v1beta20210501.ManagedClusterAgentPoolProfile instead
type ManagedClusterAgentPoolProfile_ARM struct {
	AvailabilityZones         []string                                               `json:"availabilityZones,omitempty"`
	Count                     *int                                                   `json:"count,omitempty"`
	EnableAutoScaling         *bool                                                  `json:"enableAutoScaling,omitempty"`
	EnableEncryptionAtHost    *bool                                                  `json:"enableEncryptionAtHost,omitempty"`
	EnableFIPS                *bool                                                  `json:"enableFIPS,omitempty"`
	EnableNodePublicIP        *bool                                                  `json:"enableNodePublicIP,omitempty"`
	EnableUltraSSD            *bool                                                  `json:"enableUltraSSD,omitempty"`
	GpuInstanceProfile        *ManagedClusterAgentPoolProfile_GpuInstanceProfile     `json:"gpuInstanceProfile,omitempty"`
	KubeletConfig             *KubeletConfig_ARM                                     `json:"kubeletConfig,omitempty"`
	KubeletDiskType           *ManagedClusterAgentPoolProfile_KubeletDiskType        `json:"kubeletDiskType,omitempty"`
	LinuxOSConfig             *LinuxOSConfig_ARM                                     `json:"linuxOSConfig,omitempty"`
	MaxCount                  *int                                                   `json:"maxCount,omitempty"`
	MaxPods                   *int                                                   `json:"maxPods,omitempty"`
	MinCount                  *int                                                   `json:"minCount,omitempty"`
	Mode                      *ManagedClusterAgentPoolProfile_Mode                   `json:"mode,omitempty"`
	Name                      *string                                                `json:"name,omitempty"`
	NodeLabels                map[string]string                                      `json:"nodeLabels,omitempty"`
	NodePublicIPPrefixID      *string                                                `json:"nodePublicIPPrefixID,omitempty"`
	NodeTaints                []string                                               `json:"nodeTaints,omitempty"`
	OrchestratorVersion       *string                                                `json:"orchestratorVersion,omitempty"`
	OsDiskSizeGB              *int                                                   `json:"osDiskSizeGB,omitempty"`
	OsDiskType                *ManagedClusterAgentPoolProfile_OsDiskType             `json:"osDiskType,omitempty"`
	OsSKU                     *ManagedClusterAgentPoolProfile_OsSKU                  `json:"osSKU,omitempty"`
	OsType                    *ManagedClusterAgentPoolProfile_OsType                 `json:"osType,omitempty"`
	PodSubnetID               *string                                                `json:"podSubnetID,omitempty"`
	ProximityPlacementGroupID *string                                                `json:"proximityPlacementGroupID,omitempty"`
	ScaleSetEvictionPolicy    *ManagedClusterAgentPoolProfile_ScaleSetEvictionPolicy `json:"scaleSetEvictionPolicy,omitempty"`
	ScaleSetPriority          *ManagedClusterAgentPoolProfile_ScaleSetPriority       `json:"scaleSetPriority,omitempty"`
	SpotMaxPrice              *float64                                               `json:"spotMaxPrice,omitempty"`
	Tags                      map[string]string                                      `json:"tags,omitempty"`
	Type                      *ManagedClusterAgentPoolProfile_Type                   `json:"type,omitempty"`
	UpgradeSettings           *AgentPoolUpgradeSettings_ARM                          `json:"upgradeSettings,omitempty"`
	VmSize                    *string                                                `json:"vmSize,omitempty"`
	VnetSubnetID              *string                                                `json:"vnetSubnetID,omitempty"`
}

// Deprecated version of ManagedClusterAPIServerAccessProfile. Use v1beta20210501.ManagedClusterAPIServerAccessProfile instead
type ManagedClusterAPIServerAccessProfile_ARM struct {
	AuthorizedIPRanges             []string `json:"authorizedIPRanges,omitempty"`
	EnablePrivateCluster           *bool    `json:"enablePrivateCluster,omitempty"`
	EnablePrivateClusterPublicFQDN *bool    `json:"enablePrivateClusterPublicFQDN,omitempty"`
	PrivateDNSZone                 *string  `json:"privateDNSZone,omitempty"`
}

// Deprecated version of ManagedClusterAutoUpgradeProfile. Use v1beta20210501.ManagedClusterAutoUpgradeProfile instead
type ManagedClusterAutoUpgradeProfile_ARM struct {
	UpgradeChannel *ManagedClusterAutoUpgradeProfile_UpgradeChannel `json:"upgradeChannel,omitempty"`
}

// Deprecated version of ManagedClusterHTTPProxyConfig. Use v1beta20210501.ManagedClusterHTTPProxyConfig instead
type ManagedClusterHTTPProxyConfig_ARM struct {
	HttpProxy  *string  `json:"httpProxy,omitempty"`
	HttpsProxy *string  `json:"httpsProxy,omitempty"`
	NoProxy    []string `json:"noProxy,omitempty"`
	TrustedCa  *string  `json:"trustedCa,omitempty"`
}

// Deprecated version of ManagedClusterIdentity_Type. Use v1beta20210501.ManagedClusterIdentity_Type instead
// +kubebuilder:validation:Enum={"None","SystemAssigned","UserAssigned"}
type ManagedClusterIdentity_Type string

const (
	ManagedClusterIdentity_Type_None           = ManagedClusterIdentity_Type("None")
	ManagedClusterIdentity_Type_SystemAssigned = ManagedClusterIdentity_Type("SystemAssigned")
	ManagedClusterIdentity_Type_UserAssigned   = ManagedClusterIdentity_Type("UserAssigned")
)

// Deprecated version of ManagedClusterPodIdentityProfile. Use v1beta20210501.ManagedClusterPodIdentityProfile instead
type ManagedClusterPodIdentityProfile_ARM struct {
	AllowNetworkPluginKubenet      *bool                                    `json:"allowNetworkPluginKubenet,omitempty"`
	Enabled                        *bool                                    `json:"enabled,omitempty"`
	UserAssignedIdentities         []ManagedClusterPodIdentity_ARM          `json:"userAssignedIdentities,omitempty"`
	UserAssignedIdentityExceptions []ManagedClusterPodIdentityException_ARM `json:"userAssignedIdentityExceptions,omitempty"`
}

// Deprecated version of ManagedClusterPropertiesAutoScalerProfile. Use v1beta20210501.ManagedClusterPropertiesAutoScalerProfile instead
type ManagedClusterPropertiesAutoScalerProfile_ARM struct {
	BalanceSimilarNodeGroups      *string                                             `json:"balance-similar-node-groups,omitempty"`
	Expander                      *ManagedClusterPropertiesAutoScalerProfile_Expander `json:"expander,omitempty"`
	MaxEmptyBulkDelete            *string                                             `json:"max-empty-bulk-delete,omitempty"`
	MaxGracefulTerminationSec     *string                                             `json:"max-graceful-termination-sec,omitempty"`
	MaxNodeProvisionTime          *string                                             `json:"max-node-provision-time,omitempty"`
	MaxTotalUnreadyPercentage     *string                                             `json:"max-total-unready-percentage,omitempty"`
	NewPodScaleUpDelay            *string                                             `json:"new-pod-scale-up-delay,omitempty"`
	OkTotalUnreadyCount           *string                                             `json:"ok-total-unready-count,omitempty"`
	ScaleDownDelayAfterAdd        *string                                             `json:"scale-down-delay-after-add,omitempty"`
	ScaleDownDelayAfterDelete     *string                                             `json:"scale-down-delay-after-delete,omitempty"`
	ScaleDownDelayAfterFailure    *string                                             `json:"scale-down-delay-after-failure,omitempty"`
	ScaleDownUnneededTime         *string                                             `json:"scale-down-unneeded-time,omitempty"`
	ScaleDownUnreadyTime          *string                                             `json:"scale-down-unready-time,omitempty"`
	ScaleDownUtilizationThreshold *string                                             `json:"scale-down-utilization-threshold,omitempty"`
	ScanInterval                  *string                                             `json:"scan-interval,omitempty"`
	SkipNodesWithLocalStorage     *string                                             `json:"skip-nodes-with-local-storage,omitempty"`
	SkipNodesWithSystemPods       *string                                             `json:"skip-nodes-with-system-pods,omitempty"`
}

// Deprecated version of ManagedClusterServicePrincipalProfile. Use v1beta20210501.ManagedClusterServicePrincipalProfile instead
type ManagedClusterServicePrincipalProfile_ARM struct {
	ClientId *string `json:"clientId,omitempty"`
	Secret   *string `json:"secret,omitempty"`
}

// Deprecated version of ManagedClusterSKU_Name. Use v1beta20210501.ManagedClusterSKU_Name instead
// +kubebuilder:validation:Enum={"Basic"}
type ManagedClusterSKU_Name string

const ManagedClusterSKU_Name_Basic = ManagedClusterSKU_Name("Basic")

// Deprecated version of ManagedClusterSKU_Tier. Use v1beta20210501.ManagedClusterSKU_Tier instead
// +kubebuilder:validation:Enum={"Free","Paid"}
type ManagedClusterSKU_Tier string

const (
	ManagedClusterSKU_Tier_Free = ManagedClusterSKU_Tier("Free")
	ManagedClusterSKU_Tier_Paid = ManagedClusterSKU_Tier("Paid")
)

// Deprecated version of ManagedClusterWindowsProfile. Use v1beta20210501.ManagedClusterWindowsProfile instead
type ManagedClusterWindowsProfile_ARM struct {
	AdminPassword  *string                                   `json:"adminPassword,omitempty"`
	AdminUsername  *string                                   `json:"adminUsername,omitempty"`
	EnableCSIProxy *bool                                     `json:"enableCSIProxy,omitempty"`
	LicenseType    *ManagedClusterWindowsProfile_LicenseType `json:"licenseType,omitempty"`
}

// Deprecated version of PrivateLinkResource. Use v1beta20210501.PrivateLinkResource instead
type PrivateLinkResource_ARM struct {
	GroupId         *string  `json:"groupId,omitempty"`
	Id              *string  `json:"id,omitempty"`
	Name            *string  `json:"name,omitempty"`
	RequiredMembers []string `json:"requiredMembers,omitempty"`
	Type            *string  `json:"type,omitempty"`
}

// Deprecated version of ContainerServiceSshConfiguration. Use v1beta20210501.ContainerServiceSshConfiguration instead
type ContainerServiceSshConfiguration_ARM struct {
	PublicKeys []ContainerServiceSshPublicKey_ARM `json:"publicKeys,omitempty"`
}

// Deprecated version of ManagedClusterLoadBalancerProfile. Use v1beta20210501.ManagedClusterLoadBalancerProfile instead
type ManagedClusterLoadBalancerProfile_ARM struct {
	AllocatedOutboundPorts *int                                                     `json:"allocatedOutboundPorts,omitempty"`
	EffectiveOutboundIPs   []ResourceReference_ARM                                  `json:"effectiveOutboundIPs,omitempty"`
	IdleTimeoutInMinutes   *int                                                     `json:"idleTimeoutInMinutes,omitempty"`
	ManagedOutboundIPs     *ManagedClusterLoadBalancerProfileManagedOutboundIPs_ARM `json:"managedOutboundIPs,omitempty"`
	OutboundIPPrefixes     *ManagedClusterLoadBalancerProfileOutboundIPPrefixes_ARM `json:"outboundIPPrefixes,omitempty"`
	OutboundIPs            *ManagedClusterLoadBalancerProfileOutboundIPs_ARM        `json:"outboundIPs,omitempty"`
}

// Deprecated version of ManagedClusterPodIdentity. Use v1beta20210501.ManagedClusterPodIdentity instead
type ManagedClusterPodIdentity_ARM struct {
	BindingSelector *string                   `json:"bindingSelector,omitempty"`
	Identity        *UserAssignedIdentity_ARM `json:"identity,omitempty"`
	Name            *string                   `json:"name,omitempty"`
	Namespace       *string                   `json:"namespace,omitempty"`
}

// Deprecated version of ManagedClusterPodIdentityException. Use v1beta20210501.ManagedClusterPodIdentityException instead
type ManagedClusterPodIdentityException_ARM struct {
	Name      *string           `json:"name,omitempty"`
	Namespace *string           `json:"namespace,omitempty"`
	PodLabels map[string]string `json:"podLabels,omitempty"`
}

// Deprecated version of ContainerServiceSshPublicKey. Use v1beta20210501.ContainerServiceSshPublicKey instead
type ContainerServiceSshPublicKey_ARM struct {
	KeyData *string `json:"keyData,omitempty"`
}

// Deprecated version of ManagedClusterLoadBalancerProfileManagedOutboundIPs. Use v1beta20210501.ManagedClusterLoadBalancerProfileManagedOutboundIPs instead
type ManagedClusterLoadBalancerProfileManagedOutboundIPs_ARM struct {
	Count *int `json:"count,omitempty"`
}

// Deprecated version of ManagedClusterLoadBalancerProfileOutboundIPPrefixes. Use v1beta20210501.ManagedClusterLoadBalancerProfileOutboundIPPrefixes instead
type ManagedClusterLoadBalancerProfileOutboundIPPrefixes_ARM struct {
	PublicIPPrefixes []ResourceReference_ARM `json:"publicIPPrefixes,omitempty"`
}

// Deprecated version of ManagedClusterLoadBalancerProfileOutboundIPs. Use v1beta20210501.ManagedClusterLoadBalancerProfileOutboundIPs instead
type ManagedClusterLoadBalancerProfileOutboundIPs_ARM struct {
	PublicIPs []ResourceReference_ARM `json:"publicIPs,omitempty"`
}

// Deprecated version of ResourceReference. Use v1beta20210501.ResourceReference instead
type ResourceReference_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of UserAssignedIdentity. Use v1beta20210501.UserAssignedIdentity instead
type UserAssignedIdentity_ARM struct {
	ClientId   *string `json:"clientId,omitempty"`
	ObjectId   *string `json:"objectId,omitempty"`
	ResourceId *string `json:"resourceId,omitempty"`
}
