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

// LicenseNotifications struct for LicenseNotifications
type LicenseNotifications struct {
	Results []LicenseNotification `json:"results"`
	Links *[]Link `json:"_links,omitempty"`
}

// NewLicenseNotifications instantiates a new LicenseNotifications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseNotifications(results []LicenseNotification) *LicenseNotifications {
	this := LicenseNotifications{}
	this.Results = results
	return &this
}

// NewLicenseNotificationsWithDefaults instantiates a new LicenseNotifications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseNotificationsWithDefaults() *LicenseNotifications {
	this := LicenseNotifications{}
	return &this
}

// GetResults returns the Results field value
func (o *LicenseNotifications) GetResults() []LicenseNotification {
	if o == nil {
		var ret []LicenseNotification
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *LicenseNotifications) GetResultsOk() (*[]LicenseNotification, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *LicenseNotifications) SetResults(v []LicenseNotification) {
	o.Results = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *LicenseNotifications) GetLinks() []Link {
	if o == nil || o.Links == nil {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseNotifications) GetLinksOk() (*[]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *LicenseNotifications) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *LicenseNotifications) SetLinks(v []Link) {
	o.Links = &v
}

func (o LicenseNotifications) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["results"] = o.Results
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableLicenseNotifications struct {
	value *LicenseNotifications
	isSet bool
}

func (v NullableLicenseNotifications) Get() *LicenseNotifications {
	return v.value
}

func (v *NullableLicenseNotifications) Set(val *LicenseNotifications) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseNotifications) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseNotifications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseNotifications(val *LicenseNotifications) *NullableLicenseNotifications {
	return &NullableLicenseNotifications{value: val, isSet: true}
}

func (v NullableLicenseNotifications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseNotifications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


