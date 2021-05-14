# AmazonInstanceTypesByRegionPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**Results** | [**[]AmazonInstanceType**](AmazonInstanceType.md) |  | 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 

## Methods

### NewAmazonInstanceTypesByRegionPage

`func NewAmazonInstanceTypesByRegionPage(totalCount int32, results []AmazonInstanceType, ) *AmazonInstanceTypesByRegionPage`

NewAmazonInstanceTypesByRegionPage instantiates a new AmazonInstanceTypesByRegionPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonInstanceTypesByRegionPageWithDefaults

`func NewAmazonInstanceTypesByRegionPageWithDefaults() *AmazonInstanceTypesByRegionPage`

NewAmazonInstanceTypesByRegionPageWithDefaults instantiates a new AmazonInstanceTypesByRegionPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *AmazonInstanceTypesByRegionPage) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AmazonInstanceTypesByRegionPage) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AmazonInstanceTypesByRegionPage) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetResults

`func (o *AmazonInstanceTypesByRegionPage) GetResults() []AmazonInstanceType`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AmazonInstanceTypesByRegionPage) GetResultsOk() (*[]AmazonInstanceType, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AmazonInstanceTypesByRegionPage) SetResults(v []AmazonInstanceType)`

SetResults sets Results field to given value.


### GetLinks

`func (o *AmazonInstanceTypesByRegionPage) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AmazonInstanceTypesByRegionPage) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AmazonInstanceTypesByRegionPage) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AmazonInstanceTypesByRegionPage) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


