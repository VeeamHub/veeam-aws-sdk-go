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

// RegionSettings struct for RegionSettings
type RegionSettings struct {
	RegionId string `json:"regionId"`
	RegionName *string `json:"regionName,omitempty"`
	ParentReiognId *string `json:"parentReiognId,omitempty"`
	ParentRegionName *string `json:"parentRegionName,omitempty"`
	CloudNetworkId string `json:"cloudNetworkId"`
	CloudNetworkName *string `json:"cloudNetworkName,omitempty"`
	CloudSubnetworkId string `json:"cloudSubnetworkId"`
	CloudSubnetworkName *string `json:"cloudSubnetworkName,omitempty"`
	CloudSecurityGroupId string `json:"cloudSecurityGroupId"`
	CloudSecurityGroupName *string `json:"cloudSecurityGroupName,omitempty"`
	Links *[]Link `json:"_links,omitempty"`
}

// NewRegionSettings instantiates a new RegionSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionSettings(regionId string, cloudNetworkId string, cloudSubnetworkId string, cloudSecurityGroupId string) *RegionSettings {
	this := RegionSettings{}
	this.RegionId = regionId
	this.CloudNetworkId = cloudNetworkId
	this.CloudSubnetworkId = cloudSubnetworkId
	this.CloudSecurityGroupId = cloudSecurityGroupId
	return &this
}

// NewRegionSettingsWithDefaults instantiates a new RegionSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionSettingsWithDefaults() *RegionSettings {
	this := RegionSettings{}
	return &this
}

// GetRegionId returns the RegionId field value
func (o *RegionSettings) GetRegionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value
// and a boolean to check if the value has been set.
func (o *RegionSettings) GetRegionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RegionId, true
}

// SetRegionId sets field value
func (o *RegionSettings) SetRegionId(v string) {
	o.RegionId = v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *RegionSettings) GetRegionName() string {
	if o == nil || o.RegionName == nil {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionSettings) GetRegionNameOk() (*string, bool) {
	if o == nil || o.RegionName == nil {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *RegionSettings) HasRegionName() bool {
	if o != nil && o.RegionName != nil {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *RegionSettings) SetRegionName(v string) {
	o.RegionName = &v
}

// GetParentReiognId returns the ParentReiognId field value if set, zero value otherwise.
func (o *RegionSettings) GetParentReiognId() string {
	if o == nil || o.ParentReiognId == nil {
		var ret string
		return ret
	}
	return *o.ParentReiognId
}

// GetParentReiognIdOk returns a tuple with the ParentReiognId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionSettings) GetParentReiognIdOk() (*string, bool) {
	if o == nil || o.ParentReiognId == nil {
		return nil, false
	}
	return o.ParentReiognId, true
}

// HasParentReiognId returns a boolean if a field has been set.
func (o *RegionSettings) HasParentReiognId() bool {
	if o != nil && o.ParentReiognId != nil {
		return true
	}

	return false
}

// SetParentReiognId gets a reference to the given string and assigns it to the ParentReiognId field.
func (o *RegionSettings) SetParentReiognId(v string) {
	o.ParentReiognId = &v
}

// GetParentRegionName returns the ParentRegionName field value if set, zero value otherwise.
func (o *RegionSettings) GetParentRegionName() string {
	if o == nil || o.ParentRegionName == nil {
		var ret string
		return ret
	}
	return *o.ParentRegionName
}

// GetParentRegionNameOk returns a tuple with the ParentRegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionSettings) GetParentRegionNameOk() (*string, bool) {
	if o == nil || o.ParentRegionName == nil {
		return nil, false
	}
	return o.ParentRegionName, true
}

// HasParentRegionName returns a boolean if a field has been set.
func (o *RegionSettings) HasParentRegionName() bool {
	if o != nil && o.ParentRegionName != nil {
		return true
	}

	return false
}

// SetParentRegionName gets a reference to the given string and assigns it to the ParentRegionName field.
func (o *RegionSettings) SetParentRegionName(v string) {
	o.ParentRegionName = &v
}

// GetCloudNetworkId returns the CloudNetworkId field value
func (o *RegionSettings) GetCloudNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudNetworkId
}

// GetCloudNetworkIdOk returns a tuple with the CloudNetworkId field value
// and a boolean to check if the value has been set.
func (o *RegionSettings) GetCloudNetworkIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CloudNetworkId, true
}

// SetCloudNetworkId sets field value
func (o *RegionSettings) SetCloudNetworkId(v string) {
	o.CloudNetworkId = v
}

// GetCloudNetworkName returns the CloudNetworkName field value if set, zero value otherwise.
func (o *RegionSettings) GetCloudNetworkName() string {
	if o == nil || o.CloudNetworkName == nil {
		var ret string
		return ret
	}
	return *o.CloudNetworkName
}

// GetCloudNetworkNameOk returns a tuple with the CloudNetworkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionSettings) GetCloudNetworkNameOk() (*string, bool) {
	if o == nil || o.CloudNetworkName == nil {
		return nil, false
	}
	return o.CloudNetworkName, true
}

// HasCloudNetworkName returns a boolean if a field has been set.
func (o *RegionSettings) HasCloudNetworkName() bool {
	if o != nil && o.CloudNetworkName != nil {
		return true
	}

	return false
}

