# RetentionJobDeletedRestorePointsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**Results** | [**[]RetentionJobDeletedRestorePoint**](RetentionJobDeletedRestorePoint.md) |  | 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 

## Methods

### NewRetentionJobDeletedRestorePointsPage

`func NewRetentionJobDeletedRestorePointsPage(totalCount int32, results []RetentionJobDeletedRestorePoint, ) *RetentionJobDeletedRestorePointsPage`

NewRetentionJobDeletedRestorePointsPage instantiates a new RetentionJobDeletedRestorePointsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetentionJobDeletedRestorePointsPageWithDefaults

`func NewRetentionJobDeletedRestorePointsPageWithDefaults() *RetentionJobDeletedRestorePointsPage`

NewRetentionJobDeletedRestorePointsPageWithDefaults instantiates a new RetentionJobDeletedRestorePointsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *RetentionJobDeletedRestorePointsPage) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *RetentionJobDeletedRestorePointsPage) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *RetentionJobDeletedRestorePointsPage) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetResults

`func (o *RetentionJobDeletedRestorePointsPage) GetResults() []RetentionJobDeletedRestorePoint`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *RetentionJobDeletedRestorePointsPage) GetResultsOk() (*[]RetentionJobDeletedRestorePoint, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *RetentionJobDeletedRestorePointsPage) SetResults(v []RetentionJobDeletedRestorePoint)`

SetResults sets Results field to given value.


### GetLinks

`func (o *RetentionJobDeletedRestorePointsPage) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RetentionJobDeletedRestorePointsPage) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RetentionJobDeletedRestorePointsPage) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RetentionJobDeletedRestorePointsPage) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


