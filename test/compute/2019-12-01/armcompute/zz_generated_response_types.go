//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute

// AvailabilitySetsClientCreateOrUpdateResponse contains the response from method AvailabilitySetsClient.CreateOrUpdate.
type AvailabilitySetsClientCreateOrUpdateResponse struct {
	AvailabilitySet
}

// AvailabilitySetsClientDeleteResponse contains the response from method AvailabilitySetsClient.Delete.
type AvailabilitySetsClientDeleteResponse struct {
	// placeholder for future response values
}

// AvailabilitySetsClientGetResponse contains the response from method AvailabilitySetsClient.Get.
type AvailabilitySetsClientGetResponse struct {
	AvailabilitySet
}

// AvailabilitySetsClientListAvailableSizesResponse contains the response from method AvailabilitySetsClient.ListAvailableSizes.
type AvailabilitySetsClientListAvailableSizesResponse struct {
	VirtualMachineSizeListResult
}

// AvailabilitySetsClientListBySubscriptionResponse contains the response from method AvailabilitySetsClient.ListBySubscription.
type AvailabilitySetsClientListBySubscriptionResponse struct {
	AvailabilitySetListResult
}

// AvailabilitySetsClientListResponse contains the response from method AvailabilitySetsClient.List.
type AvailabilitySetsClientListResponse struct {
	AvailabilitySetListResult
}

// AvailabilitySetsClientUpdateResponse contains the response from method AvailabilitySetsClient.Update.
type AvailabilitySetsClientUpdateResponse struct {
	AvailabilitySet
}

// ContainerServicesClientCreateOrUpdateResponse contains the response from method ContainerServicesClient.CreateOrUpdate.
type ContainerServicesClientCreateOrUpdateResponse struct {
	ContainerService
}

// ContainerServicesClientDeleteResponse contains the response from method ContainerServicesClient.Delete.
type ContainerServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// ContainerServicesClientGetResponse contains the response from method ContainerServicesClient.Get.
type ContainerServicesClientGetResponse struct {
	ContainerService
}

// ContainerServicesClientListByResourceGroupResponse contains the response from method ContainerServicesClient.ListByResourceGroup.
type ContainerServicesClientListByResourceGroupResponse struct {
	ContainerServiceListResult
}

// ContainerServicesClientListResponse contains the response from method ContainerServicesClient.List.
type ContainerServicesClientListResponse struct {
	ContainerServiceListResult
}

// DedicatedHostGroupsClientCreateOrUpdateResponse contains the response from method DedicatedHostGroupsClient.CreateOrUpdate.
type DedicatedHostGroupsClientCreateOrUpdateResponse struct {
	DedicatedHostGroup
}

// DedicatedHostGroupsClientDeleteResponse contains the response from method DedicatedHostGroupsClient.Delete.
type DedicatedHostGroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// DedicatedHostGroupsClientGetResponse contains the response from method DedicatedHostGroupsClient.Get.
type DedicatedHostGroupsClientGetResponse struct {
	DedicatedHostGroup
}

// DedicatedHostGroupsClientListByResourceGroupResponse contains the response from method DedicatedHostGroupsClient.ListByResourceGroup.
type DedicatedHostGroupsClientListByResourceGroupResponse struct {
	DedicatedHostGroupListResult
}

// DedicatedHostGroupsClientListBySubscriptionResponse contains the response from method DedicatedHostGroupsClient.ListBySubscription.
type DedicatedHostGroupsClientListBySubscriptionResponse struct {
	DedicatedHostGroupListResult
}

// DedicatedHostGroupsClientUpdateResponse contains the response from method DedicatedHostGroupsClient.Update.
type DedicatedHostGroupsClientUpdateResponse struct {
	DedicatedHostGroup
}

// DedicatedHostsClientCreateOrUpdateResponse contains the response from method DedicatedHostsClient.CreateOrUpdate.
type DedicatedHostsClientCreateOrUpdateResponse struct {
	DedicatedHost
}

// DedicatedHostsClientDeleteResponse contains the response from method DedicatedHostsClient.Delete.
type DedicatedHostsClientDeleteResponse struct {
	// placeholder for future response values
}

// DedicatedHostsClientGetResponse contains the response from method DedicatedHostsClient.Get.
type DedicatedHostsClientGetResponse struct {
	DedicatedHost
}

// DedicatedHostsClientListByHostGroupResponse contains the response from method DedicatedHostsClient.ListByHostGroup.
type DedicatedHostsClientListByHostGroupResponse struct {
	DedicatedHostListResult
}

// DedicatedHostsClientUpdateResponse contains the response from method DedicatedHostsClient.Update.
type DedicatedHostsClientUpdateResponse struct {
	DedicatedHost
}