// SetCloudNetworkName gets a reference to the given string and assigns it to the CloudNetworkName field.
func (o *RegionSettings) SetCloudNetworkName(v string) {
	o.CloudNetworkName = &v
}

// GetCloudSubnetworkId returns the CloudSubnetworkId field value
func (o *RegionSettings) GetCloudSubnetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudSubnetworkId
}

// GetCloudSubnetworkIdOk returns a tuple with the CloudSubnetworkId field value
// and a boolean to check if the value has been set.
func (o *RegionSettings) GetCloudSubnetworkIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CloudSubnetworkId, true
}

// SetCloudSubnetworkId sets field value
func (o *RegionSettings) SetCloudSubnetworkId(v string) {
	o.CloudSubnetworkId = v
}

// GetCloudSubnetworkName returns the CloudSubnetworkName field value if set, zero value otherwise.
func (o *RegionSettings) GetCloudSubnetworkName() string {
	if o == nil || o.CloudSubnetworkName == nil {
		var ret string
		return ret
	}
	return *o.CloudSubnetworkName
}

// GetCloudSubnetworkNameOk returns a tuple with the CloudSubnetworkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionSettings) GetCloudSubnetworkNameOk() (*string, bool) {
	if o == nil || o.CloudSubnetworkName == nil {
		return nil, false
	}
	return o.CloudSubnetworkName, true
}

// HasCloudSubnetworkName returns a boolean if a field has been set.
func (o *RegionSettings) HasCloudSubnetworkName() bool {
	if o != nil && o.CloudSubnetworkName != nil {
		return true
	}

	return false
}

// SetCloudSubnetworkName gets a reference to the given string and assigns it to the CloudSubnetworkName field.
func (o *RegionSettings) SetCloudSubnetworkName(v string) {
	o.CloudSubnetworkName = &v
}

// GetCloudSecurityGroupId returns the CloudSecurityGroupId field value
func (o *RegionSettings) GetCloudSecurityGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudSecurityGroupId
}

// GetCloudSecurityGroupIdOk returns a tuple with the CloudSecurityGroupId field value
// and a boolean to check if the value has been set.
func (o *RegionSettings) GetCloudSecurityGroupIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CloudSecurityGroupId, true
}

// SetCloudSecurityGroupId sets field value
func (o *RegionSettings) SetCloudSecurityGroupId(v string) {
	o.CloudSecurityGroupId = v
}

// GetCloudSecurityGroupName returns the CloudSecurityGroupName field value if set, zero value otherwise.
func (o *RegionSettings) GetCloudSecurityGroupName() string {
	if o == nil || o.CloudSecurityGroupName == nil {
		var ret string
		return ret
	}
	return *o.CloudSecurityGroupName
}

// GetCloudSecurityGroupNameOk returns a tuple with the CloudSecurityGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionSettings) GetCloudSecurityGroupNameOk() (*string, bool) {
	if o == nil || o.CloudSecurityGroupName == nil {
		return nil, false
	}
	return o.CloudSecurityGroupName, true
}

// HasCloudSecurityGroupName returns a boolean if a field has been set.
func (o *RegionSettings) HasCloudSecurityGroupName() bool {
	if o != nil && o.CloudSecurityGroupName != nil {
		return true
	}

	return false
}

// SetCloudSecurityGroupName gets a reference to the given string and assigns it to the CloudSecurityGroupName field.
func (o *RegionSettings) SetCloudSecurityGroupName(v string) {
	o.CloudSecurityGroupName = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RegionSettings) GetLinks() []Link {
	if o == nil || o.Links == nil {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionSettings) GetLinksOk() (*[]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RegionSettings) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *RegionSettings) SetLinks(v []Link) {
	o.Links = &v
}

func (o RegionSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["regionId"] = o.RegionId
	}
	if o.RegionName != nil {
		toSerialize["regionName"] = o.RegionName
	}
	if o.ParentReiognId != nil {
		toSerialize["parentReiognId"] = o.ParentReiognId
	}
	if o.ParentRegionName != nil {
		toSerialize["parentRegionName"] = o.ParentRegionName
	}
	if true {
		toSerialize["cloudNetworkId"] = o.CloudNetworkId
	}
	if o.CloudNetworkName != nil {
		toSerialize["cloudNetworkName"] = o.CloudNetworkName
	}
	if true {
		toSerialize["cloudSubnetworkId"] = o.CloudSubnetworkId
	}
	if o.CloudSubnetworkName != nil {
		toSerialize["cloudSubnetworkName"] = o.CloudSubnetworkName
	}
	if true {
		toSerialize["cloudSecurityGroupId"] = o.CloudSecurityGroupId
	}
	if o.CloudSecurityGroupName != nil {
		toSerialize["cloudSecurityGroupName"] = o.CloudSecurityGroupName
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableRegionSettings struct {
	value *RegionSettings
	isSet bool
}

func (v NullableRegionSettings) Get() *RegionSettings {
	return v.value
}

func (v *NullableRegionSettings) Set(val *RegionSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionSettings(val *RegionSettings) *NullableRegionSettings {
	return &NullableRegionSettings{value: val, isSet: true}
}

func (v NullableRegionSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


