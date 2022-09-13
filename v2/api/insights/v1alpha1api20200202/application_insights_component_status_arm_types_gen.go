// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200202

import "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

// Deprecated version of ApplicationInsightsComponent_STATUS. Use v1beta20200202.ApplicationInsightsComponent_STATUS instead
type ApplicationInsightsComponent_STATUS_ARM struct {
	Etag       *string                                            `json:"etag,omitempty"`
	Id         *string                                            `json:"id,omitempty"`
	Kind       *string                                            `json:"kind,omitempty"`
	Location   *string                                            `json:"location,omitempty"`
	Name       *string                                            `json:"name,omitempty"`
	Properties *ApplicationInsightsComponentProperties_STATUS_ARM `json:"properties,omitempty"`
	Tags       *v1.JSON                                           `json:"tags,omitempty"`
	Type       *string                                            `json:"type,omitempty"`
}

// Deprecated version of ApplicationInsightsComponentProperties_STATUS. Use v1beta20200202.ApplicationInsightsComponentProperties_STATUS instead
type ApplicationInsightsComponentProperties_STATUS_ARM struct {
	AppId                           *string                                                         `json:"AppId,omitempty"`
	ApplicationId                   *string                                                         `json:"ApplicationId,omitempty"`
	Application_Type                *ApplicationInsightsComponentProperties_Application_Type_STATUS `json:"Application_Type,omitempty"`
	ConnectionString                *string                                                         `json:"ConnectionString,omitempty"`
	CreationDate                    *string                                                         `json:"CreationDate,omitempty"`
	DisableIpMasking                *bool                                                           `json:"DisableIpMasking,omitempty"`
	DisableLocalAuth                *bool                                                           `json:"DisableLocalAuth,omitempty"`
	Flow_Type                       *ApplicationInsightsComponentProperties_Flow_Type_STATUS        `json:"Flow_Type,omitempty"`
	ForceCustomerStorageForProfiler *bool                                                           `json:"ForceCustomerStorageForProfiler,omitempty"`
	HockeyAppId                     *string                                                         `json:"HockeyAppId,omitempty"`
	HockeyAppToken                  *string                                                         `json:"HockeyAppToken,omitempty"`
	ImmediatePurgeDataOn30Days      *bool                                                           `json:"ImmediatePurgeDataOn30Days,omitempty"`
	IngestionMode                   *ApplicationInsightsComponentProperties_IngestionMode_STATUS    `json:"IngestionMode,omitempty"`
	InstrumentationKey              *string                                                         `json:"InstrumentationKey,omitempty"`
	LaMigrationDate                 *string                                                         `json:"LaMigrationDate,omitempty"`
	Name                            *string                                                         `json:"Name,omitempty"`
	PrivateLinkScopedResources      []PrivateLinkScopedResource_STATUS_ARM                          `json:"PrivateLinkScopedResources,omitempty"`
	ProvisioningState               *string                                                         `json:"provisioningState,omitempty"`
	PublicNetworkAccessForIngestion *PublicNetworkAccessType_STATUS                                 `json:"publicNetworkAccessForIngestion,omitempty"`
	PublicNetworkAccessForQuery     *PublicNetworkAccessType_STATUS                                 `json:"publicNetworkAccessForQuery,omitempty"`
	Request_Source                  *ApplicationInsightsComponentProperties_Request_Source_STATUS   `json:"Request_Source,omitempty"`
	RetentionInDays                 *int                                                            `json:"RetentionInDays,omitempty"`
	SamplingPercentage              *float64                                                        `json:"SamplingPercentage,omitempty"`
	TenantId                        *string                                                         `json:"TenantId,omitempty"`
	WorkspaceResourceId             *string                                                         `json:"WorkspaceResourceId,omitempty"`
}

// Deprecated version of ApplicationInsightsComponentProperties_Application_Type_STATUS. Use
// v1beta20200202.ApplicationInsightsComponentProperties_Application_Type_STATUS instead
type ApplicationInsightsComponentProperties_Application_Type_STATUS string

const (
	ApplicationInsightsComponentProperties_Application_Type_STATUS_Other = ApplicationInsightsComponentProperties_Application_Type_STATUS("other")
	ApplicationInsightsComponentProperties_Application_Type_STATUS_Web   = ApplicationInsightsComponentProperties_Application_Type_STATUS("web")
)

// Deprecated version of ApplicationInsightsComponentProperties_Flow_Type_STATUS. Use
// v1beta20200202.ApplicationInsightsComponentProperties_Flow_Type_STATUS instead
type ApplicationInsightsComponentProperties_Flow_Type_STATUS string

const ApplicationInsightsComponentProperties_Flow_Type_STATUS_Bluefield = ApplicationInsightsComponentProperties_Flow_Type_STATUS("Bluefield")

// Deprecated version of ApplicationInsightsComponentProperties_IngestionMode_STATUS. Use
// v1beta20200202.ApplicationInsightsComponentProperties_IngestionMode_STATUS instead
type ApplicationInsightsComponentProperties_IngestionMode_STATUS string

const (
	ApplicationInsightsComponentProperties_IngestionMode_STATUS_ApplicationInsights                       = ApplicationInsightsComponentProperties_IngestionMode_STATUS("ApplicationInsights")
	ApplicationInsightsComponentProperties_IngestionMode_STATUS_ApplicationInsightsWithDiagnosticSettings = ApplicationInsightsComponentProperties_IngestionMode_STATUS("ApplicationInsightsWithDiagnosticSettings")
	ApplicationInsightsComponentProperties_IngestionMode_STATUS_LogAnalytics                              = ApplicationInsightsComponentProperties_IngestionMode_STATUS("LogAnalytics")
)

// Deprecated version of ApplicationInsightsComponentProperties_Request_Source_STATUS. Use
// v1beta20200202.ApplicationInsightsComponentProperties_Request_Source_STATUS instead
type ApplicationInsightsComponentProperties_Request_Source_STATUS string

const ApplicationInsightsComponentProperties_Request_Source_STATUS_Rest = ApplicationInsightsComponentProperties_Request_Source_STATUS("rest")

// Deprecated version of PrivateLinkScopedResource_STATUS. Use v1beta20200202.PrivateLinkScopedResource_STATUS instead
type PrivateLinkScopedResource_STATUS_ARM struct {
	ResourceId *string `json:"ResourceId,omitempty"`
	ScopeId    *string `json:"ScopeId,omitempty"`
}

// Deprecated version of PublicNetworkAccessType_STATUS. Use v1beta20200202.PublicNetworkAccessType_STATUS instead
type PublicNetworkAccessType_STATUS string

const (
	PublicNetworkAccessType_STATUS_Disabled = PublicNetworkAccessType_STATUS("Disabled")
	PublicNetworkAccessType_STATUS_Enabled  = PublicNetworkAccessType_STATUS("Enabled")
)
