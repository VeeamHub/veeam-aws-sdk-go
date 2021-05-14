# ProtectedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Status** | Pointer to [**SessionItemsStatuses**](SessionItemsStatuses.md) |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Embedded** | Pointer to [**ProtectedItemEmbedded**](ProtectedItemEmbedded.md) |  | [optional] 

## Methods

### NewProtectedItem

`func NewProtectedItem(id string, name string, ) *ProtectedItem`

NewProtectedItem instantiates a new ProtectedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedItemWithDefaults

`func NewProtectedItemWithDefaults() *ProtectedItem`

NewProtectedItemWithDefaults instantiates a new ProtectedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProtectedItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProtectedItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProtectedItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ProtectedItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProtectedItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProtectedItem) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *ProtectedItem) GetStatus() SessionItemsStatuses`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProtectedItem) GetStatusOk() (*SessionItemsStatuses, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProtectedItem) SetStatus(v SessionItemsStatuses)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProtectedItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRegion

`func (o *ProtectedItem) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ProtectedItem) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ProtectedItem) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ProtectedItem) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetEmbedded

`func (o *ProtectedItem) GetEmbedded() ProtectedItemEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ProtectedItem) GetEmbeddedOk() (*ProtectedItemEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ProtectedItem) SetEmbedded(v ProtectedItemEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *ProtectedItem) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


