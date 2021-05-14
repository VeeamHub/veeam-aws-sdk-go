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

// StandardAccountCreateSpec struct for StandardAccountCreateSpec
type StandardAccountCreateSpec struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Username string `json:"username"`
	Password *string `json:"password,omitempty"`
}

// NewStandardAccountCreateSpec instantiates a new StandardAccountCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardAccountCreateSpec(name string, username string) *StandardAccountCreateSpec {
	this := StandardAccountCreateSpec{}
	this.Name = name
	this.Username = username
	return &this
}

// NewStandardAccountCreateSpecWithDefaults instantiates a new StandardAccountCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardAccountCreateSpecWithDefaults() *StandardAccountCreateSpec {
	this := StandardAccountCreateSpec{}
	return &this
}

// GetName returns the Name field value
func (o *StandardAccountCreateSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StandardAccountCreateSpec) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StandardAccountCreateSpec) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StandardAccountCreateSpec) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardAccountCreateSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StandardAccountCreateSpec) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StandardAccountCreateSpec) SetDescription(v string) {
	o.Description = &v
}

// GetUsername returns the Username field value
func (o *StandardAccountCreateSpec) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *StandardAccountCreateSpec) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *StandardAccountCreateSpec) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *StandardAccountCreateSpec) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardAccountCreateSpec) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *StandardAccountCreateSpec) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *StandardAccountCreateSpec) SetPassword(v string) {
	o.Password = &v
}

func (o StandardAccountCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableStandardAccountCreateSpec struct {
	value *StandardAccountCreateSpec
	isSet bool
}

func (v NullableStandardAccountCreateSpec) Get() *StandardAccountCreateSpec {
	return v.value
}

func (v *NullableStandardAccountCreateSpec) Set(val *StandardAccountCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardAccountCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardAccountCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardAccountCreateSpec(val *StandardAccountCreateSpec) *NullableStandardAccountCreateSpec {
	return &NullableStandardAccountCreateSpec{value: val, isSet: true}
}

func (v NullableStandardAccountCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandardAccountCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


