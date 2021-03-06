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

// AmazonAccountCreateSpec struct for AmazonAccountCreateSpec
type AmazonAccountCreateSpec struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	AccessKeys *AmazonAccountAccessKeysCreateSpec `json:"accessKeys,omitempty"`
	IAMRole *AmazonAccountIAMRoleCreateSpec `json:"IAMRole,omitempty"`
	IAMRoleFromAnotherAccount *AmazonAccountIAMRoleFromAnotherAccountCreateSpec `json:"IAMRoleFromAnotherAccount,omitempty"`
}

// NewAmazonAccountCreateSpec instantiates a new AmazonAccountCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonAccountCreateSpec(name string) *AmazonAccountCreateSpec {
	this := AmazonAccountCreateSpec{}
	this.Name = name
	return &this
}

// NewAmazonAccountCreateSpecWithDefaults instantiates a new AmazonAccountCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonAccountCreateSpecWithDefaults() *AmazonAccountCreateSpec {
	this := AmazonAccountCreateSpec{}
	return &this
}

// GetName returns the Name field value
func (o *AmazonAccountCreateSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AmazonAccountCreateSpec) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AmazonAccountCreateSpec) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AmazonAccountCreateSpec) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonAccountCreateSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AmazonAccountCreateSpec) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AmazonAccountCreateSpec) SetDescription(v string) {
	o.Description = &v
}

// GetAccessKeys returns the AccessKeys field value if set, zero value otherwise.
func (o *AmazonAccountCreateSpec) GetAccessKeys() AmazonAccountAccessKeysCreateSpec {
	if o == nil || o.AccessKeys == nil {
		var ret AmazonAccountAccessKeysCreateSpec
		return ret
	}
	return *o.AccessKeys
}

// GetAccessKeysOk returns a tuple with the AccessKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonAccountCreateSpec) GetAccessKeysOk() (*AmazonAccountAccessKeysCreateSpec, bool) {
	if o == nil || o.AccessKeys == nil {
		return nil, false
	}
	return o.AccessKeys, true
}

// HasAccessKeys returns a boolean if a field has been set.
func (o *AmazonAccountCreateSpec) HasAccessKeys() bool {
	if o != nil && o.AccessKeys != nil {
		return true
	}

	return false
}

// SetAccessKeys gets a reference to the given AmazonAccountAccessKeysCreateSpec and assigns it to the AccessKeys field.
func (o *AmazonAccountCreateSpec) SetAccessKeys(v AmazonAccountAccessKeysCreateSpec) {
	o.AccessKeys = &v
}

// GetIAMRole returns the IAMRole field value if set, zero value otherwise.
func (o *AmazonAccountCreateSpec) GetIAMRole() AmazonAccountIAMRoleCreateSpec {
	if o == nil || o.IAMRole == nil {
		var ret AmazonAccountIAMRoleCreateSpec
		return ret
	}
	return *o.IAMRole
}

// GetIAMRoleOk returns a tuple with the IAMRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonAccountCreateSpec) GetIAMRoleOk() (*AmazonAccountIAMRoleCreateSpec, bool) {
	if o == nil || o.IAMRole == nil {
		return nil, false
	}
	return o.IAMRole, true
}

// HasIAMRole returns a boolean if a field has been set.
func (o *AmazonAccountCreateSpec) HasIAMRole() bool {
	if o != nil && o.IAMRole != nil {
		return true
	}

	return false
}

// SetIAMRole gets a reference to the given AmazonAccountIAMRoleCreateSpec and assigns it to the IAMRole field.
func (o *AmazonAccountCreateSpec) SetIAMRole(v AmazonAccountIAMRoleCreateSpec) {
	o.IAMRole = &v
}

// GetIAMRoleFromAnotherAccount returns the IAMRoleFromAnotherAccount field value if set, zero value otherwise.
func (o *AmazonAccountCreateSpec) GetIAMRoleFromAnotherAccount() AmazonAccountIAMRoleFromAnotherAccountCreateSpec {
	if o == nil || o.IAMRoleFromAnotherAccount == nil {
		var ret AmazonAccountIAMRoleFromAnotherAccountCreateSpec
		return ret
	}
	return *o.IAMRoleFromAnotherAccount
}

// GetIAMRoleFromAnotherAccountOk returns a tuple with the IAMRoleFromAnotherAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonAccountCreateSpec) GetIAMRoleFromAnotherAccountOk() (*AmazonAccountIAMRoleFromAnotherAccountCreateSpec, bool) {
	if o == nil || o.IAMRoleFromAnotherAccount == nil {
		return nil, false
	}
	return o.IAMRoleFromAnotherAccount, true
}

// HasIAMRoleFromAnotherAccount returns a boolean if a field has been set.
func (o *AmazonAccountCreateSpec) HasIAMRoleFromAnotherAccount() bool {
	if o != nil && o.IAMRoleFromAnotherAccount != nil {
		return true
	}

	return false
}

// SetIAMRoleFromAnotherAccount gets a reference to the given AmazonAccountIAMRoleFromAnotherAccountCreateSpec and assigns it to the IAMRoleFromAnotherAccount field.
func (o *AmazonAccountCreateSpec) SetIAMRoleFromAnotherAccount(v AmazonAccountIAMRoleFromAnotherAccountCreateSpec) {
	o.IAMRoleFromAnotherAccount = &v
}

func (o AmazonAccountCreateSpec) MarshalJSON() ([]byte, error) {
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

type NullableAmazonAccountCreateSpec struct {
	value *AmazonAccountCreateSpec
	isSet bool
}

func (v NullableAmazonAccountCreateSpec) Get() *AmazonAccountCreateSpec {
	return v.value
}

func (v *NullableAmazonAccountCreateSpec) Set(val *AmazonAccountCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonAccountCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonAccountCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonAccountCreateSpec(val *AmazonAccountCreateSpec) *NullableAmazonAccountCreateSpec {
	return &NullableAmazonAccountCreateSpec{value: val, isSet: true}
}

func (v NullableAmazonAccountCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonAccountCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


