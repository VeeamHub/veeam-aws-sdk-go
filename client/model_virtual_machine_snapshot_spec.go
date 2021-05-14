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

// VirtualMachineSnapshotSpec struct for VirtualMachineSnapshotSpec
type VirtualMachineSnapshotSpec struct {
	AmazonAccountId string `json:"amazonAccountId"`
	TargetAmazonAccountIdToCopySnapshots *string `json:"targetAmazonAccountIdToCopySnapshots,omitempty"`
	TargetRegionIdToCopySnapshots *string `json:"targetRegionIdToCopySnapshots,omitempty"`
	EncryptionKey *string `json:"encryptionKey,omitempty"`
	AdditionalTags *[]TagSpec `json:"additionalTags,omitempty"`
	CopyTagsFromVolumeEnabled *bool `json:"copyTagsFromVolumeEnabled,omitempty"`
}

// NewVirtualMachineSnapshotSpec instantiates a new VirtualMachineSnapshotSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualMachineSnapshotSpec(amazonAccountId string) *VirtualMachineSnapshotSpec {
	this := VirtualMachineSnapshotSpec{}
	this.AmazonAccountId = amazonAccountId
	return &this
}

// NewVirtualMachineSnapshotSpecWithDefaults instantiates a new VirtualMachineSnapshotSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualMachineSnapshotSpecWithDefaults() *VirtualMachineSnapshotSpec {
	this := VirtualMachineSnapshotSpec{}
	return &this
}

// GetAmazonAccountId returns the AmazonAccountId field value
func (o *VirtualMachineSnapshotSpec) GetAmazonAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmazonAccountId
}

// GetAmazonAccountIdOk returns a tuple with the AmazonAccountId field value
// and a boolean to check if the value has been set.
func (o *VirtualMachineSnapshotSpec) GetAmazonAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AmazonAccountId, true
}

// SetAmazonAccountId sets field value
func (o *VirtualMachineSnapshotSpec) SetAmazonAccountId(v string) {
	o.AmazonAccountId = v
}

// GetTargetAmazonAccountIdToCopySnapshots returns the TargetAmazonAccountIdToCopySnapshots field value if set, zero value otherwise.
func (o *VirtualMachineSnapshotSpec) GetTargetAmazonAccountIdToCopySnapshots() string {
	if o == nil || o.TargetAmazonAccountIdToCopySnapshots == nil {
		var ret string
		return ret
	}
	return *o.TargetAmazonAccountIdToCopySnapshots
}

// GetTargetAmazonAccountIdToCopySnapshotsOk returns a tuple with the TargetAmazonAccountIdToCopySnapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineSnapshotSpec) GetTargetAmazonAccountIdToCopySnapshotsOk() (*string, bool) {
	if o == nil || o.TargetAmazonAccountIdToCopySnapshots == nil {
		return nil, false
	}
	return o.TargetAmazonAccountIdToCopySnapshots, true
}

// HasTargetAmazonAccountIdToCopySnapshots returns a boolean if a field has been set.
func (o *VirtualMachineSnapshotSpec) HasTargetAmazonAccountIdToCopySnapshots() bool {
	if o != nil && o.TargetAmazonAccountIdToCopySnapshots != nil {
		return true
	}

	return false
}

// SetTargetAmazonAccountIdToCopySnapshots gets a reference to the given string and assigns it to the TargetAmazonAccountIdToCopySnapshots field.
func (o *VirtualMachineSnapshotSpec) SetTargetAmazonAccountIdToCopySnapshots(v string) {
	o.TargetAmazonAccountIdToCopySnapshots = &v
}

// GetTargetRegionIdToCopySnapshots returns the TargetRegionIdToCopySnapshots field value if set, zero value otherwise.
func (o *VirtualMachineSnapshotSpec) GetTargetRegionIdToCopySnapshots() string {
	if o == nil || o.TargetRegionIdToCopySnapshots == nil {
		var ret string
		return ret
	}
	return *o.TargetRegionIdToCopySnapshots
}

// GetTargetRegionIdToCopySnapshotsOk returns a tuple with the TargetRegionIdToCopySnapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineSnapshotSpec) GetTargetRegionIdToCopySnapshotsOk() (*string, bool) {
	if o == nil || o.TargetRegionIdToCopySnapshots == nil {
		return nil, false
	}
	return o.TargetRegionIdToCopySnapshots, true
}

// HasTargetRegionIdToCopySnapshots returns a boolean if a field has been set.
func (o *VirtualMachineSnapshotSpec) HasTargetRegionIdToCopySnapshots() bool {
	if o != nil && o.TargetRegionIdToCopySnapshots != nil {
		return true
	}

	return false
}

// SetTargetRegionIdToCopySnapshots gets a reference to the given string and assigns it to the TargetRegionIdToCopySnapshots field.
func (o *VirtualMachineSnapshotSpec) SetTargetRegionIdToCopySnapshots(v string) {
	o.TargetRegionIdToCopySnapshots = &v
}