// DiskEncryptionSetsClientCreateOrUpdateResponse contains the response from method DiskEncryptionSetsClient.CreateOrUpdate.
type DiskEncryptionSetsClientCreateOrUpdateResponse struct {
	DiskEncryptionSet
}

// DiskEncryptionSetsClientDeleteResponse contains the response from method DiskEncryptionSetsClient.Delete.
type DiskEncryptionSetsClientDeleteResponse struct {
	// placeholder for future response values
}

// DiskEncryptionSetsClientGetResponse contains the response from method DiskEncryptionSetsClient.Get.
type DiskEncryptionSetsClientGetResponse struct {
	DiskEncryptionSet
}

// DiskEncryptionSetsClientListByResourceGroupResponse contains the response from method DiskEncryptionSetsClient.ListByResourceGroup.
type DiskEncryptionSetsClientListByResourceGroupResponse struct {
	DiskEncryptionSetList
}

// DiskEncryptionSetsClientListResponse contains the response from method DiskEncryptionSetsClient.List.
type DiskEncryptionSetsClientListResponse struct {
	DiskEncryptionSetList
}

// DiskEncryptionSetsClientUpdateResponse contains the response from method DiskEncryptionSetsClient.Update.
type DiskEncryptionSetsClientUpdateResponse struct {
	DiskEncryptionSet
}

// DisksClientCreateOrUpdateResponse contains the response from method DisksClient.CreateOrUpdate.
type DisksClientCreateOrUpdateResponse struct {
	Disk
}

// DisksClientDeleteResponse contains the response from method DisksClient.Delete.
type DisksClientDeleteResponse struct {
	// placeholder for future response values
}

// DisksClientGetResponse contains the response from method DisksClient.Get.
type DisksClientGetResponse struct {
	Disk
}

// DisksClientGrantAccessResponse contains the response from method DisksClient.GrantAccess.
type DisksClientGrantAccessResponse struct {
	AccessURI
}

// DisksClientListByResourceGroupResponse contains the response from method DisksClient.ListByResourceGroup.
type DisksClientListByResourceGroupResponse struct {
	DiskList
}

// DisksClientListResponse contains the response from method DisksClient.List.
type DisksClientListResponse struct {
	DiskList
}

// DisksClientRevokeAccessResponse contains the response from method DisksClient.RevokeAccess.
type DisksClientRevokeAccessResponse struct {
	// placeholder for future response values
}

// DisksClientUpdateResponse contains the response from method DisksClient.Update.
type DisksClientUpdateResponse struct {
	Disk
}

// GalleriesClientCreateOrUpdateResponse contains the response from method GalleriesClient.CreateOrUpdate.
type GalleriesClientCreateOrUpdateResponse struct {
	Gallery
}

// GalleriesClientDeleteResponse contains the response from method GalleriesClient.Delete.
type GalleriesClientDeleteResponse struct {
	// placeholder for future response values
}

// GalleriesClientGetResponse contains the response from method GalleriesClient.Get.
type GalleriesClientGetResponse struct {
	Gallery
}

// GalleriesClientListByResourceGroupResponse contains the response from method GalleriesClient.ListByResourceGroup.
type GalleriesClientListByResourceGroupResponse struct {
	GalleryList
}

// GalleriesClientListResponse contains the response from method GalleriesClient.List.
type GalleriesClientListResponse struct {
	GalleryList
}

// GalleriesClientUpdateResponse contains the response from method GalleriesClient.Update.
type GalleriesClientUpdateResponse struct {
	Gallery
}

// GalleryApplicationVersionsClientCreateOrUpdateResponse contains the response from method GalleryApplicationVersionsClient.CreateOrUpdate.
type GalleryApplicationVersionsClientCreateOrUpdateResponse struct {
	GalleryApplicationVersion
}

// GalleryApplicationVersionsClientDeleteResponse contains the response from method GalleryApplicationVersionsClient.Delete.
type GalleryApplicationVersionsClientDeleteResponse struct {
	// placeholder for future response values
}

// GalleryApplicationVersionsClientGetResponse contains the response from method GalleryApplicationVersionsClient.Get.
type GalleryApplicationVersionsClientGetResponse struct {
	GalleryApplicationVersion
}

// GalleryApplicationVersionsClientListByGalleryApplicationResponse contains the response from method GalleryApplicationVersionsClient.ListByGalleryApplication.
type GalleryApplicationVersionsClientListByGalleryApplicationResponse struct {
	GalleryApplicationVersionList
}

// GalleryApplicationVersionsClientUpdateResponse contains the response from method GalleryApplicationVersionsClient.Update.
type GalleryApplicationVersionsClientUpdateResponse struct {
	GalleryApplicationVersion
}

// GalleryApplicationsClientCreateOrUpdateResponse contains the response from method GalleryApplicationsClient.CreateOrUpdate.
type GalleryApplicationsClientCreateOrUpdateResponse struct {
	GalleryApplication
}

