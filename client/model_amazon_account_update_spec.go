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

// AmazonAccountUpdateSpec struct for AmazonAccountUpdateSpec
type AmazonAccountUpdateSpec struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	AccessKeys *AmazonAccountAccessKeysUpdateSpec `json:"accessKeys,omitempty"`
	IAMRole *AmazonAccountIAMRoleUpdateSpec `json:"IAMRole,omitempty"`
	IAMRoleFromAnotherAccount *AmazonAccountIAMRoleFromAnotherAccountUpdateSpec `json:"IAMRoleFromAnotherAccount,omitempty"`
}

// NewAmazonAccountUpdateSpec instantiates a new AmazonAccountUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonAccountUpdateSpec(name string) *AmazonAccountUpdateSpec {
	this := AmazonAccountUpdateSpec{}
	this.Name = name
	return &this
}

// NewAmazonAccountUpdateSpecWithDefaults instantiates a new AmazonAccountUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonAccountUpdateSpecWithDefaults() *AmazonAccountUpdateSpec {
	this := AmazonAccountUpdateSpec{}
	return &this
}

// GetName returns the Name field value
func (o *AmazonAccountUpdateSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AmazonAccountUpdateSpec) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AmazonAccountUpdateSpec) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AmazonAccountUpdateSpec) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonAccountUpdateSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AmazonAccountUpdateSpec) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AmazonAccountUpdateSpec) SetDescription(v string) {
	o.Description = &v
}

// GetAccessKeys returns the AccessKeys field value if set, zero value otherwise.
func (o *AmazonAccountUpdateSpec) GetAccessKeys() AmazonAccountAccessKeysUpdateSpec {
	if o == nil || o.AccessKeys == nil {
		var ret AmazonAccountAccessKeysUpdateSpec
		return ret
	}
	return *o.AccessKeys
}

// GetAccessKeysOk returns a tuple with the AccessKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonAccountUpdateSpec) GetAccessKeysOk() (*AmazonAccountAccessKeysUpdateSpec, bool) {
	if o == nil || o.AccessKeys == nil {
		return nil, false
	}
	return o.AccessKeys, true
}

// HasAccessKeys returns a boolean if a field has been set.
func (o *AmazonAccountUpdateSpec) HasAccessKeys() bool {
	if o != nil && o.AccessKeys != nil {
		return true
	}

	return false
}

// SetAccessKeys gets a reference to the given AmazonAccountAccessKeysUpdateSpec and assigns it to the AccessKeys field.
func (o *AmazonAccountUpdateSpec) SetAccessKeys(v AmazonAccountAccessKeysUpdateSpec) {
	o.AccessKeys = &v
}

// GetIAMRole returns the IAMRole field value if set, zero value otherwise.
func (o *AmazonAccountUpdateSpec) GetIAMRole() AmazonAccountIAMRoleUpdateSpec {
	if o == nil || o.IAMRole == nil {
		var ret AmazonAccountIAMRoleUpdateSpec
		return ret
	}
	return *o.IAMRole
}

// GetIAMRoleOk returns a tuple with the IAMRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonAccountUpdateSpec) GetIAMRoleOk() (*AmazonAccountIAMRoleUpdateSpec, bool) {
	if o == nil || o.IAMRole == nil {
		return nil, false
	}
	return o.IAMRole, true
}

// HasIAMRole returns a boolean if a field has been set.
func (o *AmazonAccountUpdateSpec) HasIAMRole() bool {
	if o != nil && o.IAMRole != nil {
		return true
	}

	return false
}

// SetIAMRole gets a reference to the given AmazonAccountIAMRoleUpdateSpec and assigns it to the IAMRole field.
func (o *AmazonAccountUpdateSpec) SetIAMRole(v AmazonAccountIAMRoleUpdateSpec) {
	o.IAMRole = &v
}

// GetIAMRoleFromAnotherAccount returns the IAMRoleFromAnotherAccount field value if set, zero value otherwise.
func (o *AmazonAccountUpdateSpec) GetIAMRoleFromAnotherAccount() AmazonAccountIAMRoleFromAnotherAccountUpdateSpec {
	if o == nil || o.IAMRoleFromAnotherAccount == nil {
		var ret AmazonAccountIAMRoleFromAnotherAccountUpdateSpec
		return ret
	}
	return *o.IAMRoleFromAnotherAccount
}

// GetIAMRoleFromAnotherAccountOk returns a tuple with the IAMRoleFromAnotherAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonAccountUpdateSpec) GetIAMRoleFromAnotherAccountOk() (*AmazonAccountIAMRoleFromAnotherAccountUpdateSpec, bool) {
	if o == nil || o.IAMRoleFromAnotherAccount == nil {
		return nil, false
	}
	return o.IAMRoleFromAnotherAccount, true
}

// HasIAMRoleFromAnotherAccount returns a boolean if a field has been set.
func (o *AmazonAccountUpdateSpec) HasIAMRoleFromAnotherAccount() bool {
	if o != nil && o.IAMRoleFromAnotherAccount != nil {
		return true
	}

	return false
}

// SetIAMRoleFromAnotherAccount gets a reference to the given AmazonAccountIAMRoleFromAnotherAccountUpdateSpec and assigns it to the IAMRoleFromAnotherAccount field.
func (o *AmazonAccountUpdateSpec) SetIAMRoleFromAnotherAccount(v AmazonAccountIAMRoleFromAnotherAccountUpdateSpec) {
	o.IAMRoleFromAnotherAccount = &v
}

func (o AmazonAccountUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.AccessKeys != nil {
		toSerialize["accessKeys"] = o.AccessKeys
	}
	if o.IAMRole != nil {
		toSerialize["IAMRole"] = o.IAMRole
	}
	if o.IAMRoleFromAnotherAccount != nil {
		toSerialize["IAMRoleFromAnotherAccount"] = o.IAMRoleFromAnotherAccount
	}
	return json.Marshal(toSerialize)
}

type NullableAmazonAccountUpdateSpec struct {
	value *AmazonAccountUpdateSpec
	isSet bool
}

func (v NullableAmazonAccountUpdateSpec) Get() *AmazonAccountUpdateSpec {
	return v.value
}

func (v *NullableAmazonAccountUpdateSpec) Set(val *AmazonAccountUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonAccountUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonAccountUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonAccountUpdateSpec(val *AmazonAccountUpdateSpec) *NullableAmazonAccountUpdateSpec {
	return &NullableAmazonAccountUpdateSpec{value: val, isSet: true}
}

func (v NullableAmazonAccountUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonAccountUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


