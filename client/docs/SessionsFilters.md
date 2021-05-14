# SessionsFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**PolicyId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**[]SessionStatuses**](SessionStatuses.md) |  | [optional] 
**Types** | Pointer to [**[]SessionTypes**](SessionTypes.md) |  | [optional] 
**FromUtc** | Pointer to **time.Time** |  | [optional] 
**Usn** | Pointer to **int64** |  | [optional] 
**Sort** | Pointer to [**[]SessionsSortColumns**](SessionsSortColumns.md) |  | [optional] 

## Methods

### NewSessionsFilters

`func NewSessionsFilters() *SessionsFilters`

NewSessionsFilters instantiates a new SessionsFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionsFiltersWithDefaults

`func NewSessionsFiltersWithDefaults() *SessionsFilters`

NewSessionsFiltersWithDefaults instantiates a new SessionsFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *SessionsFilters) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SessionsFilters) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SessionsFilters) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *SessionsFilters) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *SessionsFilters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SessionsFilters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SessionsFilters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SessionsFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetPolicyId

`func (o *SessionsFilters) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *SessionsFilters) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *SessionsFilters) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *SessionsFilters) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetStatus

`func (o *SessionsFilters) GetStatus() []SessionStatuses`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SessionsFilters) GetStatusOk() (*[]SessionStatuses, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SessionsFilters) SetStatus(v []SessionStatuses)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SessionsFilters) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTypes

`func (o *SessionsFilters) GetTypes() []SessionTypes`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *SessionsFilters) GetTypesOk() (*[]SessionTypes, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *SessionsFilters) SetTypes(v []SessionTypes)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *SessionsFilters) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetFromUtc

`func (o *SessionsFilters) GetFromUtc() time.Time`

GetFromUtc returns the FromUtc field if non-nil, zero value otherwise.

### GetFromUtcOk

`func (o *SessionsFilters) GetFromUtcOk() (*time.Time, bool)`

GetFromUtcOk returns a tuple with the FromUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUtc

`func (o *SessionsFilters) SetFromUtc(v time.Time)`

SetFromUtc sets FromUtc field to given value.

### HasFromUtc

`func (o *SessionsFilters) HasFromUtc() bool`

HasFromUtc returns a boolean if a field has been set.

### GetUsn

`func (o *SessionsFilters) GetUsn() int64`

GetUsn returns the Usn field if non-nil, zero value otherwise.

### GetUsnOk

`func (o *SessionsFilters) GetUsnOk() (*int64, bool)`

GetUsnOk returns a tuple with the Usn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsn

`func (o *SessionsFilters) SetUsn(v int64)`

SetUsn sets Usn field to given value.

### HasUsn

`func (o *SessionsFilters) HasUsn() bool`

HasUsn returns a boolean if a field has been set.

### GetSort

`func (o *SessionsFilters) GetSort() []SessionsSortColumns`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *SessionsFilters) GetSortOk() (*[]SessionsSortColumns, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *SessionsFilters) SetSort(v []SessionsSortColumns)`

SetSort sets Sort field to given value.

### HasSort

`func (o *SessionsFilters) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


