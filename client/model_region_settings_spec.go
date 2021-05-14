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

// RegionSettingsSpec struct for RegionSettingsSpec
type RegionSettingsSpec struct {
	AvailableZoneId string `json:"availableZoneId"`
	CloudNetworkId string `json:"cloudNetworkId"`
	CloudSubnetworkId string `json:"cloudSubnetworkId"`
	CloudSecurityGroupId string `json:"cloudSecurityGroupId"`
}

// NewRegionSettingsSpec instantiates a new RegionSettingsSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionSettingsSpec(availableZoneId string, cloudNetworkId string, cloudSubnetworkId string, cloudSecurityGroupId string) *RegionSettingsSpec {
	this := RegionSettingsSpec{}
	this.AvailableZoneId = availableZoneId
	this.CloudNetworkId = cloudNetworkId
	this.CloudSubnetworkId = cloudSubnetworkId
	this.CloudSecurityGroupId = cloudSecurityGroupId
	return &this
}

// NewRegionSettingsSpecWithDefaults instantiates a new RegionSettingsSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionSettingsSpecWithDefaults() *RegionSettingsSpec {
	this := RegionSettingsSpec{}
	return &this
}

// GetAvailableZoneId returns the AvailableZoneId field value
func (o *RegionSettingsSpec) GetAvailableZoneId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AvailableZoneId
}

// GetAvailableZoneIdOk returns a tuple with the AvailableZoneId field value
// and a boolean to check if the value has been set.
func (o *RegionSettingsSpec) GetAvailableZoneIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AvailableZoneId, true
}

// SetAvailableZoneId sets field value
func (o *RegionSettingsSpec) SetAvailableZoneId(v string) {
	o.AvailableZoneId = v
}

// GetCloudNetworkId returns the CloudNetworkId field value
func (o *RegionSettingsSpec) GetCloudNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudNetworkId
}

// GetCloudNetworkIdOk returns a tuple with the CloudNetworkId field value
// and a boolean to check if the value has been set.
func (o *RegionSettingsSpec) GetCloudNetworkIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CloudNetworkId, true
}

// SetCloudNetworkId sets field value
func (o *RegionSettingsSpec) SetCloudNetworkId(v string) {
	o.CloudNetworkId = v
}

// GetCloudSubnetworkId returns the CloudSubnetworkId field value
func (o *RegionSettingsSpec) GetCloudSubnetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudSubnetworkId
}

// GetCloudSubnetworkIdOk returns a tuple with the CloudSubnetworkId field value
// and a boolean to check if the value has been set.
func (o *RegionSettingsSpec) GetCloudSubnetworkIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CloudSubnetworkId, true
}

// SetCloudSubnetworkId sets field value
func (o *RegionSettingsSpec) SetCloudSubnetworkId(v string) {
	o.CloudSubnetworkId = v
}

// GetCloudSecurityGroupId returns the CloudSecurityGroupId field value
func (o *RegionSettingsSpec) GetCloudSecurityGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudSecurityGroupId
}

// GetCloudSecurityGroupIdOk returns a tuple with the CloudSecurityGroupId field value
// and a boolean to check if the value has been set.
func (o *RegionSettingsSpec) GetCloudSecurityGroupIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CloudSecurityGroupId, true
}

// SetCloudSecurityGroupId sets field value
func (o *RegionSettingsSpec) SetCloudSecurityGroupId(v string) {
	o.CloudSecurityGroupId = v
}

func (o RegionSettingsSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["availableZoneId"] = o.AvailableZoneId
	}
	if true {
		toSerialize["cloudNetworkId"] = o.CloudNetworkId
	}
	if true {
		toSerialize["cloudSubnetworkId"] = o.CloudSubnetworkId
	}
	if true {
		toSerialize["cloudSecurityGroupId"] = o.CloudSecurityGroupId
	}
	return json.Marshal(toSerialize)
}

type NullableRegionSettingsSpec struct {
	value *RegionSettingsSpec
	isSet bool
}

func (v NullableRegionSettingsSpec) Get() *RegionSettingsSpec {
	return v.value
}

func (v *NullableRegionSettingsSpec) Set(val *RegionSettingsSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionSettingsSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionSettingsSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionSettingsSpec(val *RegionSettingsSpec) *NullableRegionSettingsSpec {
	return &NullableRegionSettingsSpec{value: val, isSet: true}
}

func (v NullableRegionSettingsSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionSettingsSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

