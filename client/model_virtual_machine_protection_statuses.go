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

// VirtualMachineProtectionStatuses the model 'VirtualMachineProtectionStatuses'
type VirtualMachineProtectionStatuses string

// List of VirtualMachineProtectionStatuses
const (
	VIRTUALMACHINEPROTECTIONSTATUSES_PROTECTED VirtualMachineProtectionStatuses = "Protected"
	VIRTUALMACHINEPROTECTIONSTATUSES_UNPROTECTED VirtualMachineProtectionStatuses = "Unprotected"
)

func (v *VirtualMachineProtectionStatuses) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VirtualMachineProtectionStatuses(value)
	for _, existing := range []VirtualMachineProtectionStatuses{ "Protected", "Unprotected",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VirtualMachineProtectionStatuses", value)
}

// Ptr returns reference to VirtualMachineProtectionStatuses value
func (v VirtualMachineProtectionStatuses) Ptr() *VirtualMachineProtectionStatuses {
	return &v
}

type NullableVirtualMachineProtectionStatuses struct {
	value *VirtualMachineProtectionStatuses
	isSet bool
}

func (v NullableVirtualMachineProtectionStatuses) Get() *VirtualMachineProtectionStatuses {
	return v.value
}

func (v *NullableVirtualMachineProtectionStatuses) Set(val *VirtualMachineProtectionStatuses) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualMachineProtectionStatuses) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualMachineProtectionStatuses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualMachineProtectionStatuses(val *VirtualMachineProtectionStatuses) *NullableVirtualMachineProtectionStatuses {
	return &NullableVirtualMachineProtectionStatuses{value: val, isSet: true}
}

func (v NullableVirtualMachineProtectionStatuses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualMachineProtectionStatuses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

