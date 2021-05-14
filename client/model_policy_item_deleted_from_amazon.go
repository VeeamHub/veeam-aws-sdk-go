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

// PolicyItemDeletedFromAmazon struct for PolicyItemDeletedFromAmazon
type PolicyItemDeletedFromAmazon struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// NewPolicyItemDeletedFromAmazon instantiates a new PolicyItemDeletedFromAmazon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyItemDeletedFromAmazon(id string, name string, type_ string) *PolicyItemDeletedFromAmazon {
	this := PolicyItemDeletedFromAmazon{}
	this.Id = id
	this.Name = name
	this.Type = type_
	return &this
}

// NewPolicyItemDeletedFromAmazonWithDefaults instantiates a new PolicyItemDeletedFromAmazon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyItemDeletedFromAmazonWithDefaults() *PolicyItemDeletedFromAmazon {
	this := PolicyItemDeletedFromAmazon{}
	return &this
}

// GetId returns the Id field value
func (o *PolicyItemDeletedFromAmazon) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PolicyItemDeletedFromAmazon) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PolicyItemDeletedFromAmazon) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *PolicyItemDeletedFromAmazon) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PolicyItemDeletedFromAmazon) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PolicyItemDeletedFromAmazon) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *PolicyItemDeletedFromAmazon) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PolicyItemDeletedFromAmazon) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PolicyItemDeletedFromAmazon) SetType(v string) {
	o.Type = v
}

func (o PolicyItemDeletedFromAmazon) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyItemDeletedFromAmazon struct {
	value *PolicyItemDeletedFromAmazon
	isSet bool
}

func (v NullablePolicyItemDeletedFromAmazon) Get() *PolicyItemDeletedFromAmazon {
	return v.value
}

func (v *NullablePolicyItemDeletedFromAmazon) Set(val *PolicyItemDeletedFromAmazon) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyItemDeletedFromAmazon) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyItemDeletedFromAmazon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyItemDeletedFromAmazon(val *PolicyItemDeletedFromAmazon) *NullablePolicyItemDeletedFromAmazon {
	return &NullablePolicyItemDeletedFromAmazon{value: val, isSet: true}
}

func (v NullablePolicyItemDeletedFromAmazon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyItemDeletedFromAmazon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

