# Policy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SnapshotSettings** | Pointer to [**PolicySnapshotSettings**](PolicySnapshotSettings.md) |  | [optional] 
**ReplicaSettings** | Pointer to [**PolicyReplicaSettings**](PolicyReplicaSettings.md) |  | [optional] 
**BackupSettings** | Pointer to [**PolicyBackupSettings**](PolicyBackupSettings.md) |  | [optional] 
**RegionIds** | **[]string** |  | 
**SelectedItems** | Pointer to [**PolicyBackupItems**](PolicyBackupItems.md) |  | [optional] 
**ExcludedItems** | Pointer to [**PolicyBackupItems**](PolicyBackupItems.md) |  | [optional] 
**ScheduleSettings** | Pointer to [**ScheduleSettings**](ScheduleSettings.md) |  | [optional] 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Priority** | **int64** |  | 
**AmazonAccountId** | **string** |  | 
**RetrySettings** | Pointer to [**RetrySettings**](RetrySettings.md) |  | [optional] 
**PolicyNotificationSettings** | Pointer to [**PolicyNotificationSettings**](PolicyNotificationSettings.md) |  | [optional] 
**IsEnabled** | **bool** |  | 
**BackupType** | [**PolicySelectionTypes**](PolicySelectionTypes.md) |  | 
**CreatedBy** | Pointer to **string** |  | [optional] 
**ModifiedBy** | Pointer to **string** |  | [optional] 
**LastPolicySessionStatus** | Pointer to [**SessionStatuses**](SessionStatuses.md) |  | [optional] 
**Warning** | Pointer to **string** |  | [optional] 
**Usn** | Pointer to **int64** |  | [optional] 
**EmbeddedResources** | Pointer to [**PolicyEmbeddedResources**](PolicyEmbeddedResources.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 

## Methods

### NewPolicy

`func NewPolicy(id string, regionIds []string, name string, description string, priority int64, amazonAccountId string, isEnabled bool, backupType PolicySelectionTypes, ) *Policy`

NewPolicy instantiates a new Policy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyWithDefaults

`func NewPolicyWithDefaults() *Policy`

NewPolicyWithDefaults instantiates a new Policy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Policy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Policy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Policy) SetId(v string)`

SetId sets Id field to given value.


### GetSnapshotSettings

`func (o *Policy) GetSnapshotSettings() PolicySnapshotSettings`

GetSnapshotSettings returns the SnapshotSettings field if non-nil, zero value otherwise.

### GetSnapshotSettingsOk

`func (o *Policy) GetSnapshotSettingsOk() (*PolicySnapshotSettings, bool)`

GetSnapshotSettingsOk returns a tuple with the SnapshotSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotSettings

`func (o *Policy) SetSnapshotSettings(v PolicySnapshotSettings)`

SetSnapshotSettings sets SnapshotSettings field to given value.

### HasSnapshotSettings

`func (o *Policy) HasSnapshotSettings() bool`

HasSnapshotSettings returns a boolean if a field has been set.

### GetReplicaSettings

`func (o *Policy) GetReplicaSettings() PolicyReplicaSettings`

GetReplicaSettings returns the ReplicaSettings field if non-nil, zero value otherwise.

### GetReplicaSettingsOk

`func (o *Policy) GetReplicaSettingsOk() (*PolicyReplicaSettings, bool)`

GetReplicaSettingsOk returns a tuple with the ReplicaSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSettings

`func (o *Policy) SetReplicaSettings(v PolicyReplicaSettings)`

SetReplicaSettings sets ReplicaSettings field to given value.

### HasReplicaSettings

`func (o *Policy) HasReplicaSettings() bool`

HasReplicaSettings returns a boolean if a field has been set.

### GetBackupSettings

`func (o *Policy) GetBackupSettings() PolicyBackupSettings`

GetBackupSettings returns the BackupSettings field if non-nil, zero value otherwise.

### GetBackupSettingsOk

`func (o *Policy) GetBackupSettingsOk() (*PolicyBackupSettings, bool)`

GetBackupSettingsOk returns a tuple with the BackupSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSettings

`func (o *Policy) SetBackupSettings(v PolicyBackupSettings)`

SetBackupSettings sets BackupSettings field to given value.

### HasBackupSettings

`func (o *Policy) HasBackupSettings() bool`

HasBackupSettings returns a boolean if a field has been set.

### GetRegionIds

`func (o *Policy) GetRegionIds() []string`

GetRegionIds returns the RegionIds field if non-nil, zero value otherwise.

### GetRegionIdsOk

`func (o *Policy) GetRegionIdsOk() (*[]string, bool)`

GetRegionIdsOk returns a tuple with the RegionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionIds

`func (o *Policy) SetRegionIds(v []string)`

SetRegionIds sets RegionIds field to given value.


### GetSelectedItems

`func (o *Policy) GetSelectedItems() PolicyBackupItems`

GetSelectedItems returns the SelectedItems field if non-nil, zero value otherwise.

### GetSelectedItemsOk

`func (o *Policy) GetSelectedItemsOk() (*PolicyBackupItems, bool)`

GetSelectedItemsOk returns a tuple with the SelectedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedItems

`func (o *Policy) SetSelectedItems(v PolicyBackupItems)`

SetSelectedItems sets SelectedItems field to given value.

### HasSelectedItems

`func (o *Policy) HasSelectedItems() bool`

HasSelectedItems returns a boolean if a field has been set.

### GetExcludedItems

`func (o *Policy) GetExcludedItems() PolicyBackupItems`

GetExcludedItems returns the ExcludedItems field if non-nil, zero value otherwise.

### GetExcludedItemsOk

`func (o *Policy) GetExcludedItemsOk() (*PolicyBackupItems, bool)`

GetExcludedItemsOk returns a tuple with the ExcludedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedItems

`func (o *Policy) SetExcludedItems(v PolicyBackupItems)`

SetExcludedItems sets ExcludedItems field to given value.

### HasExcludedItems

`func (o *Policy) HasExcludedItems() bool`

HasExcludedItems returns a boolean if a field has been set.

### GetScheduleSettings

`func (o *Policy) GetScheduleSettings() ScheduleSettings`

GetScheduleSettings returns the ScheduleSettings field if non-nil, zero value otherwise.

### GetScheduleSettingsOk

`func (o *Policy) GetScheduleSettingsOk() (*ScheduleSettings, bool)`

GetScheduleSettingsOk returns a tuple with the ScheduleSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleSettings

`func (o *Policy) SetScheduleSettings(v ScheduleSettings)`

SetScheduleSettings sets ScheduleSettings field to given value.

### HasScheduleSettings

`func (o *Policy) HasScheduleSettings() bool`

HasScheduleSettings returns a boolean if a field has been set.

### GetName

`func (o *Policy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Policy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Policy) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Policy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Policy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Policy) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPriority

