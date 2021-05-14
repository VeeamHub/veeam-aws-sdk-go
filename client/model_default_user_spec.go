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

// DefaultUserSpec struct for DefaultUserSpec
type DefaultUserSpec struct {
	Instance CheckInstanceIdSpec `json:"instance"`
	UserSpec UserCreateSpec `json:"userSpec"`
}

// NewDefaultUserSpec instantiates a new DefaultUserSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultUserSpec(instance CheckInstanceIdSpec, userSpec UserCreateSpec) *DefaultUserSpec {
	this := DefaultUserSpec{}
	this.Instance = instance
	this.UserSpec = userSpec
	return &this
}

// NewDefaultUserSpecWithDefaults instantiates a new DefaultUserSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultUserSpecWithDefaults() *DefaultUserSpec {
	this := DefaultUserSpec{}
	return &this
}

// GetInstance returns the Instance field value
func (o *DefaultUserSpec) GetInstance() CheckInstanceIdSpec {
	if o == nil {
		var ret CheckInstanceIdSpec
		return ret
	}

	return o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value
// and a boolean to check if the value has been set.
func (o *DefaultUserSpec) GetInstanceOk() (*CheckInstanceIdSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Instance, true
}

// SetInstance sets field value
func (o *DefaultUserSpec) SetInstance(v CheckInstanceIdSpec) {
	o.Instance = v
}

// GetUserSpec returns the UserSpec field value
func (o *DefaultUserSpec) GetUserSpec() UserCreateSpec {
	if o == nil {
		var ret UserCreateSpec
		return ret
	}

	return o.UserSpec
}

// GetUserSpecOk returns a tuple with the UserSpec field value
// and a boolean to check if the value has been set.
func (o *DefaultUserSpec) GetUserSpecOk() (*UserCreateSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserSpec, true
}

// SetUserSpec sets field value
func (o *DefaultUserSpec) SetUserSpec(v UserCreateSpec) {
	o.UserSpec = v
}

func (o DefaultUserSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["instance"] = o.Instance
	}
	if true {
		toSerialize["userSpec"] = o.UserSpec
	}
	return json.Marshal(toSerialize)
}

type NullableDefaultUserSpec struct {
	value *DefaultUserSpec
	isSet bool
}

func (v NullableDefaultUserSpec) Get() *DefaultUserSpec {
	return v.value
}

func (v *NullableDefaultUserSpec) Set(val *DefaultUserSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultUserSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultUserSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultUserSpec(val *DefaultUserSpec) *NullableDefaultUserSpec {
	return &NullableDefaultUserSpec{value: val, isSet: true}
}

func (v NullableDefaultUserSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultUserSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

