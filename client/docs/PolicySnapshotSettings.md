# PolicySnapshotSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalTags** | Pointer to [**[]TagSpec**](TagSpec.md) |  | [optional] 
**CopyTagsFromVolumeEnabled** | Pointer to **bool** |  | [optional] 
**TryCreateVSSSnapshot** | Pointer to **bool** |  | [optional] 

## Methods

### NewPolicySnapshotSettings

`func NewPolicySnapshotSettings() *PolicySnapshotSettings`

NewPolicySnapshotSettings instantiates a new PolicySnapshotSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicySnapshotSettingsWithDefaults

`func NewPolicySnapshotSettingsWithDefaults() *PolicySnapshotSettings`

NewPolicySnapshotSettingsWithDefaults instantiates a new PolicySnapshotSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalTags

`func (o *PolicySnapshotSettings) GetAdditionalTags() []TagSpec`

GetAdditionalTags returns the AdditionalTags field if non-nil, zero value otherwise.

### GetAdditionalTagsOk

`func (o *PolicySnapshotSettings) GetAdditionalTagsOk() (*[]TagSpec, bool)`

GetAdditionalTagsOk returns a tuple with the AdditionalTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalTags

`func (o *PolicySnapshotSettings) SetAdditionalTags(v []TagSpec)`

SetAdditionalTags sets AdditionalTags field to given value.

### HasAdditionalTags

`func (o *PolicySnapshotSettings) HasAdditionalTags() bool`

HasAdditionalTags returns a boolean if a field has been set.

### GetCopyTagsFromVolumeEnabled

`func (o *PolicySnapshotSettings) GetCopyTagsFromVolumeEnabled() bool`

GetCopyTagsFromVolumeEnabled returns the CopyTagsFromVolumeEnabled field if non-nil, zero value otherwise.

### GetCopyTagsFromVolumeEnabledOk

`func (o *PolicySnapshotSettings) GetCopyTagsFromVolumeEnabledOk() (*bool, bool)`

GetCopyTagsFromVolumeEnabledOk returns a tuple with the CopyTagsFromVolumeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyTagsFromVolumeEnabled

`func (o *PolicySnapshotSettings) SetCopyTagsFromVolumeEnabled(v bool)`

SetCopyTagsFromVolumeEnabled sets CopyTagsFromVolumeEnabled field to given value.

### HasCopyTagsFromVolumeEnabled

`func (o *PolicySnapshotSettings) HasCopyTagsFromVolumeEnabled() bool`

HasCopyTagsFromVolumeEnabled returns a boolean if a field has been set.

### GetTryCreateVSSSnapshot

`func (o *PolicySnapshotSettings) GetTryCreateVSSSnapshot() bool`

GetTryCreateVSSSnapshot returns the TryCreateVSSSnapshot field if non-nil, zero value otherwise.

### GetTryCreateVSSSnapshotOk

`func (o *PolicySnapshotSettings) GetTryCreateVSSSnapshotOk() (*bool, bool)`

GetTryCreateVSSSnapshotOk returns a tuple with the TryCreateVSSSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTryCreateVSSSnapshot

`func (o *PolicySnapshotSettings) SetTryCreateVSSSnapshot(v bool)`

SetTryCreateVSSSnapshot sets TryCreateVSSSnapshot field to given value.

### HasTryCreateVSSSnapshot

`func (o *PolicySnapshotSettings) HasTryCreateVSSSnapshot() bool`

HasTryCreateVSSSnapshot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


