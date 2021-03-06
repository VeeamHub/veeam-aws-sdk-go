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

// BucketFoldersFiltersOrderColumns the model 'BucketFoldersFiltersOrderColumns'
type BucketFoldersFiltersOrderColumns string

// List of BucketFoldersFiltersOrderColumns
const (
	BUCKETFOLDERSFILTERSORDERCOLUMNS_NAME_ASC BucketFoldersFiltersOrderColumns = "NameAsc"
	BUCKETFOLDERSFILTERSORDERCOLUMNS_NAME_DESC BucketFoldersFiltersOrderColumns = "NameDesc"
)

func (v *BucketFoldersFiltersOrderColumns) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BucketFoldersFiltersOrderColumns(value)
	for _, existing := range []BucketFoldersFiltersOrderColumns{ "NameAsc", "NameDesc",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BucketFoldersFiltersOrderColumns", value)
}

// Ptr returns reference to BucketFoldersFiltersOrderColumns value
func (v BucketFoldersFiltersOrderColumns) Ptr() *BucketFoldersFiltersOrderColumns {
	return &v
}

type NullableBucketFoldersFiltersOrderColumns struct {
	value *BucketFoldersFiltersOrderColumns
	isSet bool
}

func (v NullableBucketFoldersFiltersOrderColumns) Get() *BucketFoldersFiltersOrderColumns {
	return v.value
}

func (v *NullableBucketFoldersFiltersOrderColumns) Set(val *BucketFoldersFiltersOrderColumns) {
	v.value = val
	v.isSet = true
}

func (v NullableBucketFoldersFiltersOrderColumns) IsSet() bool {
	return v.isSet
}

func (v *NullableBucketFoldersFiltersOrderColumns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBucketFoldersFiltersOrderColumns(val *BucketFoldersFiltersOrderColumns) *NullableBucketFoldersFiltersOrderColumns {
	return &NullableBucketFoldersFiltersOrderColumns{value: val, isSet: true}
}

func (v NullableBucketFoldersFiltersOrderColumns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBucketFoldersFiltersOrderColumns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

