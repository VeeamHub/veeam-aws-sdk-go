# RestoredItemsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**Results** | [**[]RestoredItem**](RestoredItem.md) |  | 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 

## Methods

### NewRestoredItemsPage

`func NewRestoredItemsPage(totalCount int32, results []RestoredItem, ) *RestoredItemsPage`

NewRestoredItemsPage instantiates a new RestoredItemsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoredItemsPageWithDefaults

`func NewRestoredItemsPageWithDefaults() *RestoredItemsPage`

NewRestoredItemsPageWithDefaults instantiates a new RestoredItemsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *RestoredItemsPage) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *RestoredItemsPage) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *RestoredItemsPage) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetResults

`func (o *RestoredItemsPage) GetResults() []RestoredItem`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *RestoredItemsPage) GetResultsOk() (*[]RestoredItem, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *RestoredItemsPage) SetResults(v []RestoredItem)`

SetResults sets Results field to given value.


### GetLinks

`func (o *RestoredItemsPage) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RestoredItemsPage) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RestoredItemsPage) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RestoredItemsPage) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


