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

// DaysOfWeek the model 'DaysOfWeek'
type DaysOfWeek string

// List of DaysOfWeek
const (
	DAYSOFWEEK_SUNDAY DaysOfWeek = "Sunday"
	DAYSOFWEEK_MONDAY DaysOfWeek = "Monday"
	DAYSOFWEEK_TUESDAY DaysOfWeek = "Tuesday"
	DAYSOFWEEK_WEDNESDAY DaysOfWeek = "Wednesday"
	DAYSOFWEEK_THURSDAY DaysOfWeek = "Thursday"
	DAYSOFWEEK_FRIDAY DaysOfWeek = "Friday"
	DAYSOFWEEK_SATURDAY DaysOfWeek = "Saturday"
)

func (v *DaysOfWeek) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DaysOfWeek(value)
	for _, existing := range []DaysOfWeek{ "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DaysOfWeek", value)
}

// Ptr returns reference to DaysOfWeek value
func (v DaysOfWeek) Ptr() *DaysOfWeek {
	return &v
}

type NullableDaysOfWeek struct {
	value *DaysOfWeek
	isSet bool
}

func (v NullableDaysOfWeek) Get() *DaysOfWeek {
	return v.value
}

func (v *NullableDaysOfWeek) Set(val *DaysOfWeek) {
	v.value = val
	v.isSet = true
}

func (v NullableDaysOfWeek) IsSet() bool {
	return v.isSet
}

func (v *NullableDaysOfWeek) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDaysOfWeek(val *DaysOfWeek) *NullableDaysOfWeek {
	return &NullableDaysOfWeek{value: val, isSet: true}
}

func (v NullableDaysOfWeek) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDaysOfWeek) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