`func (o *Policy) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Policy) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Policy) SetPriority(v int64)`

SetPriority sets Priority field to given value.


### GetAmazonAccountId

`func (o *Policy) GetAmazonAccountId() string`

GetAmazonAccountId returns the AmazonAccountId field if non-nil, zero value otherwise.

### GetAmazonAccountIdOk

`func (o *Policy) GetAmazonAccountIdOk() (*string, bool)`

GetAmazonAccountIdOk returns a tuple with the AmazonAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonAccountId

`func (o *Policy) SetAmazonAccountId(v string)`

SetAmazonAccountId sets AmazonAccountId field to given value.


### GetRetrySettings

`func (o *Policy) GetRetrySettings() RetrySettings`

GetRetrySettings returns the RetrySettings field if non-nil, zero value otherwise.

### GetRetrySettingsOk

`func (o *Policy) GetRetrySettingsOk() (*RetrySettings, bool)`

GetRetrySettingsOk returns a tuple with the RetrySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrySettings

`func (o *Policy) SetRetrySettings(v RetrySettings)`

SetRetrySettings sets RetrySettings field to given value.

### HasRetrySettings

`func (o *Policy) HasRetrySettings() bool`

HasRetrySettings returns a boolean if a field has been set.

### GetPolicyNotificationSettings

