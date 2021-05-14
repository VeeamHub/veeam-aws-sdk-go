# SessionLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | **string** |  | 
**Log** | Pointer to [**[]LogItem**](LogItem.md) |  | [optional] 

## Methods

### NewSessionLog

`func NewSessionLog(sessionId string, ) *SessionLog`

NewSessionLog instantiates a new SessionLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionLogWithDefaults

`func NewSessionLogWithDefaults() *SessionLog`

NewSessionLogWithDefaults instantiates a new SessionLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *SessionLog) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *SessionLog) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *SessionLog) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetLog

`func (o *SessionLog) GetLog() []LogItem`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *SessionLog) GetLogOk() (*[]LogItem, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *SessionLog) SetLog(v []LogItem)`

SetLog sets Log field to given value.

### HasLog

`func (o *SessionLog) HasLog() bool`

HasLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


