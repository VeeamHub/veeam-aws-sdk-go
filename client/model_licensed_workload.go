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

// LicensedWorkload struct for LicensedWorkload
type LicensedWorkload struct {
	Id             *string `json:"id,omitempty"`
	Name           *string `json:"name,omitempty"`
	LastBackupTime *Time   `json:"lastBackupTime,omitempty"`
	Links          *[]Link `json:"_links,omitempty"`
}

// NewLicensedWorkload instantiates a new LicensedWorkload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicensedWorkload() *LicensedWorkload {
	this := LicensedWorkload{}
	return &this
}

// NewLicensedWorkloadWithDefaults instantiates a new LicensedWorkload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicensedWorkloadWithDefaults() *LicensedWorkload {
	this := LicensedWorkload{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LicensedWorkload) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicensedWorkload) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LicensedWorkload) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LicensedWorkload) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LicensedWorkload) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicensedWorkload) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LicensedWorkload) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LicensedWorkload) SetName(v string) {
	o.Name = &v
}

// GetLastBackupTime returns the LastBackupTime field value if set, zero value otherwise.
func (o *LicensedWorkload) GetLastBackupTime() Time {
	if o == nil || o.LastBackupTime == nil {
		var ret Time
		return ret
	}
	return *o.LastBackupTime
}

// GetLastBackupTimeOk returns a tuple with the LastBackupTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicensedWorkload) GetLastBackupTimeOk() (*Time, bool) {
	if o == nil || o.LastBackupTime == nil {
		return nil, false
	}
	return o.LastBackupTime, true
}

// HasLastBackupTime returns a boolean if a field has been set.
func (o *LicensedWorkload) HasLastBackupTime() bool {
	if o != nil && o.LastBackupTime != nil {
		return true
	}

	return false
}

// SetLastBackupTime gets a reference to the given Time and assigns it to the LastBackupTime field.
func (o *LicensedWorkload) SetLastBackupTime(v Time) {
	o.LastBackupTime = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *LicensedWorkload) GetLinks() []Link {
	if o == nil || o.Links == nil {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicensedWorkload) GetLinksOk() (*[]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *LicensedWorkload) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *LicensedWorkload) SetLinks(v []Link) {
	o.Links = &v
}

func (o LicensedWorkload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.LastBackupTime != nil {
		toSerialize["lastBackupTime"] = o.LastBackupTime
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableLicensedWorkload struct {
	value *LicensedWorkload
	isSet bool
}

func (v NullableLicensedWorkload) Get() *LicensedWorkload {
	return v.value
}

func (v *NullableLicensedWorkload) Set(val *LicensedWorkload) {
	v.value = val
	v.isSet = true
}

func (v NullableLicensedWorkload) IsSet() bool {
	return v.isSet
}

func (v *NullableLicensedWorkload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicensedWorkload(val *LicensedWorkload) *NullableLicensedWorkload {
	return &NullableLicensedWorkload{value: val, isSet: true}
}

func (v NullableLicensedWorkload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicensedWorkload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
