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
	"fmt"
)

// LicenseAgreementTypes the model 'LicenseAgreementTypes'
type LicenseAgreementTypes string

// List of LicenseAgreementTypes
const (
	LICENSEAGREEMENTTYPES_EULA LicenseAgreementTypes = "Eula"
	LICENSEAGREEMENTTYPES_THIRD_PARTY_AGREEMENT LicenseAgreementTypes = "ThirdPartyAgreement"
)

func (v *LicenseAgreementTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LicenseAgreementTypes(value)
	for _, existing := range []LicenseAgreementTypes{ "Eula", "ThirdPartyAgreement",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LicenseAgreementTypes", value)
}

// Ptr returns reference to LicenseAgreementTypes value
func (v LicenseAgreementTypes) Ptr() *LicenseAgreementTypes {
	return &v
}

type NullableLicenseAgreementTypes struct {
	value *LicenseAgreementTypes
	isSet bool
}

func (v NullableLicenseAgreementTypes) Get() *LicenseAgreementTypes {
	return v.value
}

func (v *NullableLicenseAgreementTypes) Set(val *LicenseAgreementTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseAgreementTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseAgreementTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseAgreementTypes(val *LicenseAgreementTypes) *NullableLicenseAgreementTypes {
	return &NullableLicenseAgreementTypes{value: val, isSet: true}
}

func (v NullableLicenseAgreementTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseAgreementTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

