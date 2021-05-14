# Session

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**SessionTypes**](SessionTypes.md) |  | 
**ExecutionStartTime** | Pointer to **time.Time** |  | [optional] 
**ExecutionStopTime** | Pointer to **time.Time** |  | [optional] 
**ExecutionDuration** | Pointer to **int64** |  | [optional] 
**Status** | [**SessionStatuses**](SessionStatuses.md) |  | 
**Result** | Pointer to [**SessionResults**](SessionResults.md) |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Usn** | **int64** |  | 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 
**Embedded** | Pointer to [**SessionEmbeddedResources**](SessionEmbeddedResources.md) |  | [optional] 

## Methods

### NewSession

`func NewSession(id string, type_ SessionTypes, status SessionStatuses, usn int64, ) *Session`

NewSession instantiates a new Session object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionWithDefaults

`func NewSessionWithDefaults() *Session`

NewSessionWithDefaults instantiates a new Session object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Session) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Session) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Session) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *Session) GetType() SessionTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Session) GetTypeOk() (*SessionTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Session) SetType(v SessionTypes)`

SetType sets Type field to given value.


### GetExecutionStartTime

`func (o *Session) GetExecutionStartTime() time.Time`

GetExecutionStartTime returns the ExecutionStartTime field if non-nil, zero value otherwise.

### GetExecutionStartTimeOk

`func (o *Session) GetExecutionStartTimeOk() (*time.Time, bool)`

GetExecutionStartTimeOk returns a tuple with the ExecutionStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionStartTime

`func (o *Session) SetExecutionStartTime(v time.Time)`

SetExecutionStartTime sets ExecutionStartTime field to given value.

### HasExecutionStartTime

`func (o *Session) HasExecutionStartTime() bool`

HasExecutionStartTime returns a boolean if a field has been set.

### GetExecutionStopTime

`func (o *Session) GetExecutionStopTime() time.Time`

GetExecutionStopTime returns the ExecutionStopTime field if non-nil, zero value otherwise.

### GetExecutionStopTimeOk

`func (o *Session) GetExecutionStopTimeOk() (*time.Time, bool)`

GetExecutionStopTimeOk returns a tuple with the ExecutionStopTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionStopTime

`func (o *Session) SetExecutionStopTime(v time.Time)`

SetExecutionStopTime sets ExecutionStopTime field to given value.

### HasExecutionStopTime

`func (o *Session) HasExecutionStopTime() bool`

HasExecutionStopTime returns a boolean if a field has been set.

### GetExecutionDuration

`func (o *Session) GetExecutionDuration() int64`

GetExecutionDuration returns the ExecutionDuration field if non-nil, zero value otherwise.

### GetExecutionDurationOk

`func (o *Session) GetExecutionDurationOk() (*int64, bool)`

GetExecutionDurationOk returns a tuple with the ExecutionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDuration

`func (o *Session) SetExecutionDuration(v int64)`

SetExecutionDuration sets ExecutionDuration field to given value.

### HasExecutionDuration

`func (o *Session) HasExecutionDuration() bool`

HasExecutionDuration returns a boolean if a field has been set.

### GetStatus

`func (o *Session) GetStatus() SessionStatuses`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Session) GetStatusOk() (*SessionStatuses, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Session) SetStatus(v SessionStatuses)`

SetStatus sets Status field to given value.


### GetResult

`func (o *Session) GetResult() SessionResults`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Session) GetResultOk() (*SessionResults, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Session) SetResult(v SessionResults)`

SetResult sets Result field to given value.

### HasResult

`func (o *Session) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetReason

`func (o *Session) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Session) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Session) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Session) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetUsn

`func (o *Session) GetUsn() int64`

GetUsn returns the Usn field if non-nil, zero value otherwise.

### GetUsnOk

`func (o *Session) GetUsnOk() (*int64, bool)`

GetUsnOk returns a tuple with the Usn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsn

`func (o *Session) SetUsn(v int64)`

SetUsn sets Usn field to given value.


### GetLinks

`func (o *Session) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Session) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Session) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Session) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEmbedded

`func (o *Session) GetEmbedded() SessionEmbeddedResources`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *Session) GetEmbeddedOk() (*SessionEmbeddedResources, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *Session) SetEmbedded(v SessionEmbeddedResources)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *Session) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


