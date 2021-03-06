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

// Months the model 'Months'
type Months string

// List of Months
const (
	MONTHS_JANUARY Months = "January"
	MONTHS_FEBRUARY Months = "February"
	MONTHS_MARCH Months = "March"
	MONTHS_APRIL Months = "April"
	MONTHS_MAY Months = "May"
	MONTHS_JUNE Months = "June"
	MONTHS_JULY Months = "July"
	MONTHS_AUGUST Months = "August"
	MONTHS_SEPTEMBER Months = "September"
	MONTHS_OCTOBER Months = "October"
	MONTHS_NOVEMBER Months = "November"
	MONTHS_DECEMBER Months = "December"
)

func (v *Months) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Months(value)
	for _, existing := range []Months{ "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Months", value)
}

// Ptr returns reference to Months value
func (v Months) Ptr() *Months {
	return &v
}

type NullableMonths struct {
	value *Months
	isSet bool
}

func (v NullableMonths) Get() *Months {
	return v.value
}

func (v *NullableMonths) Set(val *Months) {
	v.value = val
	v.isSet = true
}

func (v NullableMonths) IsSet() bool {
	return v.isSet
}

func (v *NullableMonths) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonths(val *Months) *NullableMonths {
	return &NullableMonths{value: val, isSet: true}
}

func (v NullableMonths) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonths) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

