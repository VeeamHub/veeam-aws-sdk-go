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

// ReplicaMapping struct for ReplicaMapping
type ReplicaMapping struct {
	SourceRegionId string `json:"sourceRegionId"`
	TargetRegionId string `json:"targetRegionId"`
	TargetAmazonAccountId string `json:"targetAmazonAccountId"`
	EncryptionKey *string `json:"encryptionKey,omitempty"`
}

// NewReplicaMapping instantiates a new ReplicaMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplicaMapping(sourceRegionId string, targetRegionId string, targetAmazonAccountId string) *ReplicaMapping {
	this := ReplicaMapping{}
	this.SourceRegionId = sourceRegionId
	this.TargetRegionId = targetRegionId
	this.TargetAmazonAccountId = targetAmazonAccountId
	return &this
}

// NewReplicaMappingWithDefaults instantiates a new ReplicaMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplicaMappingWithDefaults() *ReplicaMapping {
	this := ReplicaMapping{}
	return &this
}

// GetSourceRegionId returns the SourceRegionId field value
func (o *ReplicaMapping) GetSourceRegionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceRegionId
}

// GetSourceRegionIdOk returns a tuple with the SourceRegionId field value
// and a boolean to check if the value has been set.
func (o *ReplicaMapping) GetSourceRegionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SourceRegionId, true
}

// SetSourceRegionId sets field value
func (o *ReplicaMapping) SetSourceRegionId(v string) {
	o.SourceRegionId = v
}

// GetTargetRegionId returns the TargetRegionId field value
func (o *ReplicaMapping) GetTargetRegionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetRegionId
}

// GetTargetRegionIdOk returns a tuple with the TargetRegionId field value
// and a boolean to check if the value has been set.
func (o *ReplicaMapping) GetTargetRegionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetRegionId, true
}

// SetTargetRegionId sets field value
func (o *ReplicaMapping) SetTargetRegionId(v string) {
	o.TargetRegionId = v
}

// GetTargetAmazonAccountId returns the TargetAmazonAccountId field value
func (o *ReplicaMapping) GetTargetAmazonAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetAmazonAccountId
}

// GetTargetAmazonAccountIdOk returns a tuple with the TargetAmazonAccountId field value
// and a boolean to check if the value has been set.
func (o *ReplicaMapping) GetTargetAmazonAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetAmazonAccountId, true
}

// SetTargetAmazonAccountId sets field value
func (o *ReplicaMapping) SetTargetAmazonAccountId(v string) {
	o.TargetAmazonAccountId = v
}

// GetEncryptionKey returns the EncryptionKey field value if set, zero value otherwise.
func (o *ReplicaMapping) GetEncryptionKey() string {
	if o == nil || o.EncryptionKey == nil {
		var ret string
		return ret
	}
	return *o.EncryptionKey
}

// GetEncryptionKeyOk returns a tuple with the EncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicaMapping) GetEncryptionKeyOk() (*string, bool) {
	if o == nil || o.EncryptionKey == nil {
		return nil, false
	}
	return o.EncryptionKey, true
}

// HasEncryptionKey returns a boolean if a field has been set.
func (o *ReplicaMapping) HasEncryptionKey() bool {
	if o != nil && o.EncryptionKey != nil {
		return true
	}

	return false
}

// SetEncryptionKey gets a reference to the given string and assigns it to the EncryptionKey field.
func (o *ReplicaMapping) SetEncryptionKey(v string) {
	o.EncryptionKey = &v
}

func (o ReplicaMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceRegionId"] = o.SourceRegionId
	}
	if true {
		toSerialize["targetRegionId"] = o.TargetRegionId
	}
	if true {
		toSerialize["targetAmazonAccountId"] = o.TargetAmazonAccountId
	}
	if o.EncryptionKey != nil {
		toSerialize["encryptionKey"] = o.EncryptionKey
	}
	return json.Marshal(toSerialize)
}

type NullableReplicaMapping struct {
	value *ReplicaMapping
	isSet bool
}

func (v NullableReplicaMapping) Get() *ReplicaMapping {
	return v.value
}

func (v *NullableReplicaMapping) Set(val *ReplicaMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableReplicaMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableReplicaMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplicaMapping(val *ReplicaMapping) *NullableReplicaMapping {
	return &NullableReplicaMapping{value: val, isSet: true}
}

func (v NullableReplicaMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplicaMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


