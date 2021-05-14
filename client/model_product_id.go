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

// ProductId struct for ProductId
type ProductId struct {
	Awsid string `json:"awsid"`
	Supportid string `json:"supportid"`
}

// NewProductId instantiates a new ProductId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductId(awsid string, supportid string) *ProductId {
	this := ProductId{}
	this.Awsid = awsid
	this.Supportid = supportid
	return &this
}

// NewProductIdWithDefaults instantiates a new ProductId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductIdWithDefaults() *ProductId {
	this := ProductId{}
	return &this
}

// GetAwsid returns the Awsid field value
func (o *ProductId) GetAwsid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Awsid
}

// GetAwsidOk returns a tuple with the Awsid field value
// and a boolean to check if the value has been set.
func (o *ProductId) GetAwsidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Awsid, true
}

// SetAwsid sets field value
func (o *ProductId) SetAwsid(v string) {
	o.Awsid = v
}

// GetSupportid returns the Supportid field value
func (o *ProductId) GetSupportid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Supportid
}

// GetSupportidOk returns a tuple with the Supportid field value
// and a boolean to check if the value has been set.
func (o *ProductId) GetSupportidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Supportid, true
}

// SetSupportid sets field value
func (o *ProductId) SetSupportid(v string) {
	o.Supportid = v
}

func (o ProductId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["awsid"] = o.Awsid
	}
	if true {
		toSerialize["supportid"] = o.Supportid
	}
	return json.Marshal(toSerialize)
}

type NullableProductId struct {
	value *ProductId
	isSet bool
}

func (v NullableProductId) Get() *ProductId {
	return v.value
}

func (v *NullableProductId) Set(val *ProductId) {
	v.value = val
	v.isSet = true
}

func (v NullableProductId) IsSet() bool {
	return v.isSet
}

func (v *NullableProductId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductId(val *ProductId) *NullableProductId {
	return &NullableProductId{value: val, isSet: true}
}

func (v NullableProductId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


