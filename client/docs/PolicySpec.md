# PolicySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotSettings** | Pointer to [**PolicySnapshotSettings**](PolicySnapshotSettings.md) |  | [optional] 
**ReplicaSettings** | Pointer to [**PolicyReplicaSettings**](PolicyReplicaSettings.md) |  | [optional] 
**BackupSettings** | Pointer to [**PolicyBackupSettings**](PolicyBackupSettings.md) |  | [optional] 
**RegionIds** | Pointer to **[]string** |  | [optional] 
**SelectedItems** | Pointer to [**PolicyBackupItems**](PolicyBackupItems.md) |  | [optional] 
**ExcludedItems** | Pointer to [**PolicyBackupItems**](PolicyBackupItems.md) |  | [optional] 
**ScheduleSettings** | Pointer to [**ScheduleSettings**](ScheduleSettings.md) |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**AmazonAccountId** | **string** |  | 
**RetrySettings** | Pointer to [**RetrySettings**](RetrySettings.md) |  | [optional] 
**PolicyNotificationSettings** | Pointer to [**PolicyNotificationSettings**](PolicyNotificationSettings.md) |  | [optional] 
**BackupType** | [**PolicySelectionTypes**](PolicySelectionTypes.md) |  | 

## Methods

### NewPolicySpec

`func NewPolicySpec(name string, amazonAccountId string, backupType PolicySelectionTypes, ) *PolicySpec`

NewPolicySpec instantiates a new PolicySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicySpecWithDefaults

`func NewPolicySpecWithDefaults() *PolicySpec`

NewPolicySpecWithDefaults instantiates a new PolicySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotSettings

`func (o *PolicySpec) GetSnapshotSettings() PolicySnapshotSettings`

GetSnapshotSettings returns the SnapshotSettings field if non-nil, zero value otherwise.

### GetSnapshotSettingsOk

`func (o *PolicySpec) GetSnapshotSettingsOk() (*PolicySnapshotSettings, bool)`

GetSnapshotSettingsOk returns a tuple with the SnapshotSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotSettings

`func (o *PolicySpec) SetSnapshotSettings(v PolicySnapshotSettings)`

SetSnapshotSettings sets SnapshotSettings field to given value.

### HasSnapshotSettings

`func (o *PolicySpec) HasSnapshotSettings() bool`

HasSnapshotSettings returns a boolean if a field has been set.

### GetReplicaSettings

`func (o *PolicySpec) GetReplicaSettings() PolicyReplicaSettings`

GetReplicaSettings returns the ReplicaSettings field if non-nil, zero value otherwise.

### GetReplicaSettingsOk

`func (o *PolicySpec) GetReplicaSettingsOk() (*PolicyReplicaSettings, bool)`

GetReplicaSettingsOk returns a tuple with the ReplicaSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSettings

`func (o *PolicySpec) SetReplicaSettings(v PolicyReplicaSettings)`

SetReplicaSettings sets ReplicaSettings field to given value.

### HasReplicaSettings

`func (o *PolicySpec) HasReplicaSettings() bool`

HasReplicaSettings returns a boolean if a field has been set.

### GetBackupSettings

`func (o *PolicySpec) GetBackupSettings() PolicyBackupSettings`

GetBackupSettings returns the BackupSettings field if non-nil, zero value otherwise.

### GetBackupSettingsOk

`func (o *PolicySpec) GetBackupSettingsOk() (*PolicyBackupSettings, bool)`

GetBackupSettingsOk returns a tuple with the BackupSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSettings

`func (o *PolicySpec) SetBackupSettings(v PolicyBackupSettings)`

SetBackupSettings sets BackupSettings field to given value.

### HasBackupSettings

`func (o *PolicySpec) HasBackupSettings() bool`

HasBackupSettings returns a boolean if a field has been set.

### GetRegionIds

`func (o *PolicySpec) GetRegionIds() []string`

GetRegionIds returns the RegionIds field if non-nil, zero value otherwise.

### GetRegionIdsOk

`func (o *PolicySpec) GetRegionIdsOk() (*[]string, bool)`

GetRegionIdsOk returns a tuple with the RegionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionIds

`func (o *PolicySpec) SetRegionIds(v []string)`

SetRegionIds sets RegionIds field to given value.

### HasRegionIds

`func (o *PolicySpec) HasRegionIds() bool`

HasRegionIds returns a boolean if a field has been set.

### GetSelectedItems

`func (o *PolicySpec) GetSelectedItems() PolicyBackupItems`

