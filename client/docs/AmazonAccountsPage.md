# AmazonAccountsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]AmazonAccount**](AmazonAccount.md) |  | 
**TotalCount** | **int32** |  | 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 

## Methods

### NewAmazonAccountsPage

`func NewAmazonAccountsPage(results []AmazonAccount, totalCount int32, ) *AmazonAccountsPage`

NewAmazonAccountsPage instantiates a new AmazonAccountsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonAccountsPageWithDefaults

`func NewAmazonAccountsPageWithDefaults() *AmazonAccountsPage`

NewAmazonAccountsPageWithDefaults instantiates a new AmazonAccountsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *AmazonAccountsPage) GetResults() []AmazonAccount`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AmazonAccountsPage) GetResultsOk() (*[]AmazonAccount, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AmazonAccountsPage) SetResults(v []AmazonAccount)`

SetResults sets Results field to given value.


### GetTotalCount

`func (o *AmazonAccountsPage) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AmazonAccountsPage) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AmazonAccountsPage) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetLinks

`func (o *AmazonAccountsPage) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AmazonAccountsPage) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AmazonAccountsPage) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AmazonAccountsPage) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


