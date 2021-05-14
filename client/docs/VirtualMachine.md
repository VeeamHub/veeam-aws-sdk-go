# VirtualMachine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**AwsResourceId** | **string** |  | 
**ProtectionStatus** | [**VirtualMachineProtectionStatuses**](VirtualMachineProtectionStatuses.md) |  | 
**BackupDestinations** | Pointer to [**[]BackupDestinations**](BackupDestinations.md) |  | [optional] 
**Location** | Pointer to [**VirtualMachineLocation**](VirtualMachineLocation.md) |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 

## Methods

### NewVirtualMachine

`func NewVirtualMachine(id string, awsResourceId string, protectionStatus VirtualMachineProtectionStatuses, ) *VirtualMachine`

NewVirtualMachine instantiates a new VirtualMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualMachineWithDefaults

`func NewVirtualMachineWithDefaults() *VirtualMachine`

NewVirtualMachineWithDefaults instantiates a new VirtualMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualMachine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualMachine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualMachine) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *VirtualMachine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualMachine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualMachine) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualMachine) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAwsResourceId

`func (o *VirtualMachine) GetAwsResourceId() string`

GetAwsResourceId returns the AwsResourceId field if non-nil, zero value otherwise.

### GetAwsResourceIdOk

`func (o *VirtualMachine) GetAwsResourceIdOk() (*string, bool)`

GetAwsResourceIdOk returns a tuple with the AwsResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsResourceId

`func (o *VirtualMachine) SetAwsResourceId(v string)`

SetAwsResourceId sets AwsResourceId field to given value.


### GetProtectionStatus

`func (o *VirtualMachine) GetProtectionStatus() VirtualMachineProtectionStatuses`

GetProtectionStatus returns the ProtectionStatus field if non-nil, zero value otherwise.

### GetProtectionStatusOk

`func (o *VirtualMachine) GetProtectionStatusOk() (*VirtualMachineProtectionStatuses, bool)`

GetProtectionStatusOk returns a tuple with the ProtectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionStatus

`func (o *VirtualMachine) SetProtectionStatus(v VirtualMachineProtectionStatuses)`

SetProtectionStatus sets ProtectionStatus field to given value.


### GetBackupDestinations

`func (o *VirtualMachine) GetBackupDestinations() []BackupDestinations`

GetBackupDestinations returns the BackupDestinations field if non-nil, zero value otherwise.

### GetBackupDestinationsOk

`func (o *VirtualMachine) GetBackupDestinationsOk() (*[]BackupDestinations, bool)`

GetBackupDestinationsOk returns a tuple with the BackupDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDestinations

`func (o *VirtualMachine) SetBackupDestinations(v []BackupDestinations)`

SetBackupDestinations sets BackupDestinations field to given value.

### HasBackupDestinations

`func (o *VirtualMachine) HasBackupDestinations() bool`

HasBackupDestinations returns a boolean if a field has been set.

### GetLocation

`func (o *VirtualMachine) GetLocation() VirtualMachineLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *VirtualMachine) GetLocationOk() (*VirtualMachineLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *VirtualMachine) SetLocation(v VirtualMachineLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *VirtualMachine) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetIsDeleted

`func (o *VirtualMachine) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *VirtualMachine) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *VirtualMachine) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *VirtualMachine) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetLinks

`func (o *VirtualMachine) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VirtualMachine) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VirtualMachine) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VirtualMachine) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


