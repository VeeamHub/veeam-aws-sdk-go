# PolicyReplicaSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mapping** | [**[]ReplicaMapping**](ReplicaMapping.md) |  | 
**AdditionalTags** | Pointer to [**[]TagSpec**](TagSpec.md) |  | [optional] 
**CopyTagsFromVolumeEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewPolicyReplicaSettings

`func NewPolicyReplicaSettings(mapping []ReplicaMapping, ) *PolicyReplicaSettings`

NewPolicyReplicaSettings instantiates a new PolicyReplicaSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyReplicaSettingsWithDefaults

`func NewPolicyReplicaSettingsWithDefaults() *PolicyReplicaSettings`

NewPolicyReplicaSettingsWithDefaults instantiates a new PolicyReplicaSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMapping

`func (o *PolicyReplicaSettings) GetMapping() []ReplicaMapping`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *PolicyReplicaSettings) GetMappingOk() (*[]ReplicaMapping, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *PolicyReplicaSettings) SetMapping(v []ReplicaMapping)`

SetMapping sets Mapping field to given value.


### GetAdditionalTags

`func (o *PolicyReplicaSettings) GetAdditionalTags() []TagSpec`

GetAdditionalTags returns the AdditionalTags field if non-nil, zero value otherwise.

### GetAdditionalTagsOk

`func (o *PolicyReplicaSettings) GetAdditionalTagsOk() (*[]TagSpec, bool)`

GetAdditionalTagsOk returns a tuple with the AdditionalTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalTags

`func (o *PolicyReplicaSettings) SetAdditionalTags(v []TagSpec)`

SetAdditionalTags sets AdditionalTags field to given value.

### HasAdditionalTags

`func (o *PolicyReplicaSettings) HasAdditionalTags() bool`

HasAdditionalTags returns a boolean if a field has been set.

### GetCopyTagsFromVolumeEnabled

`func (o *PolicyReplicaSettings) GetCopyTagsFromVolumeEnabled() bool`

GetCopyTagsFromVolumeEnabled returns the CopyTagsFromVolumeEnabled field if non-nil, zero value otherwise.

### GetCopyTagsFromVolumeEnabledOk

`func (o *PolicyReplicaSettings) GetCopyTagsFromVolumeEnabledOk() (*bool, bool)`

GetCopyTagsFromVolumeEnabledOk returns a tuple with the CopyTagsFromVolumeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyTagsFromVolumeEnabled

`func (o *PolicyReplicaSettings) SetCopyTagsFromVolumeEnabled(v bool)`

SetCopyTagsFromVolumeEnabled sets CopyTagsFromVolumeEnabled field to given value.

### HasCopyTagsFromVolumeEnabled

`func (o *PolicyReplicaSettings) HasCopyTagsFromVolumeEnabled() bool`

HasCopyTagsFromVolumeEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


