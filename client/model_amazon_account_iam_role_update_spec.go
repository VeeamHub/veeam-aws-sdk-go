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

// AmazonAccountIAMRoleUpdateSpec struct for AmazonAccountIAMRoleUpdateSpec
type AmazonAccountIAMRoleUpdateSpec struct {
	RoleName string `json:"roleName"`
}

// NewAmazonAccountIAMRoleUpdateSpec instantiates a new AmazonAccountIAMRoleUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonAccountIAMRoleUpdateSpec(roleName string) *AmazonAccountIAMRoleUpdateSpec {
	this := AmazonAccountIAMRoleUpdateSpec{}
	this.RoleName = roleName
	return &this
}

// NewAmazonAccountIAMRoleUpdateSpecWithDefaults instantiates a new AmazonAccountIAMRoleUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonAccountIAMRoleUpdateSpecWithDefaults() *AmazonAccountIAMRoleUpdateSpec {
	this := AmazonAccountIAMRoleUpdateSpec{}
	return &this
}

// GetRoleName returns the RoleName field value
func (o *AmazonAccountIAMRoleUpdateSpec) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *AmazonAccountIAMRoleUpdateSpec) GetRoleNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value
func (o *AmazonAccountIAMRoleUpdateSpec) SetRoleName(v string) {
	o.RoleName = v
}

func (o AmazonAccountIAMRoleUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["roleName"] = o.RoleName
	}
	return json.Marshal(toSerialize)
}

type NullableAmazonAccountIAMRoleUpdateSpec struct {
	value *AmazonAccountIAMRoleUpdateSpec
	isSet bool
}

func (v NullableAmazonAccountIAMRoleUpdateSpec) Get() *AmazonAccountIAMRoleUpdateSpec {
	return v.value
}

func (v *NullableAmazonAccountIAMRoleUpdateSpec) Set(val *AmazonAccountIAMRoleUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonAccountIAMRoleUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonAccountIAMRoleUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonAccountIAMRoleUpdateSpec(val *AmazonAccountIAMRoleUpdateSpec) *NullableAmazonAccountIAMRoleUpdateSpec {
	return &NullableAmazonAccountIAMRoleUpdateSpec{value: val, isSet: true}
}

func (v NullableAmazonAccountIAMRoleUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonAccountIAMRoleUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


