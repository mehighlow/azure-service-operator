// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211001

type Alias_STATUS_ARM struct {
	// Id: Fully qualified ID for the alias resource.
	Id *string `json:"id,omitempty"`

	// Name: Alias ID.
	Name *string `json:"name,omitempty"`

	// Properties: Subscription Alias response properties.
	Properties *SubscriptionAliasResponseProperties_STATUS_ARM `json:"properties,omitempty"`

	// SystemData: Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`

	// Type: Resource type, Microsoft.Subscription/aliases.
	Type *string `json:"type,omitempty"`
}

// Put subscription creation result properties.
type SubscriptionAliasResponseProperties_STATUS_ARM struct {
	// AcceptOwnershipState: The accept ownership state of the resource.
	AcceptOwnershipState *AcceptOwnershipState_STATUS `json:"acceptOwnershipState,omitempty"`

	// AcceptOwnershipUrl: Url to accept ownership of the subscription.
	AcceptOwnershipUrl *string `json:"acceptOwnershipUrl,omitempty"`
	BillingScope       *string `json:"billingScope,omitempty"`

	// CreatedTime: Created Time
	CreatedTime *string `json:"createdTime,omitempty"`

	// DisplayName: The display name of the subscription.
	DisplayName *string `json:"displayName,omitempty"`

	// ManagementGroupId: The Management Group Id.
	ManagementGroupId *string `json:"managementGroupId,omitempty"`

	// ProvisioningState: The provisioning state of the resource.
	ProvisioningState *SubscriptionAliasResponseProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// ResellerId: Reseller Id
	ResellerId *string `json:"resellerId,omitempty"`

	// SubscriptionId: Newly created subscription Id.
	SubscriptionId *string `json:"subscriptionId,omitempty"`

	// SubscriptionOwnerId: Owner Id of the subscription
	SubscriptionOwnerId *string `json:"subscriptionOwnerId,omitempty"`

	// Tags: Tags for the subscription
	Tags map[string]string `json:"tags,omitempty"`

	// Workload: The workload type of the subscription. It can be either Production or DevTest.
	Workload *Workload_STATUS `json:"workload,omitempty"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS_ARM struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemData_CreatedByType_STATUS `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemData_LastModifiedByType_STATUS `json:"lastModifiedByType,omitempty"`
}

// The accept ownership state of the resource.
type AcceptOwnershipState_STATUS string

const (
	AcceptOwnershipState_STATUS_Completed = AcceptOwnershipState_STATUS("Completed")
	AcceptOwnershipState_STATUS_Expired   = AcceptOwnershipState_STATUS("Expired")
	AcceptOwnershipState_STATUS_Pending   = AcceptOwnershipState_STATUS("Pending")
)

type SubscriptionAliasResponseProperties_ProvisioningState_STATUS string

const (
	SubscriptionAliasResponseProperties_ProvisioningState_STATUS_Accepted  = SubscriptionAliasResponseProperties_ProvisioningState_STATUS("Accepted")
	SubscriptionAliasResponseProperties_ProvisioningState_STATUS_Failed    = SubscriptionAliasResponseProperties_ProvisioningState_STATUS("Failed")
	SubscriptionAliasResponseProperties_ProvisioningState_STATUS_Succeeded = SubscriptionAliasResponseProperties_ProvisioningState_STATUS("Succeeded")
)

type SystemData_CreatedByType_STATUS string

const (
	SystemData_CreatedByType_STATUS_Application     = SystemData_CreatedByType_STATUS("Application")
	SystemData_CreatedByType_STATUS_Key             = SystemData_CreatedByType_STATUS("Key")
	SystemData_CreatedByType_STATUS_ManagedIdentity = SystemData_CreatedByType_STATUS("ManagedIdentity")
	SystemData_CreatedByType_STATUS_User            = SystemData_CreatedByType_STATUS("User")
)

type SystemData_LastModifiedByType_STATUS string

const (
	SystemData_LastModifiedByType_STATUS_Application     = SystemData_LastModifiedByType_STATUS("Application")
	SystemData_LastModifiedByType_STATUS_Key             = SystemData_LastModifiedByType_STATUS("Key")
	SystemData_LastModifiedByType_STATUS_ManagedIdentity = SystemData_LastModifiedByType_STATUS("ManagedIdentity")
	SystemData_LastModifiedByType_STATUS_User            = SystemData_LastModifiedByType_STATUS("User")
)

// The workload type of the subscription. It can be either Production or DevTest.
type Workload_STATUS string

const (
	Workload_STATUS_DevTest    = Workload_STATUS("DevTest")
	Workload_STATUS_Production = Workload_STATUS("Production")
)