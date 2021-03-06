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

// UserCreateSpec struct for UserCreateSpec
type UserCreateSpec struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Password string `json:"password"`
}

// NewUserCreateSpec instantiates a new UserCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCreateSpec(name string, description string, password string) *UserCreateSpec {
	this := UserCreateSpec{}
	this.Name = name
	this.Description = description
	this.Password = password
	return &this
}

// NewUserCreateSpecWithDefaults instantiates a new UserCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCreateSpecWithDefaults() *UserCreateSpec {
	this := UserCreateSpec{}
	return &this
}

// GetName returns the Name field value
func (o *UserCreateSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserCreateSpec) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserCreateSpec) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *UserCreateSpec) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *UserCreateSpec) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *UserCreateSpec) SetDescription(v string) {
	o.Description = v
}

// GetPassword returns the Password field value
func (o *UserCreateSpec) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *UserCreateSpec) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *UserCreateSpec) SetPassword(v string) {
	o.Password = v
}

func (o UserCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableUserCreateSpec struct {
	value *UserCreateSpec
	isSet bool
}

func (v NullableUserCreateSpec) Get() *UserCreateSpec {
	return v.value
}

func (v *NullableUserCreateSpec) Set(val *UserCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCreateSpec(val *UserCreateSpec) *NullableUserCreateSpec {
	return &NullableUserCreateSpec{value: val, isSet: true}
}

func (v NullableUserCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


