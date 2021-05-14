# RestoredItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**SessionItemsStatuses**](SessionItemsStatuses.md) |  | [optional] 
**RegionName** | Pointer to **string** |  | [optional] 
**Embedded** | Pointer to [**RestoreItemEmbedded**](RestoreItemEmbedded.md) |  | [optional] 

## Methods

### NewRestoredItem

`func NewRestoredItem() *RestoredItem`

NewRestoredItem instantiates a new RestoredItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoredItemWithDefaults

`func NewRestoredItemWithDefaults() *RestoredItem`

NewRestoredItemWithDefaults instantiates a new RestoredItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RestoredItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestoredItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestoredItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RestoredItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *RestoredItem) GetStatus() SessionItemsStatuses`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RestoredItem) GetStatusOk() (*SessionItemsStatuses, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RestoredItem) SetStatus(v SessionItemsStatuses)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RestoredItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRegionName

`func (o *RestoredItem) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *RestoredItem) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *RestoredItem) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *RestoredItem) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetEmbedded

`func (o *RestoredItem) GetEmbedded() RestoreItemEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *RestoredItem) GetEmbeddedOk() (*RestoreItemEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *RestoredItem) SetEmbedded(v RestoreItemEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *RestoredItem) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


