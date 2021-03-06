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

// AmazonAccountIAMRoleFromAnotherAccount struct for AmazonAccountIAMRoleFromAnotherAccount
type AmazonAccountIAMRoleFromAnotherAccount struct {
	ParentAmazonAccountId string `json:"parentAmazonAccountId"`
	AccountId string `json:"accountId"`
	RoleName string `json:"roleName"`
}

// NewAmazonAccountIAMRoleFromAnotherAccount instantiates a new AmazonAccountIAMRoleFromAnotherAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonAccountIAMRoleFromAnotherAccount(parentAmazonAccountId string, accountId string, roleName string) *AmazonAccountIAMRoleFromAnotherAccount {
	this := AmazonAccountIAMRoleFromAnotherAccount{}
	this.ParentAmazonAccountId = parentAmazonAccountId
	this.AccountId = accountId
	this.RoleName = roleName
	return &this
}

// NewAmazonAccountIAMRoleFromAnotherAccountWithDefaults instantiates a new AmazonAccountIAMRoleFromAnotherAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonAccountIAMRoleFromAnotherAccountWithDefaults() *AmazonAccountIAMRoleFromAnotherAccount {
	this := AmazonAccountIAMRoleFromAnotherAccount{}
	return &this
}

// GetParentAmazonAccountId returns the ParentAmazonAccountId field value
func (o *AmazonAccountIAMRoleFromAnotherAccount) GetParentAmazonAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentAmazonAccountId
}

// GetParentAmazonAccountIdOk returns a tuple with the ParentAmazonAccountId field value
// and a boolean to check if the value has been set.
func (o *AmazonAccountIAMRoleFromAnotherAccount) GetParentAmazonAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ParentAmazonAccountId, true
}

// SetParentAmazonAccountId sets field value
func (o *AmazonAccountIAMRoleFromAnotherAccount) SetParentAmazonAccountId(v string) {
	o.ParentAmazonAccountId = v
}

// GetAccountId returns the AccountId field value
func (o *AmazonAccountIAMRoleFromAnotherAccount) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AmazonAccountIAMRoleFromAnotherAccount) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *AmazonAccountIAMRoleFromAnotherAccount) SetAccountId(v string) {
	o.AccountId = v
}

// GetRoleName returns the RoleName field value
func (o *AmazonAccountIAMRoleFromAnotherAccount) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *AmazonAccountIAMRoleFromAnotherAccount) GetRoleNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value
func (o *AmazonAccountIAMRoleFromAnotherAccount) SetRoleName(v string) {
	o.RoleName = v
}

func (o AmazonAccountIAMRoleFromAnotherAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["parentAmazonAccountId"] = o.ParentAmazonAccountId
	}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["roleName"] = o.RoleName
	}
	return json.Marshal(toSerialize)
}

type NullableAmazonAccountIAMRoleFromAnotherAccount struct {
	value *AmazonAccountIAMRoleFromAnotherAccount
	isSet bool
}

func (v NullableAmazonAccountIAMRoleFromAnotherAccount) Get() *AmazonAccountIAMRoleFromAnotherAccount {
	return v.value
}

func (v *NullableAmazonAccountIAMRoleFromAnotherAccount) Set(val *AmazonAccountIAMRoleFromAnotherAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonAccountIAMRoleFromAnotherAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonAccountIAMRoleFromAnotherAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonAccountIAMRoleFromAnotherAccount(val *AmazonAccountIAMRoleFromAnotherAccount) *NullableAmazonAccountIAMRoleFromAnotherAccount {
	return &NullableAmazonAccountIAMRoleFromAnotherAccount{value: val, isSet: true}
}

func (v NullableAmazonAccountIAMRoleFromAnotherAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonAccountIAMRoleFromAnotherAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


