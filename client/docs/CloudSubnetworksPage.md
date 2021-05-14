# CloudSubnetworksPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**Results** | [**[]CloudSubnetwork**](CloudSubnetwork.md) |  | 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 

## Methods

### NewCloudSubnetworksPage

`func NewCloudSubnetworksPage(totalCount int32, results []CloudSubnetwork, ) *CloudSubnetworksPage`

NewCloudSubnetworksPage instantiates a new CloudSubnetworksPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSubnetworksPageWithDefaults

`func NewCloudSubnetworksPageWithDefaults() *CloudSubnetworksPage`

NewCloudSubnetworksPageWithDefaults instantiates a new CloudSubnetworksPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *CloudSubnetworksPage) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CloudSubnetworksPage) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CloudSubnetworksPage) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetResults

`func (o *CloudSubnetworksPage) GetResults() []CloudSubnetwork`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CloudSubnetworksPage) GetResultsOk() (*[]CloudSubnetwork, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CloudSubnetworksPage) SetResults(v []CloudSubnetwork)`

SetResults sets Results field to given value.


### GetLinks

`func (o *CloudSubnetworksPage) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CloudSubnetworksPage) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CloudSubnetworksPage) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CloudSubnetworksPage) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


