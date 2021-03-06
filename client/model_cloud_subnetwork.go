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

// CloudSubnetwork struct for CloudSubnetwork
type CloudSubnetwork struct {
	ResourceId string `json:"resourceId"`
	Name string `json:"name"`
	CloudNetworkResourceId string `json:"cloudNetworkResourceId"`
	CidrBlock string `json:"cidrBlock"`
	AvailabilityZone *string `json:"availabilityZone,omitempty"`
}

// NewCloudSubnetwork instantiates a new CloudSubnetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudSubnetwork(resourceId string, name string, cloudNetworkResourceId string, cidrBlock string) *CloudSubnetwork {
	this := CloudSubnetwork{}
	this.ResourceId = resourceId
	this.Name = name
	this.CloudNetworkResourceId = cloudNetworkResourceId
	this.CidrBlock = cidrBlock
	return &this
}

// NewCloudSubnetworkWithDefaults instantiates a new CloudSubnetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudSubnetworkWithDefaults() *CloudSubnetwork {
	this := CloudSubnetwork{}
	return &this
}

// GetResourceId returns the ResourceId field value
func (o *CloudSubnetwork) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *CloudSubnetwork) GetResourceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *CloudSubnetwork) SetResourceId(v string) {
	o.ResourceId = v
}

// GetName returns the Name field value
func (o *CloudSubnetwork) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CloudSubnetwork) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CloudSubnetwork) SetName(v string) {
	o.Name = v
}

// GetCloudNetworkResourceId returns the CloudNetworkResourceId field value
func (o *CloudSubnetwork) GetCloudNetworkResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudNetworkResourceId
}

// GetCloudNetworkResourceIdOk returns a tuple with the CloudNetworkResourceId field value
// and a boolean to check if the value has been set.
func (o *CloudSubnetwork) GetCloudNetworkResourceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CloudNetworkResourceId, true
}

// SetCloudNetworkResourceId sets field value
func (o *CloudSubnetwork) SetCloudNetworkResourceId(v string) {
	o.CloudNetworkResourceId = v
}

// GetCidrBlock returns the CidrBlock field value
func (o *CloudSubnetwork) GetCidrBlock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CidrBlock
}

// GetCidrBlockOk returns a tuple with the CidrBlock field value
// and a boolean to check if the value has been set.
func (o *CloudSubnetwork) GetCidrBlockOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CidrBlock, true
}

// SetCidrBlock sets field value
func (o *CloudSubnetwork) SetCidrBlock(v string) {
	o.CidrBlock = v
}

// GetAvailabilityZone returns the AvailabilityZone field value if set, zero value otherwise.
func (o *CloudSubnetwork) GetAvailabilityZone() string {
	if o == nil || o.AvailabilityZone == nil {
		var ret string
		return ret
	}
	return *o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSubnetwork) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil || o.AvailabilityZone == nil {
		return nil, false
	}
	return o.AvailabilityZone, true
}

// HasAvailabilityZone returns a boolean if a field has been set.
func (o *CloudSubnetwork) HasAvailabilityZone() bool {
	if o != nil && o.AvailabilityZone != nil {
		return true
	}

	return false
}

// SetAvailabilityZone gets a reference to the given string and assigns it to the AvailabilityZone field.
func (o *CloudSubnetwork) SetAvailabilityZone(v string) {
	o.AvailabilityZone = &v
}

func (o CloudSubnetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resourceId"] = o.ResourceId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["cloudNetworkResourceId"] = o.CloudNetworkResourceId
	}
	if true {
		toSerialize["cidrBlock"] = o.CidrBlock
	}
	if o.AvailabilityZone != nil {
		toSerialize["availabilityZone"] = o.AvailabilityZone
	}
	return json.Marshal(toSerialize)
}

type NullableCloudSubnetwork struct {
	value *CloudSubnetwork
	isSet bool
}

func (v NullableCloudSubnetwork) Get() *CloudSubnetwork {
	return v.value
}

func (v *NullableCloudSubnetwork) Set(val *CloudSubnetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudSubnetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudSubnetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudSubnetwork(val *CloudSubnetwork) *NullableCloudSubnetwork {
	return &NullableCloudSubnetwork{value: val, isSet: true}
}

func (v NullableCloudSubnetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudSubnetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