GetSelectedItems returns the SelectedItems field if non-nil, zero value otherwise.

### GetSelectedItemsOk

`func (o *PolicySpec) GetSelectedItemsOk() (*PolicyBackupItems, bool)`

GetSelectedItemsOk returns a tuple with the SelectedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedItems

`func (o *PolicySpec) SetSelectedItems(v PolicyBackupItems)`

SetSelectedItems sets SelectedItems field to given value.

### HasSelectedItems

`func (o *PolicySpec) HasSelectedItems() bool`

HasSelectedItems returns a boolean if a field has been set.

### GetExcludedItems

`func (o *PolicySpec) GetExcludedItems() PolicyBackupItems`

GetExcludedItems returns the ExcludedItems field if non-nil, zero value otherwise.

### GetExcludedItemsOk

`func (o *PolicySpec) GetExcludedItemsOk() (*PolicyBackupItems, bool)`

GetExcludedItemsOk returns a tuple with the ExcludedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedItems

`func (o *PolicySpec) SetExcludedItems(v PolicyBackupItems)`

SetExcludedItems sets ExcludedItems field to given value.

### HasExcludedItems

`func (o *PolicySpec) HasExcludedItems() bool`

HasExcludedItems returns a boolean if a field has been set.

### GetScheduleSettings

`func (o *PolicySpec) GetScheduleSettings() ScheduleSettings`

GetScheduleSettings returns the ScheduleSettings field if non-nil, zero value otherwise.

### GetScheduleSettingsOk

`func (o *PolicySpec) GetScheduleSettingsOk() (*ScheduleSettings, bool)`

GetScheduleSettingsOk returns a tuple with the ScheduleSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleSettings

`func (o *PolicySpec) SetScheduleSettings(v ScheduleSettings)`

SetScheduleSettings sets ScheduleSettings field to given value.

### HasScheduleSettings

`func (o *PolicySpec) HasScheduleSettings() bool`

HasScheduleSettings returns a boolean if a field has been set.

### GetName

`func (o *PolicySpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicySpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicySpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PolicySpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicySpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicySpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicySpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAmazonAccountId

`func (o *PolicySpec) GetAmazonAccountId() string`

GetAmazonAccountId returns the AmazonAccountId field if non-nil, zero value otherwise.

### GetAmazonAccountIdOk

`func (o *PolicySpec) GetAmazonAccountIdOk() (*string, bool)`

GetAmazonAccountIdOk returns a tuple with the AmazonAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonAccountId

`func (o *PolicySpec) SetAmazonAccountId(v string)`

SetAmazonAccountId sets AmazonAccountId field to given value.


### GetRetrySettings

`func (o *PolicySpec) GetRetrySettings() RetrySettings`

GetRetrySettings returns the RetrySettings field if non-nil, zero value otherwise.

### GetRetrySettingsOk

`func (o *PolicySpec) GetRetrySettingsOk() (*RetrySettings, bool)`

GetRetrySettingsOk returns a tuple with the RetrySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrySettings

`func (o *PolicySpec) SetRetrySettings(v RetrySettings)`

SetRetrySettings sets RetrySettings field to given value.

### HasRetrySettings

`func (o *PolicySpec) HasRetrySettings() bool`

HasRetrySettings returns a boolean if a field has been set.

### GetPolicyNotificationSettings

`func (o *PolicySpec) GetPolicyNotificationSettings() PolicyNotificationSettings`

GetPolicyNotificationSettings returns the PolicyNotificationSettings field if non-nil, zero value otherwise.

### GetPolicyNotificationSettingsOk

`func (o *PolicySpec) GetPolicyNotificationSettingsOk() (*PolicyNotificationSettings, bool)`

GetPolicyNotificationSettingsOk returns a tuple with the PolicyNotificationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyNotificationSettings

`func (o *PolicySpec) SetPolicyNotificationSettings(v PolicyNotificationSettings)`

SetPolicyNotificationSettings sets PolicyNotificationSettings field to given value.

### HasPolicyNotificationSettings

`func (o *PolicySpec) HasPolicyNotificationSettings() bool`

HasPolicyNotificationSettings returns a boolean if a field has been set.

### GetBackupType

`func (o *PolicySpec) GetBackupType() PolicySelectionTypes`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *PolicySpec) GetBackupTypeOk() (*PolicySelectionTypes, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *PolicySpec) SetBackupType(v PolicySelectionTypes)`

SetBackupType sets BackupType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


