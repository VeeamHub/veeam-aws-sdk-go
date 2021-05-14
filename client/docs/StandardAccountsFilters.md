# StandardAccountsFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchPattern** | Pointer to **string** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Sort** | Pointer to [**[]StandardAccountsSortColumns**](StandardAccountsSortColumns.md) |  | [optional] 

## Methods

### NewStandardAccountsFilters

`func NewStandardAccountsFilters() *StandardAccountsFilters`

NewStandardAccountsFilters instantiates a new StandardAccountsFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandardAccountsFiltersWithDefaults

`func NewStandardAccountsFiltersWithDefaults() *StandardAccountsFilters`

NewStandardAccountsFiltersWithDefaults instantiates a new StandardAccountsFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchPattern

`func (o *StandardAccountsFilters) GetSearchPattern() string`

GetSearchPattern returns the SearchPattern field if non-nil, zero value otherwise.

### GetSearchPatternOk

`func (o *StandardAccountsFilters) GetSearchPatternOk() (*string, bool)`

GetSearchPatternOk returns a tuple with the SearchPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchPattern

`func (o *StandardAccountsFilters) SetSearchPattern(v string)`

SetSearchPattern sets SearchPattern field to given value.

### HasSearchPattern

`func (o *StandardAccountsFilters) HasSearchPattern() bool`

HasSearchPattern returns a boolean if a field has been set.

### GetOffset

`func (o *StandardAccountsFilters) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *StandardAccountsFilters) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *StandardAccountsFilters) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *StandardAccountsFilters) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *StandardAccountsFilters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *StandardAccountsFilters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *StandardAccountsFilters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *StandardAccountsFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetSort

`func (o *StandardAccountsFilters) GetSort() []StandardAccountsSortColumns`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *StandardAccountsFilters) GetSortOk() (*[]StandardAccountsSortColumns, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *StandardAccountsFilters) SetSort(v []StandardAccountsSortColumns)`

SetSort sets Sort field to given value.

### HasSort

`func (o *StandardAccountsFilters) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


