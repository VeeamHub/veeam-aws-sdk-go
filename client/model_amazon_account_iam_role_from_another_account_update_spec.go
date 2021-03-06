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

// AmazonAccountIAMRoleFromAnotherAccountUpdateSpec struct for AmazonAccountIAMRoleFromAnotherAccountUpdateSpec
type AmazonAccountIAMRoleFromAnotherAccountUpdateSpec struct {
	AccountId string `json:"accountId"`
	RoleName string `json:"roleName"`
	ExternalId *string `json:"externalId,omitempty"`
}

// NewAmazonAccountIAMRoleFromAnotherAccountUpdateSpec instantiates a new AmazonAccountIAMRoleFromAnotherAccountUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonAccountIAMRoleFromAnotherAccountUpdateSpec(accountId string, roleName string) *AmazonAccountIAMRoleFromAnotherAccountUpdateSpec {
	this := AmazonAccountIAMRoleFromAnotherAccountUpdateSpec{}
	this.AccountId = accountId
	this.RoleName = roleName
	return &this
}

// NewAmazonAccountIAMRoleFromAnotherAccountUpdateSpecWithDefaults instantiates a new AmazonAccountIAMRoleFromAnotherAccountUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonAccountIAMRoleFromAnotherAccountUpdateSpecWithDefaults() *AmazonAccountIAMRoleFromAnotherAccountUpdateSpec {
	this := AmazonAccountIAMRoleFromAnotherAccountUpdateSpec{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *AmazonAccountIAMRoleFromAnotherAccountUpdateSpec) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AmazonAccountIAMRoleFromAnotherAccountUpdateSpec) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *AmazonAccountIAMRoleFromAnotherAccountUpdateSpec) SetAccountId(v string) {
	o.AccountId = v
}

// GetRoleName returns the RoleName field value
func (o *AmazonAccountIAMRoleFromAnotherAccountUpdateSpec) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *AmazonAccountIAMRoleFromAnotherAccountUpdateSpec) GetRoleNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value
func (o *AmazonAccountIAMRoleFromAnotherAccountUpdateSpec) SetRoleName(v string) {
	o.RoleName = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *AmazonAccountIAMRoleFromAnotherAccountUpdateSpec) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonAccountIAMRoleFromAnotherAccountUpdateSpec) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *AmazonAccountIAMRoleFromAnotherAccountUpdateSpec) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *AmazonAccountIAMRoleFromAnotherAccountUpdateSpec) SetExternalId(v string) {
	o.ExternalId = &v
}

func (o AmazonAccountIAMRoleFromAnotherAccountUpdateSpec) MarshalJSON() ([]byte, error) {
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

type NullableAmazonAccountIAMRoleFromAnotherAccountUpdateSpec struct {
	value *AmazonAccountIAMRoleFromAnotherAccountUpdateSpec
	isSet bool
}

func (v NullableAmazonAccountIAMRoleFromAnotherAccountUpdateSpec) Get() *AmazonAccountIAMRoleFromAnotherAccountUpdateSpec {
	return v.value
}

func (v *NullableAmazonAccountIAMRoleFromAnotherAccountUpdateSpec) Set(val *AmazonAccountIAMRoleFromAnotherAccountUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonAccountIAMRoleFromAnotherAccountUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonAccountIAMRoleFromAnotherAccountUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonAccountIAMRoleFromAnotherAccountUpdateSpec(val *AmazonAccountIAMRoleFromAnotherAccountUpdateSpec) *NullableAmazonAccountIAMRoleFromAnotherAccountUpdateSpec {
	return &NullableAmazonAccountIAMRoleFromAnotherAccountUpdateSpec{value: val, isSet: true}
}

func (v NullableAmazonAccountIAMRoleFromAnotherAccountUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonAccountIAMRoleFromAnotherAccountUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


