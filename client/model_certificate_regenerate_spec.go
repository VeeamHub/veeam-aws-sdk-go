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

// CertificateRegenerateSpec struct for CertificateRegenerateSpec
type CertificateRegenerateSpec struct {
	ValidBy *Time `json:"validBy,omitempty"`
}

// NewCertificateRegenerateSpec instantiates a new CertificateRegenerateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateRegenerateSpec() *CertificateRegenerateSpec {
	this := CertificateRegenerateSpec{}
	return &this
}

// NewCertificateRegenerateSpecWithDefaults instantiates a new CertificateRegenerateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateRegenerateSpecWithDefaults() *CertificateRegenerateSpec {
	this := CertificateRegenerateSpec{}
	return &this
}

// GetValidBy returns the ValidBy field value if set, zero value otherwise.
func (o *CertificateRegenerateSpec) GetValidBy() Time {
	if o == nil || o.ValidBy == nil {
		var ret Time
		return ret
	}
	return *o.ValidBy
}

// GetValidByOk returns a tuple with the ValidBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateRegenerateSpec) GetValidByOk() (*Time, bool) {
	if o == nil || o.ValidBy == nil {
		return nil, false
	}
	return o.ValidBy, true
}

// HasValidBy returns a boolean if a field has been set.
func (o *CertificateRegenerateSpec) HasValidBy() bool {
	if o != nil && o.ValidBy != nil {
		return true
	}

	return false
}

// SetValidBy gets a reference to the given Time and assigns it to the ValidBy field.
func (o *CertificateRegenerateSpec) SetValidBy(v Time) {
	o.ValidBy = &v
}

func (o CertificateRegenerateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ValidBy != nil {
		toSerialize["validBy"] = o.ValidBy
	}
	return json.Marshal(toSerialize)
}

type NullableCertificateRegenerateSpec struct {
	value *CertificateRegenerateSpec
	isSet bool
}

func (v NullableCertificateRegenerateSpec) Get() *CertificateRegenerateSpec {
	return v.value
}

func (v *NullableCertificateRegenerateSpec) Set(val *CertificateRegenerateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateRegenerateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateRegenerateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateRegenerateSpec(val *CertificateRegenerateSpec) *NullableCertificateRegenerateSpec {
	return &NullableCertificateRegenerateSpec{value: val, isSet: true}
}

func (v NullableCertificateRegenerateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateRegenerateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
