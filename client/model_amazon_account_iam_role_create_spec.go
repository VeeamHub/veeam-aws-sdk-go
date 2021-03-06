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

// AmazonAccountIAMRoleCreateSpec struct for AmazonAccountIAMRoleCreateSpec
type AmazonAccountIAMRoleCreateSpec struct {
	RoleName string `json:"roleName"`
}

// NewAmazonAccountIAMRoleCreateSpec instantiates a new AmazonAccountIAMRoleCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonAccountIAMRoleCreateSpec(roleName string) *AmazonAccountIAMRoleCreateSpec {
	this := AmazonAccountIAMRoleCreateSpec{}
	this.RoleName = roleName
	return &this
}

// NewAmazonAccountIAMRoleCreateSpecWithDefaults instantiates a new AmazonAccountIAMRoleCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonAccountIAMRoleCreateSpecWithDefaults() *AmazonAccountIAMRoleCreateSpec {
	this := AmazonAccountIAMRoleCreateSpec{}
	return &this
}

// GetRoleName returns the RoleName field value
func (o *AmazonAccountIAMRoleCreateSpec) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *AmazonAccountIAMRoleCreateSpec) GetRoleNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value
func (o *AmazonAccountIAMRoleCreateSpec) SetRoleName(v string) {
	o.RoleName = v
}

func (o AmazonAccountIAMRoleCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["roleName"] = o.RoleName
	}
	return json.Marshal(toSerialize)
}

type NullableAmazonAccountIAMRoleCreateSpec struct {
	value *AmazonAccountIAMRoleCreateSpec
	isSet bool
}

func (v NullableAmazonAccountIAMRoleCreateSpec) Get() *AmazonAccountIAMRoleCreateSpec {
	return v.value
}

func (v *NullableAmazonAccountIAMRoleCreateSpec) Set(val *AmazonAccountIAMRoleCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonAccountIAMRoleCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonAccountIAMRoleCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonAccountIAMRoleCreateSpec(val *AmazonAccountIAMRoleCreateSpec) *NullableAmazonAccountIAMRoleCreateSpec {
	return &NullableAmazonAccountIAMRoleCreateSpec{value: val, isSet: true}
}

func (v NullableAmazonAccountIAMRoleCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonAccountIAMRoleCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


