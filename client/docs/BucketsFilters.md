# BucketsFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmazonAccountId** | Pointer to **string** |  | [optional] 
**SearchPattern** | Pointer to **string** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Sort** | Pointer to [**[]BucketsSortColumns**](BucketsSortColumns.md) |  | [optional] 

## Methods

### NewBucketsFilters

`func NewBucketsFilters() *BucketsFilters`

NewBucketsFilters instantiates a new BucketsFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketsFiltersWithDefaults

`func NewBucketsFiltersWithDefaults() *BucketsFilters`

NewBucketsFiltersWithDefaults instantiates a new BucketsFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmazonAccountId

`func (o *BucketsFilters) GetAmazonAccountId() string`

GetAmazonAccountId returns the AmazonAccountId field if non-nil, zero value otherwise.

### GetAmazonAccountIdOk

`func (o *BucketsFilters) GetAmazonAccountIdOk() (*string, bool)`

GetAmazonAccountIdOk returns a tuple with the AmazonAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonAccountId

`func (o *BucketsFilters) SetAmazonAccountId(v string)`

SetAmazonAccountId sets AmazonAccountId field to given value.

### HasAmazonAccountId

`func (o *BucketsFilters) HasAmazonAccountId() bool`

HasAmazonAccountId returns a boolean if a field has been set.

### GetSearchPattern

`func (o *BucketsFilters) GetSearchPattern() string`

GetSearchPattern returns the SearchPattern field if non-nil, zero value otherwise.

### GetSearchPatternOk

`func (o *BucketsFilters) GetSearchPatternOk() (*string, bool)`

GetSearchPatternOk returns a tuple with the SearchPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchPattern

`func (o *BucketsFilters) SetSearchPattern(v string)`

SetSearchPattern sets SearchPattern field to given value.

### HasSearchPattern

`func (o *BucketsFilters) HasSearchPattern() bool`

HasSearchPattern returns a boolean if a field has been set.

### GetOffset

`func (o *BucketsFilters) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *BucketsFilters) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *BucketsFilters) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *BucketsFilters) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *BucketsFilters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *BucketsFilters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *BucketsFilters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *BucketsFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetSort

`func (o *BucketsFilters) GetSort() []BucketsSortColumns`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *BucketsFilters) GetSortOk() (*[]BucketsSortColumns, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *BucketsFilters) SetSort(v []BucketsSortColumns)`

SetSort sets Sort field to given value.

### HasSort

`func (o *BucketsFilters) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


