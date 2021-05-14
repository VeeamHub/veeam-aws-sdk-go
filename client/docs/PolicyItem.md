# PolicyItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to [**Link**](Link.md) |  | [optional] 
**VirtualMachine** | Pointer to [**Link**](Link.md) |  | [optional] 
**DeletedItem** | Pointer to [**PolicyItemDeletedFromAmazon**](PolicyItemDeletedFromAmazon.md) |  | [optional] 

## Methods

### NewPolicyItem

`func NewPolicyItem() *PolicyItem`

NewPolicyItem instantiates a new PolicyItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyItemWithDefaults

`func NewPolicyItemWithDefaults() *PolicyItem`

NewPolicyItemWithDefaults instantiates a new PolicyItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PolicyItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTag

`func (o *PolicyItem) GetTag() Link`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *PolicyItem) GetTagOk() (*Link, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *PolicyItem) SetTag(v Link)`

SetTag sets Tag field to given value.

### HasTag

`func (o *PolicyItem) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetVirtualMachine

`func (o *PolicyItem) GetVirtualMachine() Link`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *PolicyItem) GetVirtualMachineOk() (*Link, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *PolicyItem) SetVirtualMachine(v Link)`

SetVirtualMachine sets VirtualMachine field to given value.

### HasVirtualMachine

`func (o *PolicyItem) HasVirtualMachine() bool`

HasVirtualMachine returns a boolean if a field has been set.

### GetDeletedItem

`func (o *PolicyItem) GetDeletedItem() PolicyItemDeletedFromAmazon`

GetDeletedItem returns the DeletedItem field if non-nil, zero value otherwise.

### GetDeletedItemOk

`func (o *PolicyItem) GetDeletedItemOk() (*PolicyItemDeletedFromAmazon, bool)`

GetDeletedItemOk returns a tuple with the DeletedItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedItem

`func (o *PolicyItem) SetDeletedItem(v PolicyItemDeletedFromAmazon)`

SetDeletedItem sets DeletedItem field to given value.

### HasDeletedItem

`func (o *PolicyItem) HasDeletedItem() bool`

HasDeletedItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


