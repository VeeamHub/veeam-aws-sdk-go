# VirtualMachineRestoreOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** |  | [optional] 
**AmazonAccountId** | Pointer to **string** |  | [optional] 
**AccessKey** | Pointer to **string** |  | [optional] 
**SecretKey** | Pointer to **string** |  | [optional] 
**StartInstance** | Pointer to **bool** |  | [optional] 
**ToAlternative** | Pointer to [**VirtualMachineToAlternativeRestoreOptions**](VirtualMachineToAlternativeRestoreOptions.md) |  | [optional] 

## Methods

### NewVirtualMachineRestoreOptions

`func NewVirtualMachineRestoreOptions() *VirtualMachineRestoreOptions`

NewVirtualMachineRestoreOptions instantiates a new VirtualMachineRestoreOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualMachineRestoreOptionsWithDefaults

`func NewVirtualMachineRestoreOptionsWithDefaults() *VirtualMachineRestoreOptions`

NewVirtualMachineRestoreOptionsWithDefaults instantiates a new VirtualMachineRestoreOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *VirtualMachineRestoreOptions) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *VirtualMachineRestoreOptions) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *VirtualMachineRestoreOptions) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *VirtualMachineRestoreOptions) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetAmazonAccountId

`func (o *VirtualMachineRestoreOptions) GetAmazonAccountId() string`

GetAmazonAccountId returns the AmazonAccountId field if non-nil, zero value otherwise.

### GetAmazonAccountIdOk

`func (o *VirtualMachineRestoreOptions) GetAmazonAccountIdOk() (*string, bool)`

GetAmazonAccountIdOk returns a tuple with the AmazonAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonAccountId

`func (o *VirtualMachineRestoreOptions) SetAmazonAccountId(v string)`

SetAmazonAccountId sets AmazonAccountId field to given value.

### HasAmazonAccountId

`func (o *VirtualMachineRestoreOptions) HasAmazonAccountId() bool`

HasAmazonAccountId returns a boolean if a field has been set.

### GetAccessKey

`func (o *VirtualMachineRestoreOptions) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *VirtualMachineRestoreOptions) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *VirtualMachineRestoreOptions) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *VirtualMachineRestoreOptions) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *VirtualMachineRestoreOptions) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *VirtualMachineRestoreOptions) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *VirtualMachineRestoreOptions) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *VirtualMachineRestoreOptions) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetStartInstance

`func (o *VirtualMachineRestoreOptions) GetStartInstance() bool`

GetStartInstance returns the StartInstance field if non-nil, zero value otherwise.

### GetStartInstanceOk

`func (o *VirtualMachineRestoreOptions) GetStartInstanceOk() (*bool, bool)`

GetStartInstanceOk returns a tuple with the StartInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartInstance

`func (o *VirtualMachineRestoreOptions) SetStartInstance(v bool)`

SetStartInstance sets StartInstance field to given value.

### HasStartInstance

`func (o *VirtualMachineRestoreOptions) HasStartInstance() bool`

HasStartInstance returns a boolean if a field has been set.

### GetToAlternative

`func (o *VirtualMachineRestoreOptions) GetToAlternative() VirtualMachineToAlternativeRestoreOptions`

GetToAlternative returns the ToAlternative field if non-nil, zero value otherwise.

### GetToAlternativeOk

`func (o *VirtualMachineRestoreOptions) GetToAlternativeOk() (*VirtualMachineToAlternativeRestoreOptions, bool)`

GetToAlternativeOk returns a tuple with the ToAlternative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAlternative

`func (o *VirtualMachineRestoreOptions) SetToAlternative(v VirtualMachineToAlternativeRestoreOptions)`

SetToAlternative sets ToAlternative field to given value.

### HasToAlternative

`func (o *VirtualMachineRestoreOptions) HasToAlternative() bool`

HasToAlternative returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


