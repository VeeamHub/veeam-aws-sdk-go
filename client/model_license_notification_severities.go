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

// LicenseNotificationSeverities the model 'LicenseNotificationSeverities'
type LicenseNotificationSeverities string

// List of LicenseNotificationSeverities
const (
	LICENSENOTIFICATIONSEVERITIES_INFORMATION LicenseNotificationSeverities = "Information"
	LICENSENOTIFICATIONSEVERITIES_WARNING LicenseNotificationSeverities = "Warning"
	LICENSENOTIFICATIONSEVERITIES_ERROR LicenseNotificationSeverities = "Error"
)

func (v *LicenseNotificationSeverities) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LicenseNotificationSeverities(value)
	for _, existing := range []LicenseNotificationSeverities{ "Information", "Warning", "Error",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LicenseNotificationSeverities", value)
}

// Ptr returns reference to LicenseNotificationSeverities value
func (v LicenseNotificationSeverities) Ptr() *LicenseNotificationSeverities {
	return &v
}

type NullableLicenseNotificationSeverities struct {
	value *LicenseNotificationSeverities
	isSet bool
}

func (v NullableLicenseNotificationSeverities) Get() *LicenseNotificationSeverities {
	return v.value
}

func (v *NullableLicenseNotificationSeverities) Set(val *LicenseNotificationSeverities) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseNotificationSeverities) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseNotificationSeverities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseNotificationSeverities(val *LicenseNotificationSeverities) *NullableLicenseNotificationSeverities {
	return &NullableLicenseNotificationSeverities{value: val, isSet: true}
}

func (v NullableLicenseNotificationSeverities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseNotificationSeverities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

