# StandardAccountsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]StandardAccount**](StandardAccount.md) |  | [optional] 
**TotalCount** | **int32** |  | 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 

## Methods

### NewStandardAccountsPage

`func NewStandardAccountsPage(totalCount int32, ) *StandardAccountsPage`

NewStandardAccountsPage instantiates a new StandardAccountsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandardAccountsPageWithDefaults

`func NewStandardAccountsPageWithDefaults() *StandardAccountsPage`

NewStandardAccountsPageWithDefaults instantiates a new StandardAccountsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *StandardAccountsPage) GetResults() []StandardAccount`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *StandardAccountsPage) GetResultsOk() (*[]StandardAccount, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *StandardAccountsPage) SetResults(v []StandardAccount)`

SetResults sets Results field to given value.

### HasResults

`func (o *StandardAccountsPage) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotalCount

`func (o *StandardAccountsPage) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *StandardAccountsPage) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *StandardAccountsPage) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetLinks

`func (o *StandardAccountsPage) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StandardAccountsPage) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StandardAccountsPage) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StandardAccountsPage) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


