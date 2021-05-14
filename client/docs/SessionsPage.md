# SessionsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**Results** | [**[]Session**](Session.md) |  | 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 

## Methods

### NewSessionsPage

`func NewSessionsPage(totalCount int32, results []Session, ) *SessionsPage`

NewSessionsPage instantiates a new SessionsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionsPageWithDefaults

`func NewSessionsPageWithDefaults() *SessionsPage`

NewSessionsPageWithDefaults instantiates a new SessionsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *SessionsPage) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SessionsPage) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SessionsPage) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetResults

`func (o *SessionsPage) GetResults() []Session`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SessionsPage) GetResultsOk() (*[]Session, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SessionsPage) SetResults(v []Session)`

SetResults sets Results field to given value.


### GetLinks

`func (o *SessionsPage) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SessionsPage) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SessionsPage) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SessionsPage) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


