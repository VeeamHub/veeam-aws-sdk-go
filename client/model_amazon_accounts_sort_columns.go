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

// AmazonAccountsSortColumns the model 'AmazonAccountsSortColumns'
type AmazonAccountsSortColumns string

// List of AmazonAccountsSortColumns
const (
	AMAZONACCOUNTSSORTCOLUMNS_NAME_ASC AmazonAccountsSortColumns = "nameAsc"
	AMAZONACCOUNTSSORTCOLUMNS_NAME_DESC AmazonAccountsSortColumns = "nameDesc"
)

func (v *AmazonAccountsSortColumns) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AmazonAccountsSortColumns(value)
	for _, existing := range []AmazonAccountsSortColumns{ "nameAsc", "nameDesc",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AmazonAccountsSortColumns", value)
}

// Ptr returns reference to AmazonAccountsSortColumns value
func (v AmazonAccountsSortColumns) Ptr() *AmazonAccountsSortColumns {
	return &v
}

type NullableAmazonAccountsSortColumns struct {
	value *AmazonAccountsSortColumns
	isSet bool
}

func (v NullableAmazonAccountsSortColumns) Get() *AmazonAccountsSortColumns {
	return v.value
}

func (v *NullableAmazonAccountsSortColumns) Set(val *AmazonAccountsSortColumns) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonAccountsSortColumns) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonAccountsSortColumns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonAccountsSortColumns(val *AmazonAccountsSortColumns) *NullableAmazonAccountsSortColumns {
	return &NullableAmazonAccountsSortColumns{value: val, isSet: true}
}

func (v NullableAmazonAccountsSortColumns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonAccountsSortColumns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
