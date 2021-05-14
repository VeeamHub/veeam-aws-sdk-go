# AmazonConnectionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmazonAccountId** | Pointer to **string** |  | [optional] 
**AccessKey** | Pointer to **string** |  | [optional] 
**SecretKey** | Pointer to **string** |  | [optional] 
**RegionId** | **string** |  | 

## Methods

### NewAmazonConnectionSpec

`func NewAmazonConnectionSpec(regionId string, ) *AmazonConnectionSpec`

NewAmazonConnectionSpec instantiates a new AmazonConnectionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonConnectionSpecWithDefaults

`func NewAmazonConnectionSpecWithDefaults() *AmazonConnectionSpec`

NewAmazonConnectionSpecWithDefaults instantiates a new AmazonConnectionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmazonAccountId

`func (o *AmazonConnectionSpec) GetAmazonAccountId() string`

GetAmazonAccountId returns the AmazonAccountId field if non-nil, zero value otherwise.

### GetAmazonAccountIdOk

`func (o *AmazonConnectionSpec) GetAmazonAccountIdOk() (*string, bool)`

GetAmazonAccountIdOk returns a tuple with the AmazonAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonAccountId

`func (o *AmazonConnectionSpec) SetAmazonAccountId(v string)`

SetAmazonAccountId sets AmazonAccountId field to given value.

### HasAmazonAccountId

`func (o *AmazonConnectionSpec) HasAmazonAccountId() bool`

HasAmazonAccountId returns a boolean if a field has been set.

### GetAccessKey

`func (o *AmazonConnectionSpec) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *AmazonConnectionSpec) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *AmazonConnectionSpec) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *AmazonConnectionSpec) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *AmazonConnectionSpec) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *AmazonConnectionSpec) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *AmazonConnectionSpec) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *AmazonConnectionSpec) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetRegionId

`func (o *AmazonConnectionSpec) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *AmazonConnectionSpec) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *AmazonConnectionSpec) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


