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

// AmazonConnectionSpec struct for AmazonConnectionSpec
type AmazonConnectionSpec struct {
	AmazonAccountId *string `json:"AmazonAccountId,omitempty"`
	AccessKey *string `json:"accessKey,omitempty"`
	SecretKey *string `json:"secretKey,omitempty"`
	RegionId string `json:"RegionId"`
}

// NewAmazonConnectionSpec instantiates a new AmazonConnectionSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonConnectionSpec(regionId string) *AmazonConnectionSpec {
	this := AmazonConnectionSpec{}
	this.RegionId = regionId
	return &this
}

// NewAmazonConnectionSpecWithDefaults instantiates a new AmazonConnectionSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonConnectionSpecWithDefaults() *AmazonConnectionSpec {
	this := AmazonConnectionSpec{}
	return &this
}

// GetAmazonAccountId returns the AmazonAccountId field value if set, zero value otherwise.
func (o *AmazonConnectionSpec) GetAmazonAccountId() string {
	if o == nil || o.AmazonAccountId == nil {
		var ret string
		return ret
	}
	return *o.AmazonAccountId
}

// GetAmazonAccountIdOk returns a tuple with the AmazonAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonConnectionSpec) GetAmazonAccountIdOk() (*string, bool) {
	if o == nil || o.AmazonAccountId == nil {
		return nil, false
	}
	return o.AmazonAccountId, true
}

// HasAmazonAccountId returns a boolean if a field has been set.
func (o *AmazonConnectionSpec) HasAmazonAccountId() bool {
	if o != nil && o.AmazonAccountId != nil {
		return true
	}

	return false
}

// SetAmazonAccountId gets a reference to the given string and assigns it to the AmazonAccountId field.
func (o *AmazonConnectionSpec) SetAmazonAccountId(v string) {
	o.AmazonAccountId = &v
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *AmazonConnectionSpec) GetAccessKey() string {
	if o == nil || o.AccessKey == nil {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonConnectionSpec) GetAccessKeyOk() (*string, bool) {
	if o == nil || o.AccessKey == nil {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *AmazonConnectionSpec) HasAccessKey() bool {
	if o != nil && o.AccessKey != nil {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *AmazonConnectionSpec) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *AmazonConnectionSpec) GetSecretKey() string {
	if o == nil || o.SecretKey == nil {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonConnectionSpec) GetSecretKeyOk() (*string, bool) {
	if o == nil || o.SecretKey == nil {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *AmazonConnectionSpec) HasSecretKey() bool {
	if o != nil && o.SecretKey != nil {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *AmazonConnectionSpec) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetRegionId returns the RegionId field value
func (o *AmazonConnectionSpec) GetRegionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value
// and a boolean to check if the value has been set.
func (o *AmazonConnectionSpec) GetRegionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RegionId, true
}

// SetRegionId sets field value
func (o *AmazonConnectionSpec) SetRegionId(v string) {
	o.RegionId = v
}

func (o AmazonConnectionSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AmazonAccountId != nil {
		toSerialize["AmazonAccountId"] = o.AmazonAccountId
	}
	if o.AccessKey != nil {
		toSerialize["accessKey"] = o.AccessKey
	}
	if o.SecretKey != nil {
		toSerialize["secretKey"] = o.SecretKey
	}
	if true {
		toSerialize["RegionId"] = o.RegionId
	}
	return json.Marshal(toSerialize)
}

type NullableAmazonConnectionSpec struct {
	value *AmazonConnectionSpec
	isSet bool
}

func (v NullableAmazonConnectionSpec) Get() *AmazonConnectionSpec {
	return v.value
}

func (v *NullableAmazonConnectionSpec) Set(val *AmazonConnectionSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonConnectionSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonConnectionSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonConnectionSpec(val *AmazonConnectionSpec) *NullableAmazonConnectionSpec {
	return &NullableAmazonConnectionSpec{value: val, isSet: true}
}

func (v NullableAmazonConnectionSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonConnectionSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


