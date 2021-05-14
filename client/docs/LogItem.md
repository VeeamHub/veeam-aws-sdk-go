# LogItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogTime** | **time.Time** |  | 
**Status** | **string** |  | 
**Message** | Pointer to **string** |  | [optional] 
**ExecutionStartTime** | Pointer to **time.Time** |  | [optional] 
**ExecutionDuration** | Pointer to **int64** |  | [optional] 

## Methods

### NewLogItem

`func NewLogItem(logTime time.Time, status string, ) *LogItem`

NewLogItem instantiates a new LogItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogItemWithDefaults

`func NewLogItemWithDefaults() *LogItem`

NewLogItemWithDefaults instantiates a new LogItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogTime

`func (o *LogItem) GetLogTime() time.Time`

GetLogTime returns the LogTime field if non-nil, zero value otherwise.

### GetLogTimeOk

`func (o *LogItem) GetLogTimeOk() (*time.Time, bool)`

GetLogTimeOk returns a tuple with the LogTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogTime

`func (o *LogItem) SetLogTime(v time.Time)`

SetLogTime sets LogTime field to given value.


### GetStatus

`func (o *LogItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LogItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LogItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *LogItem) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LogItem) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LogItem) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LogItem) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetExecutionStartTime

`func (o *LogItem) GetExecutionStartTime() time.Time`

GetExecutionStartTime returns the ExecutionStartTime field if non-nil, zero value otherwise.

### GetExecutionStartTimeOk

`func (o *LogItem) GetExecutionStartTimeOk() (*time.Time, bool)`

GetExecutionStartTimeOk returns a tuple with the ExecutionStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionStartTime

`func (o *LogItem) SetExecutionStartTime(v time.Time)`

SetExecutionStartTime sets ExecutionStartTime field to given value.

### HasExecutionStartTime

`func (o *LogItem) HasExecutionStartTime() bool`

HasExecutionStartTime returns a boolean if a field has been set.

### GetExecutionDuration

`func (o *LogItem) GetExecutionDuration() int64`

GetExecutionDuration returns the ExecutionDuration field if non-nil, zero value otherwise.

### GetExecutionDurationOk

`func (o *LogItem) GetExecutionDurationOk() (*int64, bool)`

GetExecutionDurationOk returns a tuple with the ExecutionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDuration

`func (o *LogItem) SetExecutionDuration(v int64)`

SetExecutionDuration sets ExecutionDuration field to given value.

### HasExecutionDuration

`func (o *LogItem) HasExecutionDuration() bool`

HasExecutionDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