// GalleryApplicationsClientDeleteResponse contains the response from method GalleryApplicationsClient.Delete.
type GalleryApplicationsClientDeleteResponse struct {
	// placeholder for future response values
}

// GalleryApplicationsClientGetResponse contains the response from method GalleryApplicationsClient.Get.
type GalleryApplicationsClientGetResponse struct {
	GalleryApplication
}

// GalleryApplicationsClientListByGalleryResponse contains the response from method GalleryApplicationsClient.ListByGallery.
type GalleryApplicationsClientListByGalleryResponse struct {
	GalleryApplicationList
}

// GalleryApplicationsClientUpdateResponse contains the response from method GalleryApplicationsClient.Update.
type GalleryApplicationsClientUpdateResponse struct {
	GalleryApplication
}

// GalleryImageVersionsClientCreateOrUpdateResponse contains the response from method GalleryImageVersionsClient.CreateOrUpdate.
type GalleryImageVersionsClientCreateOrUpdateResponse struct {
	GalleryImageVersion
}

// GalleryImageVersionsClientDeleteResponse contains the response from method GalleryImageVersionsClient.Delete.
type GalleryImageVersionsClientDeleteResponse struct {
	// placeholder for future response values
}

// GalleryImageVersionsClientGetResponse contains the response from method GalleryImageVersionsClient.Get.
type GalleryImageVersionsClientGetResponse struct {
	GalleryImageVersion
}

// GalleryImageVersionsClientListByGalleryImageResponse contains the response from method GalleryImageVersionsClient.ListByGalleryImage.
type GalleryImageVersionsClientListByGalleryImageResponse struct {
	GalleryImageVersionList
}

// GalleryImageVersionsClientUpdateResponse contains the response from method GalleryImageVersionsClient.Update.
type GalleryImageVersionsClientUpdateResponse struct {
	GalleryImageVersion
}

// GalleryImagesClientCreateOrUpdateResponse contains the response from method GalleryImagesClient.CreateOrUpdate.
type GalleryImagesClientCreateOrUpdateResponse struct {
	GalleryImage
}

// GalleryImagesClientDeleteResponse contains the response from method GalleryImagesClient.Delete.
type GalleryImagesClientDeleteResponse struct {
	// placeholder for future response values
}

// GalleryImagesClientGetResponse contains the response from method GalleryImagesClient.Get.
type GalleryImagesClientGetResponse struct {
	GalleryImage
}

// GalleryImagesClientListByGalleryResponse contains the response from method GalleryImagesClient.ListByGallery.
type GalleryImagesClientListByGalleryResponse struct {
	GalleryImageList
}

// GalleryImagesClientUpdateResponse contains the response from method GalleryImagesClient.Update.
type GalleryImagesClientUpdateResponse struct {
	GalleryImage
}

// ImagesClientCreateOrUpdateResponse contains the response from method ImagesClient.CreateOrUpdate.
type ImagesClientCreateOrUpdateResponse struct {
	Image
}

// ImagesClientDeleteResponse contains the response from method ImagesClient.Delete.
type ImagesClientDeleteResponse struct {
	// placeholder for future response values
}

// ImagesClientGetResponse contains the response from method ImagesClient.Get.
type ImagesClientGetResponse struct {
	Image
}

// ImagesClientListByResourceGroupResponse contains the response from method ImagesClient.ListByResourceGroup.
type ImagesClientListByResourceGroupResponse struct {
	ImageListResult
}

// ImagesClientListResponse contains the response from method ImagesClient.List.
type ImagesClientListResponse struct {
	ImageListResult
}

// ImagesClientUpdateResponse contains the response from method ImagesClient.Update.
type ImagesClientUpdateResponse struct {
	Image
}

// LogAnalyticsClientExportRequestRateByIntervalResponse contains the response from method LogAnalyticsClient.ExportRequestRateByInterval.
type LogAnalyticsClientExportRequestRateByIntervalResponse struct {
	LogAnalyticsOperationResult
}

