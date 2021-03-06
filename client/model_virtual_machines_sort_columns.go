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

// VirtualMachinesSortColumns the model 'VirtualMachinesSortColumns'
type VirtualMachinesSortColumns string

// List of VirtualMachinesSortColumns
const (
	VIRTUALMACHINESSORTCOLUMNS_NAME_ASC VirtualMachinesSortColumns = "nameAsc"
	VIRTUALMACHINESSORTCOLUMNS_NAME_DESC VirtualMachinesSortColumns = "nameDesc"
	VIRTUALMACHINESSORTCOLUMNS_LOCATION_ASC VirtualMachinesSortColumns = "locationAsc"
	VIRTUALMACHINESSORTCOLUMNS_LOCATION_DESC VirtualMachinesSortColumns = "locationDesc"
)

func (v *VirtualMachinesSortColumns) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VirtualMachinesSortColumns(value)
	for _, existing := range []VirtualMachinesSortColumns{ "nameAsc", "nameDesc", "locationAsc", "locationDesc",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VirtualMachinesSortColumns", value)
}

// Ptr returns reference to VirtualMachinesSortColumns value
func (v VirtualMachinesSortColumns) Ptr() *VirtualMachinesSortColumns {
	return &v
}

type NullableVirtualMachinesSortColumns struct {
	value *VirtualMachinesSortColumns
	isSet bool
}

func (v NullableVirtualMachinesSortColumns) Get() *VirtualMachinesSortColumns {
	return v.value
}

func (v *NullableVirtualMachinesSortColumns) Set(val *VirtualMachinesSortColumns) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualMachinesSortColumns) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualMachinesSortColumns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualMachinesSortColumns(val *VirtualMachinesSortColumns) *NullableVirtualMachinesSortColumns {
	return &NullableVirtualMachinesSortColumns{value: val, isSet: true}
}

func (v NullableVirtualMachinesSortColumns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualMachinesSortColumns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

