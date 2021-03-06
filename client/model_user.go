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

// User struct for User
type User struct {
	Name string `json:"name"`
	Description string `json:"description"`
	MfaEnabled bool `json:"mfaEnabled"`
	IsDefault bool `json:"isDefault"`
	Links *[]Link `json:"_links,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser(name string, description string, mfaEnabled bool, isDefault bool) *User {
	this := User{}
	this.Name = name
	this.Description = description
	this.MfaEnabled = mfaEnabled
	this.IsDefault = isDefault
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetName returns the Name field value
func (o *User) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *User) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *User) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *User) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *User) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *User) SetDescription(v string) {
	o.Description = v
}

// GetMfaEnabled returns the MfaEnabled field value
func (o *User) GetMfaEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MfaEnabled
}

// GetMfaEnabledOk returns a tuple with the MfaEnabled field value
// and a boolean to check if the value has been set.
func (o *User) GetMfaEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MfaEnabled, true
}

// SetMfaEnabled sets field value
func (o *User) SetMfaEnabled(v bool) {
	o.MfaEnabled = v
}

// GetIsDefault returns the IsDefault field value
func (o *User) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *User) GetIsDefaultOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value
func (o *User) SetIsDefault(v bool) {
	o.IsDefault = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *User) GetLinks() []Link {
	if o == nil || o.Links == nil {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLinksOk() (*[]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *User) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *User) SetLinks(v []Link) {
	o.Links = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["mfaEnabled"] = o.MfaEnabled
	}
	if true {
		toSerialize["isDefault"] = o.IsDefault
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


