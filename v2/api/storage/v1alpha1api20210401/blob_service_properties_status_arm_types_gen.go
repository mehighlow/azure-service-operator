// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

// Deprecated version of BlobServiceProperties_STATUS. Use v1beta20210401.BlobServiceProperties_STATUS instead
type BlobServiceProperties_STATUS_ARM struct {
	Id         *string                                      `json:"id,omitempty"`
	Name       *string                                      `json:"name,omitempty"`
	Properties *BlobServiceProperties_Properties_STATUS_ARM `json:"properties,omitempty"`
	Sku        *Sku_STATUS_ARM                              `json:"sku,omitempty"`
	Type       *string                                      `json:"type,omitempty"`
}

// Deprecated version of BlobServiceProperties_Properties_STATUS. Use v1beta20210401.BlobServiceProperties_Properties_STATUS instead
type BlobServiceProperties_Properties_STATUS_ARM struct {
	AutomaticSnapshotPolicyEnabled *bool                                    `json:"automaticSnapshotPolicyEnabled,omitempty"`
	ChangeFeed                     *ChangeFeed_STATUS_ARM                   `json:"changeFeed,omitempty"`
	ContainerDeleteRetentionPolicy *DeleteRetentionPolicy_STATUS_ARM        `json:"containerDeleteRetentionPolicy,omitempty"`
	Cors                           *CorsRules_STATUS_ARM                    `json:"cors,omitempty"`
	DefaultServiceVersion          *string                                  `json:"defaultServiceVersion,omitempty"`
	DeleteRetentionPolicy          *DeleteRetentionPolicy_STATUS_ARM        `json:"deleteRetentionPolicy,omitempty"`
	IsVersioningEnabled            *bool                                    `json:"isVersioningEnabled,omitempty"`
	LastAccessTimeTrackingPolicy   *LastAccessTimeTrackingPolicy_STATUS_ARM `json:"lastAccessTimeTrackingPolicy,omitempty"`
	RestorePolicy                  *RestorePolicyProperties_STATUS_ARM      `json:"restorePolicy,omitempty"`
}

// Deprecated version of Sku_STATUS. Use v1beta20210401.Sku_STATUS instead
type Sku_STATUS_ARM struct {
	Name *SkuName_STATUS `json:"name,omitempty"`
	Tier *Tier_STATUS    `json:"tier,omitempty"`
}

// Deprecated version of ChangeFeed_STATUS. Use v1beta20210401.ChangeFeed_STATUS instead
type ChangeFeed_STATUS_ARM struct {
	Enabled         *bool `json:"enabled,omitempty"`
	RetentionInDays *int  `json:"retentionInDays,omitempty"`
}

// Deprecated version of CorsRules_STATUS. Use v1beta20210401.CorsRules_STATUS instead
type CorsRules_STATUS_ARM struct {
	CorsRules []CorsRule_STATUS_ARM `json:"corsRules,omitempty"`
}

// Deprecated version of DeleteRetentionPolicy_STATUS. Use v1beta20210401.DeleteRetentionPolicy_STATUS instead
type DeleteRetentionPolicy_STATUS_ARM struct {
	Days    *int  `json:"days,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

// Deprecated version of LastAccessTimeTrackingPolicy_STATUS. Use v1beta20210401.LastAccessTimeTrackingPolicy_STATUS instead
type LastAccessTimeTrackingPolicy_STATUS_ARM struct {
	BlobType                  []string                                  `json:"blobType,omitempty"`
	Enable                    *bool                                     `json:"enable,omitempty"`
	Name                      *LastAccessTimeTrackingPolicy_Name_STATUS `json:"name,omitempty"`
	TrackingGranularityInDays *int                                      `json:"trackingGranularityInDays,omitempty"`
}

// Deprecated version of RestorePolicyProperties_STATUS. Use v1beta20210401.RestorePolicyProperties_STATUS instead
type RestorePolicyProperties_STATUS_ARM struct {
	Days            *int    `json:"days,omitempty"`
	Enabled         *bool   `json:"enabled,omitempty"`
	LastEnabledTime *string `json:"lastEnabledTime,omitempty"`
	MinRestoreTime  *string `json:"minRestoreTime,omitempty"`
}

// Deprecated version of SkuName_STATUS. Use v1beta20210401.SkuName_STATUS instead
type SkuName_STATUS string

const (
	SkuName_STATUS_Premium_LRS     = SkuName_STATUS("Premium_LRS")
	SkuName_STATUS_Premium_ZRS     = SkuName_STATUS("Premium_ZRS")
	SkuName_STATUS_Standard_GRS    = SkuName_STATUS("Standard_GRS")
	SkuName_STATUS_Standard_GZRS   = SkuName_STATUS("Standard_GZRS")
	SkuName_STATUS_Standard_LRS    = SkuName_STATUS("Standard_LRS")
	SkuName_STATUS_Standard_RAGRS  = SkuName_STATUS("Standard_RAGRS")
	SkuName_STATUS_Standard_RAGZRS = SkuName_STATUS("Standard_RAGZRS")
	SkuName_STATUS_Standard_ZRS    = SkuName_STATUS("Standard_ZRS")
)

// Deprecated version of Tier_STATUS. Use v1beta20210401.Tier_STATUS instead
type Tier_STATUS string

const (
	Tier_STATUS_Premium  = Tier_STATUS("Premium")
	Tier_STATUS_Standard = Tier_STATUS("Standard")
)

// Deprecated version of CorsRule_STATUS. Use v1beta20210401.CorsRule_STATUS instead
type CorsRule_STATUS_ARM struct {
	AllowedHeaders  []string                         `json:"allowedHeaders,omitempty"`
	AllowedMethods  []CorsRule_AllowedMethods_STATUS `json:"allowedMethods,omitempty"`
	AllowedOrigins  []string                         `json:"allowedOrigins,omitempty"`
	ExposedHeaders  []string                         `json:"exposedHeaders,omitempty"`
	MaxAgeInSeconds *int                             `json:"maxAgeInSeconds,omitempty"`
}

// Deprecated version of LastAccessTimeTrackingPolicy_Name_STATUS. Use
// v1beta20210401.LastAccessTimeTrackingPolicy_Name_STATUS instead
type LastAccessTimeTrackingPolicy_Name_STATUS string

const LastAccessTimeTrackingPolicy_Name_STATUS_AccessTimeTracking = LastAccessTimeTrackingPolicy_Name_STATUS("AccessTimeTracking")

// Deprecated version of CorsRule_AllowedMethods_STATUS. Use v1beta20210401.CorsRule_AllowedMethods_STATUS instead
type CorsRule_AllowedMethods_STATUS string

const (
	CorsRule_AllowedMethods_STATUS_DELETE  = CorsRule_AllowedMethods_STATUS("DELETE")
	CorsRule_AllowedMethods_STATUS_GET     = CorsRule_AllowedMethods_STATUS("GET")
	CorsRule_AllowedMethods_STATUS_HEAD    = CorsRule_AllowedMethods_STATUS("HEAD")
	CorsRule_AllowedMethods_STATUS_MERGE   = CorsRule_AllowedMethods_STATUS("MERGE")
	CorsRule_AllowedMethods_STATUS_OPTIONS = CorsRule_AllowedMethods_STATUS("OPTIONS")
	CorsRule_AllowedMethods_STATUS_POST    = CorsRule_AllowedMethods_STATUS("POST")
	CorsRule_AllowedMethods_STATUS_PUT     = CorsRule_AllowedMethods_STATUS("PUT")
)