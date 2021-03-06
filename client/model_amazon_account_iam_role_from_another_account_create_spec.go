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

// AmazonAccountIAMRoleFromAnotherAccountCreateSpec struct for AmazonAccountIAMRoleFromAnotherAccountCreateSpec
type AmazonAccountIAMRoleFromAnotherAccountCreateSpec struct {
	AccountId string `json:"accountId"`
	RoleName string `json:"roleName"`
	ExternalId *string `json:"externalId,omitempty"`
}

// NewAmazonAccountIAMRoleFromAnotherAccountCreateSpec instantiates a new AmazonAccountIAMRoleFromAnotherAccountCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonAccountIAMRoleFromAnotherAccountCreateSpec(accountId string, roleName string) *AmazonAccountIAMRoleFromAnotherAccountCreateSpec {
	this := AmazonAccountIAMRoleFromAnotherAccountCreateSpec{}
	this.AccountId = accountId
	this.RoleName = roleName
	return &this
}

// NewAmazonAccountIAMRoleFromAnotherAccountCreateSpecWithDefaults instantiates a new AmazonAccountIAMRoleFromAnotherAccountCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonAccountIAMRoleFromAnotherAccountCreateSpecWithDefaults() *AmazonAccountIAMRoleFromAnotherAccountCreateSpec {
	this := AmazonAccountIAMRoleFromAnotherAccountCreateSpec{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *AmazonAccountIAMRoleFromAnotherAccountCreateSpec) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AmazonAccountIAMRoleFromAnotherAccountCreateSpec) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *AmazonAccountIAMRoleFromAnotherAccountCreateSpec) SetAccountId(v string) {
	o.AccountId = v
}

// GetRoleName returns the RoleName field value
func (o *AmazonAccountIAMRoleFromAnotherAccountCreateSpec) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *AmazonAccountIAMRoleFromAnotherAccountCreateSpec) GetRoleNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value
func (o *AmazonAccountIAMRoleFromAnotherAccountCreateSpec) SetRoleName(v string) {
	o.RoleName = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *AmazonAccountIAMRoleFromAnotherAccountCreateSpec) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonAccountIAMRoleFromAnotherAccountCreateSpec) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *AmazonAccountIAMRoleFromAnotherAccountCreateSpec) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *AmazonAccountIAMRoleFromAnotherAccountCreateSpec) SetExternalId(v string) {
	o.ExternalId = &v
}

func (o AmazonAccountIAMRoleFromAnotherAccountCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["roleName"] = o.RoleName
	}
	if o.ExternalId != nil {
		toSerialize["externalId"] = o.ExternalId
	}
	return json.Marshal(toSerialize)
}

type NullableAmazonAccountIAMRoleFromAnotherAccountCreateSpec struct {
	value *AmazonAccountIAMRoleFromAnotherAccountCreateSpec
	isSet bool
}

func (v NullableAmazonAccountIAMRoleFromAnotherAccountCreateSpec) Get() *AmazonAccountIAMRoleFromAnotherAccountCreateSpec {
	return v.value
}

func (v *NullableAmazonAccountIAMRoleFromAnotherAccountCreateSpec) Set(val *AmazonAccountIAMRoleFromAnotherAccountCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonAccountIAMRoleFromAnotherAccountCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonAccountIAMRoleFromAnotherAccountCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonAccountIAMRoleFromAnotherAccountCreateSpec(val *AmazonAccountIAMRoleFromAnotherAccountCreateSpec) *NullableAmazonAccountIAMRoleFromAnotherAccountCreateSpec {
	return &NullableAmazonAccountIAMRoleFromAnotherAccountCreateSpec{value: val, isSet: true}
}

func (v NullableAmazonAccountIAMRoleFromAnotherAccountCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonAccountIAMRoleFromAnotherAccountCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


