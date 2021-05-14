# BucketFoldersFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Sort** | Pointer to [**[]BucketFoldersFiltersOrderColumns**](BucketFoldersFiltersOrderColumns.md) |  | [optional] 

## Methods

### NewBucketFoldersFilters

`func NewBucketFoldersFilters() *BucketFoldersFilters`

NewBucketFoldersFilters instantiates a new BucketFoldersFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketFoldersFiltersWithDefaults

`func NewBucketFoldersFiltersWithDefaults() *BucketFoldersFilters`

NewBucketFoldersFiltersWithDefaults instantiates a new BucketFoldersFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *BucketFoldersFilters) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *BucketFoldersFilters) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *BucketFoldersFilters) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *BucketFoldersFilters) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *BucketFoldersFilters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *BucketFoldersFilters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *BucketFoldersFilters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *BucketFoldersFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetSort

`func (o *BucketFoldersFilters) GetSort() []BucketFoldersFiltersOrderColumns`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *BucketFoldersFilters) GetSortOk() (*[]BucketFoldersFiltersOrderColumns, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *BucketFoldersFilters) SetSort(v []BucketFoldersFiltersOrderColumns)`

SetSort sets Sort field to given value.

### HasSort

`func (o *BucketFoldersFilters) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


