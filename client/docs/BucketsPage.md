# BucketsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]Bucket**](Bucket.md) |  | 
**TotalCount** | **int32** |  | 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 

## Methods

### NewBucketsPage

`func NewBucketsPage(results []Bucket, totalCount int32, ) *BucketsPage`

NewBucketsPage instantiates a new BucketsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketsPageWithDefaults

`func NewBucketsPageWithDefaults() *BucketsPage`

NewBucketsPageWithDefaults instantiates a new BucketsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *BucketsPage) GetResults() []Bucket`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BucketsPage) GetResultsOk() (*[]Bucket, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BucketsPage) SetResults(v []Bucket)`

SetResults sets Results field to given value.


### GetTotalCount

`func (o *BucketsPage) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *BucketsPage) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *BucketsPage) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetLinks

`func (o *BucketsPage) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BucketsPage) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BucketsPage) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BucketsPage) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


