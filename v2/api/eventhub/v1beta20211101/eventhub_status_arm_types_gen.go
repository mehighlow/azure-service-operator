// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

type Eventhub_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties supplied to the Create Or Update Event Hub operation.
	Properties *Eventhub_Properties_STATUS_ARM `json:"properties,omitempty"`

	// SystemData: The system meta data relating to this resource.
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type *string `json:"type,omitempty"`
}

type Eventhub_Properties_STATUS_ARM struct {
	// CaptureDescription: Properties of capture description
	CaptureDescription *CaptureDescription_STATUS_ARM `json:"captureDescription,omitempty"`

	// CreatedAt: Exact time the Event Hub was created.
	CreatedAt *string `json:"createdAt,omitempty"`

	// MessageRetentionInDays: Number of days to retain the events for this Event Hub, value should be 1 to 7 days
	MessageRetentionInDays *int `json:"messageRetentionInDays,omitempty"`

	// PartitionCount: Number of partitions created for the Event Hub, allowed values are from 1 to 32 partitions.
	PartitionCount *int `json:"partitionCount,omitempty"`

	// PartitionIds: Current number of shards on the Event Hub.
	PartitionIds []string `json:"partitionIds,omitempty"`

	// Status: Enumerates the possible values for the status of the Event Hub.
	Status *Eventhub_Properties_Status_STATUS `json:"status,omitempty"`

	// UpdatedAt: The exact time the message was updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type CaptureDescription_STATUS_ARM struct {
	// Destination: Properties of Destination where capture will be stored. (Storage Account, Blob Names)
	Destination *Destination_STATUS_ARM `json:"destination,omitempty"`

	// Enabled: A value that indicates whether capture description is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Encoding: Enumerates the possible values for the encoding format of capture description. Note: 'AvroDeflate' will be
	// deprecated in New API Version
	Encoding *CaptureDescription_Encoding_STATUS `json:"encoding,omitempty"`

	// IntervalInSeconds: The time window allows you to set the frequency with which the capture to Azure Blobs will happen,
	// value should between 60 to 900 seconds
	IntervalInSeconds *int `json:"intervalInSeconds,omitempty"`

	// SizeLimitInBytes: The size window defines the amount of data built up in your Event Hub before an capture operation,
	// value should be between 10485760 to 524288000 bytes
	SizeLimitInBytes *int `json:"sizeLimitInBytes,omitempty"`

	// SkipEmptyArchives: A value that indicates whether to Skip Empty Archives
	SkipEmptyArchives *bool `json:"skipEmptyArchives,omitempty"`
}

type Eventhub_Properties_Status_STATUS string

const (
	Eventhub_Properties_Status_STATUS_Active          = Eventhub_Properties_Status_STATUS("Active")
	Eventhub_Properties_Status_STATUS_Creating        = Eventhub_Properties_Status_STATUS("Creating")
	Eventhub_Properties_Status_STATUS_Deleting        = Eventhub_Properties_Status_STATUS("Deleting")
	Eventhub_Properties_Status_STATUS_Disabled        = Eventhub_Properties_Status_STATUS("Disabled")
	Eventhub_Properties_Status_STATUS_ReceiveDisabled = Eventhub_Properties_Status_STATUS("ReceiveDisabled")
	Eventhub_Properties_Status_STATUS_Renaming        = Eventhub_Properties_Status_STATUS("Renaming")
	Eventhub_Properties_Status_STATUS_Restoring       = Eventhub_Properties_Status_STATUS("Restoring")
	Eventhub_Properties_Status_STATUS_SendDisabled    = Eventhub_Properties_Status_STATUS("SendDisabled")
	Eventhub_Properties_Status_STATUS_Unknown         = Eventhub_Properties_Status_STATUS("Unknown")
)

type CaptureDescription_Encoding_STATUS string

const (
	CaptureDescription_Encoding_STATUS_Avro        = CaptureDescription_Encoding_STATUS("Avro")
	CaptureDescription_Encoding_STATUS_AvroDeflate = CaptureDescription_Encoding_STATUS("AvroDeflate")
)

type Destination_STATUS_ARM struct {
	// Name: Name for capture destination
	Name *string `json:"name,omitempty"`

	// Properties: Properties describing the storage account, blob container and archive name format for capture destination
	Properties *Destination_Properties_STATUS_ARM `json:"properties,omitempty"`
}

type Destination_Properties_STATUS_ARM struct {
	// ArchiveNameFormat: Blob naming convention for archive, e.g.
	// {Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}. Here all the parameters
	// (Namespace,EventHub .. etc) are mandatory irrespective of order
	ArchiveNameFormat *string `json:"archiveNameFormat,omitempty"`

	// BlobContainer: Blob container Name
	BlobContainer *string `json:"blobContainer,omitempty"`

	// DataLakeAccountName: The Azure Data Lake Store name for the captured events
	DataLakeAccountName *string `json:"dataLakeAccountName,omitempty"`

	// DataLakeFolderPath: The destination folder path for the captured events
	DataLakeFolderPath *string `json:"dataLakeFolderPath,omitempty"`

	// DataLakeSubscriptionId: Subscription Id of Azure Data Lake Store
	DataLakeSubscriptionId *string `json:"dataLakeSubscriptionId,omitempty"`

	// StorageAccountResourceId: Resource id of the storage account to be used to create the blobs
	StorageAccountResourceId *string `json:"storageAccountResourceId,omitempty"`
}