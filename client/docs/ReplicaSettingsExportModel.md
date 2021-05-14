# ReplicaSettingsExportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mapping** | [**[]RemoteMappingExportModel**](RemoteMappingExportModel.md) |  | 
**AdditionalTags** | Pointer to [**[]TagSpec**](TagSpec.md) |  | [optional] 
**CopyTagsFromVolumeEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewReplicaSettingsExportModel

`func NewReplicaSettingsExportModel(mapping []RemoteMappingExportModel, ) *ReplicaSettingsExportModel`

NewReplicaSettingsExportModel instantiates a new ReplicaSettingsExportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicaSettingsExportModelWithDefaults

`func NewReplicaSettingsExportModelWithDefaults() *ReplicaSettingsExportModel`

NewReplicaSettingsExportModelWithDefaults instantiates a new ReplicaSettingsExportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMapping

`func (o *ReplicaSettingsExportModel) GetMapping() []RemoteMappingExportModel`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *ReplicaSettingsExportModel) GetMappingOk() (*[]RemoteMappingExportModel, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *ReplicaSettingsExportModel) SetMapping(v []RemoteMappingExportModel)`

SetMapping sets Mapping field to given value.


### GetAdditionalTags

`func (o *ReplicaSettingsExportModel) GetAdditionalTags() []TagSpec`

GetAdditionalTags returns the AdditionalTags field if non-nil, zero value otherwise.

### GetAdditionalTagsOk

`func (o *ReplicaSettingsExportModel) GetAdditionalTagsOk() (*[]TagSpec, bool)`

GetAdditionalTagsOk returns a tuple with the AdditionalTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalTags

`func (o *ReplicaSettingsExportModel) SetAdditionalTags(v []TagSpec)`

SetAdditionalTags sets AdditionalTags field to given value.

### HasAdditionalTags

`func (o *ReplicaSettingsExportModel) HasAdditionalTags() bool`

HasAdditionalTags returns a boolean if a field has been set.

### GetCopyTagsFromVolumeEnabled

`func (o *ReplicaSettingsExportModel) GetCopyTagsFromVolumeEnabled() bool`

GetCopyTagsFromVolumeEnabled returns the CopyTagsFromVolumeEnabled field if non-nil, zero value otherwise.

### GetCopyTagsFromVolumeEnabledOk

`func (o *ReplicaSettingsExportModel) GetCopyTagsFromVolumeEnabledOk() (*bool, bool)`

GetCopyTagsFromVolumeEnabledOk returns a tuple with the CopyTagsFromVolumeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyTagsFromVolumeEnabled

`func (o *ReplicaSettingsExportModel) SetCopyTagsFromVolumeEnabled(v bool)`

SetCopyTagsFromVolumeEnabled sets CopyTagsFromVolumeEnabled field to given value.

### HasCopyTagsFromVolumeEnabled

`func (o *ReplicaSettingsExportModel) HasCopyTagsFromVolumeEnabled() bool`

HasCopyTagsFromVolumeEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


