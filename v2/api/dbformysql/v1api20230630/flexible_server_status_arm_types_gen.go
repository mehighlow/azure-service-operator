// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230630

import "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

type FlexibleServer_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. E.g.
	// "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id *string `json:"id,omitempty"`

	// Identity: The cmk identity for the server.
	Identity *MySQLServerIdentity_STATUS_ARM `json:"identity,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the server.
	Properties *ServerProperties_STATUS_ARM `json:"properties,omitempty"`

	// Sku: The SKU (pricing tier) of the server.
	Sku *MySQLServerSku_STATUS_ARM `json:"sku,omitempty"`

	// SystemData: Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// Properties to configure Identity for Bring your Own Keys
type MySQLServerIdentity_STATUS_ARM struct {
	// PrincipalId: ObjectId from the KeyVault
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: TenantId from the KeyVault
	TenantId *string `json:"tenantId,omitempty"`

	// Type: Type of managed service identity.
	Type *MySQLServerIdentity_Type_STATUS `json:"type,omitempty"`

	// UserAssignedIdentities: Metadata of user assigned identity.
	UserAssignedIdentities map[string]v1.JSON `json:"userAssignedIdentities,omitempty"`
}

// Billing information related properties of a server.
type MySQLServerSku_STATUS_ARM struct {
	// Name: The name of the sku, e.g. Standard_D32s_v3.
	Name *string `json:"name,omitempty"`

	// Tier: The tier of the particular SKU, e.g. GeneralPurpose.
	Tier *MySQLServerSku_Tier_STATUS `json:"tier,omitempty"`
}

// The properties of a server.
type ServerProperties_STATUS_ARM struct {
	// AdministratorLogin: The administrator's login name of a server. Can only be specified when the server is being created
	// (and is required for creation).
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	// AvailabilityZone: availability Zone information of the server.
	AvailabilityZone *string `json:"availabilityZone,omitempty"`

	// Backup: Backup related properties of a server.
	Backup *Backup_STATUS_ARM `json:"backup,omitempty"`

	// CreateMode: The mode to create a new MySQL server.
	CreateMode *ServerProperties_CreateMode_STATUS `json:"createMode,omitempty"`

	// DataEncryption: The Data Encryption for CMK.
	DataEncryption *DataEncryption_STATUS_ARM `json:"dataEncryption,omitempty"`

	// FullyQualifiedDomainName: The fully qualified domain name of a server.
	FullyQualifiedDomainName *string `json:"fullyQualifiedDomainName,omitempty"`

	// HighAvailability: High availability related properties of a server.
	HighAvailability *HighAvailability_STATUS_ARM `json:"highAvailability,omitempty"`

	// ImportSourceProperties: Source properties for import from storage.
	ImportSourceProperties *ImportSourceProperties_STATUS_ARM `json:"importSourceProperties,omitempty"`

	// MaintenanceWindow: Maintenance window of a server.
	MaintenanceWindow *MaintenanceWindow_STATUS_ARM `json:"maintenanceWindow,omitempty"`

	// Network: Network related properties of a server.
	Network *Network_STATUS_ARM `json:"network,omitempty"`

	// PrivateEndpointConnections: PrivateEndpointConnections related properties of a server.
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS_ARM `json:"privateEndpointConnections,omitempty"`

	// ReplicaCapacity: The maximum number of replicas that a primary server can have.
	ReplicaCapacity *int `json:"replicaCapacity,omitempty"`

	// ReplicationRole: The replication role.
	ReplicationRole *ReplicationRole_STATUS `json:"replicationRole,omitempty"`

	// RestorePointInTime: Restore point creation time (ISO8601 format), specifying the time to restore from.
	RestorePointInTime *string `json:"restorePointInTime,omitempty"`

	// SourceServerResourceId: The source MySQL server id.
	SourceServerResourceId *string `json:"sourceServerResourceId,omitempty"`

	// State: The state of a server.
	State *ServerProperties_State_STATUS `json:"state,omitempty"`

	// Storage: Storage related properties of a server.
	Storage *Storage_STATUS_ARM `json:"storage,omitempty"`

	// Version: Server version.
	Version *ServerVersion_STATUS `json:"version,omitempty"`
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

// Storage Profile properties of a server
type Backup_STATUS_ARM struct {
	// BackupRetentionDays: Backup retention days for the server.
	BackupRetentionDays *int `json:"backupRetentionDays,omitempty"`

	// EarliestRestoreDate: Earliest restore point creation time (ISO8601 format)
	EarliestRestoreDate *string `json:"earliestRestoreDate,omitempty"`

	// GeoRedundantBackup: Whether or not geo redundant backup is enabled.
	GeoRedundantBackup *EnableStatusEnum_STATUS `json:"geoRedundantBackup,omitempty"`
}

// The date encryption for cmk.
type DataEncryption_STATUS_ARM struct {
	// GeoBackupKeyURI: Geo backup key uri as key vault can't cross region, need cmk in same region as geo backup
	GeoBackupKeyURI *string `json:"geoBackupKeyURI,omitempty"`

	// GeoBackupUserAssignedIdentityId: Geo backup user identity resource id as identity can't cross region, need identity in
	// same region as geo backup
	GeoBackupUserAssignedIdentityId *string `json:"geoBackupUserAssignedIdentityId,omitempty"`

	// PrimaryKeyURI: Primary key uri
	PrimaryKeyURI *string `json:"primaryKeyURI,omitempty"`

	// PrimaryUserAssignedIdentityId: Primary user identity resource id
	PrimaryUserAssignedIdentityId *string `json:"primaryUserAssignedIdentityId,omitempty"`

	// Type: The key type, AzureKeyVault for enable cmk, SystemManaged for disable cmk.
	Type *DataEncryption_Type_STATUS `json:"type,omitempty"`
}

// Network related properties of a server
type HighAvailability_STATUS_ARM struct {
	// Mode: High availability mode for a server.
	Mode *HighAvailability_Mode_STATUS `json:"mode,omitempty"`

	// StandbyAvailabilityZone: Availability zone of the standby server.
	StandbyAvailabilityZone *string `json:"standbyAvailabilityZone,omitempty"`

	// State: The state of server high availability.
	State *HighAvailability_State_STATUS `json:"state,omitempty"`
}

// Import source related properties.
type ImportSourceProperties_STATUS_ARM struct {
	// DataDirPath: Relative path of data directory in storage.
	DataDirPath *string `json:"dataDirPath,omitempty"`

	// StorageType: Storage type of import source.
	StorageType *ImportSourceProperties_StorageType_STATUS `json:"storageType,omitempty"`

	// StorageUrl: Uri of the import source storage.
	StorageUrl *string `json:"storageUrl,omitempty"`
}

// Maintenance window of a server.
type MaintenanceWindow_STATUS_ARM struct {
	// CustomWindow: indicates whether custom window is enabled or disabled
	CustomWindow *string `json:"customWindow,omitempty"`

	// DayOfWeek: day of week for maintenance window
	DayOfWeek *int `json:"dayOfWeek,omitempty"`

	// StartHour: start hour for maintenance window
	StartHour *int `json:"startHour,omitempty"`

	// StartMinute: start minute for maintenance window
	StartMinute *int `json:"startMinute,omitempty"`
}

type MySQLServerIdentity_Type_STATUS string

const MySQLServerIdentity_Type_STATUS_UserAssigned = MySQLServerIdentity_Type_STATUS("UserAssigned")

// Mapping from string to MySQLServerIdentity_Type_STATUS
var mySQLServerIdentity_Type_STATUS_Values = map[string]MySQLServerIdentity_Type_STATUS{
	"userassigned": MySQLServerIdentity_Type_STATUS_UserAssigned,
}

type MySQLServerSku_Tier_STATUS string

const (
	MySQLServerSku_Tier_STATUS_Burstable       = MySQLServerSku_Tier_STATUS("Burstable")
	MySQLServerSku_Tier_STATUS_GeneralPurpose  = MySQLServerSku_Tier_STATUS("GeneralPurpose")
	MySQLServerSku_Tier_STATUS_MemoryOptimized = MySQLServerSku_Tier_STATUS("MemoryOptimized")
)

// Mapping from string to MySQLServerSku_Tier_STATUS
var mySQLServerSku_Tier_STATUS_Values = map[string]MySQLServerSku_Tier_STATUS{
	"burstable":       MySQLServerSku_Tier_STATUS_Burstable,
	"generalpurpose":  MySQLServerSku_Tier_STATUS_GeneralPurpose,
	"memoryoptimized": MySQLServerSku_Tier_STATUS_MemoryOptimized,
}

// Network related properties of a server
type Network_STATUS_ARM struct {
	// DelegatedSubnetResourceId: Delegated subnet resource id used to setup vnet for a server.
	DelegatedSubnetResourceId *string `json:"delegatedSubnetResourceId,omitempty"`

	// PrivateDnsZoneResourceId: Private DNS zone resource id.
	PrivateDnsZoneResourceId *string `json:"privateDnsZoneResourceId,omitempty"`

	// PublicNetworkAccess: Whether or not public network access is allowed for this server. Value is 'Disabled' when server
	// has VNet integration.
	PublicNetworkAccess *EnableStatusEnum_STATUS `json:"publicNetworkAccess,omitempty"`
}

// The private endpoint connection resource.
type PrivateEndpointConnection_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. E.g.
	// "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id *string `json:"id,omitempty"`
}

// Storage Profile properties of a server
type Storage_STATUS_ARM struct {
	// AutoGrow: Enable Storage Auto Grow or not.
	AutoGrow *EnableStatusEnum_STATUS `json:"autoGrow,omitempty"`

	// AutoIoScaling: Enable IO Auto Scaling or not.
	AutoIoScaling *EnableStatusEnum_STATUS `json:"autoIoScaling,omitempty"`

	// Iops: Storage IOPS for a server.
	Iops *int `json:"iops,omitempty"`

	// LogOnDisk: Enable Log On Disk or not.
	LogOnDisk *EnableStatusEnum_STATUS `json:"logOnDisk,omitempty"`

	// StorageSizeGB: Max storage size allowed for a server.
	StorageSizeGB *int `json:"storageSizeGB,omitempty"`

	// StorageSku: The sku name of the server storage.
	StorageSku *string `json:"storageSku,omitempty"`
}

type SystemData_CreatedByType_STATUS string

const (
	SystemData_CreatedByType_STATUS_Application     = SystemData_CreatedByType_STATUS("Application")
	SystemData_CreatedByType_STATUS_Key             = SystemData_CreatedByType_STATUS("Key")
	SystemData_CreatedByType_STATUS_ManagedIdentity = SystemData_CreatedByType_STATUS("ManagedIdentity")
	SystemData_CreatedByType_STATUS_User            = SystemData_CreatedByType_STATUS("User")
)

// Mapping from string to SystemData_CreatedByType_STATUS
var systemData_CreatedByType_STATUS_Values = map[string]SystemData_CreatedByType_STATUS{
	"application":     SystemData_CreatedByType_STATUS_Application,
	"key":             SystemData_CreatedByType_STATUS_Key,
	"managedidentity": SystemData_CreatedByType_STATUS_ManagedIdentity,
	"user":            SystemData_CreatedByType_STATUS_User,
}

type SystemData_LastModifiedByType_STATUS string

const (
	SystemData_LastModifiedByType_STATUS_Application     = SystemData_LastModifiedByType_STATUS("Application")
	SystemData_LastModifiedByType_STATUS_Key             = SystemData_LastModifiedByType_STATUS("Key")
	SystemData_LastModifiedByType_STATUS_ManagedIdentity = SystemData_LastModifiedByType_STATUS("ManagedIdentity")
	SystemData_LastModifiedByType_STATUS_User            = SystemData_LastModifiedByType_STATUS("User")
)

// Mapping from string to SystemData_LastModifiedByType_STATUS
var systemData_LastModifiedByType_STATUS_Values = map[string]SystemData_LastModifiedByType_STATUS{
	"application":     SystemData_LastModifiedByType_STATUS_Application,
	"key":             SystemData_LastModifiedByType_STATUS_Key,
	"managedidentity": SystemData_LastModifiedByType_STATUS_ManagedIdentity,
	"user":            SystemData_LastModifiedByType_STATUS_User,
}
