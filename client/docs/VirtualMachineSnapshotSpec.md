# VirtualMachineSnapshotSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmazonAccountId** | **string** |  | 
**TargetAmazonAccountIdToCopySnapshots** | Pointer to **string** |  | [optional] 
**TargetRegionIdToCopySnapshots** | Pointer to **string** |  | [optional] 
**EncryptionKey** | Pointer to **string** |  | [optional] 
**AdditionalTags** | Pointer to [**[]TagSpec**](TagSpec.md) |  | [optional] 
**CopyTagsFromVolumeEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewVirtualMachineSnapshotSpec

`func NewVirtualMachineSnapshotSpec(amazonAccountId string, ) *VirtualMachineSnapshotSpec`

NewVirtualMachineSnapshotSpec instantiates a new VirtualMachineSnapshotSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualMachineSnapshotSpecWithDefaults

`func NewVirtualMachineSnapshotSpecWithDefaults() *VirtualMachineSnapshotSpec`

NewVirtualMachineSnapshotSpecWithDefaults instantiates a new VirtualMachineSnapshotSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmazonAccountId

`func (o *VirtualMachineSnapshotSpec) GetAmazonAccountId() string`

GetAmazonAccountId returns the AmazonAccountId field if non-nil, zero value otherwise.

### GetAmazonAccountIdOk

`func (o *VirtualMachineSnapshotSpec) GetAmazonAccountIdOk() (*string, bool)`

GetAmazonAccountIdOk returns a tuple with the AmazonAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonAccountId

`func (o *VirtualMachineSnapshotSpec) SetAmazonAccountId(v string)`

SetAmazonAccountId sets AmazonAccountId field to given value.


### GetTargetAmazonAccountIdToCopySnapshots

`func (o *VirtualMachineSnapshotSpec) GetTargetAmazonAccountIdToCopySnapshots() string`

GetTargetAmazonAccountIdToCopySnapshots returns the TargetAmazonAccountIdToCopySnapshots field if non-nil, zero value otherwise.

### GetTargetAmazonAccountIdToCopySnapshotsOk

`func (o *VirtualMachineSnapshotSpec) GetTargetAmazonAccountIdToCopySnapshotsOk() (*string, bool)`

GetTargetAmazonAccountIdToCopySnapshotsOk returns a tuple with the TargetAmazonAccountIdToCopySnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAmazonAccountIdToCopySnapshots

`func (o *VirtualMachineSnapshotSpec) SetTargetAmazonAccountIdToCopySnapshots(v string)`

SetTargetAmazonAccountIdToCopySnapshots sets TargetAmazonAccountIdToCopySnapshots field to given value.

### HasTargetAmazonAccountIdToCopySnapshots

`func (o *VirtualMachineSnapshotSpec) HasTargetAmazonAccountIdToCopySnapshots() bool`

HasTargetAmazonAccountIdToCopySnapshots returns a boolean if a field has been set.

### GetTargetRegionIdToCopySnapshots

`func (o *VirtualMachineSnapshotSpec) GetTargetRegionIdToCopySnapshots() string`

GetTargetRegionIdToCopySnapshots returns the TargetRegionIdToCopySnapshots field if non-nil, zero value otherwise.

### GetTargetRegionIdToCopySnapshotsOk

`func (o *VirtualMachineSnapshotSpec) GetTargetRegionIdToCopySnapshotsOk() (*string, bool)`

GetTargetRegionIdToCopySnapshotsOk returns a tuple with the TargetRegionIdToCopySnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRegionIdToCopySnapshots

`func (o *VirtualMachineSnapshotSpec) SetTargetRegionIdToCopySnapshots(v string)`

SetTargetRegionIdToCopySnapshots sets TargetRegionIdToCopySnapshots field to given value.

### HasTargetRegionIdToCopySnapshots

`func (o *VirtualMachineSnapshotSpec) HasTargetRegionIdToCopySnapshots() bool`

HasTargetRegionIdToCopySnapshots returns a boolean if a field has been set.

### GetEncryptionKey

`func (o *VirtualMachineSnapshotSpec) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *VirtualMachineSnapshotSpec) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *VirtualMachineSnapshotSpec) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *VirtualMachineSnapshotSpec) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### GetAdditionalTags

`func (o *VirtualMachineSnapshotSpec) GetAdditionalTags() []TagSpec`

GetAdditionalTags returns the AdditionalTags field if non-nil, zero value otherwise.

### GetAdditionalTagsOk

`func (o *VirtualMachineSnapshotSpec) GetAdditionalTagsOk() (*[]TagSpec, bool)`

GetAdditionalTagsOk returns a tuple with the AdditionalTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalTags

`func (o *VirtualMachineSnapshotSpec) SetAdditionalTags(v []TagSpec)`

SetAdditionalTags sets AdditionalTags field to given value.

### HasAdditionalTags

`func (o *VirtualMachineSnapshotSpec) HasAdditionalTags() bool`

HasAdditionalTags returns a boolean if a field has been set.

### GetCopyTagsFromVolumeEnabled

`func (o *VirtualMachineSnapshotSpec) GetCopyTagsFromVolumeEnabled() bool`

GetCopyTagsFromVolumeEnabled returns the CopyTagsFromVolumeEnabled field if non-nil, zero value otherwise.

### GetCopyTagsFromVolumeEnabledOk

`func (o *VirtualMachineSnapshotSpec) GetCopyTagsFromVolumeEnabledOk() (*bool, bool)`

GetCopyTagsFromVolumeEnabledOk returns a tuple with the CopyTagsFromVolumeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyTagsFromVolumeEnabled

`func (o *VirtualMachineSnapshotSpec) SetCopyTagsFromVolumeEnabled(v bool)`

SetCopyTagsFromVolumeEnabled sets CopyTagsFromVolumeEnabled field to given value.

### HasCopyTagsFromVolumeEnabled

`func (o *VirtualMachineSnapshotSpec) HasCopyTagsFromVolumeEnabled() bool`

HasCopyTagsFromVolumeEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


