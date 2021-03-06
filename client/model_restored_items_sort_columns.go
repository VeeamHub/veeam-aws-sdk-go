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

// RestoredItemsSortColumns the model 'RestoredItemsSortColumns'
type RestoredItemsSortColumns string

// List of RestoredItemsSortColumns
const (
	RESTOREDITEMSSORTCOLUMNS_NAME_ASC RestoredItemsSortColumns = "nameAsc"
	RESTOREDITEMSSORTCOLUMNS_NAME_DESC RestoredItemsSortColumns = "nameDesc"
	RESTOREDITEMSSORTCOLUMNS_REGION_ASC RestoredItemsSortColumns = "regionAsc"
	RESTOREDITEMSSORTCOLUMNS_REGION_DESC RestoredItemsSortColumns = "regionDesc"
)

func (v *RestoredItemsSortColumns) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RestoredItemsSortColumns(value)
	for _, existing := range []RestoredItemsSortColumns{ "nameAsc", "nameDesc", "regionAsc", "regionDesc",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RestoredItemsSortColumns", value)
}

// Ptr returns reference to RestoredItemsSortColumns value
func (v RestoredItemsSortColumns) Ptr() *RestoredItemsSortColumns {
	return &v
}

type NullableRestoredItemsSortColumns struct {
	value *RestoredItemsSortColumns
	isSet bool
}

func (v NullableRestoredItemsSortColumns) Get() *RestoredItemsSortColumns {
	return v.value
}

func (v *NullableRestoredItemsSortColumns) Set(val *RestoredItemsSortColumns) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoredItemsSortColumns) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoredItemsSortColumns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoredItemsSortColumns(val *RestoredItemsSortColumns) *NullableRestoredItemsSortColumns {
	return &NullableRestoredItemsSortColumns{value: val, isSet: true}
}

func (v NullableRestoredItemsSortColumns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoredItemsSortColumns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

