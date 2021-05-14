/*
 * Veeam Backup for AWS public API 1.0
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0-rev0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ValidationMessage struct for ValidationMessage
type ValidationMessage struct {
	Severity *string `json:"severity,omitempty"`
	Message *string `json:"message,omitempty"`
	ContextId *string `json:"contextId,omitempty"`
}

// NewValidationMessage instantiates a new ValidationMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidationMessage() *ValidationMessage {
	this := ValidationMessage{}
	return &this
}

// NewValidationMessageWithDefaults instantiates a new ValidationMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationMessageWithDefaults() *ValidationMessage {
	this := ValidationMessage{}
	return &this
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *ValidationMessage) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationMessage) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *ValidationMessage) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *ValidationMessage) SetSeverity(v string) {
	o.Severity = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ValidationMessage) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationMessage) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ValidationMessage) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ValidationMessage) SetMessage(v string) {
	o.Message = &v
}

// GetContextId returns the ContextId field value if set, zero value otherwise.
func (o *ValidationMessage) GetContextId() string {
	if o == nil || o.ContextId == nil {
		var ret string
		return ret
	}
	return *o.ContextId
}

// GetContextIdOk returns a tuple with the ContextId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationMessage) GetContextIdOk() (*string, bool) {
	if o == nil || o.ContextId == nil {
		return nil, false
	}
	return o.ContextId, true
}

// HasContextId returns a boolean if a field has been set.
func (o *ValidationMessage) HasContextId() bool {
	if o != nil && o.ContextId != nil {
		return true
	}

	return false
}

// SetContextId gets a reference to the given string and assigns it to the ContextId field.
func (o *ValidationMessage) SetContextId(v string) {
	o.ContextId = &v
}

func (o ValidationMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.ContextId != nil {
		toSerialize["contextId"] = o.ContextId
	}
	return json.Marshal(toSerialize)
}

type NullableValidationMessage struct {
	value *ValidationMessage
	isSet bool
}

func (v NullableValidationMessage) Get() *ValidationMessage {
	return v.value
}

func (v *NullableValidationMessage) Set(val *ValidationMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationMessage(val *ValidationMessage) *NullableValidationMessage {
	return &NullableValidationMessage{value: val, isSet: true}
}

func (v NullableValidationMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

