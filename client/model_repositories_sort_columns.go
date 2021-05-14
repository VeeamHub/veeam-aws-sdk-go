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

// RepositoriesSortColumns the model 'RepositoriesSortColumns'
type RepositoriesSortColumns string

// List of RepositoriesSortColumns
const (
	REPOSITORIESSORTCOLUMNS_NAME_ASC RepositoriesSortColumns = "nameAsc"
	REPOSITORIESSORTCOLUMNS_NAME_DESC RepositoriesSortColumns = "nameDesc"
	REPOSITORIESSORTCOLUMNS_DESCRIPTION_ASC RepositoriesSortColumns = "descriptionAsc"
	REPOSITORIESSORTCOLUMNS_DESCRIPTION_DESC RepositoriesSortColumns = "descriptionDesc"
	REPOSITORIESSORTCOLUMNS_AMAZON_ACCOUNT_NAME_ASC RepositoriesSortColumns = "amazonAccountNameAsc"
	REPOSITORIESSORTCOLUMNS_AMAZON_ACCOUNT_NAME_DESC RepositoriesSortColumns = "amazonAccountNameDesc"
	REPOSITORIESSORTCOLUMNS_AMAZON_BUCKET_ASC RepositoriesSortColumns = "amazonBucketAsc"
	REPOSITORIESSORTCOLUMNS_AMAZON_BUCKET_DESC RepositoriesSortColumns = "amazonBucketDesc"
	REPOSITORIESSORTCOLUMNS_AMAZON_STORAGE_FOLDER_ASC RepositoriesSortColumns = "amazonStorageFolderAsc"
	REPOSITORIESSORTCOLUMNS_AMAZON_STORAGE_FOLDER_DESC RepositoriesSortColumns = "amazonStorageFolderDesc"
	REPOSITORIESSORTCOLUMNS_REGION_ASC RepositoriesSortColumns = "regionAsc"
	REPOSITORIESSORTCOLUMNS_REGION_DESC RepositoriesSortColumns = "regionDesc"
	REPOSITORIESSORTCOLUMNS_ENCRYPTION_ASC RepositoriesSortColumns = "encryptionAsc"
	REPOSITORIESSORTCOLUMNS_ENCRYPTION_DESC RepositoriesSortColumns = "encryptionDesc"
)

func (v *RepositoriesSortColumns) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RepositoriesSortColumns(value)
	for _, existing := range []RepositoriesSortColumns{ "nameAsc", "nameDesc", "descriptionAsc", "descriptionDesc", "amazonAccountNameAsc", "amazonAccountNameDesc", "amazonBucketAsc", "amazonBucketDesc", "amazonStorageFolderAsc", "amazonStorageFolderDesc", "regionAsc", "regionDesc", "encryptionAsc", "encryptionDesc",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RepositoriesSortColumns", value)
}

// Ptr returns reference to RepositoriesSortColumns value
func (v RepositoriesSortColumns) Ptr() *RepositoriesSortColumns {
	return &v
}

type NullableRepositoriesSortColumns struct {
	value *RepositoriesSortColumns
	isSet bool
}

func (v NullableRepositoriesSortColumns) Get() *RepositoriesSortColumns {
	return v.value
}

func (v *NullableRepositoriesSortColumns) Set(val *RepositoriesSortColumns) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoriesSortColumns) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoriesSortColumns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoriesSortColumns(val *RepositoriesSortColumns) *NullableRepositoriesSortColumns {
	return &NullableRepositoriesSortColumns{value: val, isSet: true}
}

func (v NullableRepositoriesSortColumns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoriesSortColumns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