// LogAnalyticsClientExportThrottledRequestsResponse contains the response from method LogAnalyticsClient.ExportThrottledRequests.
type LogAnalyticsClientExportThrottledRequestsResponse struct {
	LogAnalyticsOperationResult
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// ProximityPlacementGroupsClientCreateOrUpdateResponse contains the response from method ProximityPlacementGroupsClient.CreateOrUpdate.
type ProximityPlacementGroupsClientCreateOrUpdateResponse struct {
	ProximityPlacementGroup
}

// ProximityPlacementGroupsClientDeleteResponse contains the response from method ProximityPlacementGroupsClient.Delete.
type ProximityPlacementGroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// ProximityPlacementGroupsClientGetResponse contains the response from method ProximityPlacementGroupsClient.Get.
type ProximityPlacementGroupsClientGetResponse struct {
	ProximityPlacementGroup
}

// ProximityPlacementGroupsClientListByResourceGroupResponse contains the response from method ProximityPlacementGroupsClient.ListByResourceGroup.
type ProximityPlacementGroupsClientListByResourceGroupResponse struct {
	ProximityPlacementGroupListResult
}

// ProximityPlacementGroupsClientListBySubscriptionResponse contains the response from method ProximityPlacementGroupsClient.ListBySubscription.
type ProximityPlacementGroupsClientListBySubscriptionResponse struct {
	ProximityPlacementGroupListResult
}

// ProximityPlacementGroupsClientUpdateResponse contains the response from method ProximityPlacementGroupsClient.Update.
type ProximityPlacementGroupsClientUpdateResponse struct {
	ProximityPlacementGroup
}

// ResourceSKUsClientListResponse contains the response from method ResourceSKUsClient.List.
type ResourceSKUsClientListResponse struct {
	ResourceSKUsResult
}

// SSHPublicKeysClientCreateResponse contains the response from method SSHPublicKeysClient.Create.
type SSHPublicKeysClientCreateResponse struct {
	SSHPublicKeyResource
}

// SSHPublicKeysClientDeleteResponse contains the response from method SSHPublicKeysClient.Delete.
type SSHPublicKeysClientDeleteResponse struct {
	// placeholder for future response values
}

// SSHPublicKeysClientGenerateKeyPairResponse contains the response from method SSHPublicKeysClient.GenerateKeyPair.
type SSHPublicKeysClientGenerateKeyPairResponse struct {
	SSHPublicKeyGenerateKeyPairResult
}

// SSHPublicKeysClientGetResponse contains the response from method SSHPublicKeysClient.Get.
type SSHPublicKeysClientGetResponse struct {
	SSHPublicKeyResource
}

// SSHPublicKeysClientListByResourceGroupResponse contains the response from method SSHPublicKeysClient.ListByResourceGroup.
type SSHPublicKeysClientListByResourceGroupResponse struct {
	SSHPublicKeysGroupListResult
}

// SSHPublicKeysClientListBySubscriptionResponse contains the response from method SSHPublicKeysClient.ListBySubscription.
type SSHPublicKeysClientListBySubscriptionResponse struct {
	SSHPublicKeysGroupListResult
}

// SSHPublicKeysClientUpdateResponse contains the response from method SSHPublicKeysClient.Update.
type SSHPublicKeysClientUpdateResponse struct {
	SSHPublicKeyResource
}

// SnapshotsClientCreateOrUpdateResponse contains the response from method SnapshotsClient.CreateOrUpdate.
type SnapshotsClientCreateOrUpdateResponse struct {
	Snapshot
}

// SnapshotsClientDeleteResponse contains the response from method SnapshotsClient.Delete.
type SnapshotsClientDeleteResponse struct {
	// placeholder for future response values
}

// SnapshotsClientGetResponse contains the response from method SnapshotsClient.Get.
type SnapshotsClientGetResponse struct {
	Snapshot
}

// SnapshotsClientGrantAccessResponse contains the response from method SnapshotsClient.GrantAccess.
type SnapshotsClientGrantAccessResponse struct {
	AccessURI
}

// SnapshotsClientListByResourceGroupResponse contains the response from method SnapshotsClient.ListByResourceGroup.
type SnapshotsClientListByResourceGroupResponse struct {
	SnapshotList
}

// SnapshotsClientListResponse contains the response from method SnapshotsClient.List.
type SnapshotsClientListResponse struct {
	SnapshotList
}

// SnapshotsClientRevokeAccessResponse contains the response from method SnapshotsClient.RevokeAccess.
type SnapshotsClientRevokeAccessResponse struct {
	// placeholder for future response values
}

// SnapshotsClientUpdateResponse contains the response from method SnapshotsClient.Update.
type SnapshotsClientUpdateResponse struct {
	Snapshot
}

// UsageClientListResponse contains the response from method UsageClient.List.
type UsageClientListResponse struct {
	ListUsagesResult
}

// VirtualMachineExtensionImagesClientGetResponse contains the response from method VirtualMachineExtensionImagesClient.Get.
type VirtualMachineExtensionImagesClientGetResponse struct {
	VirtualMachineExtensionImage
}

// VirtualMachineExtensionImagesClientListTypesResponse contains the response from method VirtualMachineExtensionImagesClient.ListTypes.
type VirtualMachineExtensionImagesClientListTypesResponse struct {
	// Array of VirtualMachineExtensionImage
	VirtualMachineExtensionImageArray []*VirtualMachineExtensionImage
}

// VirtualMachineExtensionImagesClientListVersionsResponse contains the response from method VirtualMachineExtensionImagesClient.ListVersions.
type VirtualMachineExtensionImagesClientListVersionsResponse struct {
	// Array of VirtualMachineExtensionImage
	VirtualMachineExtensionImageArray []*VirtualMachineExtensionImage
}

// VirtualMachineExtensionsClientCreateOrUpdateResponse contains the response from method VirtualMachineExtensionsClient.CreateOrUpdate.
type VirtualMachineExtensionsClientCreateOrUpdateResponse struct {
	VirtualMachineExtension
}

// VirtualMachineExtensionsClientDeleteResponse contains the response from method VirtualMachineExtensionsClient.Delete.
type VirtualMachineExtensionsClientDeleteResponse struct {
	// placeholder for future response values
}

// VirtualMachineExtensionsClientGetResponse contains the response from method VirtualMachineExtensionsClient.Get.
type VirtualMachineExtensionsClientGetResponse struct {
	VirtualMachineExtension
}

// VirtualMachineExtensionsClientListResponse contains the response from method VirtualMachineExtensionsClient.List.
type VirtualMachineExtensionsClientListResponse struct {
	VirtualMachineExtensionsListResult
}

// VirtualMachineExtensionsClientUpdateResponse contains the response from method VirtualMachineExtensionsClient.Update.
type VirtualMachineExtensionsClientUpdateResponse struct {
	VirtualMachineExtension
}

// VirtualMachineImagesClientGetResponse contains the response from method VirtualMachineImagesClient.Get.
type VirtualMachineImagesClientGetResponse struct {
	VirtualMachineImage
}

// VirtualMachineImagesClientListOffersResponse contains the response from method VirtualMachineImagesClient.ListOffers.
type VirtualMachineImagesClientListOffersResponse struct {
	// Array of VirtualMachineImageResource
	VirtualMachineImageResourceArray []*VirtualMachineImageResource
}

// VirtualMachineImagesClientListPublishersResponse contains the response from method VirtualMachineImagesClient.ListPublishers.
type VirtualMachineImagesClientListPublishersResponse struct {
	// Array of VirtualMachineImageResource
	VirtualMachineImageResourceArray []*VirtualMachineImageResource
}

// VirtualMachineImagesClientListResponse contains the response from method VirtualMachineImagesClient.List.
type VirtualMachineImagesClientListResponse struct {
	// Array of VirtualMachineImageResource
	VirtualMachineImageResourceArray []*VirtualMachineImageResource
}

// VirtualMachineImagesClientListSKUsResponse contains the response from method VirtualMachineImagesClient.ListSKUs.
type VirtualMachineImagesClientListSKUsResponse struct {
	// Array of VirtualMachineImageResource
	VirtualMachineImageResourceArray []*VirtualMachineImageResource
}

// VirtualMachineRunCommandsClientGetResponse contains the response from method VirtualMachineRunCommandsClient.Get.
type VirtualMachineRunCommandsClientGetResponse struct {
	RunCommandDocument
}

// VirtualMachineRunCommandsClientListResponse contains the response from method VirtualMachineRunCommandsClient.List.
type VirtualMachineRunCommandsClientListResponse struct {
	RunCommandListResult
}

// VirtualMachineScaleSetExtensionsClientCreateOrUpdateResponse contains the response from method VirtualMachineScaleSetExtensionsClient.CreateOrUpdate.
type VirtualMachineScaleSetExtensionsClientCreateOrUpdateResponse struct {
	VirtualMachineScaleSetExtension
}

// VirtualMachineScaleSetExtensionsClientDeleteResponse contains the response from method VirtualMachineScaleSetExtensionsClient.Delete.
type VirtualMachineScaleSetExtensionsClientDeleteResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetExtensionsClientGetResponse contains the response from method VirtualMachineScaleSetExtensionsClient.Get.
type VirtualMachineScaleSetExtensionsClientGetResponse struct {
	VirtualMachineScaleSetExtension
}

// VirtualMachineScaleSetExtensionsClientListResponse contains the response from method VirtualMachineScaleSetExtensionsClient.List.
type VirtualMachineScaleSetExtensionsClientListResponse struct {
	VirtualMachineScaleSetExtensionListResult
}

// VirtualMachineScaleSetExtensionsClientUpdateResponse contains the response from method VirtualMachineScaleSetExtensionsClient.Update.
type VirtualMachineScaleSetExtensionsClientUpdateResponse struct {
	VirtualMachineScaleSetExtension
}

// VirtualMachineScaleSetRollingUpgradesClientCancelResponse contains the response from method VirtualMachineScaleSetRollingUpgradesClient.Cancel.
type VirtualMachineScaleSetRollingUpgradesClientCancelResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetRollingUpgradesClientGetLatestResponse contains the response from method VirtualMachineScaleSetRollingUpgradesClient.GetLatest.
type VirtualMachineScaleSetRollingUpgradesClientGetLatestResponse struct {
	RollingUpgradeStatusInfo
}

// VirtualMachineScaleSetRollingUpgradesClientStartExtensionUpgradeResponse contains the response from method VirtualMachineScaleSetRollingUpgradesClient.StartExtensionUpgrade.
type VirtualMachineScaleSetRollingUpgradesClientStartExtensionUpgradeResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetRollingUpgradesClientStartOSUpgradeResponse contains the response from method VirtualMachineScaleSetRollingUpgradesClient.StartOSUpgrade.
type VirtualMachineScaleSetRollingUpgradesClientStartOSUpgradeResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetVMExtensionsClientCreateOrUpdateResponse contains the response from method VirtualMachineScaleSetVMExtensionsClient.CreateOrUpdate.
type VirtualMachineScaleSetVMExtensionsClientCreateOrUpdateResponse struct {
	VirtualMachineExtension
}

// VirtualMachineScaleSetVMExtensionsClientDeleteResponse contains the response from method VirtualMachineScaleSetVMExtensionsClient.Delete.
type VirtualMachineScaleSetVMExtensionsClientDeleteResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetVMExtensionsClientGetResponse contains the response from method VirtualMachineScaleSetVMExtensionsClient.Get.
type VirtualMachineScaleSetVMExtensionsClientGetResponse struct {
	VirtualMachineExtension
}

// VirtualMachineScaleSetVMExtensionsClientListResponse contains the response from method VirtualMachineScaleSetVMExtensionsClient.List.
type VirtualMachineScaleSetVMExtensionsClientListResponse struct {
	VirtualMachineExtensionsListResult
}

// VirtualMachineScaleSetVMExtensionsClientUpdateResponse contains the response from method VirtualMachineScaleSetVMExtensionsClient.Update.
type VirtualMachineScaleSetVMExtensionsClientUpdateResponse struct {
	VirtualMachineExtension
}

// VirtualMachineScaleSetVMsClientDeallocateResponse contains the response from method VirtualMachineScaleSetVMsClient.Deallocate.
type VirtualMachineScaleSetVMsClientDeallocateResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetVMsClientDeleteResponse contains the response from method VirtualMachineScaleSetVMsClient.Delete.
type VirtualMachineScaleSetVMsClientDeleteResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetVMsClientGetInstanceViewResponse contains the response from method VirtualMachineScaleSetVMsClient.GetInstanceView.
type VirtualMachineScaleSetVMsClientGetInstanceViewResponse struct {
	VirtualMachineScaleSetVMInstanceView
}

// VirtualMachineScaleSetVMsClientGetResponse contains the response from method VirtualMachineScaleSetVMsClient.Get.
type VirtualMachineScaleSetVMsClientGetResponse struct {
	VirtualMachineScaleSetVM
}

// VirtualMachineScaleSetVMsClientListResponse contains the response from method VirtualMachineScaleSetVMsClient.List.
type VirtualMachineScaleSetVMsClientListResponse struct {
	VirtualMachineScaleSetVMListResult
}

// VirtualMachineScaleSetVMsClientPerformMaintenanceResponse contains the response from method VirtualMachineScaleSetVMsClient.PerformMaintenance.
type VirtualMachineScaleSetVMsClientPerformMaintenanceResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetVMsClientPowerOffResponse contains the response from method VirtualMachineScaleSetVMsClient.PowerOff.
type VirtualMachineScaleSetVMsClientPowerOffResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetVMsClientRedeployResponse contains the response from method VirtualMachineScaleSetVMsClient.Redeploy.
type VirtualMachineScaleSetVMsClientRedeployResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetVMsClientReimageAllResponse contains the response from method VirtualMachineScaleSetVMsClient.ReimageAll.
type VirtualMachineScaleSetVMsClientReimageAllResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetVMsClientReimageResponse contains the response from method VirtualMachineScaleSetVMsClient.Reimage.
type VirtualMachineScaleSetVMsClientReimageResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetVMsClientRestartResponse contains the response from method VirtualMachineScaleSetVMsClient.Restart.
type VirtualMachineScaleSetVMsClientRestartResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetVMsClientRunCommandResponse contains the response from method VirtualMachineScaleSetVMsClient.RunCommand.
type VirtualMachineScaleSetVMsClientRunCommandResponse struct {
	RunCommandResult
}

// VirtualMachineScaleSetVMsClientSimulateEvictionResponse contains the response from method VirtualMachineScaleSetVMsClient.SimulateEviction.
type VirtualMachineScaleSetVMsClientSimulateEvictionResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetVMsClientStartResponse contains the response from method VirtualMachineScaleSetVMsClient.Start.
type VirtualMachineScaleSetVMsClientStartResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetVMsClientUpdateResponse contains the response from method VirtualMachineScaleSetVMsClient.Update.
type VirtualMachineScaleSetVMsClientUpdateResponse struct {
	VirtualMachineScaleSetVM
}

// VirtualMachineScaleSetsClientConvertToSinglePlacementGroupResponse contains the response from method VirtualMachineScaleSetsClient.ConvertToSinglePlacementGroup.
type VirtualMachineScaleSetsClientConvertToSinglePlacementGroupResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetsClientCreateOrUpdateResponse contains the response from method VirtualMachineScaleSetsClient.CreateOrUpdate.
type VirtualMachineScaleSetsClientCreateOrUpdateResponse struct {
	VirtualMachineScaleSet
}

// VirtualMachineScaleSetsClientDeallocateResponse contains the response from method VirtualMachineScaleSetsClient.Deallocate.
type VirtualMachineScaleSetsClientDeallocateResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetsClientDeleteInstancesResponse contains the response from method VirtualMachineScaleSetsClient.DeleteInstances.
type VirtualMachineScaleSetsClientDeleteInstancesResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetsClientDeleteResponse contains the response from method VirtualMachineScaleSetsClient.Delete.
type VirtualMachineScaleSetsClientDeleteResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetsClientForceRecoveryServiceFabricPlatformUpdateDomainWalkResponse contains the response from method
// VirtualMachineScaleSetsClient.ForceRecoveryServiceFabricPlatformUpdateDomainWalk.
type VirtualMachineScaleSetsClientForceRecoveryServiceFabricPlatformUpdateDomainWalkResponse struct {
	RecoveryWalkResponse
}

// VirtualMachineScaleSetsClientGetInstanceViewResponse contains the response from method VirtualMachineScaleSetsClient.GetInstanceView.
type VirtualMachineScaleSetsClientGetInstanceViewResponse struct {
	VirtualMachineScaleSetInstanceView
}

// VirtualMachineScaleSetsClientGetOSUpgradeHistoryResponse contains the response from method VirtualMachineScaleSetsClient.GetOSUpgradeHistory.
type VirtualMachineScaleSetsClientGetOSUpgradeHistoryResponse struct {
	VirtualMachineScaleSetListOSUpgradeHistory
}

// VirtualMachineScaleSetsClientGetResponse contains the response from method VirtualMachineScaleSetsClient.Get.
type VirtualMachineScaleSetsClientGetResponse struct {
	VirtualMachineScaleSet
}

// VirtualMachineScaleSetsClientListAllResponse contains the response from method VirtualMachineScaleSetsClient.ListAll.
type VirtualMachineScaleSetsClientListAllResponse struct {
	VirtualMachineScaleSetListWithLinkResult
}

// VirtualMachineScaleSetsClientListResponse contains the response from method VirtualMachineScaleSetsClient.List.
type VirtualMachineScaleSetsClientListResponse struct {
	VirtualMachineScaleSetListResult
}

// VirtualMachineScaleSetsClientListSKUsResponse contains the response from method VirtualMachineScaleSetsClient.ListSKUs.
type VirtualMachineScaleSetsClientListSKUsResponse struct {
	VirtualMachineScaleSetListSKUsResult
}

// VirtualMachineScaleSetsClientPerformMaintenanceResponse contains the response from method VirtualMachineScaleSetsClient.PerformMaintenance.
type VirtualMachineScaleSetsClientPerformMaintenanceResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetsClientPowerOffResponse contains the response from method VirtualMachineScaleSetsClient.PowerOff.
type VirtualMachineScaleSetsClientPowerOffResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetsClientRedeployResponse contains the response from method VirtualMachineScaleSetsClient.Redeploy.
type VirtualMachineScaleSetsClientRedeployResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetsClientReimageAllResponse contains the response from method VirtualMachineScaleSetsClient.ReimageAll.
type VirtualMachineScaleSetsClientReimageAllResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetsClientReimageResponse contains the response from method VirtualMachineScaleSetsClient.Reimage.
type VirtualMachineScaleSetsClientReimageResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetsClientRestartResponse contains the response from method VirtualMachineScaleSetsClient.Restart.
type VirtualMachineScaleSetsClientRestartResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetsClientSetOrchestrationServiceStateResponse contains the response from method VirtualMachineScaleSetsClient.SetOrchestrationServiceState.
type VirtualMachineScaleSetsClientSetOrchestrationServiceStateResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetsClientStartResponse contains the response from method VirtualMachineScaleSetsClient.Start.
type VirtualMachineScaleSetsClientStartResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetsClientUpdateInstancesResponse contains the response from method VirtualMachineScaleSetsClient.UpdateInstances.
type VirtualMachineScaleSetsClientUpdateInstancesResponse struct {
	// placeholder for future response values
}

// VirtualMachineScaleSetsClientUpdateResponse contains the response from method VirtualMachineScaleSetsClient.Update.
type VirtualMachineScaleSetsClientUpdateResponse struct {
	VirtualMachineScaleSet
}

// VirtualMachineSizesClientListResponse contains the response from method VirtualMachineSizesClient.List.
type VirtualMachineSizesClientListResponse struct {
	VirtualMachineSizeListResult
}

// VirtualMachinesClientCaptureResponse contains the response from method VirtualMachinesClient.Capture.
type VirtualMachinesClientCaptureResponse struct {
	VirtualMachineCaptureResult
}

// VirtualMachinesClientConvertToManagedDisksResponse contains the response from method VirtualMachinesClient.ConvertToManagedDisks.
type VirtualMachinesClientConvertToManagedDisksResponse struct {
	// placeholder for future response values
}

// VirtualMachinesClientCreateOrUpdateResponse contains the response from method VirtualMachinesClient.CreateOrUpdate.
type VirtualMachinesClientCreateOrUpdateResponse struct {
	VirtualMachine
}

// VirtualMachinesClientDeallocateResponse contains the response from method VirtualMachinesClient.Deallocate.
type VirtualMachinesClientDeallocateResponse struct {
	// placeholder for future response values
}

// VirtualMachinesClientDeleteResponse contains the response from method VirtualMachinesClient.Delete.
type VirtualMachinesClientDeleteResponse struct {
	// placeholder for future response values
}

// VirtualMachinesClientGeneralizeResponse contains the response from method VirtualMachinesClient.Generalize.
type VirtualMachinesClientGeneralizeResponse struct {
	// placeholder for future response values
}

// VirtualMachinesClientGetResponse contains the response from method VirtualMachinesClient.Get.
type VirtualMachinesClientGetResponse struct {
	VirtualMachine
}

// VirtualMachinesClientInstanceViewResponse contains the response from method VirtualMachinesClient.InstanceView.
type VirtualMachinesClientInstanceViewResponse struct {
	VirtualMachineInstanceView
}

// VirtualMachinesClientListAllResponse contains the response from method VirtualMachinesClient.ListAll.
type VirtualMachinesClientListAllResponse struct {
	VirtualMachineListResult
}

// VirtualMachinesClientListAvailableSizesResponse contains the response from method VirtualMachinesClient.ListAvailableSizes.
type VirtualMachinesClientListAvailableSizesResponse struct {
	VirtualMachineSizeListResult
}

// VirtualMachinesClientListByLocationResponse contains the response from method VirtualMachinesClient.ListByLocation.
type VirtualMachinesClientListByLocationResponse struct {
	VirtualMachineListResult
}

// VirtualMachinesClientListResponse contains the response from method VirtualMachinesClient.List.
type VirtualMachinesClientListResponse struct {
	VirtualMachineListResult
}

// VirtualMachinesClientPerformMaintenanceResponse contains the response from method VirtualMachinesClient.PerformMaintenance.
type VirtualMachinesClientPerformMaintenanceResponse struct {
	// placeholder for future response values
}

// VirtualMachinesClientPowerOffResponse contains the response from method VirtualMachinesClient.PowerOff.
type VirtualMachinesClientPowerOffResponse struct {
	// placeholder for future response values
}

// VirtualMachinesClientReapplyResponse contains the response from method VirtualMachinesClient.Reapply.
type VirtualMachinesClientReapplyResponse struct {
	// placeholder for future response values
}

// VirtualMachinesClientRedeployResponse contains the response from method VirtualMachinesClient.Redeploy.
type VirtualMachinesClientRedeployResponse struct {
	// placeholder for future response values
}

// VirtualMachinesClientReimageResponse contains the response from method VirtualMachinesClient.Reimage.
type VirtualMachinesClientReimageResponse struct {
	// placeholder for future response values
}

// VirtualMachinesClientRestartResponse contains the response from method VirtualMachinesClient.Restart.
type VirtualMachinesClientRestartResponse struct {
	// placeholder for future response values
}

// VirtualMachinesClientRunCommandResponse contains the response from method VirtualMachinesClient.RunCommand.
type VirtualMachinesClientRunCommandResponse struct {
	RunCommandResult
}

// VirtualMachinesClientSimulateEvictionResponse contains the response from method VirtualMachinesClient.SimulateEviction.
type VirtualMachinesClientSimulateEvictionResponse struct {
	// placeholder for future response values
}

// VirtualMachinesClientStartResponse contains the response from method VirtualMachinesClient.Start.
type VirtualMachinesClientStartResponse struct {
	// placeholder for future response values
}

// VirtualMachinesClientUpdateResponse contains the response from method VirtualMachinesClient.Update.
type VirtualMachinesClientUpdateResponse struct {
	VirtualMachine
}