// GetEncryptionKey returns the EncryptionKey field value if set, zero value otherwise.
func (o *VirtualMachineSnapshotSpec) GetEncryptionKey() string {
	if o == nil || o.EncryptionKey == nil {
		var ret string
		return ret
	}
	return *o.EncryptionKey
}

// GetEncryptionKeyOk returns a tuple with the EncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineSnapshotSpec) GetEncryptionKeyOk() (*string, bool) {
	if o == nil || o.EncryptionKey == nil {
		return nil, false
	}
	return o.EncryptionKey, true
}

// HasEncryptionKey returns a boolean if a field has been set.
func (o *VirtualMachineSnapshotSpec) HasEncryptionKey() bool {
	if o != nil && o.EncryptionKey != nil {
		return true
	}

	return false
}

// SetEncryptionKey gets a reference to the given string and assigns it to the EncryptionKey field.
func (o *VirtualMachineSnapshotSpec) SetEncryptionKey(v string) {
	o.EncryptionKey = &v
}

// GetAdditionalTags returns the AdditionalTags field value if set, zero value otherwise.
func (o *VirtualMachineSnapshotSpec) GetAdditionalTags() []TagSpec {
	if o == nil || o.AdditionalTags == nil {
		var ret []TagSpec
		return ret
	}
	return *o.AdditionalTags
}

// GetAdditionalTagsOk returns a tuple with the AdditionalTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineSnapshotSpec) GetAdditionalTagsOk() (*[]TagSpec, bool) {
	if o == nil || o.AdditionalTags == nil {
		return nil, false
	}
	return o.AdditionalTags, true
}

// HasAdditionalTags returns a boolean if a field has been set.
func (o *VirtualMachineSnapshotSpec) HasAdditionalTags() bool {
	if o != nil && o.AdditionalTags != nil {
		return true
	}

	return false
}

// SetAdditionalTags gets a reference to the given []TagSpec and assigns it to the AdditionalTags field.
func (o *VirtualMachineSnapshotSpec) SetAdditionalTags(v []TagSpec) {
	o.AdditionalTags = &v
}

// GetCopyTagsFromVolumeEnabled returns the CopyTagsFromVolumeEnabled field value if set, zero value otherwise.
func (o *VirtualMachineSnapshotSpec) GetCopyTagsFromVolumeEnabled() bool {
	if o == nil || o.CopyTagsFromVolumeEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CopyTagsFromVolumeEnabled
}

// GetCopyTagsFromVolumeEnabledOk returns a tuple with the CopyTagsFromVolumeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineSnapshotSpec) GetCopyTagsFromVolumeEnabledOk() (*bool, bool) {
	if o == nil || o.CopyTagsFromVolumeEnabled == nil {
		return nil, false
	}
	return o.CopyTagsFromVolumeEnabled, true
}

// HasCopyTagsFromVolumeEnabled returns a boolean if a field has been set.
func (o *VirtualMachineSnapshotSpec) HasCopyTagsFromVolumeEnabled() bool {
	if o != nil && o.CopyTagsFromVolumeEnabled != nil {
		return true
	}

	return false
}

// SetCopyTagsFromVolumeEnabled gets a reference to the given bool and assigns it to the CopyTagsFromVolumeEnabled field.
func (o *VirtualMachineSnapshotSpec) SetCopyTagsFromVolumeEnabled(v bool) {
	o.CopyTagsFromVolumeEnabled = &v
}

func (o VirtualMachineSnapshotSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amazonAccountId"] = o.AmazonAccountId
	}
	if o.TargetAmazonAccountIdToCopySnapshots != nil {
		toSerialize["targetAmazonAccountIdToCopySnapshots"] = o.TargetAmazonAccountIdToCopySnapshots
	}
	if o.TargetRegionIdToCopySnapshots != nil {
		toSerialize["targetRegionIdToCopySnapshots"] = o.TargetRegionIdToCopySnapshots
	}
	if o.EncryptionKey != nil {
		toSerialize["encryptionKey"] = o.EncryptionKey
	}
	if o.AdditionalTags != nil {
		toSerialize["additionalTags"] = o.AdditionalTags
	}
	if o.CopyTagsFromVolumeEnabled != nil {
		toSerialize["copyTagsFromVolumeEnabled"] = o.CopyTagsFromVolumeEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualMachineSnapshotSpec struct {
	value *VirtualMachineSnapshotSpec
	isSet bool
}

func (v NullableVirtualMachineSnapshotSpec) Get() *VirtualMachineSnapshotSpec {
	return v.value
}

func (v *NullableVirtualMachineSnapshotSpec) Set(val *VirtualMachineSnapshotSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualMachineSnapshotSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualMachineSnapshotSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualMachineSnapshotSpec(val *VirtualMachineSnapshotSpec) *NullableVirtualMachineSnapshotSpec {
	return &NullableVirtualMachineSnapshotSpec{value: val, isSet: true}
}

func (v NullableVirtualMachineSnapshotSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualMachineSnapshotSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

