# PolicyExportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotSettings** | Pointer to [**PolicySnapshotSettings**](PolicySnapshotSettings.md) |  | [optional] 
**ReplicaSettings** | Pointer to [**ReplicaSettingsExportModel**](ReplicaSettingsExportModel.md) |  | [optional] 
**BackupSettings** | Pointer to [**BackupSettingsExportModel**](BackupSettingsExportModel.md) |  | [optional] 
**Regions** | **[]string** |  | 
**SelectedItems** | Pointer to [**PolicyBackupItemsExportModel**](PolicyBackupItemsExportModel.md) |  | [optional] 
**ExcludedItems** | Pointer to [**PolicyBackupItemsExportModel**](PolicyBackupItemsExportModel.md) |  | [optional] 
**ScheduleSettings** | Pointer to [**ScheduleSettings**](ScheduleSettings.md) |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int64** |  | [optional] 
**AmazonAccountName** | **string** |  | 
**RetrySettings** | Pointer to [**RetrySettings**](RetrySettings.md) |  | [optional] 
**PolicyNotificationSettings** | Pointer to [**PolicyNotificationSettings**](PolicyNotificationSettings.md) |  | [optional] 
**IsEnabled** | **bool** |  | 
**BackupType** | [**PolicySelectionTypes**](PolicySelectionTypes.md) |  | 

## Methods

### NewPolicyExportModel

`func NewPolicyExportModel(regions []string, name string, amazonAccountName string, isEnabled bool, backupType PolicySelectionTypes, ) *PolicyExportModel`

NewPolicyExportModel instantiates a new PolicyExportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyExportModelWithDefaults

`func NewPolicyExportModelWithDefaults() *PolicyExportModel`

NewPolicyExportModelWithDefaults instantiates a new PolicyExportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotSettings

`func (o *PolicyExportModel) GetSnapshotSettings() PolicySnapshotSettings`

GetSnapshotSettings returns the SnapshotSettings field if non-nil, zero value otherwise.

### GetSnapshotSettingsOk

`func (o *PolicyExportModel) GetSnapshotSettingsOk() (*PolicySnapshotSettings, bool)`

GetSnapshotSettingsOk returns a tuple with the SnapshotSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotSettings

`func (o *PolicyExportModel) SetSnapshotSettings(v PolicySnapshotSettings)`

SetSnapshotSettings sets SnapshotSettings field to given value.

### HasSnapshotSettings

`func (o *PolicyExportModel) HasSnapshotSettings() bool`

HasSnapshotSettings returns a boolean if a field has been set.

### GetReplicaSettings

`func (o *PolicyExportModel) GetReplicaSettings() ReplicaSettingsExportModel`

GetReplicaSettings returns the ReplicaSettings field if non-nil, zero value otherwise.

### GetReplicaSettingsOk

`func (o *PolicyExportModel) GetReplicaSettingsOk() (*ReplicaSettingsExportModel, bool)`

GetReplicaSettingsOk returns a tuple with the ReplicaSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSettings

`func (o *PolicyExportModel) SetReplicaSettings(v ReplicaSettingsExportModel)`

SetReplicaSettings sets ReplicaSettings field to given value.

### HasReplicaSettings

`func (o *PolicyExportModel) HasReplicaSettings() bool`

HasReplicaSettings returns a boolean if a field has been set.

### GetBackupSettings

`func (o *PolicyExportModel) GetBackupSettings() BackupSettingsExportModel`

GetBackupSettings returns the BackupSettings field if non-nil, zero value otherwise.

### GetBackupSettingsOk

`func (o *PolicyExportModel) GetBackupSettingsOk() (*BackupSettingsExportModel, bool)`

GetBackupSettingsOk returns a tuple with the BackupSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSettings

`func (o *PolicyExportModel) SetBackupSettings(v BackupSettingsExportModel)`

SetBackupSettings sets BackupSettings field to given value.

### HasBackupSettings

`func (o *PolicyExportModel) HasBackupSettings() bool`

HasBackupSettings returns a boolean if a field has been set.

### GetRegions

`func (o *PolicyExportModel) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *PolicyExportModel) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *PolicyExportModel) SetRegions(v []string)`

SetRegions sets Regions field to given value.


### GetSelectedItems

`func (o *PolicyExportModel) GetSelectedItems() PolicyBackupItemsExportModel`

GetSelectedItems returns the SelectedItems field if non-nil, zero value otherwise.

### GetSelectedItemsOk

`func (o *PolicyExportModel) GetSelectedItemsOk() (*PolicyBackupItemsExportModel, bool)`

GetSelectedItemsOk returns a tuple with the SelectedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedItems

`func (o *PolicyExportModel) SetSelectedItems(v PolicyBackupItemsExportModel)`

SetSelectedItems sets SelectedItems field to given value.

### HasSelectedItems

`func (o *PolicyExportModel) HasSelectedItems() bool`

HasSelectedItems returns a boolean if a field has been set.

### GetExcludedItems

`func (o *PolicyExportModel) GetExcludedItems() PolicyBackupItemsExportModel`

GetExcludedItems returns the ExcludedItems field if non-nil, zero value otherwise.

### GetExcludedItemsOk

`func (o *PolicyExportModel) GetExcludedItemsOk() (*PolicyBackupItemsExportModel, bool)`

GetExcludedItemsOk returns a tuple with the ExcludedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedItems

`func (o *PolicyExportModel) SetExcludedItems(v PolicyBackupItemsExportModel)`

SetExcludedItems sets ExcludedItems field to given value.

### HasExcludedItems

`func (o *PolicyExportModel) HasExcludedItems() bool`

HasExcludedItems returns a boolean if a field has been set.

### GetScheduleSettings

`func (o *PolicyExportModel) GetScheduleSettings() ScheduleSettings`

GetScheduleSettings returns the ScheduleSettings field if non-nil, zero value otherwise.

### GetScheduleSettingsOk

`func (o *PolicyExportModel) GetScheduleSettingsOk() (*ScheduleSettings, bool)`

GetScheduleSettingsOk returns a tuple with the ScheduleSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleSettings

`func (o *PolicyExportModel) SetScheduleSettings(v ScheduleSettings)`

SetScheduleSettings sets ScheduleSettings field to given value.

### HasScheduleSettings

`func (o *PolicyExportModel) HasScheduleSettings() bool`

HasScheduleSettings returns a boolean if a field has been set.

### GetName

`func (o *PolicyExportModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyExportModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyExportModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PolicyExportModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyExportModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyExportModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyExportModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPriority

`func (o *PolicyExportModel) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PolicyExportModel) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PolicyExportModel) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PolicyExportModel) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetAmazonAccountName

`func (o *PolicyExportModel) GetAmazonAccountName() string`

GetAmazonAccountName returns the AmazonAccountName field if non-nil, zero value otherwise.

### GetAmazonAccountNameOk

`func (o *PolicyExportModel) GetAmazonAccountNameOk() (*string, bool)`

GetAmazonAccountNameOk returns a tuple with the AmazonAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonAccountName

`func (o *PolicyExportModel) SetAmazonAccountName(v string)`

SetAmazonAccountName sets AmazonAccountName field to given value.


### GetRetrySettings

`func (o *PolicyExportModel) GetRetrySettings() RetrySettings`

GetRetrySettings returns the RetrySettings field if non-nil, zero value otherwise.

### GetRetrySettingsOk

`func (o *PolicyExportModel) GetRetrySettingsOk() (*RetrySettings, bool)`

GetRetrySettingsOk returns a tuple with the RetrySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrySettings

`func (o *PolicyExportModel) SetRetrySettings(v RetrySettings)`

SetRetrySettings sets RetrySettings field to given value.

### HasRetrySettings

`func (o *PolicyExportModel) HasRetrySettings() bool`

HasRetrySettings returns a boolean if a field has been set.

### GetPolicyNotificationSettings

`func (o *PolicyExportModel) GetPolicyNotificationSettings() PolicyNotificationSettings`

GetPolicyNotificationSettings returns the PolicyNotificationSettings field if non-nil, zero value otherwise.

### GetPolicyNotificationSettingsOk

`func (o *PolicyExportModel) GetPolicyNotificationSettingsOk() (*PolicyNotificationSettings, bool)`

GetPolicyNotificationSettingsOk returns a tuple with the PolicyNotificationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyNotificationSettings

`func (o *PolicyExportModel) SetPolicyNotificationSettings(v PolicyNotificationSettings)`

SetPolicyNotificationSettings sets PolicyNotificationSettings field to given value.

### HasPolicyNotificationSettings

`func (o *PolicyExportModel) HasPolicyNotificationSettings() bool`

HasPolicyNotificationSettings returns a boolean if a field has been set.

### GetIsEnabled

`func (o *PolicyExportModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *PolicyExportModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *PolicyExportModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetBackupType

`func (o *PolicyExportModel) GetBackupType() PolicySelectionTypes`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *PolicyExportModel) GetBackupTypeOk() (*PolicySelectionTypes, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *PolicyExportModel) SetBackupType(v PolicySelectionTypes)`

SetBackupType sets BackupType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


