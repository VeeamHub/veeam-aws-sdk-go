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

// RetrySettings struct for RetrySettings
type RetrySettings struct {
	RetryTimes int32 `json:"retryTimes"`
}

// NewRetrySettings instantiates a new RetrySettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetrySettings(retryTimes int32) *RetrySettings {
	this := RetrySettings{}
	this.RetryTimes = retryTimes
	return &this
}

// NewRetrySettingsWithDefaults instantiates a new RetrySettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetrySettingsWithDefaults() *RetrySettings {
	this := RetrySettings{}
	return &this
}

// GetRetryTimes returns the RetryTimes field value
func (o *RetrySettings) GetRetryTimes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RetryTimes
}

// GetRetryTimesOk returns a tuple with the RetryTimes field value
// and a boolean to check if the value has been set.
func (o *RetrySettings) GetRetryTimesOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RetryTimes, true
}

// SetRetryTimes sets field value
func (o *RetrySettings) SetRetryTimes(v int32) {
	o.RetryTimes = v
}

func (o RetrySettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["retryTimes"] = o.RetryTimes
	}
	return json.Marshal(toSerialize)
}

type NullableRetrySettings struct {
	value *RetrySettings
	isSet bool
}

func (v NullableRetrySettings) Get() *RetrySettings {
	return v.value
}

func (v *NullableRetrySettings) Set(val *RetrySettings) {
	v.value = val
	v.isSet = true
}

func (v NullableRetrySettings) IsSet() bool {
	return v.isSet
}

func (v *NullableRetrySettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetrySettings(val *RetrySettings) *NullableRetrySettings {
	return &NullableRetrySettings{value: val, isSet: true}
}

func (v NullableRetrySettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetrySettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