`func (o *Policy) GetPolicyNotificationSettings() PolicyNotificationSettings`

GetPolicyNotificationSettings returns the PolicyNotificationSettings field if non-nil, zero value otherwise.

### GetPolicyNotificationSettingsOk

`func (o *Policy) GetPolicyNotificationSettingsOk() (*PolicyNotificationSettings, bool)`

GetPolicyNotificationSettingsOk returns a tuple with the PolicyNotificationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyNotificationSettings

`func (o *Policy) SetPolicyNotificationSettings(v PolicyNotificationSettings)`

SetPolicyNotificationSettings sets PolicyNotificationSettings field to given value.

### HasPolicyNotificationSettings

`func (o *Policy) HasPolicyNotificationSettings() bool`

HasPolicyNotificationSettings returns a boolean if a field has been set.

### GetIsEnabled

`func (o *Policy) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *Policy) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *Policy) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetBackupType

`func (o *Policy) GetBackupType() PolicySelectionTypes`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *Policy) GetBackupTypeOk() (*PolicySelectionTypes, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *Policy) SetBackupType(v PolicySelectionTypes)`

SetBackupType sets BackupType field to given value.


### GetCreatedBy

`func (o *Policy) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Policy) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Policy) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Policy) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetModifiedBy

`func (o *Policy) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Policy) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Policy) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *Policy) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetLastPolicySessionStatus

`func (o *Policy) GetLastPolicySessionStatus() SessionStatuses`

GetLastPolicySessionStatus returns the LastPolicySessionStatus field if non-nil, zero value otherwise.

### GetLastPolicySessionStatusOk

`func (o *Policy) GetLastPolicySessionStatusOk() (*SessionStatuses, bool)`

GetLastPolicySessionStatusOk returns a tuple with the LastPolicySessionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPolicySessionStatus

`func (o *Policy) SetLastPolicySessionStatus(v SessionStatuses)`

SetLastPolicySessionStatus sets LastPolicySessionStatus field to given value.

### HasLastPolicySessionStatus

`func (o *Policy) HasLastPolicySessionStatus() bool`

HasLastPolicySessionStatus returns a boolean if a field has been set.

### GetWarning

`func (o *Policy) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *Policy) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *Policy) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *Policy) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### GetUsn

`func (o *Policy) GetUsn() int64`

GetUsn returns the Usn field if non-nil, zero value otherwise.

### GetUsnOk

`func (o *Policy) GetUsnOk() (*int64, bool)`

GetUsnOk returns a tuple with the Usn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsn

`func (o *Policy) SetUsn(v int64)`

SetUsn sets Usn field to given value.

### HasUsn

`func (o *Policy) HasUsn() bool`

HasUsn returns a boolean if a field has been set.

### GetEmbeddedResources

`func (o *Policy) GetEmbeddedResources() PolicyEmbeddedResources`

GetEmbeddedResources returns the EmbeddedResources field if non-nil, zero value otherwise.

### GetEmbeddedResourcesOk

`func (o *Policy) GetEmbeddedResourcesOk() (*PolicyEmbeddedResources, bool)`

GetEmbeddedResourcesOk returns a tuple with the EmbeddedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddedResources

`func (o *Policy) SetEmbeddedResources(v PolicyEmbeddedResources)`

SetEmbeddedResources sets EmbeddedResources field to given value.

### HasEmbeddedResources

`func (o *Policy) HasEmbeddedResources() bool`

HasEmbeddedResources returns a boolean if a field has been set.

### GetLinks

`func (o *Policy) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Policy) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Policy) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Policy) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


