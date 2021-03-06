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

// ServiceAccountSettings struct for ServiceAccountSettings
type ServiceAccountSettings struct {
	ServiceAmazonAccountId string `json:"serviceAmazonAccountId"`
	ServiceAmazonAccountName *string `json:"serviceAmazonAccountName,omitempty"`
	ServiceAmazonAccountRegionType *RegionTypes `json:"serviceAmazonAccountRegionType,omitempty"`
	ServiceRegion *string `json:"serviceRegion,omitempty"`
	Links *[]Link `json:"_links,omitempty"`
}

// NewServiceAccountSettings instantiates a new ServiceAccountSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountSettings(serviceAmazonAccountId string) *ServiceAccountSettings {
	this := ServiceAccountSettings{}
	this.ServiceAmazonAccountId = serviceAmazonAccountId
	return &this
}

// NewServiceAccountSettingsWithDefaults instantiates a new ServiceAccountSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountSettingsWithDefaults() *ServiceAccountSettings {
	this := ServiceAccountSettings{}
	return &this
}

// GetServiceAmazonAccountId returns the ServiceAmazonAccountId field value
func (o *ServiceAccountSettings) GetServiceAmazonAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceAmazonAccountId
}

// GetServiceAmazonAccountIdOk returns a tuple with the ServiceAmazonAccountId field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountSettings) GetServiceAmazonAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServiceAmazonAccountId, true
}

// SetServiceAmazonAccountId sets field value
func (o *ServiceAccountSettings) SetServiceAmazonAccountId(v string) {
	o.ServiceAmazonAccountId = v
}

// GetServiceAmazonAccountName returns the ServiceAmazonAccountName field value if set, zero value otherwise.
func (o *ServiceAccountSettings) GetServiceAmazonAccountName() string {
	if o == nil || o.ServiceAmazonAccountName == nil {
		var ret string
		return ret
	}
	return *o.ServiceAmazonAccountName
}

// GetServiceAmazonAccountNameOk returns a tuple with the ServiceAmazonAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountSettings) GetServiceAmazonAccountNameOk() (*string, bool) {
	if o == nil || o.ServiceAmazonAccountName == nil {
		return nil, false
	}
	return o.ServiceAmazonAccountName, true
}

// HasServiceAmazonAccountName returns a boolean if a field has been set.
func (o *ServiceAccountSettings) HasServiceAmazonAccountName() bool {
	if o != nil && o.ServiceAmazonAccountName != nil {
		return true
	}

	return false
}

// SetServiceAmazonAccountName gets a reference to the given string and assigns it to the ServiceAmazonAccountName field.
func (o *ServiceAccountSettings) SetServiceAmazonAccountName(v string) {
	o.ServiceAmazonAccountName = &v
}

// GetServiceAmazonAccountRegionType returns the ServiceAmazonAccountRegionType field value if set, zero value otherwise.
func (o *ServiceAccountSettings) GetServiceAmazonAccountRegionType() RegionTypes {
	if o == nil || o.ServiceAmazonAccountRegionType == nil {
		var ret RegionTypes
		return ret
	}
	return *o.ServiceAmazonAccountRegionType
}

// GetServiceAmazonAccountRegionTypeOk returns a tuple with the ServiceAmazonAccountRegionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountSettings) GetServiceAmazonAccountRegionTypeOk() (*RegionTypes, bool) {
	if o == nil || o.ServiceAmazonAccountRegionType == nil {
		return nil, false
	}
	return o.ServiceAmazonAccountRegionType, true
}

// HasServiceAmazonAccountRegionType returns a boolean if a field has been set.
func (o *ServiceAccountSettings) HasServiceAmazonAccountRegionType() bool {
	if o != nil && o.ServiceAmazonAccountRegionType != nil {
		return true
	}

	return false
}

// SetServiceAmazonAccountRegionType gets a reference to the given RegionTypes and assigns it to the ServiceAmazonAccountRegionType field.
func (o *ServiceAccountSettings) SetServiceAmazonAccountRegionType(v RegionTypes) {
	o.ServiceAmazonAccountRegionType = &v
}

// GetServiceRegion returns the ServiceRegion field value if set, zero value otherwise.
func (o *ServiceAccountSettings) GetServiceRegion() string {
	if o == nil || o.ServiceRegion == nil {
		var ret string
		return ret
	}
	return *o.ServiceRegion
}

// GetServiceRegionOk returns a tuple with the ServiceRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountSettings) GetServiceRegionOk() (*string, bool) {
	if o == nil || o.ServiceRegion == nil {
		return nil, false
	}
	return o.ServiceRegion, true
}

// HasServiceRegion returns a boolean if a field has been set.
func (o *ServiceAccountSettings) HasServiceRegion() bool {
	if o != nil && o.ServiceRegion != nil {
		return true
	}

	return false
}

// SetServiceRegion gets a reference to the given string and assigns it to the ServiceRegion field.
func (o *ServiceAccountSettings) SetServiceRegion(v string) {
	o.ServiceRegion = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ServiceAccountSettings) GetLinks() []Link {
	if o == nil || o.Links == nil {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountSettings) GetLinksOk() (*[]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ServiceAccountSettings) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *ServiceAccountSettings) SetLinks(v []Link) {
	o.Links = &v
}

func (o ServiceAccountSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serviceAmazonAccountId"] = o.ServiceAmazonAccountId
	}
	if o.ServiceAmazonAccountName != nil {
		toSerialize["serviceAmazonAccountName"] = o.ServiceAmazonAccountName
	}
	if o.ServiceAmazonAccountRegionType != nil {
		toSerialize["serviceAmazonAccountRegionType"] = o.ServiceAmazonAccountRegionType
	}
	if o.ServiceRegion != nil {
		toSerialize["serviceRegion"] = o.ServiceRegion
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableServiceAccountSettings struct {
	value *ServiceAccountSettings
	isSet bool
}

func (v NullableServiceAccountSettings) Get() *ServiceAccountSettings {
	return v.value
}

func (v *NullableServiceAccountSettings) Set(val *ServiceAccountSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountSettings(val *ServiceAccountSettings) *NullableServiceAccountSettings {
	return &NullableServiceAccountSettings{value: val, isSet: true}
}

func (v NullableServiceAccountSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


