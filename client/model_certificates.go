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

// Certificates struct for Certificates
type Certificates struct {
	Data *[]Certificate `json:"data,omitempty"`
}

// NewCertificates instantiates a new Certificates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificates() *Certificates {
	this := Certificates{}
	return &this
}

// NewCertificatesWithDefaults instantiates a new Certificates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificatesWithDefaults() *Certificates {
	this := Certificates{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Certificates) GetData() []Certificate {
	if o == nil || o.Data == nil {
		var ret []Certificate
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificates) GetDataOk() (*[]Certificate, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Certificates) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Certificate and assigns it to the Data field.
func (o *Certificates) SetData(v []Certificate) {
	o.Data = &v
}

func (o Certificates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCertificates struct {
	value *Certificates
	isSet bool
}

func (v NullableCertificates) Get() *Certificates {
	return v.value
}

func (v *NullableCertificates) Set(val *Certificates) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificates) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificates(val *Certificates) *NullableCertificates {
	return &NullableCertificates{value: val, isSet: true}
}

func (v NullableCertificates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

