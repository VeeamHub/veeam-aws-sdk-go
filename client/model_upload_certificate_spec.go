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
	"os"
)

// UploadCertificateSpec struct for UploadCertificateSpec
type UploadCertificateSpec struct {
	// The .PFX certificate to upload.
	CertificateFile *os.File `json:"certificateFile"`
	// .PFX certificate password
	CertificatePassword string `json:"certificatePassword"`
}

// NewUploadCertificateSpec instantiates a new UploadCertificateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadCertificateSpec(certificateFile *os.File, certificatePassword string) *UploadCertificateSpec {
	this := UploadCertificateSpec{}
	this.CertificateFile = certificateFile
	this.CertificatePassword = certificatePassword
	return &this
}

// NewUploadCertificateSpecWithDefaults instantiates a new UploadCertificateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadCertificateSpecWithDefaults() *UploadCertificateSpec {
	this := UploadCertificateSpec{}
	return &this
}

// GetCertificateFile returns the CertificateFile field value
func (o *UploadCertificateSpec) GetCertificateFile() *os.File {
	if o == nil {
		var ret *os.File
		return ret
	}

	return o.CertificateFile
}

// GetCertificateFileOk returns a tuple with the CertificateFile field value
// and a boolean to check if the value has been set.
func (o *UploadCertificateSpec) GetCertificateFileOk() (**os.File, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CertificateFile, true
}

// SetCertificateFile sets field value
func (o *UploadCertificateSpec) SetCertificateFile(v *os.File) {
	o.CertificateFile = v
}

// GetCertificatePassword returns the CertificatePassword field value
func (o *UploadCertificateSpec) GetCertificatePassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CertificatePassword
}

// GetCertificatePasswordOk returns a tuple with the CertificatePassword field value
// and a boolean to check if the value has been set.
func (o *UploadCertificateSpec) GetCertificatePasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CertificatePassword, true
}

// SetCertificatePassword sets field value
func (o *UploadCertificateSpec) SetCertificatePassword(v string) {
	o.CertificatePassword = v
}

func (o UploadCertificateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["certificateFile"] = o.CertificateFile
	}
	if true {
		toSerialize["certificatePassword"] = o.CertificatePassword
	}
	return json.Marshal(toSerialize)
}

type NullableUploadCertificateSpec struct {
	value *UploadCertificateSpec
	isSet bool
}

func (v NullableUploadCertificateSpec) Get() *UploadCertificateSpec {
	return v.value
}

func (v *NullableUploadCertificateSpec) Set(val *UploadCertificateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadCertificateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadCertificateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadCertificateSpec(val *UploadCertificateSpec) *NullableUploadCertificateSpec {
	return &NullableUploadCertificateSpec{value: val, isSet: true}
}

func (v NullableUploadCertificateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadCertificateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


