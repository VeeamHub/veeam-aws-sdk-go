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

// CheckInstanceIdResult struct for CheckInstanceIdResult
type CheckInstanceIdResult struct {
	IsValid bool `json:"isValid"`
}

// NewCheckInstanceIdResult instantiates a new CheckInstanceIdResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckInstanceIdResult(isValid bool) *CheckInstanceIdResult {
	this := CheckInstanceIdResult{}
	this.IsValid = isValid
	return &this
}

// NewCheckInstanceIdResultWithDefaults instantiates a new CheckInstanceIdResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckInstanceIdResultWithDefaults() *CheckInstanceIdResult {
	this := CheckInstanceIdResult{}
	return &this
}

// GetIsValid returns the IsValid field value
func (o *CheckInstanceIdResult) GetIsValid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsValid
}

// GetIsValidOk returns a tuple with the IsValid field value
// and a boolean to check if the value has been set.
func (o *CheckInstanceIdResult) GetIsValidOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsValid, true
}

// SetIsValid sets field value
func (o *CheckInstanceIdResult) SetIsValid(v bool) {
	o.IsValid = v
}

func (o CheckInstanceIdResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isValid"] = o.IsValid
	}
	return json.Marshal(toSerialize)
}

type NullableCheckInstanceIdResult struct {
	value *CheckInstanceIdResult
	isSet bool
}

func (v NullableCheckInstanceIdResult) Get() *CheckInstanceIdResult {
	return v.value
}

func (v *NullableCheckInstanceIdResult) Set(val *CheckInstanceIdResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckInstanceIdResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckInstanceIdResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckInstanceIdResult(val *CheckInstanceIdResult) *NullableCheckInstanceIdResult {
	return &NullableCheckInstanceIdResult{value: val, isSet: true}
}

func (v NullableCheckInstanceIdResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckInstanceIdResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


