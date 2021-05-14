# DiskBulkRestoreOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** |  | [optional] 
**AmazonAccountId** | Pointer to **string** |  | [optional] 
**AccessKey** | Pointer to **string** |  | [optional] 
**SecretKey** | Pointer to **string** |  | [optional] 
**Volumes** | [**[]DiskRestoreOptions**](DiskRestoreOptions.md) |  | 
**ToAlternative** | Pointer to [**DiskBulkRestoreToAlternativeOptions**](DiskBulkRestoreToAlternativeOptions.md) |  | [optional] 

## Methods

### NewDiskBulkRestoreOptions

`func NewDiskBulkRestoreOptions(volumes []DiskRestoreOptions, ) *DiskBulkRestoreOptions`

NewDiskBulkRestoreOptions instantiates a new DiskBulkRestoreOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskBulkRestoreOptionsWithDefaults

`func NewDiskBulkRestoreOptionsWithDefaults() *DiskBulkRestoreOptions`

NewDiskBulkRestoreOptionsWithDefaults instantiates a new DiskBulkRestoreOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *DiskBulkRestoreOptions) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DiskBulkRestoreOptions) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DiskBulkRestoreOptions) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DiskBulkRestoreOptions) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetAmazonAccountId

`func (o *DiskBulkRestoreOptions) GetAmazonAccountId() string`

GetAmazonAccountId returns the AmazonAccountId field if non-nil, zero value otherwise.

### GetAmazonAccountIdOk

`func (o *DiskBulkRestoreOptions) GetAmazonAccountIdOk() (*string, bool)`

GetAmazonAccountIdOk returns a tuple with the AmazonAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonAccountId

`func (o *DiskBulkRestoreOptions) SetAmazonAccountId(v string)`

SetAmazonAccountId sets AmazonAccountId field to given value.

### HasAmazonAccountId

`func (o *DiskBulkRestoreOptions) HasAmazonAccountId() bool`

HasAmazonAccountId returns a boolean if a field has been set.

### GetAccessKey

`func (o *DiskBulkRestoreOptions) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *DiskBulkRestoreOptions) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *DiskBulkRestoreOptions) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *DiskBulkRestoreOptions) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *DiskBulkRestoreOptions) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *DiskBulkRestoreOptions) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *DiskBulkRestoreOptions) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *DiskBulkRestoreOptions) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetVolumes

`func (o *DiskBulkRestoreOptions) GetVolumes() []DiskRestoreOptions`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *DiskBulkRestoreOptions) GetVolumesOk() (*[]DiskRestoreOptions, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *DiskBulkRestoreOptions) SetVolumes(v []DiskRestoreOptions)`

SetVolumes sets Volumes field to given value.


### GetToAlternative

`func (o *DiskBulkRestoreOptions) GetToAlternative() DiskBulkRestoreToAlternativeOptions`

GetToAlternative returns the ToAlternative field if non-nil, zero value otherwise.

### GetToAlternativeOk

`func (o *DiskBulkRestoreOptions) GetToAlternativeOk() (*DiskBulkRestoreToAlternativeOptions, bool)`

GetToAlternativeOk returns a tuple with the ToAlternative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAlternative

`func (o *DiskBulkRestoreOptions) SetToAlternative(v DiskBulkRestoreToAlternativeOptions)`

SetToAlternative sets ToAlternative field to given value.

### HasToAlternative

`func (o *DiskBulkRestoreOptions) HasToAlternative() bool`

HasToAlternative returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


