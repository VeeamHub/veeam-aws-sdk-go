# EmailNotificationsTestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewEmailNotificationsTestResult

`func NewEmailNotificationsTestResult(success bool, ) *EmailNotificationsTestResult`

NewEmailNotificationsTestResult instantiates a new EmailNotificationsTestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailNotificationsTestResultWithDefaults

`func NewEmailNotificationsTestResultWithDefaults() *EmailNotificationsTestResult`

NewEmailNotificationsTestResultWithDefaults instantiates a new EmailNotificationsTestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *EmailNotificationsTestResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *EmailNotificationsTestResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *EmailNotificationsTestResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetMessage

`func (o *EmailNotificationsTestResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EmailNotificationsTestResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EmailNotificationsTestResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EmailNotificationsTestResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


