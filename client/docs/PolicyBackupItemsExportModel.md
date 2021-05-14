# PolicyBackupItemsExportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualMachines** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to [**[]TagSpec**](TagSpec.md) |  | [optional] 

## Methods

### NewPolicyBackupItemsExportModel

`func NewPolicyBackupItemsExportModel() *PolicyBackupItemsExportModel`

NewPolicyBackupItemsExportModel instantiates a new PolicyBackupItemsExportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyBackupItemsExportModelWithDefaults

`func NewPolicyBackupItemsExportModelWithDefaults() *PolicyBackupItemsExportModel`

NewPolicyBackupItemsExportModelWithDefaults instantiates a new PolicyBackupItemsExportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualMachines

`func (o *PolicyBackupItemsExportModel) GetVirtualMachines() []string`

GetVirtualMachines returns the VirtualMachines field if non-nil, zero value otherwise.

### GetVirtualMachinesOk

`func (o *PolicyBackupItemsExportModel) GetVirtualMachinesOk() (*[]string, bool)`

GetVirtualMachinesOk returns a tuple with the VirtualMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachines

`func (o *PolicyBackupItemsExportModel) SetVirtualMachines(v []string)`

SetVirtualMachines sets VirtualMachines field to given value.

### HasVirtualMachines

`func (o *PolicyBackupItemsExportModel) HasVirtualMachines() bool`

HasVirtualMachines returns a boolean if a field has been set.

### GetTags

`func (o *PolicyBackupItemsExportModel) GetTags() []TagSpec`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PolicyBackupItemsExportModel) GetTagsOk() (*[]TagSpec, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PolicyBackupItemsExportModel) SetTags(v []TagSpec)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PolicyBackupItemsExportModel) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


