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

// RepositoryPasswordValidationResult struct for RepositoryPasswordValidationResult
type RepositoryPasswordValidationResult struct {
	IsSuccess bool `json:"isSuccess"`
}

// NewRepositoryPasswordValidationResult instantiates a new RepositoryPasswordValidationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryPasswordValidationResult(isSuccess bool) *RepositoryPasswordValidationResult {
	this := RepositoryPasswordValidationResult{}
	this.IsSuccess = isSuccess
	return &this
}

// NewRepositoryPasswordValidationResultWithDefaults instantiates a new RepositoryPasswordValidationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryPasswordValidationResultWithDefaults() *RepositoryPasswordValidationResult {
	this := RepositoryPasswordValidationResult{}
	return &this
}

// GetIsSuccess returns the IsSuccess field value
func (o *RepositoryPasswordValidationResult) GetIsSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSuccess
}

// GetIsSuccessOk returns a tuple with the IsSuccess field value
// and a boolean to check if the value has been set.
func (o *RepositoryPasswordValidationResult) GetIsSuccessOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsSuccess, true
}

// SetIsSuccess sets field value
func (o *RepositoryPasswordValidationResult) SetIsSuccess(v bool) {
	o.IsSuccess = v
}

func (o RepositoryPasswordValidationResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isSuccess"] = o.IsSuccess
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryPasswordValidationResult struct {
	value *RepositoryPasswordValidationResult
	isSet bool
}

func (v NullableRepositoryPasswordValidationResult) Get() *RepositoryPasswordValidationResult {
	return v.value
}

func (v *NullableRepositoryPasswordValidationResult) Set(val *RepositoryPasswordValidationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryPasswordValidationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryPasswordValidationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryPasswordValidationResult(val *RepositoryPasswordValidationResult) *NullableRepositoryPasswordValidationResult {
	return &NullableRepositoryPasswordValidationResult{value: val, isSet: true}
}

func (v NullableRepositoryPasswordValidationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryPasswordValidationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


