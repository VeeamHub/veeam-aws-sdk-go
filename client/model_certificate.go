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

// Certificate struct for Certificate
type Certificate struct {
	// certificate thumpbprint
	Thumbprint   string `json:"thumbprint"`
	SerialNumber string `json:"serialNumber"`
	// certificate key algorithm
	KeyAlgorithm string `json:"keyAlgorithm"`
	// certificate key size
	KeySize string `json:"keySize"`
	// certificate subject
	Subject string `json:"subject"`
	// certificate acquirer
	IssuedTo string `json:"issuedTo"`
	// certificate issuer
	IssuedBy               string `json:"issuedBy"`
	ValidFrom              Time   `json:"validFrom"`
	ValidBy                Time   `json:"validBy"`
	AutomaticallyGenerated bool   `json:"automaticallyGenerated"`
}

// NewCertificate instantiates a new Certificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificate(thumbprint string, serialNumber string, keyAlgorithm string, keySize string, subject string, issuedTo string, issuedBy string, validFrom Time, validBy Time, automaticallyGenerated bool) *Certificate {
	this := Certificate{}
	this.Thumbprint = thumbprint
	this.SerialNumber = serialNumber
	this.KeyAlgorithm = keyAlgorithm
	this.KeySize = keySize
	this.Subject = subject
	this.IssuedTo = issuedTo
	this.IssuedBy = issuedBy
	this.ValidFrom = validFrom
	this.ValidBy = validBy
	this.AutomaticallyGenerated = automaticallyGenerated
	return &this
}

// NewCertificateWithDefaults instantiates a new Certificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateWithDefaults() *Certificate {
	this := Certificate{}
	return &this
}

// GetThumbprint returns the Thumbprint field value
func (o *Certificate) GetThumbprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Thumbprint
}

// GetThumbprintOk returns a tuple with the Thumbprint field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetThumbprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Thumbprint, true
}

// SetThumbprint sets field value
func (o *Certificate) SetThumbprint(v string) {
	o.Thumbprint = v
}

// GetSerialNumber returns the SerialNumber field value
func (o *Certificate) GetSerialNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetSerialNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SerialNumber, true
}

// SetSerialNumber sets field value
func (o *Certificate) SetSerialNumber(v string) {
	o.SerialNumber = v
}

// GetKeyAlgorithm returns the KeyAlgorithm field value
func (o *Certificate) GetKeyAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyAlgorithm
}

// GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetKeyAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyAlgorithm, true
}

// SetKeyAlgorithm sets field value
func (o *Certificate) SetKeyAlgorithm(v string) {
	o.KeyAlgorithm = v
}

// GetKeySize returns the KeySize field value
func (o *Certificate) GetKeySize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeySize
}

// GetKeySizeOk returns a tuple with the KeySize field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetKeySizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeySize, true
}

// SetKeySize sets field value
func (o *Certificate) SetKeySize(v string) {
	o.KeySize = v
}

// GetSubject returns the Subject field value
func (o *Certificate) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *Certificate) SetSubject(v string) {
	o.Subject = v
}

// GetIssuedTo returns the IssuedTo field value
func (o *Certificate) GetIssuedTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuedTo
}

// GetIssuedToOk returns a tuple with the IssuedTo field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetIssuedToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuedTo, true
}

// SetIssuedTo sets field value
func (o *Certificate) SetIssuedTo(v string) {
	o.IssuedTo = v
}

// GetIssuedBy returns the IssuedBy field value
func (o *Certificate) GetIssuedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuedBy
}

// GetIssuedByOk returns a tuple with the IssuedBy field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetIssuedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuedBy, true
}

// SetIssuedBy sets field value
func (o *Certificate) SetIssuedBy(v string) {
	o.IssuedBy = v
}

// GetValidFrom returns the ValidFrom field value
func (o *Certificate) GetValidFrom() Time {
	if o == nil {
		var ret Time
		return ret
	}

	return o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetValidFromOk() (*Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidFrom, true
}

// SetValidFrom sets field value
func (o *Certificate) SetValidFrom(v Time) {
	o.ValidFrom = v
}

// GetValidBy returns the ValidBy field value
func (o *Certificate) GetValidBy() Time {
	if o == nil {
		var ret Time
		return ret
	}

	return o.ValidBy
}

// GetValidByOk returns a tuple with the ValidBy field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetValidByOk() (*Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidBy, true
}

// SetValidBy sets field value
func (o *Certificate) SetValidBy(v Time) {
	o.ValidBy = v
}

// GetAutomaticallyGenerated returns the AutomaticallyGenerated field value
func (o *Certificate) GetAutomaticallyGenerated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutomaticallyGenerated
}

// GetAutomaticallyGeneratedOk returns a tuple with the AutomaticallyGenerated field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetAutomaticallyGeneratedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutomaticallyGenerated, true
}

// SetAutomaticallyGenerated sets field value
func (o *Certificate) SetAutomaticallyGenerated(v bool) {
	o.AutomaticallyGenerated = v
}

func (o Certificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["thumbprint"] = o.Thumbprint
	}
	if true {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if true {
		toSerialize["keyAlgorithm"] = o.KeyAlgorithm
	}
	if true {
		toSerialize["keySize"] = o.KeySize
	}
	if true {
		toSerialize["subject"] = o.Subject
	}
	if true {
		toSerialize["issuedTo"] = o.IssuedTo
	}
	if true {
		toSerialize["issuedBy"] = o.IssuedBy
	}
	if true {
		toSerialize["validFrom"] = o.ValidFrom
	}
	if true {
		toSerialize["validBy"] = o.ValidBy
	}
	if true {
		toSerialize["automaticallyGenerated"] = o.AutomaticallyGenerated
	}
	return json.Marshal(toSerialize)
}

type NullableCertificate struct {
	value *Certificate
	isSet bool
}

func (v NullableCertificate) Get() *Certificate {
	return v.value
}

func (v *NullableCertificate) Set(val *Certificate) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificate(val *Certificate) *NullableCertificate {
	return &NullableCertificate{value: val, isSet: true}
}

func (v NullableCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
