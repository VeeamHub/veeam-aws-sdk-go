# RepositoriesFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchPattern** | Pointer to **string** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Sort** | Pointer to [**[]RepositoriesSortColumns**](RepositoriesSortColumns.md) |  | [optional] 

## Methods

### NewRepositoriesFilters

`func NewRepositoriesFilters() *RepositoriesFilters`

NewRepositoriesFilters instantiates a new RepositoriesFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoriesFiltersWithDefaults

`func NewRepositoriesFiltersWithDefaults() *RepositoriesFilters`

NewRepositoriesFiltersWithDefaults instantiates a new RepositoriesFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchPattern

`func (o *RepositoriesFilters) GetSearchPattern() string`

GetSearchPattern returns the SearchPattern field if non-nil, zero value otherwise.

### GetSearchPatternOk

`func (o *RepositoriesFilters) GetSearchPatternOk() (*string, bool)`

GetSearchPatternOk returns a tuple with the SearchPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchPattern

`func (o *RepositoriesFilters) SetSearchPattern(v string)`

SetSearchPattern sets SearchPattern field to given value.

### HasSearchPattern

`func (o *RepositoriesFilters) HasSearchPattern() bool`

HasSearchPattern returns a boolean if a field has been set.

### GetOffset

`func (o *RepositoriesFilters) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *RepositoriesFilters) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *RepositoriesFilters) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *RepositoriesFilters) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *RepositoriesFilters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RepositoriesFilters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RepositoriesFilters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *RepositoriesFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetSort

`func (o *RepositoriesFilters) GetSort() []RepositoriesSortColumns`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *RepositoriesFilters) GetSortOk() (*[]RepositoriesSortColumns, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *RepositoriesFilters) SetSort(v []RepositoriesSortColumns)`

SetSort sets Sort field to given value.

### HasSort

`func (o *RepositoriesFilters) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


