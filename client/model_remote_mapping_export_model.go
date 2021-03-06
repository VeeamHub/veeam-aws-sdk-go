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

// RemoteMappingExportModel struct for RemoteMappingExportModel
type RemoteMappingExportModel struct {
	SourceRegionName string `json:"sourceRegionName"`
	TargetRegionName string `json:"targetRegionName"`
	TargetAmazonAccountName string `json:"targetAmazonAccountName"`
	EncryptionKey *string `json:"encryptionKey,omitempty"`
}

// NewRemoteMappingExportModel instantiates a new RemoteMappingExportModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteMappingExportModel(sourceRegionName string, targetRegionName string, targetAmazonAccountName string) *RemoteMappingExportModel {
	this := RemoteMappingExportModel{}
	this.SourceRegionName = sourceRegionName
	this.TargetRegionName = targetRegionName
	this.TargetAmazonAccountName = targetAmazonAccountName
	return &this
}

// NewRemoteMappingExportModelWithDefaults instantiates a new RemoteMappingExportModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteMappingExportModelWithDefaults() *RemoteMappingExportModel {
	this := RemoteMappingExportModel{}
	return &this
}

// GetSourceRegionName returns the SourceRegionName field value
func (o *RemoteMappingExportModel) GetSourceRegionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceRegionName
}

// GetSourceRegionNameOk returns a tuple with the SourceRegionName field value
// and a boolean to check if the value has been set.
func (o *RemoteMappingExportModel) GetSourceRegionNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SourceRegionName, true
}

// SetSourceRegionName sets field value
func (o *RemoteMappingExportModel) SetSourceRegionName(v string) {
	o.SourceRegionName = v
}

// GetTargetRegionName returns the TargetRegionName field value
func (o *RemoteMappingExportModel) GetTargetRegionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetRegionName
}

// GetTargetRegionNameOk returns a tuple with the TargetRegionName field value
// and a boolean to check if the value has been set.
func (o *RemoteMappingExportModel) GetTargetRegionNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetRegionName, true
}

// SetTargetRegionName sets field value
func (o *RemoteMappingExportModel) SetTargetRegionName(v string) {
	o.TargetRegionName = v
}

// GetTargetAmazonAccountName returns the TargetAmazonAccountName field value
func (o *RemoteMappingExportModel) GetTargetAmazonAccountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetAmazonAccountName
}

// GetTargetAmazonAccountNameOk returns a tuple with the TargetAmazonAccountName field value
// and a boolean to check if the value has been set.
func (o *RemoteMappingExportModel) GetTargetAmazonAccountNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetAmazonAccountName, true
}

// SetTargetAmazonAccountName sets field value
func (o *RemoteMappingExportModel) SetTargetAmazonAccountName(v string) {
	o.TargetAmazonAccountName = v
}

// GetEncryptionKey returns the EncryptionKey field value if set, zero value otherwise.
func (o *RemoteMappingExportModel) GetEncryptionKey() string {
	if o == nil || o.EncryptionKey == nil {
		var ret string
		return ret
	}
	return *o.EncryptionKey
}

// GetEncryptionKeyOk returns a tuple with the EncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteMappingExportModel) GetEncryptionKeyOk() (*string, bool) {
	if o == nil || o.EncryptionKey == nil {
		return nil, false
	}
	return o.EncryptionKey, true
}

// HasEncryptionKey returns a boolean if a field has been set.
func (o *RemoteMappingExportModel) HasEncryptionKey() bool {
	if o != nil && o.EncryptionKey != nil {
		return true
	}

	return false
}

// SetEncryptionKey gets a reference to the given string and assigns it to the EncryptionKey field.
func (o *RemoteMappingExportModel) SetEncryptionKey(v string) {
	o.EncryptionKey = &v
}

func (o RemoteMappingExportModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceRegionName"] = o.SourceRegionName
	}
	if true {
		toSerialize["targetRegionName"] = o.TargetRegionName
	}
	if true {
		toSerialize["targetAmazonAccountName"] = o.TargetAmazonAccountName
	}
	if o.EncryptionKey != nil {
		toSerialize["encryptionKey"] = o.EncryptionKey
	}
	return json.Marshal(toSerialize)
}

type NullableRemoteMappingExportModel struct {
	value *RemoteMappingExportModel
	isSet bool
}

func (v NullableRemoteMappingExportModel) Get() *RemoteMappingExportModel {
	return v.value
}

func (v *NullableRemoteMappingExportModel) Set(val *RemoteMappingExportModel) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteMappingExportModel) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteMappingExportModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteMappingExportModel(val *RemoteMappingExportModel) *NullableRemoteMappingExportModel {
	return &NullableRemoteMappingExportModel{value: val, isSet: true}
}

func (v NullableRemoteMappingExportModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteMappingExportModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


