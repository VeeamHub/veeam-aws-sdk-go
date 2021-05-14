# ValidationMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Severity** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**ContextId** | Pointer to **string** |  | [optional] 

## Methods

### NewValidationMessage

`func NewValidationMessage() *ValidationMessage`

NewValidationMessage instantiates a new ValidationMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationMessageWithDefaults

`func NewValidationMessageWithDefaults() *ValidationMessage`

NewValidationMessageWithDefaults instantiates a new ValidationMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeverity

`func (o *ValidationMessage) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ValidationMessage) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ValidationMessage) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ValidationMessage) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetMessage

`func (o *ValidationMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ValidationMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ValidationMessage) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ValidationMessage) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetContextId

`func (o *ValidationMessage) GetContextId() string`

GetContextId returns the ContextId field if non-nil, zero value otherwise.

### GetContextIdOk

`func (o *ValidationMessage) GetContextIdOk() (*string, bool)`

GetContextIdOk returns a tuple with the ContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextId

`func (o *ValidationMessage) SetContextId(v string)`

SetContextId sets ContextId field to given value.

### HasContextId

`func (o *ValidationMessage) HasContextId() bool`

HasContextId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


